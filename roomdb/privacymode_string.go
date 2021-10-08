// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

// Code generated by "stringer -type=PrivacyMode"; DO NOT EDIT.

package roomdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ModeUnknown-0]
	_ = x[ModeOpen-1]
	_ = x[ModeCommunity-2]
	_ = x[ModeRestricted-3]
}

const _PrivacyMode_name = "ModeUnknownModeOpenModeCommunityModeRestricted"

var _PrivacyMode_index = [...]uint8{0, 11, 19, 32, 46}

func (i PrivacyMode) String() string {
	if i >= PrivacyMode(len(_PrivacyMode_index)-1) {
		return "PrivacyMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PrivacyMode_name[_PrivacyMode_index[i]:_PrivacyMode_index[i+1]]
}
