// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

use crate::*;

#[derive(Clone)]
pub struct ArrayOfImmutableDonation {
	pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableDonation {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }


	pub fn get_donation(&self, index: u32) -> ImmutableDonation {
		ImmutableDonation { proxy: self.proxy.index(index) }
	}
}

#[derive(Clone)]
pub struct ImmutableDonateWithFeedbackState {
	pub(crate) proxy: Proxy,
}

impl ImmutableDonateWithFeedbackState {
    pub fn log(&self) -> ArrayOfImmutableDonation {
		ArrayOfImmutableDonation { proxy: self.proxy.root(STATE_LOG) }
	}

    pub fn max_donation(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(STATE_MAX_DONATION))
	}

    pub fn total_donation(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(STATE_TOTAL_DONATION))
	}
}

#[derive(Clone)]
pub struct ArrayOfMutableDonation {
	pub(crate) proxy: Proxy,
}

impl ArrayOfMutableDonation {

	pub fn append_donation(&self) -> MutableDonation {
		MutableDonation { proxy: self.proxy.append() }
	}
	pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }


	pub fn get_donation(&self, index: u32) -> MutableDonation {
		MutableDonation { proxy: self.proxy.index(index) }
	}
}

#[derive(Clone)]
pub struct MutableDonateWithFeedbackState {
	pub(crate) proxy: Proxy,
}

impl MutableDonateWithFeedbackState {
    pub fn as_immutable(&self) -> ImmutableDonateWithFeedbackState {
		ImmutableDonateWithFeedbackState { proxy: self.proxy.root("") }
	}

    pub fn log(&self) -> ArrayOfMutableDonation {
		ArrayOfMutableDonation { proxy: self.proxy.root(STATE_LOG) }
	}

    pub fn max_donation(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(STATE_MAX_DONATION))
	}

    pub fn total_donation(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(STATE_TOTAL_DONATION))
	}
}