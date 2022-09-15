// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#include <pulse/pulseaudio.h>

#include "dde-pulse.h"
#include <string.h>
#include <stdio.h>

#include "_cgo_export.h" //convert int

int connect_timeout = 0;

static inline void __empty_success_cb(pa_context *c, int success, void *userdata)
{
  if (success) {
    return;
  }

  int en = pa_context_errno(c);
  fprintf(stderr, "go-lib/pulse operation failed: %s. %p %p\n",
          pa_strerror(en),
          c,
          userdata);

}

pa_context_success_cb_t get_success_cb()
{
  return __empty_success_cb;
}

void set_connect_timeout()
{
  connect_timeout = 1;
}

static void __empty_index_cb(pa_context *c, uint32_t idx, void *userdata)
{
    return;
}

pa_context_index_cb_t get_index_cb()
{
    return __empty_index_cb;
}

#define DEFINE(ID, TYPE, PA_FUNC_SUFFIX)                                \
  void receive_##TYPE##_cb(pa_context *c, const pa_##TYPE##_info *info, int eol, void *userdata) \
  {                                                                     \
    if (eol < 0) {                                                      \
      int en = pa_context_errno(c);                                     \
      fprintf(stderr, "receive_%s_cb failed: %s\n",                     \
              #TYPE,                                                    \
              pa_strerror(en));                                         \
    }                                                                   \
    go_receive_some_info((int64_t)userdata, ID, (void*)info, eol);      \
  }                                                                     \
  void _get_##TYPE##_info(pa_threaded_mainloop* loop, pa_context *c, int64_t cookie, uint32_t index) \
  {                                                                     \
    pa_operation_unref(pa_context_get_##TYPE##_info##PA_FUNC_SUFFIX(c, index, receive_##TYPE##_cb, (void*)cookie)); \
  }                                                                     \
  void _get_##TYPE##_info_list(pa_threaded_mainloop* loop, pa_context* ctx, int64_t cookie) \
  {                                                                     \
    pa_operation_unref(pa_context_get_##TYPE##_info_list(ctx, receive_##TYPE##_cb, (void*)cookie)); \
  }

DEFINE(PA_SUBSCRIPTION_EVENT_SINK, sink, _by_index);
DEFINE(PA_SUBSCRIPTION_EVENT_SOURCE, source, _by_index);
DEFINE(PA_SUBSCRIPTION_EVENT_SINK_INPUT, sink_input, );
DEFINE(PA_SUBSCRIPTION_EVENT_SOURCE_OUTPUT, source_output, );
DEFINE(PA_SUBSCRIPTION_EVENT_CARD, card, _by_index);
DEFINE(PA_SUBSCRIPTION_EVENT_CLIENT, client, );
DEFINE(PA_SUBSCRIPTION_EVENT_MODULE, module, );
DEFINE(PA_SUBSCRIPTION_EVENT_SAMPLE_CACHE, sample, _by_index);


void receive_server_info_cb(pa_context *c, const pa_server_info *info, void *userdata)
{
    if (info == NULL) {
      return;
    }
    go_receive_some_info((int64_t)userdata, PA_SUBSCRIPTION_EVENT_SERVER, (void*)info, 0);
}
void _get_server_info(pa_threaded_mainloop* loop, pa_context *c, int64_t cookie)
{
    pa_operation_unref(pa_context_get_server_info(c, receive_server_info_cb, (void*)cookie));
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
                         PA_SUBSCRIPTION_MASK_SERVER |
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


static void
_wait_context_state_change(pa_context* ctx, void* userdata)
{
  pa_threaded_mainloop* loop = userdata;
  pa_threaded_mainloop_signal(loop, 0);
}

pa_context* new_pa_context(pa_threaded_mainloop* m)
{
    pa_threaded_mainloop_lock(m);
    pa_threaded_mainloop_start(m);
    pa_mainloop_api* mlapi = pa_threaded_mainloop_get_api(m);
    pa_context* ctx = pa_context_new(mlapi, "go-pulseaudio");
    pa_threaded_mainloop_unlock(m);

    if (!ctx) {
	fprintf(stderr, "pa_context_new() failed.\n");
	return NULL;
    }

    connect_timeout = 0;

    pa_threaded_mainloop_lock(m);
    pa_context_set_state_callback(ctx, _wait_context_state_change, m);
    pa_context_connect(ctx, NULL, PA_CONTEXT_NOFAIL, NULL);
    int state = pa_context_get_state(ctx);
    while (state != PA_CONTEXT_READY) {
        // Exit condition one.
        if (state == PA_CONTEXT_FAILED || state == PA_CONTEXT_TERMINATED) {
          pa_threaded_mainloop_unlock(m);
          fprintf(stderr, "Failed Connect to pulseaudio server\n");
          return NULL;
        }

        // It must be the last line of while loop.
        // wait for connect state changed.
        pa_threaded_mainloop_wait(m);

        // Exit condition two.
        if (connect_timeout) {
          pa_threaded_mainloop_unlock(m);
          fprintf(stderr, "Failed Connect to pulseaudio server timeout\n");
          return NULL;
        }

        state = pa_context_get_state(ctx);
    }
    pa_threaded_mainloop_unlock(m);

    setup_monitor(m, ctx);
    return ctx;
}

// #define PA_INVALID_INDEX ((uint32_t) -1)
//
// If idx is PA_INVALID_INDEX, all sinks will be suspended.
void
_suspend_sink_by_id(pa_threaded_mainloop *loop, pa_context* ctx, uint32_t idx, int suspend)
{
    pa_operation* o = pa_context_suspend_sink_by_index(ctx, idx, suspend,
                                                       get_success_cb(), NULL);

    if (!o) {
        fprintf(stderr, "Failed suspend sink %u\n", idx);
        return;
    }
    pa_operation_unref(o);
}

// If idx is PA_INVALID_INDEX, all sources will be suspended.
void
_suspend_source_by_id(pa_threaded_mainloop *loop, pa_context* ctx, uint32_t idx, int suspend)
{
    pa_operation* o = pa_context_suspend_source_by_index(ctx, idx, suspend,
                                                         get_success_cb(), NULL);

    if (!o) {
        fprintf(stderr, "Failed suspend sink %u\n", idx);
        return;
    }
    pa_operation_unref(o);
}
