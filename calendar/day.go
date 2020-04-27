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

package calendar


type Day struct {
	Year, Month, Day int
}

var solarFestival = map[int]string{
	101:  "元旦",
	214:  "情人节",
	305:  "学雷锋纪念日",
	308:  "妇女节",
	312:  "植树节",
	401:  "愚人节",
	415:  "全民国家安全教育日",
	501:  "劳动节",
	504:  "青年节",
	601:  "儿童节",
	701:  "建党节,香港回归纪念日",
	801:  "建军节",
	903:  "抗日战争胜利纪念日",
	910:  "教师节",
	1001: "国庆节",
	1213: "南京大屠杀死难者国家公祭日",
	1220: "澳门回归纪念",
	1224: "平安夜",
	1225: "圣诞节",
	1226: "毛泽东诞辰纪念",
}

func (d *Day) Festival() string {
	year := d.Year
	month := d.Month
	day := d.Day
	name := ""
	if (month == 5) || (month == 6){
		name = festivalForFatherAndMother(year, month, day)
		if name != ""{
			return name
		}
	}
	key := month * 100 + day
	if name, ok := solarFestival[key]; ok {
		return name
	}
	return ""
}

func (d *Day) festivalForFatherAndMother(year, month, day int) string {
	var disparityMotherDay,disparityFatherDay,fatherDay,i,motherDay int
	var leapYear int
	for i = 1900 ; i <= year ;i++{
		if( ( i % 400 == 0 ) || ( ( i % 100 != 0 ) && ( i % 4 == 0 ) ) ){
			leapYear = leapYear + 1
		}
	}
	if month == 5 {
		disparityMotherDay = ( ( ( year - 1899 ) * 365 + leapYear ) -( 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31 ) ) % 7
		motherDay = 14 - disparityMotherDay
		if day == motherDay {
			return "母亲节"
		}else{
			return ""
		}
	}
	if month == 6 {
		disparityFatherDay = ( ( ( year - 1899 ) * 365 + leapYear ) -(  30 + 31 + 31 + 30 + 31 + 30 + 31 ) ) % 7
		fatherDay =  21 - disparityFatherDay
		if day == fatherDay {
			return "父亲节"
		}else{
			return ""
		}
	}

	return ""

}
