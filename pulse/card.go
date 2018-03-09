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

package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "fmt"
import "unsafe"

const (
	DirectionSink int = iota + 1
	DirectionSource
)

type CardPortInfo struct {
	PortInfo
	Direction int
	Profiles  ProfileInfos2
}
type CardPortInfos []CardPortInfo

//typedef struct pa_card_info {
//	uint32_t index;                      /**< Index of this card */
//	const char *name;                    /**< Name of this card */
//	uint32_t owner_module;               /**< Index of the owning module, or PA_INVALID_INDEX. */
//	const char *driver;                  /**< Driver name */
//	uint32_t n_profiles;                 /**< Number of entries in profile array */
//	pa_card_profile_info* profiles;      /**< \deprecated Superseded by profiles2 */
//	pa_card_profile_info* active_profile; /**< \deprecated Superseded by active_profile2 */
//	pa_proplist *proplist;               /**< Property list */
//	uint32_t n_ports;                    /**< Number of entries in port array */
//	pa_card_port_info **ports;           /**< Array of pointers to ports, or NULL. Array is terminated by an entry set to NULL. */
//	pa_card_profile_info2** profiles2;    /**< Array of pointers to available profiles, or NULL. Array is terminated by an entry set to NULL. \since 5.0 */
//	pa_card_profile_info2* active_profile2; /**< Pointer to active profile in the array, or NULL. \since 5.0 */
//} pa_card_info;
type Card struct {
	Index       uint32
	Name        string
	OwnerModule uint32
	Driver      string

	PropList      map[string]string
	Profiles      ProfileInfos2
	ActiveProfile ProfileInfo2
	Ports         CardPortInfos
}

func toCardInfo(info *C.pa_card_info) (c *Card) {
	c = &Card{}

	c.Index = uint32(info.index)
	c.Name = C.GoString(info.name)
	c.OwnerModule = uint32(info.owner_module)
	c.Driver = C.GoString(info.driver)
	c.PropList = toProplist(info.proplist)
	c.Profiles = toProfiles(uint32(info.n_profiles), info.profiles2)
	c.ActiveProfile = toProfile(info.active_profile2)
	c.Ports = toCardPorts(uint32(info.n_ports), info.ports)
	return
}

func (card *Card) SetProfile(name string) {
	cname := C.CString(fmt.Sprint(card.Index))
	defer C.free(unsafe.Pointer(cname))
	pname := C.CString(name)
	defer C.free(unsafe.Pointer(pname))

	c := GetContext()
	c.safeDo(func() {
		C.pa_context_set_card_profile_by_name(c.ctx, cname, pname, C.get_success_cb(), nil)
	})
}

func (infos CardPortInfos) TrySelectProfile(portName string) (string, error) {
	for _, info := range infos {
		if info.Name != portName {
			continue
		}
		return info.Profiles.SelectProfile(), nil
	}
	return "", fmt.Errorf("Invalid card port name: %s", portName)
}

func (infos CardPortInfos) Get(name string, direction int) (CardPortInfo, error) {
	for _, info := range infos {
		if info.Name == name && info.Direction == direction {
			return info, nil
		}
	}
	return CardPortInfo{}, fmt.Errorf("Invalid card port name: %s", name)
}
