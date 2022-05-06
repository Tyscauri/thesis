// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableControlAddressesResults extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultBlockIndex));
	}

	governingAddress(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ResultGoverningAddress));
	}

	stateControllerAddress(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ResultStateControllerAddress));
	}
}

export class MutableControlAddressesResults extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultBlockIndex));
	}

	governingAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ResultGoverningAddress));
	}

	stateControllerAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ResultStateControllerAddress));
	}
}

export class ImmutableGetBlockInfoResults extends wasmtypes.ScProxy {
	blockInfo(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultBlockInfo));
	}
}

export class MutableGetBlockInfoResults extends wasmtypes.ScProxy {
	blockInfo(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultBlockInfo));
	}
}

export class ArrayOfImmutableBytes extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getBytes(index: u32): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.index(index));
	}
}

export class ImmutableGetEventsForBlockResults extends wasmtypes.ScProxy {
	event(): sc.ArrayOfImmutableBytes {
		return new sc.ArrayOfImmutableBytes(this.proxy.root(sc.ResultEvent));
	}
}

export class ArrayOfMutableBytes extends wasmtypes.ScProxy {

	appendBytes(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getBytes(index: u32): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.index(index));
	}
}

export class MutableGetEventsForBlockResults extends wasmtypes.ScProxy {
	event(): sc.ArrayOfMutableBytes {
		return new sc.ArrayOfMutableBytes(this.proxy.root(sc.ResultEvent));
	}
}

export class ImmutableGetEventsForContractResults extends wasmtypes.ScProxy {
	event(): sc.ArrayOfImmutableBytes {
		return new sc.ArrayOfImmutableBytes(this.proxy.root(sc.ResultEvent));
	}
}

export class MutableGetEventsForContractResults extends wasmtypes.ScProxy {
	event(): sc.ArrayOfMutableBytes {
		return new sc.ArrayOfMutableBytes(this.proxy.root(sc.ResultEvent));
	}
}

export class ImmutableGetEventsForRequestResults extends wasmtypes.ScProxy {
	event(): sc.ArrayOfImmutableBytes {
		return new sc.ArrayOfImmutableBytes(this.proxy.root(sc.ResultEvent));
	}
}

export class MutableGetEventsForRequestResults extends wasmtypes.ScProxy {
	event(): sc.ArrayOfMutableBytes {
		return new sc.ArrayOfMutableBytes(this.proxy.root(sc.ResultEvent));
	}
}

export class ImmutableGetLatestBlockInfoResults extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultBlockIndex));
	}

	blockInfo(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultBlockInfo));
	}
}

export class MutableGetLatestBlockInfoResults extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultBlockIndex));
	}

	blockInfo(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultBlockInfo));
	}
}

export class ArrayOfImmutableRequestID extends wasmtypes.ScProxy {

	length(): u32 {
		return this.proxy.length();
	}

	getRequestID(index: u32): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.index(index));
	}
}

export class ImmutableGetRequestIDsForBlockResults extends wasmtypes.ScProxy {
	requestID(): sc.ArrayOfImmutableRequestID {
		return new sc.ArrayOfImmutableRequestID(this.proxy.root(sc.ResultRequestID));
	}
}

export class ArrayOfMutableRequestID extends wasmtypes.ScProxy {

	appendRequestID(): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.append());
	}

	clear(): void {
		this.proxy.clearArray();
	}

	length(): u32 {
		return this.proxy.length();
	}

	getRequestID(index: u32): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.index(index));
	}
}

export class MutableGetRequestIDsForBlockResults extends wasmtypes.ScProxy {
	requestID(): sc.ArrayOfMutableRequestID {
		return new sc.ArrayOfMutableRequestID(this.proxy.root(sc.ResultRequestID));
	}
}

export class ImmutableGetRequestReceiptResults extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultBlockIndex));
	}

	requestIndex(): wasmtypes.ScImmutableUint16 {
		return new wasmtypes.ScImmutableUint16(this.proxy.root(sc.ResultRequestIndex));
	}

	requestRecord(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultRequestRecord));
	}
}

export class MutableGetRequestReceiptResults extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultBlockIndex));
	}

	requestIndex(): wasmtypes.ScMutableUint16 {
		return new wasmtypes.ScMutableUint16(this.proxy.root(sc.ResultRequestIndex));
	}

	requestRecord(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultRequestRecord));
	}
}

export class ImmutableGetRequestReceiptsForBlockResults extends wasmtypes.ScProxy {
	requestRecord(): sc.ArrayOfImmutableBytes {
		return new sc.ArrayOfImmutableBytes(this.proxy.root(sc.ResultRequestRecord));
	}
}

export class MutableGetRequestReceiptsForBlockResults extends wasmtypes.ScProxy {
	requestRecord(): sc.ArrayOfMutableBytes {
		return new sc.ArrayOfMutableBytes(this.proxy.root(sc.ResultRequestRecord));
	}
}

export class ImmutableIsRequestProcessedResults extends wasmtypes.ScProxy {
	requestProcessed(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ResultRequestProcessed));
	}
}

export class MutableIsRequestProcessedResults extends wasmtypes.ScProxy {
	requestProcessed(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ResultRequestProcessed));
	}
}