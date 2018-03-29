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

#include "dde-pulse.h"
#include <string.h>
#include <stdio.h>
#include "_cgo_export.h"


static void suspended_callback(pa_stream *s, void *userdata) {
    /*MainWindow *w = static_cast<MainWindow*>(userdata);*/

    /*if (pa_stream_is_suspended(s))*/
        /*w->updateVolumeMeter(pa_stream_get_device_index(s), PA_INVALID_INDEX, -1);*/
}

static void read_callback(pa_stream *s, size_t length, void *userdata) {
    const void *data;
    double v;

    if (pa_stream_peek(s, &data, &length) < 0) {
	fprintf(stderr, "Failed to read data from stream");
	return;
    }

    if (!data) {
	pa_stream_drop(s);
	return;
    }

    assert(length > 0);
    assert(length % sizeof(float) == 0);

    v = ((const float*) data)[length / sizeof(float) -1];

    pa_stream_drop(s);

    if (v < 0)
	v = 0;
    if (v > 1)
	v = 1;

    go_update_volume_meter(pa_stream_get_device_index(s), pa_stream_get_monitor_stream(s), v);
}


pa_stream* createMonitorStreamForSource(pa_threaded_mainloop* loop, pa_context* ctx, uint32_t source_idx, uint32_t stream_idx, int suspend)
{
    stream_idx = -1;

    pa_stream *s;
    char t[16];
    pa_buffer_attr attr;
    pa_sample_spec ss;
    pa_stream_flags_t flags;

    ss.channels = 1;
    ss.format = PA_SAMPLE_FLOAT32;
    ss.rate = 25;

    memset(&attr, 0, sizeof(attr));
    attr.fragsize = sizeof(float);
    attr.maxlength = (uint32_t) -1;

    snprintf(t, sizeof(t), "%u", source_idx);

    pa_threaded_mainloop_lock(loop);
    if (!(s = pa_stream_new(ctx, "Peak detect", &ss, NULL))) {
        pa_threaded_mainloop_unlock(loop);
        fprintf(stderr, "Failed to create monitoring stream");
        return NULL;
    }

    if (stream_idx != (uint32_t) -1)
        pa_stream_set_monitor_stream(s, stream_idx);

    pa_stream_set_read_callback(s, read_callback, NULL);
    pa_stream_set_suspended_callback(s, suspended_callback, NULL);

    flags = (pa_stream_flags_t) (PA_STREAM_DONT_MOVE | PA_STREAM_PEAK_DETECT | PA_STREAM_ADJUST_LATENCY |
                                 (suspend ? PA_STREAM_DONT_INHIBIT_AUTO_SUSPEND : PA_STREAM_NOFLAGS));

    if (pa_stream_connect_record(s, t, &attr, flags) < 0) {
        fprintf(stderr, "Failed to connect monitoring stream");
        pa_stream_unref(s);
        pa_threaded_mainloop_unlock(loop);
        return NULL;
    }
    pa_threaded_mainloop_unlock(loop);
    return s;
}
