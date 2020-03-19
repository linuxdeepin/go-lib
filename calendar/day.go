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
	501:  "劳动节",
	504:  "青年节",
	601:  "儿童节",
	701:  "建党节",
	801:  "建军节",
	903:  "抗日战争胜利纪念日",
	910:  "中国教师节",
	1001: "国庆节",
	1213: "南京大屠杀死难者国家公祭日",
	1220: "澳门回归纪念",
	1224: "平安夜",
	1225: "圣诞节",
	1226: "毛泽东诞辰纪念",
}

func (d *Day) Festival() string {
	key := d.Month*100 + d.Day
	if name, ok := solarFestival[key]; ok {
		return name
	}
	return ""
}
