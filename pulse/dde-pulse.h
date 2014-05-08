#ifndef DDE_PULSE_H
#define DDE_PULSE_H

#include <pulse/pulseaudio.h>
#include <pulse/glib-mainloop.h>

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

pa_context* pa_init(pa_mainloop* ml);


void (*success_cb)(pa_context *c, int success, void *userdata);

#endif
