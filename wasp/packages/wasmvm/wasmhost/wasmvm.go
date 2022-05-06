// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

const (
	defaultTimeout   = 5 * time.Second
	FuncAbort        = "abort"
	FuncFdWrite      = "fd_write"
	FuncHostStateGet = "hostStateGet"
	FuncHostStateSet = "hostStateSet"
	ModuleEnv        = "env"
	ModuleWasi1      = "wasi_unstable"
	ModuleWasi2      = "wasi_snapshot_preview1"
	ModuleWasmLib    = "WasmLib"
)

var (
	// DisableWasmTimeout can be used to disable the annoying timeout during debugging
	DisableWasmTimeout = false

	// HostTracing turns on debug tracing for ScHost calls
	HostTracing = false

	// WasmTimeout set this to non-zero for a one-time override of the defaultTimeout
	WasmTimeout = 0 * time.Second
)

type WasmVM interface {
	Instantiate() error
	Interrupt()
	LinkHost(proc *WasmProcessor) error
	LoadWasm(wasmData []byte) error
	NewInstance() WasmVM
	RunFunction(functionName string, args ...interface{}) error
	RunScFunction(index int32) error
	UnsafeMemory() []byte
	VMGetBytes(offset int32, size int32) []byte
	VMGetSize() int32
	VMSetBytes(offset int32, size int32, bytes []byte) int32
}

type WasmVMBase struct {
	proc           *WasmProcessor
	panicErr       error
	cachedResult   []byte
	timeoutStarted bool
}

// catchPanicMessage is used in every host function to catch any panic.
// It will save the first panic it encounters in the WasmVMBase so that
// the caller of the Wasm function can retrieve the correct error.
// This is a workaround to WasmTime saving the *last* panic instead of
// the first, thereby reporting the wrong panic error sometimes
func (vm *WasmVMBase) catchPanicMessage() {
	panicMsg := recover()
	if panicMsg == nil {
		return
	}
	if vm.panicErr == nil {
		switch msg := panicMsg.(type) {
		case error:
			vm.panicErr = msg
		default:
			vm.panicErr = fmt.Errorf("%v", msg)
		}
	}
	// rethrow and let nature run its course...
	panic(panicMsg)
}

func (vm *WasmVMBase) getContext(id int32) *WasmContext {
	return vm.proc.GetContext(id)
}

func (vm *WasmVMBase) HostAbort(errMsg, fileName, line, col int32) {
	// crude implementation assumes texts to only use ASCII part of UTF-16

	defer vm.catchPanicMessage()
	impl := vm.proc.vm

	// null-terminated UTF-16 error message
	str1 := make([]byte, 0)
	ptr := impl.VMGetBytes(errMsg, 2)
	for i := errMsg; ptr[0] != 0; i += 2 {
		str1 = append(str1, ptr[0])
		ptr = impl.VMGetBytes(i, 2)
	}

	// null-terminated UTF-16 file name
	str2 := make([]byte, 0)
	ptr = impl.VMGetBytes(fileName, 2)
	for i := fileName; ptr[0] != 0; i += 2 {
		str2 = append(str2, ptr[0])
		ptr = impl.VMGetBytes(i, 2)
	}

	panic(fmt.Sprintf("AssemblyScript panic: %s (%s %d:%d)", string(str1), string(str2), line, col))
}

func (vm *WasmVMBase) HostFdWrite(_fd, iovs, _size, written int32) int32 {
	defer vm.catchPanicMessage()
	impl := vm.proc.vm

	ctx := vm.getContext(0)
	ctx.log().Debugf("HostFdWrite(...)")

	// very basic implementation that expects fd to be stdout and iovs to be only one element
	ptr := impl.VMGetBytes(iovs, 8)
	text := int32(binary.LittleEndian.Uint32(ptr[0:4]))
	size := int32(binary.LittleEndian.Uint32(ptr[4:8]))
	// msg := vm.impl.VMGetBytes(text, size)
	// fmt.Print(string(msg))
	ptr = make([]byte, 4)
	binary.LittleEndian.PutUint32(ptr, uint32(size))
	impl.VMSetBytes(written, size, ptr)

	// strip off "panic: " prefix and call sandbox panic function
	vm.HostStateGet(0, wasmlib.FnPanic, text+7, size)
	return size
}

