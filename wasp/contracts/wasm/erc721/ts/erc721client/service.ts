// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmclient from "wasmclient"
import * as events from "./events"

const ArgApproval = "approval";
const ArgApproved = "approved";
const ArgData = "data";
const ArgFrom = "from";
const ArgName = "n";
const ArgOperator = "operator";
const ArgOwner = "owner";
const ArgSymbol = "s";
const ArgTo = "to";
const ArgTokenID = "tokenID";
const ArgTokenURI = "tokenURI";

const ResAmount = "amount";
const ResApproval = "approval";
const ResApproved = "approved";
const ResName = "name";
const ResOwner = "owner";
const ResSymbol = "symbol";
const ResTokenURI = "tokenURI";

///////////////////////////// approve /////////////////////////////

export class ApproveFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public approved(v: wasmclient.AgentID): void {
		this.args.set(ArgApproved, this.args.fromAgentID(v));
	}
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgTokenID);
		return await super.post(0xa0661268, this.args);
	}
}

///////////////////////////// burn /////////////////////////////

export class BurnFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgTokenID);
		return await super.post(0x7bc1efb1, this.args);
	}
}

///////////////////////////// init /////////////////////////////

export class InitFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public name(v: string): void {
		this.args.set(ArgName, this.args.fromString(v));
	}
	
	public symbol(v: string): void {
		this.args.set(ArgSymbol, this.args.fromString(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgName);
		this.args.mandatory(ArgSymbol);
		return await super.post(0x1f44d644, this.args);
	}
}

///////////////////////////// mint /////////////////////////////

export class MintFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}
	
	public tokenURI(v: string): void {
		this.args.set(ArgTokenURI, this.args.fromString(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgTokenID);
		return await super.post(0xa29addcf, this.args);
	}
}

///////////////////////////// safeTransferFrom /////////////////////////////

export class SafeTransferFromFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public data(v: wasmclient.Bytes): void {
		this.args.set(ArgData, this.args.fromBytes(v));
	}
	
	public from(v: wasmclient.AgentID): void {
		this.args.set(ArgFrom, this.args.fromAgentID(v));
	}
	
	public to(v: wasmclient.AgentID): void {
		this.args.set(ArgTo, this.args.fromAgentID(v));
	}
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgFrom);
		this.args.mandatory(ArgTo);
		this.args.mandatory(ArgTokenID);
		return await super.post(0x130ce158, this.args);
	}
}

///////////////////////////// setApprovalForAll /////////////////////////////

export class SetApprovalForAllFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public approval(v: boolean): void {
		this.args.set(ArgApproval, this.args.fromBool(v));
	}
	
	public operator(v: wasmclient.AgentID): void {
		this.args.set(ArgOperator, this.args.fromAgentID(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgApproval);
		this.args.mandatory(ArgOperator);
		return await super.post(0xb8d8c776, this.args);
	}
}

///////////////////////////// transferFrom /////////////////////////////

export class TransferFromFunc extends wasmclient.ClientFunc {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public from(v: wasmclient.AgentID): void {
		this.args.set(ArgFrom, this.args.fromAgentID(v));
	}
	
	public to(v: wasmclient.AgentID): void {
		this.args.set(ArgTo, this.args.fromAgentID(v));
	}
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}
	
	public async post(): Promise<wasmclient.RequestID> {
		this.args.mandatory(ArgFrom);
		this.args.mandatory(ArgTo);
		this.args.mandatory(ArgTokenID);
		return await super.post(0xd5e0a602, this.args);
	}
}

///////////////////////////// balanceOf /////////////////////////////

export class BalanceOfView extends wasmclient.ClientView {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public owner(v: wasmclient.AgentID): void {
		this.args.set(ArgOwner, this.args.fromAgentID(v));
	}

	public async call(): Promise<BalanceOfResults> {
		this.args.mandatory(ArgOwner);
		const res = new BalanceOfResults();
		await this.callView("balanceOf", this.args, res);
		return res;
	}
}

export class BalanceOfResults extends wasmclient.Results {
	
	amountExists(): boolean {
		return this.exists(ResAmount)
	}

	amount(): wasmclient.Uint64 {
		return this.toUint64(this.get(ResAmount));
	}
}

///////////////////////////// getApproved /////////////////////////////

export class GetApprovedView extends wasmclient.ClientView {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}

	public async call(): Promise<GetApprovedResults> {
		this.args.mandatory(ArgTokenID);
		const res = new GetApprovedResults();
		await this.callView("getApproved", this.args, res);
		return res;
	}
}

export class GetApprovedResults extends wasmclient.Results {
	
