#include <gtk/gtk.h>

#include "cursor.h"

static void update_gtk_cursor();

static int sig_id = 0;

void
handle_gtk_cursor_changed()
{
	if (!gtk_init_check(NULL, NULL)) {
		return;
	}

	GtkSettings* s = gtk_settings_get_default();
	sig_id = g_signal_connect(s, "notify::gtk-cursor-theme-name", 
			update_gtk_cursor, NULL);
	if (sig_id <= 0) {
		return;
	}

	gtk_main();
}

static void
update_gtk_cursor()
{
	GtkSettings* s = gtk_settings_get_default();
	g_signal_handler_disconnect(s, sig_id);
	sig_id = 0;

	GdkCursor* cursor = gdk_cursor_new_for_display(
			gdk_display_get_default(), 
			GDK_LEFT_PTR);
	gdk_window_set_cursor(gdk_get_default_root_window(), cursor);
	g_object_unref(G_OBJECT(cursor));

	gtk_main_quit();
}
