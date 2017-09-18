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

package calendar

type Day struct {
	Year, Month, Day int
}

var solarFestival = map[int]string{
	101:  "元旦",
	202:  "世界湿地日",
	210:  "国际气象节",
	214:  "情人节",
	301:  "国际海豹日",
	303:  "全国爱耳日",
	305:  "学雷锋纪念日",
	308:  "妇女节",
	312:  "植树节 孙中山逝世纪念日",
	314:  "国际警察日",
	315:  "消费者权益日",
	317:  "中国国医节 国际航海日",
	321:  "世界森林日 消除种族歧视国际日 世界儿歌日",
	322:  "世界水日",
	323:  "世界气象日",
	324:  "世界防治结核病日",
	325:  "全国中小学生安全教育日",
	330:  "巴勒斯坦国土日",
	401:  "愚人节",
	407:  "世界卫生日",
	422:  "世界地球日",
	423:  "世界图书和版权日",
	424:  "亚非新闻工作者日",
	501:  "劳动节",
	504:  "青年节",
	505:  "碘缺乏病防治日",
	508:  "世界红十字日",
	512:  "国际护士节",
	515:  "国际家庭日",
	517:  "世界电信日",
	518:  "国际博物馆日",
	520:  "全国学生营养日",
	522:  "国际生物多样性日",
	523:  "国际牛奶日",
	531:  "世界无烟日",
	601:  "国际儿童节",
	605:  "世界环境日",
	606:  "全国爱眼日",
	617:  "防治荒漠化和干旱日",
	623:  "国际奥林匹克日",
	625:  "全国土地日",
	626:  "国际禁毒日",
	701:  "香港回归纪念日 中共诞辰 世界建筑日",
	702:  "国际体育记者日",
	707:  "抗日战争纪念日",
	711:  "世界人口日",
	730:  "非洲妇女日",
	801:  "建军节",
	808:  "中国男子节(爸爸节)",
	815:  "抗日战争胜利纪念",
	908:  "国际扫盲日 国际新闻工作者日",
	909:  "毛泽东逝世纪念",
	910:  "中国教师节",
	914:  "世界清洁地球日",
	916:  "国际臭氧层保护日",
	918:  "九一八事变纪念日",
	920:  "国际爱牙日",
	927:  "世界旅游日",
	928:  "孔子诞辰",
	1001: "国庆节 世界音乐日 国际老人节",
	1002: "国际和平与民主自由斗争日",
	1004: "世界动物日",
	1006: "老人节",
	1008: "全国高血压日 世界视觉日",
	1009: "世界邮政日 万国邮联日",
	1010: "辛亥革命纪念日 世界精神卫生日",
	1013: "世界保健日 国际教师节",
	1014: "世界标准日",
	1015: "国际盲人节(白手杖节)",
	1016: "世界粮食日",
	1017: "世界消除贫困日",
	1022: "世界传统医药日",
	1024: "联合国日 世界发展信息日",
	1031: "世界勤俭日",
	1107: "十月社会主义革命纪念日",
	1108: "中国记者日",
	1109: "全国消防安全宣传教育日",
	1110: "世界青年节",
	1111: "国际科学与和平周(本日所属的一周)",
	1112: "孙中山诞辰纪念日",
	1114: "世界糖尿病日",
	1117: "国际大学生节 世界学生节",
	1121: "世界问候日 世界电视日",
	1129: "国际声援巴勒斯坦人民国际日",
	1201: "世界艾滋病日",
	1203: "世界残疾人日",
	1205: "国际经济和社会发展志愿人员日",
	1208: "国际儿童电视日",
	1209: "世界足球日",
	1210: "世界人权日",
	1212: "西安事变纪念日",
	1213: "南京大屠杀纪念日",
	1220: "澳门回归纪念",
	1221: "国际篮球日",
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
