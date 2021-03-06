// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME        : &str = "erc721";
pub const SC_DESCRIPTION : &str = "ERC-721 NFT PoC for IOTA Smart Contracts";
pub const HSC_NAME       : ScHname = ScHname(0xd967c216);

pub const PARAM_APPROVAL  : &str = "approval";
pub const PARAM_APPROVED  : &str = "approved";
pub const PARAM_DATA      : &str = "data";
pub const PARAM_FROM      : &str = "from";
pub const PARAM_NAME      : &str = "n";
pub const PARAM_OPERATOR  : &str = "operator";
pub const PARAM_OWNER     : &str = "owner";
pub const PARAM_SYMBOL    : &str = "s";
pub const PARAM_TO        : &str = "to";
pub const PARAM_TOKEN_ID  : &str = "tokenID";
pub const PARAM_TOKEN_URI : &str = "tokenURI";

pub const RESULT_AMOUNT    : &str = "amount";
pub const RESULT_APPROVAL  : &str = "approval";
pub const RESULT_APPROVED  : &str = "approved";
pub const RESULT_NAME      : &str = "name";
pub const RESULT_OWNER     : &str = "owner";
pub const RESULT_SYMBOL    : &str = "symbol";
pub const RESULT_TOKEN_URI : &str = "tokenURI";

pub const STATE_APPROVED_ACCOUNTS  : &str = "approvedAccounts";
pub const STATE_APPROVED_OPERATORS : &str = "approvedOperators";
pub const STATE_BALANCES           : &str = "balances";
pub const STATE_NAME               : &str = "name";
pub const STATE_OWNERS             : &str = "owners";
pub const STATE_SYMBOL             : &str = "symbol";
pub const STATE_TOKEN_UR_IS        : &str = "tokenURIs";

pub const FUNC_APPROVE              : &str = "approve";
pub const FUNC_BURN                 : &str = "burn";
pub const FUNC_INIT                 : &str = "init";
pub const FUNC_MINT                 : &str = "mint";
pub const FUNC_SAFE_TRANSFER_FROM   : &str = "safeTransferFrom";
pub const FUNC_SET_APPROVAL_FOR_ALL : &str = "setApprovalForAll";
pub const FUNC_TRANSFER_FROM        : &str = "transferFrom";
pub const VIEW_BALANCE_OF           : &str = "balanceOf";
pub const VIEW_GET_APPROVED         : &str = "getApproved";
pub const VIEW_IS_APPROVED_FOR_ALL  : &str = "isApprovedForAll";
pub const VIEW_NAME                 : &str = "name";
pub const VIEW_OWNER_OF             : &str = "ownerOf";
pub const VIEW_SYMBOL               : &str = "symbol";
pub const VIEW_TOKEN_URI            : &str = "tokenURI";

pub const HFUNC_APPROVE              : ScHname = ScHname(0xa0661268);
pub const HFUNC_BURN                 : ScHname = ScHname(0x7bc1efb1);
pub const HFUNC_INIT                 : ScHname = ScHname(0x1f44d644);
pub const HFUNC_MINT                 : ScHname = ScHname(0xa29addcf);
pub const HFUNC_SAFE_TRANSFER_FROM   : ScHname = ScHname(0x130ce158);
pub const HFUNC_SET_APPROVAL_FOR_ALL : ScHname = ScHname(0xb8d8c776);
pub const HFUNC_TRANSFER_FROM        : ScHname = ScHname(0xd5e0a602);
pub const HVIEW_BALANCE_OF           : ScHname = ScHname(0x67ef8df4);
pub const HVIEW_GET_APPROVED         : ScHname = ScHname(0xbe34b6ba);
pub const HVIEW_IS_APPROVED_FOR_ALL  : ScHname = ScHname(0x3251b0f0);
pub const HVIEW_NAME                 : ScHname = ScHname(0x0df7da3a);
pub const HVIEW_OWNER_OF             : ScHname = ScHname(0x1246f5ad);
pub const HVIEW_SYMBOL               : ScHname = ScHname(0x3e93d19b);
pub const HVIEW_TOKEN_URI            : ScHname = ScHname(0x4e1a7397);
