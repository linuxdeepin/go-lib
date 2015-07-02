#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <glib.h>
#include <gio/gio.h>

#define N_
#define _

static char *
shorten_utf8_string (const char *base, int reduce_by_num_bytes)
{
	int len;
	char *ret;
	const char *p;

	len = strlen (base);
        fprintf(stderr,"xx%dxx\n", len);
	len -= reduce_by_num_bytes;

	if (len <= 0) {
		return NULL;
	}

	ret = g_new (char, len + 1);

	p = base;
	while (len) {
		char *next;
		next = g_utf8_next_char (p);
		if (next - p > len || *next == '\0') {
			break;
		}

		len -= next - p;
		p = next;
	}

	if (p - base == 0) {
		g_free (ret);
		return NULL;
	} else {
		memcpy (ret, base, p - base);
		ret[p - base] = '\0';
		return ret;
	}
}


/* Note that we have these two separate functions with separate format
 * strings for ease of localization.
 */

static char *
get_link_name (const char *name, int count, int max_length)
{
	const char *format;
	char *result;
	int unshortened_length;
	gboolean use_count;

	g_assert (name != NULL);

	if (count < 0) {
		g_warning ("bad count in get_link_name");
		count = 0;
	}

	if (count <= 2) {
		/* Handle special cases for low numbers.
		 * Perhaps for some locales we will need to add more.
		 */
		switch (count) {
		default:
			g_assert_not_reached ();
			/* fall through */
		case 0:
			/* duplicate original file name */
			format = "%s";
			break;
		case 1:
			/* appended to new link file */
			format = _("Link to %s");
			break;
		case 2:
			/* appended to new link file */
			format = _("Another link to %s");
			break;
		}

		use_count = FALSE;
	} else {
		/* Handle special cases for the first few numbers of each ten.
		 * For locales where getting this exactly right is difficult,
		 * these can just be made all the same as the general case below.
		 */
		switch (count % 10) {
		case 1:
			/* Localizers: Feel free to leave out the "st" suffix
			 * if there's no way to do that nicely for a
			 * particular language.
			 */
			format = _("%'dst link to %s");
			break;
		case 2:
			/* appended to new link file */
			format = _("%'dnd link to %s");
			break;
		case 3:
			/* appended to new link file */
			format = _("%'drd link to %s");
			break;
		default:
			/* appended to new link file */
			format = _("%'dth link to %s");
			break;
		}

		use_count = TRUE;
	}

	if (use_count)
		result = g_strdup_printf (format, count, name);
	else
		result = g_strdup_printf (format, name);

	if (max_length > 0 && (unshortened_length = strlen (result)) > max_length) {
		char *new_name;

		new_name = shorten_utf8_string (name, unshortened_length - max_length);
		if (new_name) {
			g_free (result);

			if (use_count)
				result = g_strdup_printf (format, count, new_name);
			else
				result = g_strdup_printf (format, new_name);

			g_assert (strlen (result) <= max_length);
			g_free (new_name);
		}
	}

	return result;
}


/* Localizers:
 * Feel free to leave out the st, nd, rd and th suffix or
 * make some or all of them match.
 */

/* localizers: tag used to detect the first copy of a file */
static const char untranslated_copy_duplicate_tag[] = N_(" (copy)");
/* localizers: tag used to detect the second copy of a file */
static const char untranslated_another_copy_duplicate_tag[] = N_(" (another copy)");

/* localizers: tag used to detect the x11th copy of a file */
static const char untranslated_x11th_copy_duplicate_tag[] = N_("th copy)");
/* localizers: tag used to detect the x12th copy of a file */
static const char untranslated_x12th_copy_duplicate_tag[] = N_("th copy)");
/* localizers: tag used to detect the x13th copy of a file */
static const char untranslated_x13th_copy_duplicate_tag[] = N_("th copy)");

/* localizers: tag used to detect the x1st copy of a file */
static const char untranslated_st_copy_duplicate_tag[] = N_("st copy)");
/* localizers: tag used to detect the x2nd copy of a file */
static const char untranslated_nd_copy_duplicate_tag[] = N_("nd copy)");
/* localizers: tag used to detect the x3rd copy of a file */
static const char untranslated_rd_copy_duplicate_tag[] = N_("rd copy)");

