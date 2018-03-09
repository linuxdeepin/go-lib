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

#ifndef DDE_PULSE_H
#define DDE_PULSE_H

#include <pulse/pulseaudio.h>

#define DECLARE(TYPE) \
  void get_##TYPE##_info(pa_threaded_mainloop*, pa_context*, int64_t, uint32_t); \
  void get_##TYPE##_info_list(pa_threaded_mainloop*, pa_context*, int64_t);

DECLARE(sink);
DECLARE(sink_input);
DECLARE(source);
DECLARE(source_output);
DECLARE(client);
DECLARE(card);
DECLARE(module);
DECLARE(sample);

void get_server_info(pa_threaded_mainloop*, pa_context *c, int64_t cookie);

pa_context* new_pa_context(pa_threaded_mainloop* ml);
void pa_finalize();

// Fixed gccgo(1.4) compile failed, becase of 'success_cb' duplicate definition
pa_context_success_cb_t get_success_cb();

pa_stream* createMonitorStreamForSource(pa_context* ctx, uint32_t source_idx, uint32_t stream_idx, int suspend);

void suspend_sink_by_id(pa_context* ctx, uint32_t idx, int suspend);
void suspend_source_by_id(pa_context* ctx, uint32_t idx, int suspend);

#endif
