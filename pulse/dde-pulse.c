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

#include <pulse/pulseaudio.h>

#include "dde-pulse.h"
#include <string.h>
#include <stdio.h>

#include "_cgo_export.h" //convert int

static inline void __empty_success_cb(pa_context *c, int success, void *userdata)
{
}
pa_context_success_cb_t get_success_cb()
{
  return __empty_success_cb;
}

#define DEFINE(ID, TYPE, PA_FUNC_SUFFIX)                                \
  void receive_##TYPE##_cb(pa_context *c, const pa_##TYPE##_info *i, int eol, void *userdata) \
  {                                                                     \
    go_receive_some_info((int64_t)userdata, ID, (void*)i, eol);            \
    if (eol < 0) {                                                      \
      if (pa_context_errno(c) == PA_ERR_NOENTITY) {                     \
        fprintf(stderr, "errno == PA_ERR_NOENTITY");                    \
        return;                                                         \
      }                                                                 \
      fprintf(stderr, "receive info failed");                           \
      return;                                                           \
    }                                                                   \
  }                                                                     \
  void get_##TYPE##_info(pa_threaded_mainloop* loop, pa_context *c, int64_t cookie, uint32_t index) \
  {                                                                     \
    pa_threaded_mainloop_lock(loop);                                    \
    pa_operation_unref(pa_context_get_##TYPE##_info##PA_FUNC_SUFFIX(c, index, receive_##TYPE##_cb, (void*)cookie)); \
    pa_threaded_mainloop_unlock(loop);                                  \
  }                                                                     \
  void get_##TYPE##_info_list(pa_threaded_mainloop* loop, pa_context* ctx, int64_t cookie) \
  {                                                                     \
    pa_threaded_mainloop_lock(loop);                                    \
    pa_operation_unref(pa_context_get_##TYPE##_info_list(ctx, receive_##TYPE##_cb, (void*)cookie)); \
    pa_threaded_mainloop_unlock(loop);                                  \
  }

DEFINE(PA_SUBSCRIPTION_EVENT_SINK, sink, _by_index);
DEFINE(PA_SUBSCRIPTION_EVENT_SOURCE, source, _by_index);
DEFINE(PA_SUBSCRIPTION_EVENT_SINK_INPUT, sink_input, );
DEFINE(PA_SUBSCRIPTION_EVENT_SOURCE_OUTPUT, source_output, );
DEFINE(PA_SUBSCRIPTION_EVENT_CARD, card, _by_index);
DEFINE(PA_SUBSCRIPTION_EVENT_CLIENT, client, );
DEFINE(PA_SUBSCRIPTION_EVENT_MODULE, module, );
DEFINE(PA_SUBSCRIPTION_EVENT_SAMPLE_CACHE, sample, _by_index);


void receive_server_info_cb(pa_context *c, const pa_server_info *i, void *userdata)
{
    if (i == NULL) {
      return;
    }
    pa_server_info *info = NULL;
    info = malloc(sizeof(pa_server_info));
    memcpy(info, i, sizeof(pa_server_info));
    go_receive_some_info((int64_t)userdata, PA_SUBSCRIPTION_EVENT_SERVER, (void*)info, 0);
}
void get_server_info(pa_threaded_mainloop* loop, pa_context *c, int64_t cookie)
{
    pa_threaded_mainloop_lock(loop);
    pa_operation_unref(pa_context_get_server_info(c, receive_server_info_cb, (void*)cookie));
    pa_threaded_mainloop_unlock(loop);
}

void dpa_context_subscribe_cb(pa_context *c, pa_subscription_event_type_t t, uint32_t idx, void *userdata)
{
    int facility = t & PA_SUBSCRIPTION_EVENT_FACILITY_MASK;
    int event_type = t & PA_SUBSCRIPTION_EVENT_TYPE_MASK;

    go_handle_changed(facility, event_type, idx);
}

static void
dpa_context_state_cb(pa_context* ctx, void* userdata)
{
    int state = pa_context_get_state (ctx);
    go_handle_state_changed(state);
}

void setup_monitor(pa_threaded_mainloop* m, pa_context *ctx)
{
    pa_threaded_mainloop_lock(m);
    pa_context_set_state_callback(ctx, dpa_context_state_cb, NULL);
    pa_context_set_subscribe_callback(ctx, dpa_context_subscribe_cb, NULL);
    pa_context_subscribe(ctx,
                         PA_SUBSCRIPTION_MASK_CARD |
                         PA_SUBSCRIPTION_MASK_SINK |
                         PA_SUBSCRIPTION_MASK_SOURCE |
                         PA_SUBSCRIPTION_MASK_SINK_INPUT |
                         PA_SUBSCRIPTION_MASK_SOURCE_OUTPUT |
                         PA_SUBSCRIPTION_MASK_SAMPLE_CACHE,
                         NULL,
                         NULL);
    pa_threaded_mainloop_unlock(m);
}


//TODO: the init_state should be protect by lock
static int init_state = 0; // O: unknown, 1: success, 2: failure

pa_context* new_pa_context(pa_threaded_mainloop* m)
{
    pa_threaded_mainloop_lock(m);
    pa_threaded_mainloop_start(m);

    pa_mainloop_api* mlapi = pa_threaded_mainloop_get_api(m);
    pa_context* ctx = pa_context_new(mlapi, "go-pulseaudio");

    pa_context_connect(ctx, NULL, PA_CONTEXT_NOFAIL, NULL);
    int	state = pa_context_get_state(ctx);

    pa_threaded_mainloop_unlock(m);

    init_state = 0;
    while(state != PA_CONTEXT_READY) {
        if (init_state != 0) {
            break;
        }
        pa_threaded_mainloop_lock(m);
        state = pa_context_get_state(ctx);
        pa_threaded_mainloop_unlock(m);
        if (state == PA_CONTEXT_FAILED || state == PA_CONTEXT_TERMINATED) {
            init_state = 2;
            fprintf(stderr, "Failed Connect to pulseaudio server");
            return NULL;
        }
    }

    if (init_state == 2) {
        fprintf(stderr, "Connect to pulseaudio timeout\n");
        return NULL;
    }
    init_state = 1;
    setup_monitor(m, ctx);
    return ctx;
}

void
pa_finalize()
{
    if (init_state != 0) {
        return;
    }
    init_state = 2;
}

// #define PA_INVALID_INDEX ((uint32_t) -1)
//
// If idx is PA_INVALID_INDEX, all sinks will be suspended.
void
suspend_sink_by_id(pa_threaded_mainloop *loop, pa_context* ctx, uint32_t idx, int suspend)
{
    pa_threaded_mainloop_lock(loop);
    pa_operation* o = pa_context_suspend_sink_by_index(ctx, idx, suspend,
                                                       get_success_cb(), NULL);
    pa_threaded_mainloop_unlock(loop);

    if (!o) {
        fprintf(stderr, "Failed suspend sink %u\n", idx);
        return;
    }
    pa_operation_unref(o);
}

// If idx is PA_INVALID_INDEX, all sources will be suspended.
void
suspend_source_by_id(pa_threaded_mainloop *loop, pa_context* ctx, uint32_t idx, int suspend)
{
    pa_threaded_mainloop_lock(loop);
    pa_operation* o = pa_context_suspend_source_by_index(ctx, idx, suspend,
                                                         get_success_cb(), NULL);
    pa_threaded_mainloop_unlock(loop);

    if (!o) {
        fprintf(stderr, "Failed suspend sink %u\n", idx);
        return;
    }
    pa_operation_unref(o);
}
