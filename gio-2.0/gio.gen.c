#include "gio.gen.h"

static void _c_callback_cleanup(void *userdata)
{
	_Gio_go_callback_cleanup(userdata);
}

extern void _GBaseFinalizeFunc_c_wrapper();
extern void _GBaseFinalizeFunc_c_wrapper_once();
extern void _GBaseInitFunc_c_wrapper();
extern void _GBaseInitFunc_c_wrapper_once();
extern void _GBindingTransformFunc_c_wrapper();
extern void _GBindingTransformFunc_c_wrapper_once();
extern void _GBoxedFreeFunc_c_wrapper();
extern void _GBoxedFreeFunc_c_wrapper_once();
extern void _GCallback_c_wrapper();
extern void _GCallback_c_wrapper_once();
extern void _GClassFinalizeFunc_c_wrapper();
extern void _GClassFinalizeFunc_c_wrapper_once();
extern void _GClassInitFunc_c_wrapper();
extern void _GClassInitFunc_c_wrapper_once();
extern void _GClosureMarshal_c_wrapper();
extern void _GClosureMarshal_c_wrapper_once();
extern void _GClosureNotify_c_wrapper();
extern void _GClosureNotify_c_wrapper_once();
extern void _GInstanceInitFunc_c_wrapper();
extern void _GInstanceInitFunc_c_wrapper_once();
extern void _GInterfaceFinalizeFunc_c_wrapper();
extern void _GInterfaceFinalizeFunc_c_wrapper_once();
extern void _GInterfaceInitFunc_c_wrapper();
extern void _GInterfaceInitFunc_c_wrapper_once();
extern void _GObjectFinalizeFunc_c_wrapper();
extern void _GObjectFinalizeFunc_c_wrapper_once();
extern void _GObjectGetPropertyFunc_c_wrapper();
extern void _GObjectGetPropertyFunc_c_wrapper_once();
extern void _GObjectSetPropertyFunc_c_wrapper();
extern void _GObjectSetPropertyFunc_c_wrapper_once();
extern void _GSignalAccumulator_c_wrapper();
extern void _GSignalAccumulator_c_wrapper_once();
extern void _GSignalEmissionHook_c_wrapper();
extern void _GSignalEmissionHook_c_wrapper_once();
extern void _GToggleNotify_c_wrapper();
extern void _GToggleNotify_c_wrapper_once();
extern void _GTypeClassCacheFunc_c_wrapper();
extern void _GTypeClassCacheFunc_c_wrapper_once();
extern void _GTypeInterfaceCheckFunc_c_wrapper();
extern void _GTypeInterfaceCheckFunc_c_wrapper_once();
extern void _GTypePluginCompleteInterfaceInfo_c_wrapper();
extern void _GTypePluginCompleteInterfaceInfo_c_wrapper_once();
extern void _GTypePluginCompleteTypeInfo_c_wrapper();
extern void _GTypePluginCompleteTypeInfo_c_wrapper_once();
extern void _GTypePluginUnuse_c_wrapper();
extern void _GTypePluginUnuse_c_wrapper_once();
extern void _GTypePluginUse_c_wrapper();
extern void _GTypePluginUse_c_wrapper_once();
extern void _GValueTransform_c_wrapper();
extern void _GValueTransform_c_wrapper_once();
extern void _GWeakNotify_c_wrapper();
extern void _GWeakNotify_c_wrapper_once();
extern void _GChildWatchFunc_c_wrapper();
extern void _GChildWatchFunc_c_wrapper_once();
extern void _GCompareDataFunc_c_wrapper();
extern void _GCompareDataFunc_c_wrapper_once();
extern void _GCompareFunc_c_wrapper();
extern void _GCompareFunc_c_wrapper_once();
extern void _GDataForeachFunc_c_wrapper();
extern void _GDataForeachFunc_c_wrapper_once();
extern void _GDestroyNotify_c_wrapper();
extern void _GDestroyNotify_c_wrapper_once();
extern void _GEqualFunc_c_wrapper();
extern void _GEqualFunc_c_wrapper_once();
extern void _GFreeFunc_c_wrapper();
extern void _GFreeFunc_c_wrapper_once();
extern void _GFunc_c_wrapper();
extern void _GFunc_c_wrapper_once();
extern void _GHFunc_c_wrapper();
extern void _GHFunc_c_wrapper_once();
extern void _GHRFunc_c_wrapper();
extern void _GHRFunc_c_wrapper_once();
extern void _GHashFunc_c_wrapper();
extern void _GHashFunc_c_wrapper_once();
extern void _GHookCheckFunc_c_wrapper();
extern void _GHookCheckFunc_c_wrapper_once();
extern void _GHookCheckMarshaller_c_wrapper();
extern void _GHookCheckMarshaller_c_wrapper_once();
extern void _GHookCompareFunc_c_wrapper();
extern void _GHookCompareFunc_c_wrapper_once();
extern void _GHookFinalizeFunc_c_wrapper();
extern void _GHookFinalizeFunc_c_wrapper_once();
extern void _GHookFindFunc_c_wrapper();
extern void _GHookFindFunc_c_wrapper_once();
extern void _GHookFunc_c_wrapper();
extern void _GHookFunc_c_wrapper_once();
extern void _GHookMarshaller_c_wrapper();
extern void _GHookMarshaller_c_wrapper_once();
extern void _GIOFunc_c_wrapper();
extern void _GIOFunc_c_wrapper_once();
extern void _GLogFunc_c_wrapper();
extern void _GLogFunc_c_wrapper_once();
extern void _GNodeForeachFunc_c_wrapper();
extern void _GNodeForeachFunc_c_wrapper_once();
extern void _GNodeTraverseFunc_c_wrapper();
extern void _GNodeTraverseFunc_c_wrapper_once();
extern void _GOptionArgFunc_c_wrapper();
extern void _GOptionArgFunc_c_wrapper_once();
extern void _GOptionErrorFunc_c_wrapper();
extern void _GOptionErrorFunc_c_wrapper_once();
extern void _GOptionParseFunc_c_wrapper();
extern void _GOptionParseFunc_c_wrapper_once();
extern void _GPollFunc_c_wrapper();
extern void _GPollFunc_c_wrapper_once();
extern void _GPrintFunc_c_wrapper();
extern void _GPrintFunc_c_wrapper_once();
extern void _GRegexEvalCallback_c_wrapper();
extern void _GRegexEvalCallback_c_wrapper_once();
extern void _GScannerMsgFunc_c_wrapper();
extern void _GScannerMsgFunc_c_wrapper_once();
extern void _GSequenceIterCompareFunc_c_wrapper();
extern void _GSequenceIterCompareFunc_c_wrapper_once();
extern void _GSourceDummyMarshal_c_wrapper();
extern void _GSourceDummyMarshal_c_wrapper_once();
extern void _GSourceFunc_c_wrapper();
extern void _GSourceFunc_c_wrapper_once();
extern void _GSpawnChildSetupFunc_c_wrapper();
extern void _GSpawnChildSetupFunc_c_wrapper_once();
extern void _GTestDataFunc_c_wrapper();
extern void _GTestDataFunc_c_wrapper_once();
extern void _GTestFixtureFunc_c_wrapper();
extern void _GTestFixtureFunc_c_wrapper_once();
extern void _GTestFunc_c_wrapper();
extern void _GTestFunc_c_wrapper_once();
extern void _GTestLogFatalFunc_c_wrapper();
extern void _GTestLogFatalFunc_c_wrapper_once();
extern void _GTranslateFunc_c_wrapper();
extern void _GTranslateFunc_c_wrapper_once();
extern void _GTraverseFunc_c_wrapper();
extern void _GTraverseFunc_c_wrapper_once();
extern void _GUnixFDSourceFunc_c_wrapper();
extern void _GUnixFDSourceFunc_c_wrapper_once();
extern void _GVoidFunc_c_wrapper();
extern void _GVoidFunc_c_wrapper_once();
extern void _GAsyncReadyCallback_c_wrapper();
extern void _GAsyncReadyCallback_c_wrapper_once();
extern void _GBusAcquiredCallback_c_wrapper();
extern void _GBusAcquiredCallback_c_wrapper_once();
extern void _GBusNameAcquiredCallback_c_wrapper();
extern void _GBusNameAcquiredCallback_c_wrapper_once();
extern void _GBusNameAppearedCallback_c_wrapper();
extern void _GBusNameAppearedCallback_c_wrapper_once();
extern void _GBusNameLostCallback_c_wrapper();
extern void _GBusNameLostCallback_c_wrapper_once();
extern void _GBusNameVanishedCallback_c_wrapper();
extern void _GBusNameVanishedCallback_c_wrapper_once();
extern void _GCancellableSourceFunc_c_wrapper();
extern void _GCancellableSourceFunc_c_wrapper_once();
extern void _GDBusInterfaceGetPropertyFunc_c_wrapper();
extern void _GDBusInterfaceGetPropertyFunc_c_wrapper_once();
extern void _GDBusInterfaceMethodCallFunc_c_wrapper();
extern void _GDBusInterfaceMethodCallFunc_c_wrapper_once();
extern void _GDBusInterfaceSetPropertyFunc_c_wrapper();
extern void _GDBusInterfaceSetPropertyFunc_c_wrapper_once();
extern void _GDBusMessageFilterFunction_c_wrapper();
extern void _GDBusMessageFilterFunction_c_wrapper_once();
extern void _GDBusProxyTypeFunc_c_wrapper();
extern void _GDBusProxyTypeFunc_c_wrapper_once();
extern void _GDBusSignalCallback_c_wrapper();
extern void _GDBusSignalCallback_c_wrapper_once();
extern void _GDBusSubtreeDispatchFunc_c_wrapper();
extern void _GDBusSubtreeDispatchFunc_c_wrapper_once();
extern void _GDBusSubtreeIntrospectFunc_c_wrapper();
extern void _GDBusSubtreeIntrospectFunc_c_wrapper_once();
extern void _GDesktopAppLaunchCallback_c_wrapper();
extern void _GDesktopAppLaunchCallback_c_wrapper_once();
extern void _GFileMeasureProgressCallback_c_wrapper();
extern void _GFileMeasureProgressCallback_c_wrapper_once();
extern void _GFileProgressCallback_c_wrapper();
extern void _GFileProgressCallback_c_wrapper_once();
extern void _GFileReadMoreCallback_c_wrapper();
extern void _GFileReadMoreCallback_c_wrapper_once();
extern void _GIOSchedulerJobFunc_c_wrapper();
extern void _GIOSchedulerJobFunc_c_wrapper_once();
extern void _GPollableSourceFunc_c_wrapper();
extern void _GPollableSourceFunc_c_wrapper_once();
extern void _GSettingsBindGetMapping_c_wrapper();
extern void _GSettingsBindGetMapping_c_wrapper_once();
extern void _GSettingsBindSetMapping_c_wrapper();
extern void _GSettingsBindSetMapping_c_wrapper_once();
extern void _GSettingsGetMapping_c_wrapper();
extern void _GSettingsGetMapping_c_wrapper_once();
extern void _GSimpleAsyncThreadFunc_c_wrapper();
extern void _GSimpleAsyncThreadFunc_c_wrapper_once();
extern void _GSocketSourceFunc_c_wrapper();
extern void _GSocketSourceFunc_c_wrapper_once();
extern void _GTaskThreadFunc_c_wrapper();
extern void _GTaskThreadFunc_c_wrapper_once();
void _g_drive_eject(GDrive* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_drive_eject(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_eject(this, arg0, arg1, 0, 0);
	}
}
void _g_drive_eject_with_operation(GDrive* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_drive_eject_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_eject_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_drive_poll_for_media(GDrive* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_drive_poll_for_media(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_poll_for_media(this, arg0, 0, 0);
	}
}
void _g_drive_start(GDrive* this, GDriveStartFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_drive_start(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_start(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_drive_stop(GDrive* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_drive_stop(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_stop(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_append_to_async(GFile* this, GFileCreateFlags arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_append_to_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_append_to_async(this, arg0, arg1, arg2, 0, 0);
	}
}
int _g_file_copy(GFile* this, GFile* arg0, GFileCopyFlags arg1, GCancellable* arg2, void* gofunc, GError** arg5) {
	if (gofunc) {
		return g_file_copy(this, arg0, arg1, arg2, _GFileProgressCallback_c_wrapper, gofunc, arg5);
	} else {
		return g_file_copy(this, arg0, arg1, arg2, 0, 0, arg5);
	}
}
void _g_file_create_async(GFile* this, GFileCreateFlags arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_create_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_create_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_create_readwrite_async(GFile* this, GFileCreateFlags arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_create_readwrite_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_create_readwrite_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_delete_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_delete_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_delete_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_eject_mountable(GFile* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_eject_mountable(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_eject_mountable(this, arg0, arg1, 0, 0);
	}
}
void _g_file_eject_mountable_with_operation(GFile* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_eject_mountable_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_eject_mountable_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_enumerate_children_async(GFile* this, char* arg0, GFileQueryInfoFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_file_enumerate_children_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_enumerate_children_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_file_find_enclosing_mount_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_find_enclosing_mount_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_find_enclosing_mount_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_load_contents_async(GFile* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_file_load_contents_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_load_contents_async(this, arg0, 0, 0);
	}
}
void _g_file_make_directory_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_make_directory_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_make_directory_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_mount_enclosing_volume(GFile* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_mount_enclosing_volume(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_mount_enclosing_volume(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_mount_mountable(GFile* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_mount_mountable(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_mount_mountable(this, arg0, arg1, arg2, 0, 0);
	}
}
int _g_file_move(GFile* this, GFile* arg0, GFileCopyFlags arg1, GCancellable* arg2, void* gofunc, GError** arg5) {
	if (gofunc) {
		return g_file_move(this, arg0, arg1, arg2, _GFileProgressCallback_c_wrapper, gofunc, arg5);
	} else {
		return g_file_move(this, arg0, arg1, arg2, 0, 0, arg5);
	}
}
void _g_file_open_readwrite_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_open_readwrite_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_open_readwrite_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_poll_mountable(GFile* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_file_poll_mountable(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_poll_mountable(this, arg0, 0, 0);
	}
}
void _g_file_query_filesystem_info_async(GFile* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_query_filesystem_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_query_filesystem_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_query_info_async(GFile* this, char* arg0, GFileQueryInfoFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_file_query_info_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_query_info_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_file_read_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_read_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_read_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_replace_async(GFile* this, char* arg0, int arg1, GFileCreateFlags arg2, int32_t arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_file_replace_async(this, arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_replace_async(this, arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
void _g_file_replace_contents_async(GFile* this, uint8_t* arg0, uint64_t arg1, char* arg2, int arg3, GFileCreateFlags arg4, GCancellable* arg5, void* gofunc) {
	if (gofunc) {
		g_file_replace_contents_async(this, arg0, arg1, arg2, arg3, arg4, arg5, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_replace_contents_async(this, arg0, arg1, arg2, arg3, arg4, arg5, 0, 0);
	}
}
void _g_file_replace_readwrite_async(GFile* this, char* arg0, int arg1, GFileCreateFlags arg2, int32_t arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_file_replace_readwrite_async(this, arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_replace_readwrite_async(this, arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
void _g_file_set_attributes_async(GFile* this, GFileInfo* arg0, GFileQueryInfoFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_file_set_attributes_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_set_attributes_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_file_set_display_name_async(GFile* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_set_display_name_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_set_display_name_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_start_mountable(GFile* this, GDriveStartFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_start_mountable(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_start_mountable(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_stop_mountable(GFile* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_stop_mountable(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_stop_mountable(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_trash_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_trash_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_trash_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_unmount_mountable(GFile* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_unmount_mountable(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_unmount_mountable(this, arg0, arg1, 0, 0);
	}
}
void _g_file_unmount_mountable_with_operation(GFile* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_unmount_mountable_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_unmount_mountable_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_enumerator_close_async(GFileEnumerator* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_enumerator_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_enumerator_close_async(this, arg0, arg1, 0, 0);
	}
}
void _g_file_enumerator_next_files_async(GFileEnumerator* this, int32_t arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_enumerator_next_files_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_enumerator_next_files_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_io_stream_query_info_async(GFileIOStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_io_stream_query_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_io_stream_query_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_input_stream_query_info_async(GFileInputStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_input_stream_query_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_input_stream_query_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_file_output_stream_query_info_async(GFileOutputStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_output_stream_query_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_output_stream_query_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_io_stream_close_async(GIOStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_io_stream_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_io_stream_close_async(this, arg0, arg1, 0, 0);
	}
}
void _g_io_stream_splice_async(GIOStream* this, GIOStream* arg0, GIOStreamSpliceFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_io_stream_splice_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_io_stream_splice_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_input_stream_close_async(GInputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_input_stream_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_close_async(this, arg0, arg1, 0, 0);
	}
}
void _g_input_stream_read_async(GInputStream* this, uint8_t* arg0, uint64_t arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_input_stream_read_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_read_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_input_stream_read_bytes_async(GInputStream* this, uint64_t arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_input_stream_read_bytes_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_read_bytes_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_input_stream_skip_async(GInputStream* this, uint64_t arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_input_stream_skip_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_skip_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_mount_eject(GMount* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_mount_eject(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_eject(this, arg0, arg1, 0, 0);
	}
}
void _g_mount_eject_with_operation(GMount* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_mount_eject_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_eject_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_mount_guess_content_type(GMount* this, int arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_mount_guess_content_type(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_guess_content_type(this, arg0, arg1, 0, 0);
	}
}
void _g_mount_remount(GMount* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_mount_remount(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_remount(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_mount_unmount(GMount* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_mount_unmount(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_unmount(this, arg0, arg1, 0, 0);
	}
}
void _g_mount_unmount_with_operation(GMount* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_mount_unmount_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_unmount_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_output_stream_close_async(GOutputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_output_stream_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_close_async(this, arg0, arg1, 0, 0);
	}
}
void _g_output_stream_flush_async(GOutputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_output_stream_flush_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_flush_async(this, arg0, arg1, 0, 0);
	}
}
void _g_output_stream_splice_async(GOutputStream* this, GInputStream* arg0, GOutputStreamSpliceFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_output_stream_splice_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_splice_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_output_stream_write_async(GOutputStream* this, uint8_t* arg0, uint64_t arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_output_stream_write_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_write_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
void _g_output_stream_write_bytes_async(GOutputStream* this, GBytes* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_output_stream_write_bytes_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_write_bytes_async(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_permission_acquire_async(GPermission* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_permission_acquire_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_permission_acquire_async(this, arg0, 0, 0);
	}
}
void _g_permission_release_async(GPermission* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_permission_release_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_permission_release_async(this, arg0, 0, 0);
	}
}
void _g_volume_eject(GVolume* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_volume_eject(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_volume_eject(this, arg0, arg1, 0, 0);
	}
}
void _g_volume_eject_with_operation(GVolume* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_volume_eject_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_volume_eject_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
void _g_volume_mount(GVolume* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_volume_mount(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_volume_mount(this, arg0, arg1, arg2, 0, 0);
	}
}