func (vm *WasmVMBase) HostStateGet(keyRef, keyLen, valRef, valLen int32) int32 {
	defer vm.catchPanicMessage()
	impl := vm.proc.vm

	ctx := vm.getContext(0)
	if HostTracing {
		vm.traceGet(ctx, keyRef, keyLen, valRef, valLen)
	}

	// only check for existence ?
	if valLen < 0 {
		key := impl.VMGetBytes(keyRef, keyLen)
		if ctx.StateExists(key) {
			return 0
		}
		// missing key is indicated by -1
		return -1
	}

	//  get value for key request, or get cached result request (keyLen == 0)
	if keyLen >= 0 {
		if keyLen > 0 {
			// retrieve value associated with key
			key := impl.VMGetBytes(keyRef, keyLen)
			vm.cachedResult = ctx.StateGet(key)
		}
		if vm.cachedResult == nil {
			return -1
		}
		return impl.VMSetBytes(valRef, valLen, vm.cachedResult)
	}

	// sandbox func call request, keyLen is func nr
	params := impl.VMGetBytes(valRef, valLen)
	vm.cachedResult = ctx.Sandbox(keyLen, params)
	return int32(len(vm.cachedResult))
}

func (vm *WasmVMBase) HostStateSet(keyRef, keyLen, valRef, valLen int32) {
	defer vm.catchPanicMessage()
	impl := vm.proc.vm

	ctx := vm.getContext(0)
	if HostTracing {
		vm.traceSet(ctx, keyRef, keyLen, valRef, valLen)
	}

	// export name?
	if keyRef == 0 {
		name := string(impl.VMGetBytes(valRef, valLen))
		if keyLen < 0 {
			// ExportWasmTag, log the wasm tag name
			ctx.proc.log.Infof(name)
			return
		}
		ctx.ExportName(keyLen, name)
		return
	}

	key := impl.VMGetBytes(keyRef, keyLen)

	// delete key ?
	if valLen < 0 {
		ctx.StateDelete(key)
		return
	}

	// set key
	value := impl.VMGetBytes(valRef, valLen)
	ctx.StateSet(key, value)
}

func (vm *WasmVMBase) Instantiate() error {
	return errors.New("cannot be cloned")
}

func (vm *WasmVMBase) LinkHost(proc *WasmProcessor) error {
	// trick vm into thinking it doesn't have to start the timeout timer
	// useful when debugging to prevent timing out on breakpoints
	vm.timeoutStarted = DisableWasmTimeout

	vm.proc = proc
	return nil
}

func (vm *WasmVMBase) Run(runner func() error) (err error) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		// could be the wrong panic message due to a WasmTime bug, so we always
		// rethrow our intercepted first panic instead of WasmTime's last panic
		if vm.panicErr != nil {
			panic(vm.panicErr)
		}
		panic(r)
	}()

	if vm.timeoutStarted {
		// no need to wrap nested calls in timeout code
		err = runner()
		if vm.panicErr != nil {
			err = vm.panicErr
			vm.panicErr = nil
		}
		return err
	}

	timeout := defaultTimeout
	if WasmTimeout != 0 {
		timeout = WasmTimeout
		WasmTimeout = 0
	}

	done := make(chan bool, 2)

	// start timeout handler
	go func() {
		select {
		case <-done: // runner was done before timeout
		case <-time.After(timeout):
			// timeout: interrupt Wasm
			vm.proc.vm.Interrupt()
			// wait for runner to finish
			<-done
		}
	}()

	vm.timeoutStarted = true
	err = runner()
	done <- true
	vm.timeoutStarted = false
	if vm.panicErr != nil {
		err = vm.panicErr
		vm.panicErr = nil
	}
	return err
}

