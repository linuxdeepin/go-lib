// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#ifndef __PLAYER_H__
#define __PLAYER_H__

#include <canberra.h>

int canberra_play_system_sound(char *theme, char *event_id,
        char *device, char* driver);
int canberra_play_sound_file(char *file, char *device, char* driver);

#endif
