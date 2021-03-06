// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc721

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableApproveParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableApproveParams) Approved() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamApproved))
}

func (s ImmutableApproveParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableApproveParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableApproveParams) Approved() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamApproved))
}

func (s MutableApproveParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

type ImmutableBurnParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBurnParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableBurnParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBurnParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

type ImmutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableInitParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableInitParams) Symbol() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamSymbol))
}

type MutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableInitParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableInitParams) Symbol() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamSymbol))
}

type ImmutableMintParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMintParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

func (s ImmutableMintParams) TokenURI() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamTokenURI))
}

type MutableMintParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMintParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

func (s MutableMintParams) TokenURI() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamTokenURI))
}

type ImmutableSafeTransferFromParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSafeTransferFromParams) Data() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamData))
}

func (s ImmutableSafeTransferFromParams) From() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamFrom))
}

func (s ImmutableSafeTransferFromParams) To() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamTo))
}

func (s ImmutableSafeTransferFromParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableSafeTransferFromParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSafeTransferFromParams) Data() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamData))
}

func (s MutableSafeTransferFromParams) From() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamFrom))
}

func (s MutableSafeTransferFromParams) To() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamTo))
}

func (s MutableSafeTransferFromParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

type ImmutableSetApprovalForAllParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSetApprovalForAllParams) Approval() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamApproval))
}

func (s ImmutableSetApprovalForAllParams) Operator() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOperator))
}

type MutableSetApprovalForAllParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSetApprovalForAllParams) Approval() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamApproval))
}

func (s MutableSetApprovalForAllParams) Operator() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOperator))
}

type ImmutableTransferFromParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTransferFromParams) From() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamFrom))
}

func (s ImmutableTransferFromParams) To() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamTo))
}

func (s ImmutableTransferFromParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableTransferFromParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTransferFromParams) From() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamFrom))
}

func (s MutableTransferFromParams) To() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamTo))
}

func (s MutableTransferFromParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

type ImmutableBalanceOfParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBalanceOfParams) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOwner))
}

type MutableBalanceOfParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBalanceOfParams) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOwner))
}

type ImmutableGetApprovedParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetApprovedParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableGetApprovedParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetApprovedParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

type ImmutableIsApprovedForAllParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableIsApprovedForAllParams) Operator() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOperator))
}

func (s ImmutableIsApprovedForAllParams) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOwner))
}

type MutableIsApprovedForAllParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableIsApprovedForAllParams) Operator() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOperator))
}

func (s MutableIsApprovedForAllParams) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOwner))
}

type ImmutableOwnerOfParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableOwnerOfParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableOwnerOfParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableOwnerOfParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}

type ImmutableTokenURIParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTokenURIParams) TokenID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamTokenID))
}

type MutableTokenURIParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTokenURIParams) TokenID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamTokenID))
}
