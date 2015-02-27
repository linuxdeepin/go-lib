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
	"encoding/json"
)

// Plan defines the unique information for each plan(gsm or cdma) in
// provider.
type Plan struct {
	IsGSM        bool
	Name         string // apn names for gsm or provider name for cdma
	ProviderName string
	APNValue     string
	APNUsageType string
}

func MarshalPlan(plan Plan) (jsonStr string) {
	bytes, _ := json.Marshal(plan)
	jsonStr = string(bytes)
	return
}

func UnmarshalPlan(jsonStr string) (plan *Plan, err error) {
	plan = &Plan{}
	err = json.Unmarshal([]byte(jsonStr), plan)
	return
}

// GetPlans return all plans(apn names for gsm and provider name for
// cdma) under target provider.
func GetPlans(countryCode, providerName string) (plans []Plan, err error) {
	provider, err := GetProvider(countryCode, providerName)
	if err != nil {
		return
	}
	if provider.GSM != nil {
		for _, a := range provider.GSM.APN {
			plans = append(plans, Plan{
				IsGSM:        true,
				Name:         GetAPNName(a),
				ProviderName: provider.Name.Body,
				APNValue:     a.Value,
				APNUsageType: GetAPNUsageType(a),
			})
		}
	}
	if provider.CDMA != nil {
		plans = append(plans, Plan{
			IsGSM:        false,
			Name:         provider.Name.Body,
			ProviderName: provider.Name.Body,
		})
	}
	return
}

// GetDefaultPlan return default plan for taget provider, usually is
// the first plan.
func GetDefaultPlan(countryCode, providerName string) (defaultPlan Plan, err error) {
	plans, err := GetPlans(countryCode, providerName)
	if err != nil {
		return
	}
	if len(plans) > 0 {
		defaultPlan = plans[0]
	} else {
		err = errPlanNotFound
	}
	return
}

// GetDefaultGSMPlanForCountry return default gsm plan in target country.
func GetDefaultGSMPlanForCountry(countryCode string) (defaultGSMPlan Plan, err error) {
	providerNames, err := GetProviderNames(countryCode)
	if err != nil {
		return
	}

	// plans which apn usage type is "internet" owns higher priority
	defaultGSMInternetPlan, err := getDefaultGSMInternetPlanForCountry(countryCode)
	if err == nil {
		defaultGSMPlan = defaultGSMInternetPlan
		return
	}

	found := false
outside:
	for _, providerName := range providerNames {
		plans, err := GetPlans(countryCode, providerName)
		if err != nil {
			continue
		}
		for _, p := range plans {
			if p.IsGSM {
				found = true
				defaultGSMPlan = p
				break outside
			}
		}
	}

	if !found {
		err = errPlanNotFound
	}
	return
}
func getDefaultGSMInternetPlanForCountry(countryCode string) (defaultGSMInternetPlan Plan, err error) {
	providerNames, err := GetProviderNames(countryCode)
	if err != nil {
		return
	}

	found := false
outside:
	for _, providerName := range providerNames {
		plans, err := GetPlans(countryCode, providerName)
		if err != nil {
			continue
		}
		for _, p := range plans {
			if p.IsGSM && p.APNUsageType == "internet" {
				found = true
				defaultGSMInternetPlan = p
				break outside
			}
		}
	}

	if !found {
		err = errPlanNotFound
	}
	return
}

// GetDefaultCDMAPlanForCountry return default gsm plan in target country.
func GetDefaultCDMAPlanForCountry(countryCode string) (defaultCDMAPlan Plan, err error) {
	providerNames, err := GetProviderNames(countryCode)
	if err != nil {
		return
	}

	found := false
	for _, providerName := range providerNames {
		plans, err := GetPlans(countryCode, providerName)
		if err != nil {
			continue
		}
		for _, p := range plans {
			if !p.IsGSM {
				found = true
				defaultCDMAPlan = p
				break
			}
		}
	}

	if !found {
		err = errPlanNotFound
	}
	return
}
