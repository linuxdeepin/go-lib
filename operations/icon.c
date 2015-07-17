#include <string.h>
#include <gtk/gtk.h>
#include <gio/gdesktopappinfo.h>

#define is_xpm(ext) (g_ascii_strcasecmp(ext, "xpm") == 0)
#define is_dataurl(ext) g_str_has_prefix(ext, "data:image")


static
char* get_data_uri_by_pixbuf(GdkPixbuf* pixbuf)
{
    gchar* buf = NULL;
    gsize size = 0;
    GError *error = NULL;

    gdk_pixbuf_save_to_buffer(pixbuf, &buf, &size, "png", &error, NULL);
    g_assert(buf != NULL);

    if (error != NULL) {
        g_warning("[%s] %s\n", __func__, error->message);
        g_error_free(error);
        g_free(buf);
        return NULL;
    }

    char* base64 = g_base64_encode((const guchar*)buf, size);
    g_free(buf);
    char* data = g_strconcat("data:image/png;base64,", base64, NULL);
    g_free(base64);

    return data;
}


static
char* get_data_uri_by_path(const char* path)
{
    GError *error = NULL;
    GdkPixbuf* pixbuf = gdk_pixbuf_new_from_file(path, &error);
    if (error != NULL) {
        g_warning("%s\n", error->message);
        g_error_free(error);
        return NULL;
    }
    char* c = get_data_uri_by_pixbuf(pixbuf);
    g_object_unref(pixbuf);
    return c;

}


static
char* icon_name_to_path(const char* name, int size)
{
    if (g_path_is_absolute(name))
        return g_strdup(name);

    g_return_val_if_fail(name != NULL, NULL);

    int pic_name_len = strlen(name);
    char* ext = strrchr(name, '.');
    if (ext != NULL) {
        if (g_ascii_strcasecmp(ext+1, "png") == 0 || g_ascii_strcasecmp(ext+1, "svg") == 0 || g_ascii_strcasecmp(ext+1, "jpg") == 0) {
            pic_name_len = ext - name;
            g_debug("desktop's Icon name should an absoulte path or an basename without extension");
        }
    }

    char* pic_name = g_strndup(name, pic_name_len);
    GtkIconTheme* them = gtk_icon_theme_get_default(); // NB: do not ref or unref it
    GtkIconInfo* info = gtk_icon_theme_lookup_icon(them, pic_name, size, GTK_ICON_LOOKUP_GENERIC_FALLBACK);

    if (info == NULL) {
        g_warning("get gtk icon theme info failed for %s", pic_name);
        g_free(pic_name);
        return NULL;
    }
    g_free(pic_name);

    char* path = g_strdup(gtk_icon_info_get_filename(info));

#if GTK_MAJOR_VERSION >= 3
    g_object_unref(info);
#elif GTK_MAJOR_VERSION == 2
    gtk_icon_info_free(info);
#endif
    g_debug("get icon from icon theme is: %s", path);
    return path;
}


static
char* check_xpm(const char* path)
{
    if (path == NULL)
        return NULL;
    char* ext = strrchr(path, '.');
    if (ext != NULL && is_xpm(ext+1)) {
        return get_data_uri_by_path(path);
    } else {
        return g_strdup(path);
    }
}


char* icon_name_to_path_with_check_xpm(const char* name, int size)
{
    char* path = icon_name_to_path(name, size);
    char* icon = check_xpm(path);
    g_free(path);
    return icon;
}


char* get_icon_for_app(char const* icon_str, int size)
{
    char* icon = icon_name_to_path_with_check_xpm(icon_str, size);
    g_debug("the final icon of app is: %s", icon);
    return icon;
}


char* get_icon_for_file(char* giconstr, int size)
{
    if (giconstr == NULL) {
        return NULL;
    }

    char* icon = NULL;
    char** icon_names = g_strsplit(giconstr, " ", -1);

    for (int i = 0; icon_names[i] != NULL && icon == NULL; ++i) {
        icon = icon_name_to_path(icon_names[i], size);
    }

    g_strfreev(icon_names);

    return icon;
}

