/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

#ifndef __PLAYER_H__
#define __PLAYER_H__

#include <canberra.h>

int canberra_play_system_sound(char *theme, char *event_id,
        char *device, char* driver);
int canberra_play_sound_file(char *file, char *device, char* driver);

#endif