/* localizers: tag used to detect the xxth copy of a file */
static const char untranslated_th_copy_duplicate_tag[] = N_("th copy)");

#define COPY_DUPLICATE_TAG _(untranslated_copy_duplicate_tag)
#define ANOTHER_COPY_DUPLICATE_TAG _(untranslated_another_copy_duplicate_tag)
#define X11TH_COPY_DUPLICATE_TAG _(untranslated_x11th_copy_duplicate_tag)
#define X12TH_COPY_DUPLICATE_TAG _(untranslated_x12th_copy_duplicate_tag)
#define X13TH_COPY_DUPLICATE_TAG _(untranslated_x13th_copy_duplicate_tag)

#define ST_COPY_DUPLICATE_TAG _(untranslated_st_copy_duplicate_tag)
#define ND_COPY_DUPLICATE_TAG _(untranslated_nd_copy_duplicate_tag)
#define RD_COPY_DUPLICATE_TAG _(untranslated_rd_copy_duplicate_tag)
#define TH_COPY_DUPLICATE_TAG _(untranslated_th_copy_duplicate_tag)

/* localizers: appended to first file copy */
static const char untranslated_first_copy_duplicate_format[] = N_("%s (copy)%s");
/* localizers: appended to second file copy */
static const char untranslated_second_copy_duplicate_format[] = N_("%s (another copy)%s");

/* localizers: appended to x11th file copy */
static const char untranslated_x11th_copy_duplicate_format[] = N_("%s (%'dth copy)%s");
/* localizers: appended to x12th file copy */
static const char untranslated_x12th_copy_duplicate_format[] = N_("%s (%'dth copy)%s");
/* localizers: appended to x13th file copy */
static const char untranslated_x13th_copy_duplicate_format[] = N_("%s (%'dth copy)%s");

/* localizers: if in your language there's no difference between 1st, 2nd, 3rd and nth
 * plurals, you can leave the st, nd, rd suffixes out and just make all the translated
 * strings look like "%s (copy %'d)%s".
 */

/* localizers: appended to x1st file copy */
static const char untranslated_st_copy_duplicate_format[] = N_("%s (%'dst copy)%s");
/* localizers: appended to x2nd file copy */
static const char untranslated_nd_copy_duplicate_format[] = N_("%s (%'dnd copy)%s");
/* localizers: appended to x3rd file copy */
static const char untranslated_rd_copy_duplicate_format[] = N_("%s (%'drd copy)%s");
/* localizers: appended to xxth file copy */
static const char untranslated_th_copy_duplicate_format[] = N_("%s (%'dth copy)%s");

#define FIRST_COPY_DUPLICATE_FORMAT _(untranslated_first_copy_duplicate_format)
#define SECOND_COPY_DUPLICATE_FORMAT _(untranslated_second_copy_duplicate_format)
#define X11TH_COPY_DUPLICATE_FORMAT _(untranslated_x11th_copy_duplicate_format)
#define X12TH_COPY_DUPLICATE_FORMAT _(untranslated_x12th_copy_duplicate_format)
#define X13TH_COPY_DUPLICATE_FORMAT _(untranslated_x13th_copy_duplicate_format)

#define ST_COPY_DUPLICATE_FORMAT _(untranslated_st_copy_duplicate_format)
#define ND_COPY_DUPLICATE_FORMAT _(untranslated_nd_copy_duplicate_format)
#define RD_COPY_DUPLICATE_FORMAT _(untranslated_rd_copy_duplicate_format)
#define TH_COPY_DUPLICATE_FORMAT _(untranslated_th_copy_duplicate_format)

static char *
extract_string_until (const char *original, const char *until_substring)
{
	char *result;

	g_assert ((int) strlen (original) >= until_substring - original);
	g_assert (until_substring - original >= 0);

	result = g_malloc (until_substring - original + 1);
	strncpy (result, original, until_substring - original);
	result[until_substring - original] = '\0';

	return result;
}

static char *
eel_filename_get_extension_offset (const char *filename)
{
	char *end, *end2;
	const char *start;

	if (filename == NULL || filename[0] == '\0') {
		return NULL;
	}

	/* basename must have at least one char */
	start = filename + 1;

	end = strrchr (start, '.');
	if (end == NULL || end[1] == '\0') {
		return NULL;
	}

	if (end != start) {
		if (strcmp (end, ".gz") == 0 ||
		    strcmp (end, ".bz2") == 0 ||
		    strcmp (end, ".sit") == 0 ||
		    strcmp (end, ".Z") == 0) {
			end2 = end - 1;
			while (end2 > start &&
			       *end2 != '.') {
				end2--;
			}
			if (end2 != start) {
				end = end2;
			}
		}
	}

	return end;
}

