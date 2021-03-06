// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableArrayLengthResults extends wasmtypes.ScProxy {
	length(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultLength));
	}
}

export class MutableArrayLengthResults extends wasmtypes.ScProxy {
	length(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultLength));
	}
}

export class ImmutableArrayValueResults extends wasmtypes.ScProxy {
	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ResultValue));
	}
}

export class MutableArrayValueResults extends wasmtypes.ScProxy {
	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ResultValue));
	}
}

export class ImmutableBlockRecordResults extends wasmtypes.ScProxy {
	record(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ResultRecord));
	}
}

export class MutableBlockRecordResults extends wasmtypes.ScProxy {
	record(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ResultRecord));
	}
}

export class ImmutableBlockRecordsResults extends wasmtypes.ScProxy {
	count(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ResultCount));
	}
}

export class MutableBlockRecordsResults extends wasmtypes.ScProxy {
	count(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ResultCount));
	}
}

export class ImmutableGetRandomResults extends wasmtypes.ScProxy {
	random(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ResultRandom));
	}
}

export class MutableGetRandomResults extends wasmtypes.ScProxy {
	random(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ResultRandom));
	}
}

export class ImmutableIotaBalanceResults extends wasmtypes.ScProxy {
	iotas(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ResultIotas));
	}
}

export class MutableIotaBalanceResults extends wasmtypes.ScProxy {
	iotas(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ResultIotas));
	}
}

export class ImmutableMapValueResults extends wasmtypes.ScProxy {
	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ResultValue));
	}
}

export class MutableMapValueResults extends wasmtypes.ScProxy {
	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ResultValue));
	}
}
