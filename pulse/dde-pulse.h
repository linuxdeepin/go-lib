// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#ifndef DDE_PULSE_H
#define DDE_PULSE_H

/*
  ANY function touch a _pa_context_ must hold the pa_threaded_mainloop lock !!!!!!
  And may object created by pa_context, like pa_stream_new also need hold the lock.
 */

#include <pulse/pulseaudio.h>

#define DECLARE(TYPE) \
  void _get_##TYPE##_info(pa_threaded_mainloop*, pa_context*, int64_t, uint32_t); \
  void _get_##TYPE##_info_list(pa_threaded_mainloop*, pa_context*, int64_t);

DECLARE(sink);
DECLARE(sink_input);
DECLARE(source);
DECLARE(source_output);
DECLARE(client);
DECLARE(card);
DECLARE(module);
DECLARE(sample);

void _get_server_info(pa_threaded_mainloop*, pa_context *c, int64_t cookie);

pa_context* new_pa_context(pa_threaded_mainloop* ml);

// Fixed gccgo(1.4) compile failed, becase of 'success_cb' duplicate definition
pa_context_success_cb_t get_success_cb();

void set_connect_timeout();

pa_context_index_cb_t get_index_cb();

pa_stream* createMonitorStreamForSource(pa_threaded_mainloop* loop, pa_context* ctx, uint32_t source_idx, uint32_t stream_idx, int suspend);

void _suspend_sink_by_id(pa_threaded_mainloop* loop, pa_context* ctx, uint32_t idx, int suspend);
void _suspend_source_by_id(pa_threaded_mainloop*  loop, pa_context* ctx, uint32_t idx, int suspend);

#endif