static int
get_filename_extension_offset(const char* start)
{
    const char* end = eel_filename_get_extension_offset(start);
    if (end == NULL) return -1;
    return end - start;
}

/* Dismantle a file name, separating the base name, the file suffix and removing any
 * (xxxcopy), etc. string. Figure out the count that corresponds to the given
 * (xxxcopy) substring.
 */
static void
parse_previous_duplicate_name (const char *name,
			       char **name_base,
			       const char **suffix,
			       int *count)
{
	const char *tag;

	g_assert (name[0] != '\0');

	*suffix = eel_filename_get_extension_offset (name);

	if (*suffix == NULL || (*suffix)[1] == '\0') {
		/* no suffix */
		*suffix = "";
	}

	tag = strstr (name, COPY_DUPLICATE_TAG);
	if (tag != NULL) {
		if (tag > *suffix) {
			/* handle case "foo. (copy)" */
			*suffix = "";
		}
		*name_base = extract_string_until (name, tag);
		*count = 1;
		return;
	}


	tag = strstr (name, ANOTHER_COPY_DUPLICATE_TAG);
	if (tag != NULL) {
		if (tag > *suffix) {
			/* handle case "foo. (another copy)" */
			*suffix = "";
		}
		*name_base = extract_string_until (name, tag);
		*count = 2;
		return;
	}


	/* Check to see if we got one of st, nd, rd, th. */
	tag = strstr (name, X11TH_COPY_DUPLICATE_TAG);

	if (tag == NULL) {
		tag = strstr (name, X12TH_COPY_DUPLICATE_TAG);
	}
	if (tag == NULL) {
		tag = strstr (name, X13TH_COPY_DUPLICATE_TAG);
	}

	if (tag == NULL) {
		tag = strstr (name, ST_COPY_DUPLICATE_TAG);
	}
	if (tag == NULL) {
		tag = strstr (name, ND_COPY_DUPLICATE_TAG);
	}
	if (tag == NULL) {
		tag = strstr (name, RD_COPY_DUPLICATE_TAG);
	}
	if (tag == NULL) {
		tag = strstr (name, TH_COPY_DUPLICATE_TAG);
	}

	/* If we got one of st, nd, rd, th, fish out the duplicate number. */
	if (tag != NULL) {
		/* localizers: opening parentheses to match the "th copy)" string */
		tag = strstr (name, _(" ("));
		if (tag != NULL) {
			if (tag > *suffix) {
				/* handle case "foo. (22nd copy)" */
				*suffix = "";
			}
			*name_base = extract_string_until (name, tag);
			/* localizers: opening parentheses of the "th copy)" string */
			if (sscanf (tag, _(" (%'d"), count) == 1) {
				if (*count < 1 || *count > 1000000) {
					/* keep the count within a reasonable range */
					*count = 0;
				}
				return;
			}
			*count = 0;
			return;
		}
	}


	*count = 0;
	if (**suffix != '\0') {
		*name_base = extract_string_until (name, *suffix);
	} else {
		*name_base = g_strdup (name);
	}
}

