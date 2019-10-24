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
	315:  "消费者权益日",
	401:  "愚人节",
	407:  "世界卫生日",
	422:  "世界地球日",
	501:  "劳动节",
	504:  "青年节",
	512:  "国际护士节",
	601:  "国际儿童节",
	626:  "国际禁毒日",
	701:  "香港回归纪念日 中共诞辰",
	707:  "抗日战争纪念日",
	801:  "建军节",
	815:  "抗日战争胜利纪念",
	909:  "毛泽东逝世纪念",
	910:  "中国教师节",
	918:  "九一八事变纪念日",
	1001: "国庆节",
	1006: "老人节",
	1010: "辛亥革命纪念日",
	1013: "国际教师节",
	1112: "孙中山诞辰纪念日",
	1213: "南京大屠杀纪念日",
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
