// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

/*
#include "dde-pulse.h"
*/
import "C"

type SourceOutput struct{}

func toSourceOutputInfo(info *C.pa_source_output_info) *SourceOutput {
	return nil
}
