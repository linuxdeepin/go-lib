// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#include <glib.h>
#include <locale.h>
#include <stdio.h>
#include <string.h>

#include "player.h"

#define DEFAULT_DRIVER "pulse"

static ca_context* connect_canberra_context(char* device, char* driver);

static uint32_t id = 0;

int
canberra_play_system_sound(char *theme, char *event_id,
        char *device, char* driver)
{
	int curid = ++id;
	setlocale(LC_ALL, "");
	ca_context *ca = connect_canberra_context(device, driver);
	if (ca == NULL) {
		return -1;
	}

	int playret = ca_context_play(ca, curid,
	                      CA_PROP_CANBERRA_XDG_THEME_NAME, theme,
	                      CA_PROP_EVENT_ID, event_id, NULL);

	if (playret != CA_SUCCESS) {
		g_warning("play: id=%d %s\n", curid, ca_strerror(playret));
		ca_context_destroy(ca);
		return playret;
	}

	// wait for end
	int playing;
	do {
		g_usleep(500 * 1000); // sleep 0.5s
		int ret = ca_context_playing(ca, curid, &playing);
		if (ret != CA_SUCCESS) {
			g_warning("ca_context_playing id=%d %s\n", curid, ca_strerror(ret));
		}
	} while (playing > 0);

	ca_context_destroy(ca);
	return playret;
}

int
canberra_play_sound_file(char *file, char *device, char* driver)
{
	int curid = ++id;
	setlocale(LC_ALL, "");
	ca_context *ca = connect_canberra_context(device, driver);
	if (ca == NULL) {
		return -1;
	}

	int playret = ca_context_play(ca, curid,
	                      CA_PROP_MEDIA_FILENAME, file, NULL);

	if (playret != CA_SUCCESS) {
		g_warning("play_sound_file filename:%s id=%d %s\n", file, curid, ca_strerror(playret));
		ca_context_destroy(ca);
		return playret;
	}

	// wait for end
	int playing;
	do {
		g_usleep(500 * 1000); // sleep 0.5s
		int ret = ca_context_playing(ca, curid, &playing);
		if (ret != CA_SUCCESS) {
			g_warning("ca_context_playing id=%d %s\n", curid, ca_strerror(ret));
		}
	} while (playing > 0);

	ca_context_destroy(ca);
	return playret;
}

static ca_context*
connect_canberra_context(char* device, char* driver)
{
	ca_context* ca = NULL;
	if (ca_context_create(&ca) != 0) {
		g_warning("Create canberra context failed");
		return NULL;
	}

	// set backend driver
	if (strlen(driver) > 0) {
		if (ca_context_set_driver(ca, driver) != 0 ) {
			g_warning("Set '%s' as backend driver failed", driver);
			ca_context_destroy(ca);
			return NULL;
		}
	}

	if (strlen(device) > 0) {
		if (ca_context_change_device(ca, device) != 0) {
			g_warning("Set '%s' as backend device failed", device);
			ca_context_destroy(ca);
			return NULL;
		}
	}

	if (ca_context_open(ca) != 0) {
		g_warning("Connect the context to sound system failed");
		ca_context_destroy(ca);
		return NULL;
	}

	return ca;
}
