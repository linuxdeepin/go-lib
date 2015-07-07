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

package iso

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	. "pkg.deepin.io/lib/gettext"
	"strings"
	"sync"
)

const iso3166XMLFile = "/usr/share/xml/iso-codes/iso_3166.xml"
const isoCodesLocalesDir = "/usr/share/locale"

// map dtd of iso_3166.xml to go structures
type CountryDatabase struct {
	Countries []Country `xml:"iso_3166_entry"`
}
type Country struct {
	Alpha2Code   string `xml:"alpha_2_code,attr"`
	Alpha3Code   string `xml:"alpha_3_code,attr"`
	NumericCode  string `xml:"numeric_code,attr"`
	CommonName   string `xml:"common_name,attr"`
	Name         string `xml:"name,attr"`
	OfficialName string `xml:"official_name,attr"`
}

var countryDatabase *CountryDatabase
var countryDatabaseLock sync.Mutex

var (
	errLanguageFormatInvalid = fmt.Errorf("invalid environment variable LANGUAGE")
	errCountryCodeInvalid    = fmt.Errorf("invalid country code")
)

// GetCountryDatabase return country database that marshaled from ISO
// 3166 xml file.
func GetCountryDatabase() (*CountryDatabase, error) {
	countryDatabaseLock.Lock()
	defer countryDatabaseLock.Unlock()

	if countryDatabase != nil {
		return countryDatabase, nil
	}

	countryDatabase = &CountryDatabase{}
	xmlContent, err := ioutil.ReadFile(iso3166XMLFile)
	if err != nil {
		return countryDatabase, err
	}
	err = xml.Unmarshal(xmlContent, countryDatabase)
	return countryDatabase, err
}

// GetLocaleCountryCode return locale country code by analysis
// environment variable "LANGUAGE".
func GetLocaleCountryCode() (code string, err error) {
	return GetCountryCodeForLanguage(getLocalLanguage())
}
func getLocalLanguage() string {
	if value := os.Getenv("LANGUAGE"); len(value) > 0 {
		return value
	}
	if value := os.Getenv("LC_ALL"); len(value) > 0 {
		return value
	}
	if value := os.Getenv("LC_MESSAGES"); len(value) > 0 {
		return value
	}
	if value := os.Getenv("LANG"); len(value) > 0 {
		return value
	}
	return "en_US.UTF-8"
}

// GetLocaleCountryName return locale country name by analysis
// environment variable "LANGUAGE".
func GetLocaleCountryName() (name string, err error) {
	code, err := GetLocaleCountryCode()
	if err != nil {
		return
	}
	return GetCountryNameForCode(code)
}

// GetCountryCodeForLanguage return country code for a language
// variable, e.g. "CN" will be return if passing "zh_CN.UTF-8".
func GetCountryCodeForLanguage(language string) (code string, err error) {
	if !strings.Contains(language, "_") {
		err = errLanguageFormatInvalid
		return
	}

	var indexFrom, indexTo int
	indexFrom = strings.Index(language, "_")
	if strings.Contains(language, ".") {
		indexTo = strings.Index(language, ".")
	} else {
		indexTo = len(language)
	}

	if indexFrom+1 >= indexTo {
		err = errLanguageFormatInvalid
		return
	}
	code = language[indexFrom+1 : indexTo]
	return
}

// GetCountryNameForCode return country name that corresponding to the
// country code.
func GetCountryNameForCode(code string) (name string, err error) {
	database, err := GetCountryDatabase()
	if err != nil {
		return
	}
	for _, entry := range database.Countries {
		if strings.ToUpper(code) == strings.ToUpper(entry.Alpha2Code) {
			name = DGettext("iso_3166", entry.Name)
			break
		}
	}
	if len(name) == 0 {
		err = errCountryCodeInvalid
	}
	return
}

// GetAllCountryCode return all country code.
func GetAllCountryCode() (codeList []string, err error) {
	database, err := GetCountryDatabase()
	if err != nil {
		return
	}
	for _, entry := range database.Countries {
		codeList = append(codeList, entry.Alpha2Code)
	}
	return
}

// GetAllCountryNames return all country names.
func GetAllCountryNames() (nameList []string, err error) {
	database, err := GetCountryDatabase()
	if err != nil {
		return
	}
	for _, entry := range database.Countries {
		nameList = append(nameList, DGettext("iso_3166", entry.Name))
	}
	return
}
