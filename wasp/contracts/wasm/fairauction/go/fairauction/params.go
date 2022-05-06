// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableFinalizeAuctionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFinalizeAuctionParams) Color() wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(s.proxy.Root(ParamColor))
}

type MutableFinalizeAuctionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFinalizeAuctionParams) Color() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(s.proxy.Root(ParamColor))
}

type ImmutablePlaceBidParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablePlaceBidParams) Color() wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(s.proxy.Root(ParamColor))
}

type MutablePlaceBidParams struct {
	proxy wasmtypes.Proxy
}

func (s MutablePlaceBidParams) Color() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(s.proxy.Root(ParamColor))
}

type ImmutableSetOwnerMarginParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSetOwnerMarginParams) OwnerMargin() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamOwnerMargin))
}

type MutableSetOwnerMarginParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSetOwnerMarginParams) OwnerMargin() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamOwnerMargin))
}

type ImmutableStartAuctionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStartAuctionParams) Color() wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(s.proxy.Root(ParamColor))
}

func (s ImmutableStartAuctionParams) Description() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamDescription))
}

func (s ImmutableStartAuctionParams) Duration() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamDuration))
}

func (s ImmutableStartAuctionParams) MinimumBid() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamMinimumBid))
}

type MutableStartAuctionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStartAuctionParams) Color() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(s.proxy.Root(ParamColor))
}

func (s MutableStartAuctionParams) Description() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamDescription))
}

func (s MutableStartAuctionParams) Duration() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamDuration))
}

func (s MutableStartAuctionParams) MinimumBid() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamMinimumBid))
}

type ImmutableGetInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetInfoParams) Color() wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(s.proxy.Root(ParamColor))
}

type MutableGetInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetInfoParams) Color() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(s.proxy.Root(ParamColor))
}
