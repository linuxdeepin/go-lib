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

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

const mobileServiceProvidersXMLFile = "/usr/share/mobile-broadband-provider-info/serviceproviders.xml"

var (
	errCountryNotFound  = fmt.Errorf("country not found")
	errProviderNotFound = fmt.Errorf("provider not found")
	errGSMNotFound      = fmt.Errorf("gsm not found")
	errCDMANotFound     = fmt.Errorf("cdma not found")
	errAPNNotFound      = fmt.Errorf("apn not found")
	errPlanNotFound     = fmt.Errorf("plan not found")
)

var mobileProviderDatabase *MobileServiceProviderDatabase
var mobileProviderDatabaseLock sync.Mutex

// GetMobileProviderDatabase return mobile service provider's database
// that marshaled from serviceproviders.xml.
func GetMobileProviderDatabase() (*MobileServiceProviderDatabase, error) {
	mobileProviderDatabaseLock.Lock()
	defer mobileProviderDatabaseLock.Unlock()

	if mobileProviderDatabase != nil {
		return mobileProviderDatabase, nil
	}

	mobileProviderDatabase = &MobileServiceProviderDatabase{}
	xmlContent, err := ioutil.ReadFile(mobileServiceProvidersXMLFile)
	if err != nil {
		return mobileProviderDatabase, err
	}
	err = xml.Unmarshal(xmlContent, mobileProviderDatabase)
	return mobileProviderDatabase, err
}

// GetAllCountryCode return all country code that provide mobile
// service.
func GetAllCountryCode() (codeList []string, err error) {
	database, err := GetMobileProviderDatabase()
	if err != nil {
		return
	}
	for _, c := range database.Countries {
		codeList = append(codeList, strings.ToUpper(c.Code))
	}
	return
}

// GetProviders return all providers in target country.
func GetProviders(countryCode string) (providers []*Provider, err error) {
	database, err := GetMobileProviderDatabase()
	if err != nil {
		return
	}
	found := false
	for _, c := range database.Countries {
		if strings.ToUpper(c.Code) == strings.ToUpper(countryCode) {
			found = true
			providers = c.Providers
			break
		}
	}
	if !found {
		err = errCountryNotFound
	}
	return
}

// GetProvider return the provider information that matched the provided country
// code and provider name.
func GetProvider(countryCode, providerName string) (provider *Provider, err error) {
	providers, err := GetProviders(countryCode)
	if err != nil {
		return
	}
	found := false
	for _, p := range providers {
		if p.Name.Body == providerName {
			found = true
			provider = p
			break
		}
	}
	if !found {
		err = errProviderNotFound
	}
	return
}

// GetGSM return the gsm information that matched the provided country
// code and provider name.
func GetGSM(countryCode, providerName string) (gsm *GSM, err error) {
	provider, err := GetProvider(countryCode, providerName)
	if err != nil {
		return
	}
	gsm = provider.GSM
	if gsm == nil {
		err = errGSMNotFound
	}
	return
}

// GetCDMA return the cdma infromation that match the provided
// country code and provider name.
func GetCDMA(countryCode, providerName string) (cdma *CDMA, err error) {
	provider, err := GetProvider(countryCode, providerName)
	if err != nil {
		return
	}
	cdma = provider.CDMA
	if cdma == nil {
		err = errCDMANotFound
	}
	return
}

// GetGSMForNetworkID return the gsm information that match the mcc
// and mnc.
func GetGSMForNetworkID(mcc, mnc string) (gsm *GSM, err error) {
	database, err := GetMobileProviderDatabase()
	if err != nil {
		return
	}

	found := false
outside:
	for _, c := range database.Countries {
		for _, p := range c.Providers {
			if p.GSM != nil {
				for _, id := range p.GSM.NetworkID {
					if id.MCC == mcc && id.MNC == mnc {
						found = true
						gsm = p.GSM
						break outside
					}
				}
			}
		}
	}
	if !found {
		err = errGSMNotFound
	}
	return
}

// GetAPN return the apn information that match the provided
// country code, provider name and apn value.
func GetAPN(countryCode, providerName, apnValue, apnUsageType string) (apn *APN, err error) {
	gsm, err := GetGSM(countryCode, providerName)
	if err != nil {
		return
	}
	found := false
	for _, a := range gsm.APN {
		if a.Value == apnValue && GetAPNUsageType(a) == apnUsageType {
			found = true
			apn = a
		}
	}
	if !found {
		err = errAPNNotFound
	}
	return
}

// GetAPNName return apn's default name, if not exist, return empty
// string.
func GetAPNName(apn *APN) (name string) {
	if len(apn.Name) == 0 {
		return
	}
	name = apn.Name[0].Body
	return
}

// GetAPNUsageType return apn's usage type, if not exist, return empty
// string.
func GetAPNUsageType(apn *APN) (usageType string) {
	if apn.Usage == nil {
		return
	}
	usageType = apn.Usage.Type
	return
}

// GetProviderNames return all provider names in target country.
func GetProviderNames(countryCode string) (names []string, err error) {
	providers, err := GetProviders(countryCode)
	if err != nil {
		return
	}
	for _, p := range providers {
		names = append(names, p.Name.Body)
	}
	return
}

// GetDefaultProvider return the default provider in target country,
// usually is the first provider.
func GetDefaultProvider(countryCode string) (defaultProvider string, err error) {
	providers, err := GetProviderNames(countryCode)
	if err != nil {
		return
	}
	if len(providers) > 0 {
		defaultProvider = providers[0]
	} else {
		err = errProviderNotFound
	}
	return
}
