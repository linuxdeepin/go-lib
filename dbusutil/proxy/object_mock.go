// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package proxy

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/mock"
	"github.com/linuxdeepin/go-lib/dbusutil"
)

type MockObject struct {
	mock.Mock
}

func (o *MockObject) Path_() dbus.ObjectPath {
	mockArgs := o.Called()

	ret, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (o *MockObject) ServiceName_() string {
	mockArgs := o.Called()

	return mockArgs.String(0)
}

func (o *MockObject) InitSignalExt(sigLoop *dbusutil.SignalLoop, ruleAuto bool) {
	o.Called(sigLoop, ruleAuto)
}

func (o *MockObject) RemoveAllHandlers() {
	o.Called()
}

func (o *MockObject) RemoveHandler(handlerId dbusutil.SignalHandlerId) {
	o.Called(handlerId)
}

func (o *MockObject) ConnectPropertiesChanged(
	cb func(interfaceName string, changedProperties map[string]dbus.Variant,
		invalidatedProperties []string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := o.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
