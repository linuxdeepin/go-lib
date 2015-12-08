/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package mobileprovider

// map mobile-broadband-provider-info/serviceproviders.2.dtd to go
// structures

type MobileServiceProviderDatabase struct {
	Format    string     `xml:"format,attr"`
	Countries []*Country `xml:"country"`
}

type Country struct {
	Code      string      `xml:"code,attr"`
	Providers []*Provider `xml:"provider"`
}

type Provider struct {
	Primary ProviderPrimaryValue `xml:"primary,attr"`
	Name    *Name                `xml:"name"`
	GSM     *GSM                 `xml:"gsm"`
	CDMA    *CDMA                `xml:"cdma"`
}

type ProviderPrimaryValue string

const (
	ProviderPrimaryValueTrue  ProviderPrimaryValue = "true"
	ProviderPrimaryValueFalse ProviderPrimaryValue = "false"
)

type GSM struct {
	NetworkID    []*NetworkID    `xml:"network-id"`
	MsisdnQuery  []*MsisdnQuery  `xml:"msisdn-query"`
	Voicemail    []string        `xml:"voicemail"`
	BalanceCheck []*BalanceCheck `xml:"balance-check"`
	BalanceTopUp []*BalanceTopUp `xml:"balance-top-up"`
	APN          []*APN          `xml:"apn"`
}

type APN struct {
	Value    string     `xml:"value,attr"`
	Plan     []*APNPlan `xml:"plan"`
	Usage    *Usage     `xml:"usage"`
	Name     []*Name    `xml:"name"`
	Gateway  string     `xml:"gateway"`
	Username string     `xml:"username"`
	Password string     `xml:"password"`
	DNS      []string   `xml:"dns"`
	MMSC     string     `xml:"mmsc"`
	MMSproxy string     `xml:"mmsproxy"`
}

type MsisdnQuery struct {
	USSD []*USSD `xml:"ussd"`
	SMS  []*SMS  `xml:"sms"`
}

type BalanceCheck struct {
	USSD         []*USSD  `xml:"ussd"`
	DTMF         []string `xml:"dtmf"`
	SMS          []*SMS   `xml:"sms"`
	USSDResponse []string `xml:"ussd-response"`
}

type USSD struct {
	Replacement string `xml:"replacement,attr"`
	Body        string `xml:",chardata"`
}

type SMS struct {
	Text string `xml:"text,attr"`
	Body string `xml:",chardata"`
}

type BalanceTopUp struct {
	USSD []*USSD `xml:"ussd"`
	SMS  []*SMS  `xml:"sms"`
}

type NetworkID struct {
	MCC string `xml:"mcc,attr"`
	MNC string `xml:"mnc,attr"`
}

type APNPlan struct {
	Type string `xml:"type,attr"`
}

const (
	PlanTypeValuePrepaid  = "prepaid"
	PlanTypeValuePostpaid = "postpaid"
)

type Usage struct {
	Type string `xml:"type,attr"`
}

const (
	UsageTypeInternet = "internet"
	UsageTypeMMS      = "mms"
	UsageTypeWAP      = "wap"
)

type CDMA struct {
	Name     []*Name  `xml:"name"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	DNS      []string `xml:"dns"`
	SID      []*SID   `xml:"sid"`
}

type SID struct {
	Value string `xml:"value,attr"`
}

type Name struct {
	XMLLang string `xml:"xml:lang,attr"`
	Body    string `xml:",chardata"`
}