static char *
make_next_duplicate_name (const char *base, const char *suffix, int count, int max_length)
{
	const char *format;
	char *result;
	int unshortened_length;
	gboolean use_count;

	if (count < 1) {
		g_warning ("bad count %d in get_duplicate_name", count);
		count = 1;
	}

	if (count <= 2) {

		/* Handle special cases for low numbers.
		 * Perhaps for some locales we will need to add more.
		 */
		switch (count) {
		default:
			g_assert_not_reached ();
			/* fall through */
		case 1:
			format = FIRST_COPY_DUPLICATE_FORMAT;
			break;
		case 2:
			format = SECOND_COPY_DUPLICATE_FORMAT;
			break;

		}

		use_count = FALSE;
	} else {

		/* Handle special cases for the first few numbers of each ten.
		 * For locales where getting this exactly right is difficult,
		 * these can just be made all the same as the general case below.
		 */

		/* Handle special cases for x11th - x20th.
		 */
		switch (count % 100) {
		case 11:
			format = X11TH_COPY_DUPLICATE_FORMAT;
			break;
		case 12:
			format = X12TH_COPY_DUPLICATE_FORMAT;
			break;
		case 13:
			format = X13TH_COPY_DUPLICATE_FORMAT;
			break;
		default:
			format = NULL;
			break;
		}

		if (format == NULL) {
			switch (count % 10) {
			case 1:
				format = ST_COPY_DUPLICATE_FORMAT;
				break;
			case 2:
				format = ND_COPY_DUPLICATE_FORMAT;
				break;
			case 3:
				format = RD_COPY_DUPLICATE_FORMAT;
				break;
			default:
				/* The general case. */
				format = TH_COPY_DUPLICATE_FORMAT;
				break;
			}
		}

		use_count = TRUE;

	}

	if (use_count)
		result = g_strdup_printf (format, base, count, suffix);
	else
		result = g_strdup_printf (format, base, suffix);

	if (max_length > 0 && (unshortened_length = strlen (result)) > max_length) {
		char *new_base;

		new_base = shorten_utf8_string (base, unshortened_length - max_length);
		if (new_base) {
			g_free (result);

			if (use_count)
				result = g_strdup_printf (format, new_base, count, suffix);
			else
				result = g_strdup_printf (format, new_base, suffix);

			g_assert (strlen (result) <= max_length);
			g_free (new_base);
		}
	}

	return result;
}

static char *
get_duplicate_name (const char *name, int count_increment, int max_length)
{
	char *result;
	char *name_base;
	const char *suffix;
	int count;

	parse_previous_duplicate_name (name, &name_base, &suffix, &count);
	result = make_next_duplicate_name (name_base, suffix, count + count_increment, max_length);

	g_free (name_base);

	return result;
}


static int
get_max_name_length (GFile *file_dir)
{
	int max_length;
	char *dir;
	long max_path;
	long max_name;

	max_length = -1;

	if (!g_file_has_uri_scheme (file_dir, "file"))
		return max_length;

	dir = g_file_get_path (file_dir);
	if (!dir)
		return max_length;

	max_path = pathconf (dir, _PC_PATH_MAX);
	max_name = pathconf (dir, _PC_NAME_MAX);

	if (max_name == -1 && max_path == -1) {
		max_length = -1;
	} else if (max_name == -1 && max_path != -1) {
		max_length = max_path - (strlen (dir) + 1);
	} else if (max_name != -1 && max_path == -1) {
		max_length = max_name;
	} else {
		int leftover;

		leftover = max_path - (strlen (dir) + 1);

		max_length = MIN (leftover, max_name);
	}

	g_free (dir);

	return max_length;
}

#define FAT_FORBIDDEN_CHARACTERS "/:;*?\"<>"

static gboolean
fat_str_replace (char *str,
		 char replacement)
{
	gboolean success;
	int i;

	success = FALSE;
	for (i = 0; str[i] != '\0'; i++) {
		if (strchr (FAT_FORBIDDEN_CHARACTERS, str[i]) ||
		    str[i] < 32) {
			success = TRUE;
			str[i] = replacement;
		}
	}

	return success;
}

static gboolean
make_file_name_valid_for_dest_fs (char *filename,
				 const char *dest_fs_type)
{
	if (dest_fs_type != NULL && filename != NULL) {
		if (!strcmp (dest_fs_type, "fat")  ||
		    !strcmp (dest_fs_type, "vfat") ||
		    !strcmp (dest_fs_type, "msdos") ||
		    !strcmp (dest_fs_type, "msdosfs")) {
			gboolean ret;
			int i, old_len;

			ret = fat_str_replace (filename, '_');

			old_len = strlen (filename);
			for (i = 0; i < old_len; i++) {
				if (filename[i] != ' ') {
					g_strchomp (filename);
					ret |= (old_len != strlen (filename));
					break;
				}
			}

			return ret;
		}
	}

	return FALSE;
}