func (vm *WasmVMBase) VMGetBytes(offset, size int32) []byte {
	ptr := vm.proc.vm.UnsafeMemory()
	bytes := make([]byte, size)
	copy(bytes, ptr[offset:offset+size])
	return bytes
}

func (vm *WasmVMBase) VMGetSize() int32 {
	ptr := vm.proc.vm.UnsafeMemory()
	return int32(len(ptr))
}

func (vm *WasmVMBase) VMSetBytes(offset, size int32, bytes []byte) int32 {
	if size != 0 {
		ptr := vm.proc.vm.UnsafeMemory()
		copy(ptr[offset:offset+size], bytes)
	}
	return int32(len(bytes))
}

func (vm *WasmVMBase) traceGet(ctx *WasmContext, keyRef, keyLen, valRef, valLen int32) {
	impl := vm.proc.vm

	// only check for existence ?
	if valLen < 0 {
		key := impl.VMGetBytes(keyRef, keyLen)
		ctx.log().Debugf("StateExists(%s) = %v", vm.traceKey(key), ctx.StateExists(key))
		return
	}

	// get value for key request, or get cached result request (keyLen == 0)
	if keyLen >= 0 {
		if keyLen == 0 {
			ctx.log().Debugf("  => %s", vm.traceVal(vm.cachedResult))
			return
		}
		// retrieve value associated with key
		key := impl.VMGetBytes(keyRef, keyLen)
		ctx.log().Debugf("StateGet(%s)", vm.traceKey(key))
		return
	}

	// sandbox func call request, keyLen is func nr
	if keyLen == wasmlib.FnLog {
		return
	}
	params := impl.VMGetBytes(valRef, valLen)
	ctx.log().Debugf("Sandbox(%s)", vm.traceSandbox(keyLen, params))
}

func (vm *WasmVMBase) traceSandbox(funcNr int32, params []byte) string {
	name := sandboxFuncNames[-funcNr]
	if name[0] == '$' {
		return name[1:] + ", " + string(params)
	}
	if name[0] != '#' {
		return name
	}
	return name[1:] + ", " + hex(params)
}

func (vm *WasmVMBase) traceSet(ctx *WasmContext, keyRef, keyLen, valRef, valLen int32) {
	impl := vm.proc.vm

	// export name?
	if keyRef == 0 {
		name := string(impl.VMGetBytes(valRef, valLen))
		ctx.log().Debugf("ExportName(%d, %s)", keyLen, name)
		return
	}

	key := impl.VMGetBytes(keyRef, keyLen)

	// delete key ?
	if valLen < 0 {
		ctx.log().Debugf("StateDelete(%s)", vm.traceKey(key))
		return
	}

	// set key
	val := impl.VMGetBytes(valRef, valLen)
	ctx.log().Debugf("StateSet(%s, %s)", vm.traceKey(key), vm.traceVal(val))
}

func (vm *WasmVMBase) traceKey(key []byte) string {
	name := ""
	for i, b := range key {
		if b == '.' {
			return string(key[:i+1]) + hex(key[i+1:])
		}
		if b == '#' {
			name = string(key[:i+1])
			j := i + 1
			for ; (key[j] & 0x80) != 0; j++ {
			}
			dec := wasmtypes.NewWasmDecoder(key[i+1 : j+1])
			index := wasmtypes.Uint64Decode(dec)
			name += wasmtypes.Uint64ToString(index)
			if j+1 == len(key) {
				return name
			}
			return name + "..." + hex(key[j+1:])
		}
	}
	return `"` + string(key) + `"`
}

func (vm *WasmVMBase) traceVal(val []byte) string {
	for _, b := range val {
		if b < ' ' || b > '~' {
			return hex(val)
		}
	}
	return string(val)
}

// hex returns a hex string representing the byte buffer
func hex(buf []byte) string {
	const hexa = "0123456789abcdef"
	res := make([]byte, len(buf)*2)
	for i, b := range buf {
		res[i*2] = hexa[b>>4]
		res[i*2+1] = hexa[b&0x0f]
	}
	return string(res)
}
