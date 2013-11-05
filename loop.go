package dlib

/*
#include <glib.h>
void _run() {
	g_main_loop_run(g_main_loop_new(0, FALSE));
}
*/
import "C"

func StartLoop() {
	C._run()
}
