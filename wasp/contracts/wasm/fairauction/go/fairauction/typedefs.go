// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ArrayOfImmutableAgentID struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableAgentID) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableAgentID) GetAgentID(index uint32) wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(a.proxy.Index(index))
}

type ImmutableBidderList = ArrayOfImmutableAgentID

type ArrayOfMutableAgentID struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableAgentID) AppendAgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(a.proxy.Append())
}

func (a ArrayOfMutableAgentID) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableAgentID) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableAgentID) GetAgentID(index uint32) wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(a.proxy.Index(index))
}

type MutableBidderList = ArrayOfMutableAgentID

type MapAgentIDToImmutableBid struct {
	proxy wasmtypes.Proxy
}

func (m MapAgentIDToImmutableBid) GetBid(key wasmtypes.ScAgentID) ImmutableBid {
	return ImmutableBid{proxy: m.proxy.Key(wasmtypes.AgentIDToBytes(key))}
}

type ImmutableBids = MapAgentIDToImmutableBid

type MapAgentIDToMutableBid struct {
	proxy wasmtypes.Proxy
}

func (m MapAgentIDToMutableBid) Clear() {
	m.proxy.ClearMap()
}

func (m MapAgentIDToMutableBid) GetBid(key wasmtypes.ScAgentID) MutableBid {
	return MutableBid{proxy: m.proxy.Key(wasmtypes.AgentIDToBytes(key))}
}

type MutableBids = MapAgentIDToMutableBid
