// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package bigrequests

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func encodeEnable() (b x.RequestBody) {
	return
}

type EnableReply struct {
	MaximumRequestLength uint32
}

func readEnableReply(r *x.Reader, v *EnableReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}

	r.ReadPad(8)

	v.MaximumRequestLength = r.Read4b() // 3

	return nil
}
