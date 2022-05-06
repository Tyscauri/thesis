// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use fairauction::*;
use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;
use crate::structs::*;
use crate::typedefs::*;

mod consts;
mod contract;
mod params;
mod results;
mod state;
mod structs;
mod typedefs;

mod fairauction;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_FINALIZE_AUCTION,
    	FUNC_PLACE_BID,
    	FUNC_SET_OWNER_MARGIN,
    	FUNC_START_AUCTION,
    	VIEW_GET_INFO,
	],
    funcs: &[
    	func_finalize_auction_thunk,
    	func_place_bid_thunk,
    	func_set_owner_margin_thunk,
    	func_start_auction_thunk,
	],
    views: &[
    	view_get_info_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	ScExports::call(index, &EXPORT_MAP);
}

#[no_mangle]
fn on_load() {
    ScExports::export(&EXPORT_MAP);
}

pub struct FinalizeAuctionContext {
	params: ImmutableFinalizeAuctionParams,
	state: MutableFairAuctionState,
}

fn func_finalize_auction_thunk(ctx: &ScFuncContext) {
	ctx.log("fairauction.funcFinalizeAuction");
	let f = FinalizeAuctionContext {
		params: ImmutableFinalizeAuctionParams { proxy: params_proxy() },
		state: MutableFairAuctionState { proxy: state_proxy() },
	};

	// only SC itself can invoke this function
	ctx.require(ctx.caller() == ctx.account_id(), "no permission");

	ctx.require(f.params.color().exists(), "missing mandatory color");
	func_finalize_auction(ctx, &f);
	ctx.log("fairauction.funcFinalizeAuction ok");
}

pub struct PlaceBidContext {
	params: ImmutablePlaceBidParams,
	state: MutableFairAuctionState,
}

fn func_place_bid_thunk(ctx: &ScFuncContext) {
	ctx.log("fairauction.funcPlaceBid");
	let f = PlaceBidContext {
		params: ImmutablePlaceBidParams { proxy: params_proxy() },
		state: MutableFairAuctionState { proxy: state_proxy() },
	};
	ctx.require(f.params.color().exists(), "missing mandatory color");
	func_place_bid(ctx, &f);
	ctx.log("fairauction.funcPlaceBid ok");
}

pub struct SetOwnerMarginContext {
	params: ImmutableSetOwnerMarginParams,
	state: MutableFairAuctionState,
}

fn func_set_owner_margin_thunk(ctx: &ScFuncContext) {
	ctx.log("fairauction.funcSetOwnerMargin");
	let f = SetOwnerMarginContext {
		params: ImmutableSetOwnerMarginParams { proxy: params_proxy() },
		state: MutableFairAuctionState { proxy: state_proxy() },
	};

	// only SC creator can set owner margin
	ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

	ctx.require(f.params.owner_margin().exists(), "missing mandatory ownerMargin");
	func_set_owner_margin(ctx, &f);
	ctx.log("fairauction.funcSetOwnerMargin ok");
}

pub struct StartAuctionContext {
	params: ImmutableStartAuctionParams,
	state: MutableFairAuctionState,
}

fn func_start_auction_thunk(ctx: &ScFuncContext) {
	ctx.log("fairauction.funcStartAuction");
	let f = StartAuctionContext {
		params: ImmutableStartAuctionParams { proxy: params_proxy() },
		state: MutableFairAuctionState { proxy: state_proxy() },
	};
	ctx.require(f.params.color().exists(), "missing mandatory color");
	ctx.require(f.params.minimum_bid().exists(), "missing mandatory minimumBid");
	func_start_auction(ctx, &f);
	ctx.log("fairauction.funcStartAuction ok");
}

pub struct GetInfoContext {
	params: ImmutableGetInfoParams,
	results: MutableGetInfoResults,
	state: ImmutableFairAuctionState,
}

fn view_get_info_thunk(ctx: &ScViewContext) {
	ctx.log("fairauction.viewGetInfo");
	let f = GetInfoContext {
		params: ImmutableGetInfoParams { proxy: params_proxy() },
		results: MutableGetInfoResults { proxy: results_proxy() },
		state: ImmutableFairAuctionState { proxy: state_proxy() },
	};
	ctx.require(f.params.color().exists(), "missing mandatory color");
	view_get_info(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("fairauction.viewGetInfo ok");
}