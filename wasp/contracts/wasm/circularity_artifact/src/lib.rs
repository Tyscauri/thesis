// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use circularity_artifact::*;
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

mod circularity_artifact;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_ADD_PP_TO_FRACTION,
    	FUNC_ADD_RECYCLER,
    	FUNC_ADD_SORTER,
    	FUNC_CREATE_FRACTION,
    	FUNC_CREATE_PP,
    	FUNC_CREATE_RECYCLATE,
    	FUNC_DELETE_PP,
    	FUNC_INIT,
    	FUNC_PAYOUT_DONATION,
    	FUNC_PAYOUT_PRODUCER,
    	FUNC_SET_DONATION_ADDRESS,
    	FUNC_SET_OWNER,
    	VIEW_GET_AMOUNT_OF_REQUIRED_FUNDS,
    	VIEW_GET_DONATION_ADDRESS,
    	VIEW_GET_FRACTION,
    	VIEW_GET_MATERIALS,
    	VIEW_GET_OWNER,
    	VIEW_GET_PP,
    	VIEW_GET_RECYCLATE,
    	VIEW_GET_TOKEN_PER_PACKAGE,
	],
    funcs: &[
    	func_add_pp_to_fraction_thunk,
    	func_add_recycler_thunk,
    	func_add_sorter_thunk,
    	func_create_fraction_thunk,
    	func_create_pp_thunk,
    	func_create_recyclate_thunk,
    	func_delete_pp_thunk,
    	func_init_thunk,
    	func_payout_donation_thunk,
    	func_payout_producer_thunk,
    	func_set_donation_address_thunk,
    	func_set_owner_thunk,
	],
    views: &[
    	view_get_amount_of_required_funds_thunk,
    	view_get_donation_address_thunk,
    	view_get_fraction_thunk,
    	view_get_materials_thunk,
    	view_get_owner_thunk,
    	view_get_pp_thunk,
    	view_get_recyclate_thunk,
    	view_get_token_per_package_thunk,
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

pub struct AddPPToFractionContext {
	params: ImmutableAddPPToFractionParams,
	results: MutableAddPPToFractionResults,
	state: Mutablecircularity_artifactState,
}

fn func_add_pp_to_fraction_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcAddPPToFraction");
	let f = AddPPToFractionContext {
		params: ImmutableAddPPToFractionParams { proxy: params_proxy() },
		results: MutableAddPPToFractionResults { proxy: results_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	ctx.require(f.params.pp_id().exists(), "missing mandatory ppID");
	func_add_pp_to_fraction(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.funcAddPPToFraction ok");
}

pub struct AddRecyclerContext {
	params: ImmutableAddRecyclerParams,
	state: Mutablecircularity_artifactState,
}

fn func_add_recycler_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcAddRecycler");
	let f = AddRecyclerContext {
		params: ImmutableAddRecyclerParams { proxy: params_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.recycler_id().exists(), "missing mandatory recyclerID");
	func_add_recycler(ctx, &f);
	ctx.log("circularity_artifact.funcAddRecycler ok");
}

pub struct AddSorterContext {
	params: ImmutableAddSorterParams,
	state: Mutablecircularity_artifactState,
}

fn func_add_sorter_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcAddSorter");
	let f = AddSorterContext {
		params: ImmutableAddSorterParams { proxy: params_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.sorter_id().exists(), "missing mandatory sorterID");
	func_add_sorter(ctx, &f);
	ctx.log("circularity_artifact.funcAddSorter ok");
}

pub struct CreateFractionContext {
	params: ImmutableCreateFractionParams,
	results: MutableCreateFractionResults,
	state: Mutablecircularity_artifactState,
}

fn func_create_fraction_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcCreateFraction");
	let f = CreateFractionContext {
		params: ImmutableCreateFractionParams { proxy: params_proxy() },
		results: MutableCreateFractionResults { proxy: results_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.purpose().exists(), "missing mandatory purpose");
	func_create_fraction(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.funcCreateFraction ok");
}

pub struct CreatePPContext {
	params: ImmutableCreatePPParams,
	results: MutableCreatePPResults,
	state: Mutablecircularity_artifactState,
}

fn func_create_pp_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcCreatePP");
	let f = CreatePPContext {
		params: ImmutableCreatePPParams { proxy: params_proxy() },
		results: MutableCreatePPResults { proxy: results_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.expiry_date().exists(), "missing mandatory expiryDate");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.package_weight().exists(), "missing mandatory packageWeight");
	ctx.require(f.params.packages_number().exists(), "missing mandatory packagesNumber");
	ctx.require(f.params.purpose().exists(), "missing mandatory purpose");
	func_create_pp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.funcCreatePP ok");
}

pub struct CreateRecyclateContext {
	params: ImmutableCreateRecyclateParams,
	results: MutableCreateRecyclateResults,
	state: Mutablecircularity_artifactState,
}

fn func_create_recyclate_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcCreateRecyclate");
	let f = CreateRecyclateContext {
		params: ImmutableCreateRecyclateParams { proxy: params_proxy() },
		results: MutableCreateRecyclateResults { proxy: results_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	func_create_recyclate(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.funcCreateRecyclate ok");
}

pub struct DeletePPContext {
	params: ImmutableDeletePPParams,
	results: MutableDeletePPResults,
	state: Mutablecircularity_artifactState,
}

fn func_delete_pp_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcDeletePP");
	let f = DeletePPContext {
		params: ImmutableDeletePPParams { proxy: params_proxy() },
		results: MutableDeletePPResults { proxy: results_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.pp_id().exists(), "missing mandatory ppID");
	func_delete_pp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.funcDeletePP ok");
}

pub struct InitContext {
	params: ImmutableInitParams,
	state: Mutablecircularity_artifactState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcInit");
	let f = InitContext {
		params: ImmutableInitParams { proxy: params_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	func_init(ctx, &f);
	ctx.log("circularity_artifact.funcInit ok");
}

pub struct PayoutDonationContext {
	state: Mutablecircularity_artifactState,
}

fn func_payout_donation_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcPayoutDonation");
	let f = PayoutDonationContext {
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	func_payout_donation(ctx, &f);
	ctx.log("circularity_artifact.funcPayoutDonation ok");
}

pub struct PayoutProducerContext {
	params: ImmutablePayoutProducerParams,
	state: Mutablecircularity_artifactState,
}

fn func_payout_producer_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcPayoutProducer");
	let f = PayoutProducerContext {
		params: ImmutablePayoutProducerParams { proxy: params_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.prod_id().exists(), "missing mandatory prodID");
	func_payout_producer(ctx, &f);
	ctx.log("circularity_artifact.funcPayoutProducer ok");
}

pub struct SetDonationAddressContext {
	params: ImmutableSetDonationAddressParams,
	state: Mutablecircularity_artifactState,
}

fn func_set_donation_address_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcSetDonationAddress");
	let f = SetDonationAddressContext {
		params: ImmutableSetDonationAddressParams { proxy: params_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.donation_address().exists(), "missing mandatory donationAddress");
	func_set_donation_address(ctx, &f);
	ctx.log("circularity_artifact.funcSetDonationAddress ok");
}

pub struct SetOwnerContext {
	params: ImmutableSetOwnerParams,
	state: Mutablecircularity_artifactState,
}

fn func_set_owner_thunk(ctx: &ScFuncContext) {
	ctx.log("circularity_artifact.funcSetOwner");
	let f = SetOwnerContext {
		params: ImmutableSetOwnerParams { proxy: params_proxy() },
		state: Mutablecircularity_artifactState { proxy: state_proxy() },
	};

	// current owner of this smart contract
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.owner().exists(), "missing mandatory owner");
	func_set_owner(ctx, &f);
	ctx.log("circularity_artifact.funcSetOwner ok");
}

pub struct GetAmountOfRequiredFundsContext {
	params: ImmutableGetAmountOfRequiredFundsParams,
	results: MutableGetAmountOfRequiredFundsResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_amount_of_required_funds_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetAmountOfRequiredFunds");
	let f = GetAmountOfRequiredFundsContext {
		params: ImmutableGetAmountOfRequiredFundsParams { proxy: params_proxy() },
		results: MutableGetAmountOfRequiredFundsResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.charge_weight().exists(), "missing mandatory chargeWeight");
	view_get_amount_of_required_funds(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetAmountOfRequiredFunds ok");
}

pub struct GetDonationAddressContext {
	results: MutableGetDonationAddressResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_donation_address_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetDonationAddress");
	let f = GetDonationAddressContext {
		results: MutableGetDonationAddressResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	view_get_donation_address(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetDonationAddress ok");
}

pub struct GetFractionContext {
	params: ImmutableGetFractionParams,
	results: MutableGetFractionResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_fraction_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetFraction");
	let f = GetFractionContext {
		params: ImmutableGetFractionParams { proxy: params_proxy() },
		results: MutableGetFractionResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	view_get_fraction(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetFraction ok");
}

pub struct GetMaterialsContext {
	params: ImmutableGetMaterialsParams,
	results: MutableGetMaterialsResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_materials_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetMaterials");
	let f = GetMaterialsContext {
		params: ImmutableGetMaterialsParams { proxy: params_proxy() },
		results: MutableGetMaterialsResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.id().exists(), "missing mandatory id");
	view_get_materials(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetMaterials ok");
}

pub struct GetOwnerContext {
	results: MutableGetOwnerResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_owner_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetOwner");
	let f = GetOwnerContext {
		results: MutableGetOwnerResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	view_get_owner(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetOwner ok");
}

pub struct GetPPContext {
	params: ImmutableGetPPParams,
	results: MutableGetPPResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_pp_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetPP");
	let f = GetPPContext {
		params: ImmutableGetPPParams { proxy: params_proxy() },
		results: MutableGetPPResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.id().exists(), "missing mandatory id");
	view_get_pp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetPP ok");
}

pub struct GetRecyclateContext {
	params: ImmutableGetRecyclateParams,
	results: MutableGetRecyclateResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_recyclate_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetRecyclate");
	let f = GetRecyclateContext {
		params: ImmutableGetRecyclateParams { proxy: params_proxy() },
		results: MutableGetRecyclateResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	ctx.require(f.params.recy_id().exists(), "missing mandatory recyID");
	view_get_recyclate(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetRecyclate ok");
}

pub struct GetTokenPerPackageContext {
	params: ImmutableGetTokenPerPackageParams,
	results: MutableGetTokenPerPackageResults,
	state: Immutablecircularity_artifactState,
}

fn view_get_token_per_package_thunk(ctx: &ScViewContext) {
	ctx.log("circularity_artifact.viewGetTokenPerPackage");
	let f = GetTokenPerPackageContext {
		params: ImmutableGetTokenPerPackageParams { proxy: params_proxy() },
		results: MutableGetTokenPerPackageResults { proxy: results_proxy() },
		state: Immutablecircularity_artifactState { proxy: state_proxy() },
	};
	view_get_token_per_package(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("circularity_artifact.viewGetTokenPerPackage ok");
}