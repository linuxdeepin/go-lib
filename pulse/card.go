/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "fmt"
import "unsafe"

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

	// TODO
	//Ports         []CardPortInfo
	PropList      map[string]string
	Profiles      []ProfileInfo2
	ActiveProfile ProfileInfo2
}

func toCardInfo(info *C.pa_card_info) (c *Card) {
	c = &Card{}

	c.Index = uint32(info.index)
	c.Name = C.GoString(info.name)
	c.OwnerModule = uint32(info.owner_module)
	c.Driver = C.GoString(info.driver)
	c.PropList = toProplist(info.proplist)
	// TODO
	//	c.Ports = toPorts(uint32(info.n_ports), info.ports)
	c.Profiles = toProfiles(uint32(info.n_profiles), info.profiles2)
	c.ActiveProfile = toProfile(info.active_profile2)
	return
}

func (card *Card) SetProfile(name string) {
	cname := C.CString(fmt.Sprint(card.Index))
	defer C.free(unsafe.Pointer(cname))
	pname := C.CString(name)
	defer C.free(unsafe.Pointer(pname))

	c := GetContext()
	c.SafeDo(func() {
		C.pa_context_set_card_profile_by_name(c.ctx, cname, pname, C.get_success_cb(), nil)
	})
}
