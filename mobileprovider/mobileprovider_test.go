// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
