#include <gtk/gtk.h>
#include <glib.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <X11/Xlib.h>

G_LOCK_DEFINE_STATIC(clipboard_info_lock);

static struct ClipboardInfo {
    char** files;
    int num;
    gboolean cut;
} clipboard_info = {
    .files = NULL,
    .num = 0,
    .cut = FALSE,
};


static GdkAtom copied_files_atom = GDK_NONE;
#define clipboard gtk_clipboard_get(GDK_SELECTION_CLIPBOARD)

static int can_paste = 0;


static
int do_get_can_paste(GdkAtom* atoms, gint n_atoms)
{
    int can_paste = 0;
    for(int i = 0; i < n_atoms; ++i) {
        if(atoms[i] == copied_files_atom) {
            can_paste = 1;
            break;
        }
    }

    return can_paste;
}


static
void targets_recived(GtkClipboard* c, GdkAtom* atoms, gint n_atoms, gpointer data)
{
    can_paste = 0;
    if(atoms == NULL) return;
    can_paste = do_get_can_paste(atoms, n_atoms);
}

static
void owner_change_callback(GtkClipboard* c, GdkEvent* ev, gpointer data)
{
    gtk_clipboard_request_targets(c, targets_recived, NULL);
}

static
void init_clipboard()
{
    if(copied_files_atom == GDK_NONE) {
        /* gdk_error_trap_push(); */
        copied_files_atom = gdk_atom_intern("x-special/gnome-copied-files", FALSE);
        g_signal_connect(clipboard, "owner-change", G_CALLBACK(owner_change_callback), NULL);
    }
}


int get_can_paste()
{
    init_clipboard();

    static gboolean inited = FALSE;
    if (!inited) {
        inited = TRUE;
        GdkAtom* targets = NULL;
        int n_targets = 0;
        gtk_clipboard_wait_for_targets(clipboard, &targets, &n_targets);
        can_paste = do_get_can_paste(targets, n_targets);;
        g_free(targets);
    }
    return can_paste;
}


char** get_clipboard_content(int *n)
{
    init_clipboard();
    GtkSelectionData* selection_data = gtk_clipboard_wait_for_contents(clipboard, copied_files_atom);
    if (selection_data == NULL) {
        g_warning("get selection data from clipboard failed");
        return NULL;
    }

    if (gtk_selection_data_get_data_type (selection_data) != copied_files_atom) {
        g_warning("has wrong data type from selection data");
        gtk_selection_data_free(selection_data);
        return NULL;
    }

    int len = 0;
    const guchar* data = gtk_selection_data_get_data_with_length(selection_data, &len);
    if(len <= 0) {
        g_warning("has no selection data for clipboard");
        gtk_selection_data_free(selection_data);
        return NULL;
    }
    len++;

    // !!! data[len] = '\0' is dangerous, copy it.
    guchar* copied_data = g_malloc0(len);
    memmove(copied_data, data, len);
    char** lines = g_strsplit(copied_data, "\n", 0);
    if (n != NULL) {
        for (*n = 0; lines[*n] != NULL; *n +=1) {/* do nothing */}
    }
    g_free(copied_data);
    gtk_selection_data_free(selection_data);

    return lines;
}


static
char* convert_file_list_to_string(struct ClipboardInfo* info, gboolean format_for_text, gsize *len)
{
    GString *uris;
    if(format_for_text) {
        uris = g_string_new(NULL);
    } else {
        uris = g_string_new(info->cut ? "cut" : "copy");
    }

    for(int i = 0; i < info->num; i++) {
        char* uri = info->files[i];

        if(format_for_text) {
            GFile* f = g_file_new_for_uri(uri);
            char* tmp = g_file_get_parse_name(f);
            g_object_unref(f);

            if(tmp != NULL) {
                g_string_append(uris, tmp);
                g_free(tmp);
            } else {
                g_string_append(uris, uri);
            }

            /* skip newline for last element */
            if(i + 1 < info->num) {
                g_string_append_c(uris, '\n');
            }
        } else {
            g_string_append_c(uris, '\n');
            g_string_append(uris, uri);
        }
    }

    g_debug("convert_file_list_to_string done");
    *len = uris->len;
    return g_string_free(uris, FALSE);
}


static
void clipboard_get_callback(GtkClipboard* c, GtkSelectionData* selection_data, guint info, gpointer clear)
{
    GdkAtom target = gtk_selection_data_get_target(selection_data);

    if (GPOINTER_TO_INT(clear)) {
        gtk_clipboard_clear(c);
        return;
    }

    if(gtk_targets_include_uri(&target, 1)) {
        G_LOCK(clipboard_info_lock);
        gtk_selection_data_set_uris(selection_data, clipboard_info.files);
        G_UNLOCK(clipboard_info_lock);
    } else if(gtk_targets_include_text(&target, 1)) {
        char *str;
        gsize len;

        G_LOCK(clipboard_info_lock);
        str = convert_file_list_to_string(&clipboard_info, TRUE, &len);
        G_UNLOCK(clipboard_info_lock);
        gtk_selection_data_set_text(selection_data, str, len);
        g_free(str);
    } else if(target == copied_files_atom) {
        char *str;
        gsize len;

        G_LOCK(clipboard_info_lock);
        str = convert_file_list_to_string(&clipboard_info, FALSE, &len);
        G_UNLOCK(clipboard_info_lock);
        gtk_selection_data_set(selection_data, copied_files_atom, 8, (guchar *) str, len);
        g_free(str);
    }
}


static
void clipboard_clear_callback(GtkClipboard* c, gpointer clear)
{
}


static
void do_set_op_content_to_clipboard(int is_cut, char** files, int n, int clear)
{
    GtkTargetList *target_list;
    GtkTargetEntry *targets;
    int n_targets;

    char** clipboard_contents = g_malloc((n+1)*sizeof(char*));
    for(int i = 0; i < n; ++i) {
        clipboard_contents[i] = g_strdup(files[i]);
        g_debug("copy %s", files[i]);
    }
    clipboard_contents[n] = NULL;
    g_debug("copy done");

    // TODO: maybe a new alloced ClipboardInfo is better.
    G_LOCK(clipboard_info_lock);
    if (clipboard_info.files != NULL) {
        g_clear_pointer(&clipboard_info.files, g_strfreev);
    }

    clipboard_info.files = clipboard_contents;
    clipboard_info.cut = is_cut;
    clipboard_info.num = n;
    G_UNLOCK(clipboard_info_lock);

    target_list = gtk_target_list_new(NULL, 0);
    gtk_target_list_add(target_list, copied_files_atom, 0, 0);
    gtk_target_list_add_uri_targets(target_list, 0);
    gtk_target_list_add_text_targets(target_list, 0);

    targets = gtk_target_table_new_from_list(target_list, &n_targets);
    gtk_target_list_unref(target_list);

    gtk_clipboard_set_with_data(clipboard,
                                targets, n_targets,
                                clipboard_get_callback,
                                clipboard_clear_callback,
                                GINT_TO_POINTER(clear));
    gtk_target_table_free(targets, n_targets);
}


void freeStrv(void* p)
{
    g_strfreev(p);
}

void set_op_content_to_clipboard(int is_cut, char** files, int n)
{
    init_clipboard();
    do_set_op_content_to_clipboard(is_cut, files, n, 0);
}


void clear_clipboard()
{
    init_clipboard();
    do_set_op_content_to_clipboard(1, NULL, 0, 1);
}


