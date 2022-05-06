// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmtypes

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

const (
	ScAddressEd25519 byte = 0 // length 32
	ScAddressNFT     byte = 1 // actually 16, length 20
	ScAddressAlias   byte = 2 // actually 8, length 20

	ScAddressLength = 33
)

type ScAddress struct {
	id [ScAddressLength]byte
}

func (o ScAddress) AsAgentID() ScAgentID {
	// agentID for address has Hname zero
	return NewScAgentID(o, 0)
}

func (o ScAddress) Bytes() []byte {
	return AddressToBytes(o)
}

func (o ScAddress) String() string {
	return AddressToString(o)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

func AddressDecode(dec *WasmDecoder) ScAddress {
	return addressFromBytesUnchecked(dec.FixedBytes(ScAddressLength))
}

func AddressEncode(enc *WasmEncoder, value ScAddress) {
	enc.FixedBytes(value.Bytes(), ScAddressLength)
}

func AddressFromBytes(buf []byte) ScAddress {
	if len(buf) == 0 {
		return ScAddress{}
	}
	if len(buf) != ScAddressLength {
		panic("invalid Address length")
	}
	if buf[0] > ScAddressAlias {
		panic("invalid Address type")
	}
	return addressFromBytesUnchecked(buf)
}

func AddressToBytes(value ScAddress) []byte {
	return value.id[:]
}

func AddressToString(value ScAddress) string {
	// TODO standardize human readable string
	return Base58Encode(value.id[:])
}

func addressFromBytesUnchecked(buf []byte) ScAddress {
	o := ScAddress{}
	copy(o.id[:], buf)
	return o
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAddress struct {
	proxy Proxy
}

func NewScImmutableAddress(proxy Proxy) ScImmutableAddress {
	return ScImmutableAddress{proxy: proxy}
}

func (o ScImmutableAddress) Exists() bool {
	return o.proxy.Exists()
}

func (o ScImmutableAddress) String() string {
	return AddressToString(o.Value())
}

func (o ScImmutableAddress) Value() ScAddress {
	return AddressFromBytes(o.proxy.Get())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddress struct {
	ScImmutableAddress
}

func NewScMutableAddress(proxy Proxy) ScMutableAddress {
	return ScMutableAddress{ScImmutableAddress{proxy: proxy}}
}

func (o ScMutableAddress) Delete() {
	o.proxy.Delete()
}

func (o ScMutableAddress) SetValue(value ScAddress) {
	o.proxy.Set(AddressToBytes(value))
}