static GFile *
get_unique_target_file (GFile *src,
			GFile *dest_dir,
			gboolean same_fs,
			const char *dest_fs_type,
			int count)
{
	const char *editname, *end;
	char *basename, *new_name;
	GFileInfo *info;
	GFile *dest;
	int max_length;

	max_length = get_max_name_length (dest_dir);

	dest = NULL;
	info = g_file_query_info (src,
				  G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME,
				  0, NULL, NULL);
	if (info != NULL) {
		editname = g_file_info_get_attribute_string (info, G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME);

		if (editname != NULL) {
			new_name = get_duplicate_name (editname, count, max_length);
			make_file_name_valid_for_dest_fs (new_name, dest_fs_type);
			dest = g_file_get_child_for_display_name (dest_dir, new_name, NULL);
			g_free (new_name);
		}

		g_object_unref (info);
	}

	if (dest == NULL) {
		basename = g_file_get_basename (src);

		if (g_utf8_validate (basename, -1, NULL)) {
			new_name = get_duplicate_name (basename, count, max_length);
			make_file_name_valid_for_dest_fs (new_name, dest_fs_type);
			dest = g_file_get_child_for_display_name (dest_dir, new_name, NULL);
			g_free (new_name);
		}

		if (dest == NULL) {
			end = strrchr (basename, '.');
			if (end != NULL) {
				count += atoi (end + 1);
			}
			new_name = g_strdup_printf ("%s.%d", basename, count);
			make_file_name_valid_for_dest_fs (new_name, dest_fs_type);
			dest = g_file_get_child (dest_dir, new_name);
			g_free (new_name);
		}

		g_free (basename);
	}

	return dest;
}

static GFile *
get_target_file_for_link (GFile *src,
			  GFile *dest_dir,
			  const char *dest_fs_type,
			  int count)
{
	const char *editname;
	char *basename, *new_name;
	GFileInfo *info;
	GFile *dest;
	int max_length;

	max_length = get_max_name_length (dest_dir);

	dest = NULL;
	info = g_file_query_info (src,
				  G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME,
				  0, NULL, NULL);
	if (info != NULL) {
		editname = g_file_info_get_attribute_string (info, G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME);

		if (editname != NULL) {
			new_name = get_link_name (editname, count, max_length);
			make_file_name_valid_for_dest_fs (new_name, dest_fs_type);
			dest = g_file_get_child_for_display_name (dest_dir, new_name, NULL);
			g_free (new_name);
		}

		g_object_unref (info);
	}

	if (dest == NULL) {
		basename = g_file_get_basename (src);
		make_file_name_valid_for_dest_fs (basename, dest_fs_type);

		if (g_utf8_validate (basename, -1, NULL)) {
			new_name = get_link_name (basename, count, max_length);
			make_file_name_valid_for_dest_fs (new_name, dest_fs_type);
			dest = g_file_get_child_for_display_name (dest_dir, new_name, NULL);
			g_free (new_name);
		}

		if (dest == NULL) {
			if (count == 1) {
				new_name = g_strdup_printf ("%s.lnk", basename);
			} else {
				new_name = g_strdup_printf ("%s.lnk%d", basename, count);
			}
			make_file_name_valid_for_dest_fs (new_name, dest_fs_type);
			dest = g_file_get_child (dest_dir, new_name);
			g_free (new_name);
		}

		g_free (basename);
	}

	return dest;
}

static GFile *
get_target_file_with_custom_name (GFile *src,
				  GFile *dest_dir,
				  const char *dest_fs_type,
				  gboolean same_fs,
				  const gchar *custom_name)
{
	char *basename;
	GFile *dest;
	GFileInfo *info;
	char *copyname;

	dest = NULL;

	if (custom_name != NULL) {
		copyname = g_strdup (custom_name);
		make_file_name_valid_for_dest_fs (copyname, dest_fs_type);
		dest = g_file_get_child_for_display_name (dest_dir, copyname, NULL);

		g_free (copyname);
	}

	if (dest == NULL && !same_fs) {
		info = g_file_query_info (src,
					  G_FILE_ATTRIBUTE_STANDARD_COPY_NAME ","
					  G_FILE_ATTRIBUTE_TRASH_ORIG_PATH,
					  0, NULL, NULL);

		if (info) {
			copyname = NULL;

			/* if file is being restored from trash make sure it uses its original name */
			if (g_file_has_uri_scheme (src, "trash")) {
				copyname = g_strdup (g_file_info_get_attribute_byte_string (info, G_FILE_ATTRIBUTE_TRASH_ORIG_PATH));
			}

			if (copyname == NULL) {
				copyname = g_strdup (g_file_info_get_attribute_string (info, G_FILE_ATTRIBUTE_STANDARD_COPY_NAME));
			}

			if (copyname) {
				make_file_name_valid_for_dest_fs (copyname, dest_fs_type);
				dest = g_file_get_child_for_display_name (dest_dir, copyname, NULL);
				g_free (copyname);
			}

			g_object_unref (info);
		}
	}

	if (dest == NULL) {
		basename = g_file_get_basename (src);
		make_file_name_valid_for_dest_fs (basename, dest_fs_type);
		dest = g_file_get_child (dest_dir, basename);
		g_free (basename);
	}

	return dest;
}