	approvedExists(): boolean {
		return this.exists(ResApproved)
	}

	approved(): wasmclient.AgentID {
		return this.toAgentID(this.get(ResApproved));
	}
}

///////////////////////////// isApprovedForAll /////////////////////////////

export class IsApprovedForAllView extends wasmclient.ClientView {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public operator(v: wasmclient.AgentID): void {
		this.args.set(ArgOperator, this.args.fromAgentID(v));
	}
	
	public owner(v: wasmclient.AgentID): void {
		this.args.set(ArgOwner, this.args.fromAgentID(v));
	}

	public async call(): Promise<IsApprovedForAllResults> {
		this.args.mandatory(ArgOperator);
		this.args.mandatory(ArgOwner);
		const res = new IsApprovedForAllResults();
		await this.callView("isApprovedForAll", this.args, res);
		return res;
	}
}

export class IsApprovedForAllResults extends wasmclient.Results {
	
	approvalExists(): boolean {
		return this.exists(ResApproval)
	}

	approval(): boolean {
		return this.toBool(this.get(ResApproval));
	}
}

///////////////////////////// name /////////////////////////////

export class NameView extends wasmclient.ClientView {

	public async call(): Promise<NameResults> {
		const res = new NameResults();
		await this.callView("name", null, res);
		return res;
	}
}

export class NameResults extends wasmclient.Results {

	name(): string {
		return this.toString(this.get(ResName));
	}
}

///////////////////////////// ownerOf /////////////////////////////

export class OwnerOfView extends wasmclient.ClientView {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}

	public async call(): Promise<OwnerOfResults> {
		this.args.mandatory(ArgTokenID);
		const res = new OwnerOfResults();
		await this.callView("ownerOf", this.args, res);
		return res;
	}
}

export class OwnerOfResults extends wasmclient.Results {
	
	ownerExists(): boolean {
		return this.exists(ResOwner)
	}

	owner(): wasmclient.AgentID {
		return this.toAgentID(this.get(ResOwner));
	}
}

///////////////////////////// symbol /////////////////////////////

export class SymbolView extends wasmclient.ClientView {

	public async call(): Promise<SymbolResults> {
		const res = new SymbolResults();
		await this.callView("symbol", null, res);
		return res;
	}
}

export class SymbolResults extends wasmclient.Results {

	symbol(): string {
		return this.toString(this.get(ResSymbol));
	}
}

///////////////////////////// tokenURI /////////////////////////////

export class TokenURIView extends wasmclient.ClientView {
	private args: wasmclient.Arguments = new wasmclient.Arguments();
	
	public tokenID(v: wasmclient.Hash): void {
		this.args.set(ArgTokenID, this.args.fromHash(v));
	}

	public async call(): Promise<TokenURIResults> {
		this.args.mandatory(ArgTokenID);
		const res = new TokenURIResults();
		await this.callView("tokenURI", this.args, res);
		return res;
	}
}

export class TokenURIResults extends wasmclient.Results {
	
	tokenURIExists(): boolean {
		return this.exists(ResTokenURI)
	}

	tokenURI(): string {
		return this.toString(this.get(ResTokenURI));
	}
}

///////////////////////////// Erc721Service /////////////////////////////

export class Erc721Service extends wasmclient.Service {

	public constructor(cl: wasmclient.ServiceClient) {
		super(cl, 0xd967c216);
	}

	public newEventHandler(): events.Erc721Events {
		return new events.Erc721Events();
	}

	public approve(): ApproveFunc {
		return new ApproveFunc(this);
	}

	public burn(): BurnFunc {
		return new BurnFunc(this);
	}

	public init(): InitFunc {
		return new InitFunc(this);
	}

	public mint(): MintFunc {
		return new MintFunc(this);
	}

	public safeTransferFrom(): SafeTransferFromFunc {
		return new SafeTransferFromFunc(this);
	}

	public setApprovalForAll(): SetApprovalForAllFunc {
		return new SetApprovalForAllFunc(this);
	}

	public transferFrom(): TransferFromFunc {
		return new TransferFromFunc(this);
	}

	public balanceOf(): BalanceOfView {
		return new BalanceOfView(this);
	}

	public getApproved(): GetApprovedView {
		return new GetApprovedView(this);
	}

	public isApprovedForAll(): IsApprovedForAllView {
		return new IsApprovedForAllView(this);
	}

	public name(): NameView {
		return new NameView(this);
	}

	public ownerOf(): OwnerOfView {
		return new OwnerOfView(this);
	}

	public symbol(): SymbolView {
		return new SymbolView(this);
	}

	public tokenURI(): TokenURIView {
		return new TokenURIView(this);
	}
}
