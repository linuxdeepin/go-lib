/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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
	C "gopkg.in/check.v1"
	"testing"
)

type testWrapper struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&testWrapper{})
}

func (*testWrapper) TestMobileProvider(c *C.C) {
	database, err := GetMobileProviderDatabase()
	c.Check(database, C.NotNil)
	c.Check(err, C.Equals, nil)
}

func (*testWrapper) TestGetProviders(c *C.C) {
	providers, err := GetProviders("cn")
	c.Check(len(providers), C.Equals, 3)
	c.Check(err, C.Equals, nil)

	providers, err = GetProviders("<invalid>")
	c.Check(providers == nil, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetProvider(c *C.C) {
	provider, err := GetProvider("cn", "China Mobile")
	c.Check(len(provider.GSM.APN), C.Equals, 3)
	c.Check(err, C.Equals, nil)

	provider, err = GetProvider("cn", "<invalid>")
	c.Check(provider == nil, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetGSM(c *C.C) {
	gsm, err := GetGSM("cn", "China Mobile")
	c.Check(len(gsm.APN), C.Equals, 3)
	c.Check(err, C.Equals, nil)

	gsm, err = GetGSM("cn", "<invalid>")
	c.Check(gsm == nil, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetCDMA(c *C.C) {
	cdma, err := GetCDMA("cn", "China Telecom")
	c.Check(cdma.Username, C.Equals, "ctnet@mycdma.cn")
	c.Check(err, C.Equals, nil)

	cdma, err = GetCDMA("cn", "<invalid>")
	c.Check(cdma == nil, C.Equals, true)
	c.Check(err, C.NotNil)

	cdma, err = GetCDMA("cn", "China Mobile")
	c.Check(cdma == nil, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetAPN(c *C.C) {
	apn, err := GetAPN("cn", "China Mobile", "cmnet", "internet")
	c.Check(GetAPNName(apn), C.Equals, "Internet")
	c.Check(err, C.Equals, nil)

	apn, err = GetAPN("cn", "China Mobile", "cmnet", "<invalid>")
	c.Check(apn == nil, C.Equals, true)
	c.Check(err, C.NotNil)

	apn, err = GetAPN("au", "Amaysim", "internet", "")
	c.Check(GetAPNName(apn), C.Equals, "")
	c.Check(GetAPNUsageType(apn), C.Equals, "")
	c.Check(err, C.Equals, nil)
}

func (*testWrapper) TestGetGSMForNetworkID(c *C.C) {
	gsm, err := GetGSMForNetworkID("460", "00")
	c.Check(len(gsm.APN), C.Equals, 3)
	c.Check(err, C.Equals, nil)

	gsm, err = GetGSMForNetworkID("460", "<invalid>")
	c.Check(gsm == nil, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetProviderNames(c *C.C) {
	names, err := GetProviderNames("cn")
	c.Check(names[0], C.Equals, "China Mobile")
	c.Check(names[1], C.Equals, "China Unicom")
	c.Check(names[2], C.Equals, "China Telecom")
	c.Check(err, C.Equals, nil)

	names, err = GetProviderNames("<invalid>")
	c.Check(len(names) == 0, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetPlans(c *C.C) {
	plans, err := GetPlans("cn", "China Mobile")
	c.Check(len(plans), C.Equals, 3)
	c.Check(err, C.Equals, nil)

	plans, err = GetPlans("ca", "Bell Mobility")
	c.Check(len(plans), C.Equals, 5)
	c.Check(err, C.Equals, nil)

	plans, err = GetPlans("cn", "<invalid>")
	c.Check(len(plans) == 0, C.Equals, true)
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetDefaultPlan(c *C.C) {
	plan, err := GetDefaultPlan("cn", "China Mobile")
	c.Check(plan.Name, C.Equals, "WAP")
	c.Check(plan.ProviderName, C.Equals, "China Mobile")
	c.Check(err, C.Equals, nil)

	plan, err = GetDefaultPlan("cn", "<invalid>")
	c.Check(plan.Name, C.Equals, "")
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetDefaultGSMPlanForCountry(c *C.C) {
	plan, err := GetDefaultGSMPlanForCountry("cn")
	c.Check(plan.Name, C.Equals, "Internet")
	c.Check(plan.ProviderName, C.Equals, "China Mobile")
	c.Check(err, C.Equals, nil)

	plan, err = GetDefaultGSMPlanForCountry("<invalid>")
	c.Check(plan.Name, C.Equals, "")
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetDefaultCDMAPlanForCountry(c *C.C) {
	plan, err := GetDefaultCDMAPlanForCountry("cn")
	c.Check(plan.Name, C.Equals, "China Telecom")
	c.Check(err, C.Equals, nil)

	plan, err = GetDefaultCDMAPlanForCountry("<invalid>")
	c.Check(plan.Name, C.Equals, "")
	c.Check(err, C.NotNil)
}