static GFile *
get_target_file (GFile *src,
		 GFile *dest_dir,
		 const char *dest_fs_type,
		 gboolean same_fs)
{
	return get_target_file_with_custom_name (src, dest_dir, dest_fs_type, same_fs, NULL);
}

static gboolean
has_fs_id (GFile *file, const char *fs_id)
{
	const char *id;
	GFileInfo *info;
	gboolean res;

	res = FALSE;
	info = g_file_query_info (file,
				  G_FILE_ATTRIBUTE_ID_FILESYSTEM,
				  G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS,
				  NULL, NULL);

	if (info) {
		id = g_file_info_get_attribute_string (info, G_FILE_ATTRIBUTE_ID_FILESYSTEM);

		if (id && strcmp (id, fs_id) == 0) {
			res = TRUE;
		}

		g_object_unref (info);
	}

	return res;
}
static gboolean
test_dir_is_parent (GFile *child, GFile *root)
{
	GFile *f, *tmp;

	f = g_file_dup (child);
	while (f) {
		if (g_file_equal (f, root)) {
			g_object_unref (f);
			return TRUE;
		}
		tmp = f;
		f = g_file_get_parent (f);
		g_object_unref (tmp);
	}
	if (f) {
		g_object_unref (f);
	}
	return FALSE;
}


static
gboolean
nautilus_is_in_system_dir (GFile *file)
{
	const char * const * data_dirs;
	char *path;
	int i;
	gboolean res;

	if (!g_file_is_native (file)) {
		return FALSE;
	}

	path = g_file_get_path (file);

	res = FALSE;

	data_dirs = g_get_system_data_dirs ();
	for (i = 0; path != NULL && data_dirs[i] != NULL; i++) {
		if (g_str_has_prefix (path, data_dirs[i])) {
			res = TRUE;
			break;
		}

	}

	g_free (path);

	return res;
}

static gboolean
is_trusted_desktop_file (GFile *file,
			 GCancellable *cancellable)
{
	char *basename;
	gboolean res;
	GFileInfo *info;

	/* Don't trust non-local files */
	if (!g_file_is_native (file)) {
		return FALSE;
	}

	basename = g_file_get_basename (file);
	if (!g_str_has_suffix (basename, ".desktop")) {
		g_free (basename);
		return FALSE;
	}
	g_free (basename);

	info = g_file_query_info (file,
				  G_FILE_ATTRIBUTE_STANDARD_TYPE ","
				  G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE,
				  G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS,
				  cancellable,
				  NULL);

	if (info == NULL) {
		return FALSE;
	}

	res = FALSE;

	/* Weird file => not trusted,
	   Already executable => no need to mark trusted */
	if (g_file_info_get_file_type (info) == G_FILE_TYPE_REGULAR &&
	    !g_file_info_get_attribute_boolean (info,
						G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE) &&
	    nautilus_is_in_system_dir (file)) {
		res = TRUE;
	}
	g_object_unref (info);

	return res;
}

static char *
get_abs_path_for_symlink (GFile *file, GFile *destination)
{
	GFile *root, *parent;
	char *relative, *abs;

	if (g_file_is_native (file) || g_file_is_native (destination)) {
		return g_file_get_path (file);
	}

	root = g_object_ref (file);
	while ((parent = g_file_get_parent (root)) != NULL) {
		g_object_unref (root);
		root = parent;
	}

	relative = g_file_get_relative_path (root, file);
	g_object_unref (root);
	abs = g_strconcat ("/", relative, NULL);
	g_free (relative);
	return abs;
}
