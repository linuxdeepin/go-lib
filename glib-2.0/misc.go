package glib

/*
#include <glib.h>
void _run() {
	g_main_loop_run(g_main_loop_new(0, FALSE));
}
#cgo pkg-config: glib-2.0
*/
import "C"

func StartLoop() {
	C._run()
}
