#ifndef DDE_PULSE_H
#define DDE_PULSE_H

#include <pulse/pulseaudio.h>

#define DECLARE(TYPE) \
void get_##TYPE##_info(pa_context*, int64_t, uint32_t);\
void get_##TYPE##_info_list(pa_context*, int64_t);

DECLARE(sink);
DECLARE(sink_input);
DECLARE(source);
DECLARE(source_output);
DECLARE(client);
DECLARE(card);
DECLARE(module);
DECLARE(sample);

void get_server_info(pa_context *c, int64_t cookie);

pa_context* pa_init(pa_threaded_mainloop* ml);

void (*success_cb)(pa_context *c, int success, void *userdata);

pa_stream* createMonitorStreamForSource(pa_context* ctx, uint32_t source_idx, uint32_t stream_idx, int suspend);

#endif
