// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableApproveParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableApproveParams) Amount() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamAmount))
}

func (s ImmutableApproveParams) Delegation() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamDelegation))
}

type MutableApproveParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableApproveParams) Amount() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamAmount))
}

func (s MutableApproveParams) Delegation() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamDelegation))
}

type ImmutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableInitParams) Creator() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamCreator))
}

func (s ImmutableInitParams) Supply() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamSupply))
}

type MutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableInitParams) Creator() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamCreator))
}

func (s MutableInitParams) Supply() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamSupply))
}

type ImmutableTransferParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTransferParams) Account() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAccount))
}

func (s ImmutableTransferParams) Amount() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamAmount))
}

type MutableTransferParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTransferParams) Account() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAccount))
}

func (s MutableTransferParams) Amount() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamAmount))
}

type ImmutableTransferFromParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTransferFromParams) Account() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAccount))
}

func (s ImmutableTransferFromParams) Amount() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamAmount))
}

func (s ImmutableTransferFromParams) Recipient() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamRecipient))
}

type MutableTransferFromParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTransferFromParams) Account() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAccount))
}

func (s MutableTransferFromParams) Amount() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamAmount))
}

func (s MutableTransferFromParams) Recipient() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamRecipient))
}

type ImmutableAllowanceParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAllowanceParams) Account() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAccount))
}

func (s ImmutableAllowanceParams) Delegation() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamDelegation))
}

type MutableAllowanceParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAllowanceParams) Account() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAccount))
}

func (s MutableAllowanceParams) Delegation() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamDelegation))
}

type ImmutableBalanceOfParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBalanceOfParams) Account() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAccount))
}

type MutableBalanceOfParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBalanceOfParams) Account() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAccount))
}