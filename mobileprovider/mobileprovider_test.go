/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package mobileprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMobileProvider(t *testing.T) {
	database, err := GetMobileProviderDatabase()
	assert.NotNil(t, database)
	assert.Equal(t, err, nil)
}

func TestGetProviders(t *testing.T) {
	providers, err := GetProviders("cn")
	require.NoError(t, err)
	assert.Equal(t, len(providers), 3)

	providers, err = GetProviders("<invalid>")
	assert.Error(t, err)
	assert.Equal(t, providers == nil, true)
}

func TestGetProvider(t *testing.T) {
	provider, err := GetProvider("cn", "China Mobile")
	require.NoError(t, err)
	assert.Equal(t, len(provider.GSM.APN), 3)

	provider, err = GetProvider("cn", "<invalid>")
	assert.Error(t, err)
	assert.Equal(t, provider == nil, true)
}

func TestGetGSM(t *testing.T) {
	gsm, err := GetGSM("cn", "China Mobile")
	require.NoError(t, err)
	assert.Equal(t, len(gsm.APN), 3)

	gsm, err = GetGSM("cn", "<invalid>")
	assert.Equal(t, gsm == nil, true)
	assert.Error(t, err)
}

func TestGetCDMA(t *testing.T) {
	cdma, err := GetCDMA("cn", "China Telecom")
	require.NoError(t, err)
	assert.Equal(t, cdma.Username, "ctnet@mycdma.cn")

	cdma, err = GetCDMA("cn", "<invalid>")
	assert.Error(t, err)
	assert.Equal(t, cdma == nil, true)

	cdma, err = GetCDMA("cn", "China Mobile")
	assert.Error(t, err)
	assert.Equal(t, cdma == nil, true)
}

func TestGetAPN(t *testing.T) {
	apn, err := GetAPN("cn", "China Mobile", "cmnet", "internet")
	require.NoError(t, err)
	assert.Equal(t, GetAPNName(apn), "Internet")

	apn, err = GetAPN("cn", "China Mobile", "cmnet", "<invalid>")
	assert.Equal(t, apn == nil, true)
	assert.Error(t, err)

	apn, err = GetAPN("au", "Amaysim", "internet", "")
	require.NoError(t, err)
	assert.Equal(t, GetAPNName(apn), "")
	assert.Equal(t, GetAPNUsageType(apn), "")
}

func TestGetGSMForNetworkID(t *testing.T) {
	gsm, err := GetGSMForNetworkID("460", "00")
	require.NoError(t, err)
	assert.Equal(t, len(gsm.APN), 3)

	gsm, err = GetGSMForNetworkID("460", "<invalid>")
	assert.Equal(t, gsm == nil, true)
	assert.Error(t, err)
}

func TestGetProviderNames(t *testing.T) {
	names, err := GetProviderNames("cn")
	require.NoError(t, err)
	assert.Equal(t, names[0], "China Mobile")
	assert.Equal(t, names[1], "China Unicom")
	assert.Equal(t, names[2], "China Telecom")

	names, err = GetProviderNames("<invalid>")
	assert.Error(t, err)
	assert.Equal(t, len(names) == 0, true)
}

func TestGetPlans(t *testing.T) {
	plans, err := GetPlans("cn", "China Mobile")
	require.NoError(t, err)
	assert.Equal(t, len(plans), 3)
	assert.Equal(t, err, nil)

	plans, err = GetPlans("ca", "Bell Mobility")
	require.NoError(t, err)
	assert.Equal(t, len(plans), 5)

	plans, err = GetPlans("cn", "<invalid>")
	assert.Equal(t, len(plans) == 0, true)
	assert.Error(t, err)
}

func TestGetDefaultPlan(t *testing.T) {
	plan, err := GetDefaultPlan("cn", "China Mobile")
	require.NoError(t, err)
	assert.Equal(t, plan.Name, "WAP")
	assert.Equal(t, plan.ProviderName, "China Mobile")

	plan, err = GetDefaultPlan("cn", "<invalid>")
	assert.Equal(t, plan.Name, "")
	assert.Error(t, err)
}

func TestGetDefaultGSMPlanForCountry(t *testing.T) {
	plan, err := GetDefaultGSMPlanForCountry("cn")
	require.NoError(t, err)
	assert.Equal(t, plan.Name, "Internet")
	assert.Equal(t, plan.ProviderName, "China Mobile")

	plan, err = GetDefaultGSMPlanForCountry("<invalid>")
	assert.Equal(t, plan.Name, "")
	assert.Error(t, err)
}

func TestGetDefaultCDMAPlanForCountry(t *testing.T) {
	plan, err := GetDefaultCDMAPlanForCountry("cn")
	require.NoError(t, err)
	assert.Equal(t, plan.Name, "China Telecom")

	plan, err = GetDefaultCDMAPlanForCountry("<invalid>")
	assert.Equal(t, plan.Name, "")
	assert.Error(t, err)
}
