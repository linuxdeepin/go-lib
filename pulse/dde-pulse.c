#include <pulse/pulseaudio.h>

#include "dde-pulse.h"
#include <string.h>
#include <stdio.h>

#include "_cgo_export.h" //convert int

static pa_threaded_mainloop* m = NULL;

#define DEFINE(ID, TYPE, PA_FUNC_SUFFIX) \
void receive_##TYPE##_cb(pa_context *c, const pa_##TYPE##_info *i, int eol, void *userdata) \
{\
    receive_some_info((int64_t)userdata, ID, (void*)i, eol); \
    if (eol < 0) { \
	if (pa_context_errno(c) == PA_ERR_NOENTITY) {\
	    fprintf(stderr, "errno == PA_ERR_NOENTITY"); \
	    return;\
	}\
	fprintf(stderr, "receive info failed");\
	return;\
    }\
}\
void get_##TYPE##_info(pa_context *c, int64_t cookie, uint32_t index) \
{\
    pa_threaded_mainloop_lock(m);\
    pa_operation_unref(pa_context_get_##TYPE##_info##PA_FUNC_SUFFIX(c, index, receive_##TYPE##_cb, (void*)cookie)); \
    pa_threaded_mainloop_unlock(m);\
}\
void get_##TYPE##_info_list(pa_context* ctx, int64_t cookie) \
{\
    pa_threaded_mainloop_lock(m);\
    pa_context_get_##TYPE##_info_list(ctx, receive_##TYPE##_cb, (void*)cookie);\
    pa_threaded_mainloop_unlock(m);\
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
    receive_some_info((int64_t)userdata, PA_SUBSCRIPTION_EVENT_SERVER, (void*)info, 0);
}
void get_server_info(pa_context *c, int64_t cookie) 
{
    pa_threaded_mainloop_lock(m);
    pa_operation_unref(pa_context_get_server_info(c, receive_server_info_cb, (void*)cookie));
    pa_threaded_mainloop_unlock(m);
}

void dpa_context_subscribe_cb(pa_context *c, pa_subscription_event_type_t t, uint32_t idx, void *userdata)
{
    int facility = t & PA_SUBSCRIPTION_EVENT_FACILITY_MASK;
    int event_type = t & PA_SUBSCRIPTION_EVENT_TYPE_MASK;

    go_handle_changed(facility, event_type, idx);
}

void setup_monitor(pa_context *ctx)
{
    pa_threaded_mainloop_lock(m);
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

void __success_cb(pa_context *c, int success, void *userdata)
{
}
pa_context* pa_init(pa_threaded_mainloop* ml)
{
    m = ml;

    pa_threaded_mainloop_lock(m);
    pa_mainloop_api* mlapi = pa_threaded_mainloop_get_api(m);
    pa_context* ctx = pa_context_new(mlapi, "go-pulseaudio");

    pa_context_connect(ctx, NULL, PA_CONTEXT_NOFAIL, NULL);
    int	state = pa_context_get_state(ctx);

    pa_threaded_mainloop_unlock(m);

    while(state != PA_CONTEXT_READY) {
	pa_threaded_mainloop_lock(m);
	state = pa_context_get_state(ctx);
	pa_threaded_mainloop_unlock(m);
	if (state == PA_CONTEXT_FAILED || state == PA_CONTEXT_TERMINATED) {
	    fprintf(stderr, "Failed Connect to pulseaudio server");
	    return NULL;
	}
    }
    success_cb = __success_cb;
    setup_monitor(ctx);
    return ctx;
}
