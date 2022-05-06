// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package helloworld

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableHelloWorldState struct {
	proxy wasmtypes.Proxy
}

type MutableHelloWorldState struct {
	proxy wasmtypes.Proxy
}

func (s MutableHelloWorldState) AsImmutable() ImmutableHelloWorldState {
	return ImmutableHelloWorldState(s)
}