/* #include "mount_job.h" */
/*  */
/* static void   d_mount_operation_finalize     (GObject          *object); */
/* static void   d_mount_operation_set_property (GObject          *object, */
/*                                               guint             prop_id, */
/*                                               const GValue     *value, */
/*                                               GParamSpec       *pspec); */
/* static void   d_mount_operation_get_property (GObject          *object, */
/*                                               guint             prop_id, */
/*                                               GValue           *value, */
/*                                               GParamSpec       *pspec); */
/*  */
/* static void   d_mount_operation_ask_password (GMountOperation *op, */
/*                                               const char      *message, */
/*                                               const char      *default_user, */
/*                                               const char      *default_domain, */
/*                                               GAskPasswordFlags flags); */
/*  */
/* static void   d_mount_operation_ask_question (GMountOperation *op, */
/*                                               const char      *message, */
/*                                               const char      *choices[]); */
/*  */
/* static void   d_mount_operation_show_processes (GMountOperation *op, */
/*                                                 const char      *message, */
/*                                                 GArray          *processes, */
/*                                                 const char      *choices[]); */
/*  */
/* static void   d_mount_operation_aborted      (GMountOperation *op); */
/*  */
/* struct _DMountOperationPrivate { */
/*     GtkWindow *parent_window; */
/*     GtkDialog *dialog; */
/*     GdkScreen *screen; */
/*  */
/*     #<{(| bus proxy |)}># */
/*     _GtkMountOperationHandler *handler; */
/*     GCancellable *cancellable; */
/*     gboolean handler_showing; */
/*  */
/*     #<{(| for the ask-password dialog |)}># */
/*     GtkWidget *grid; */
/*     GtkWidget *username_entry; */
/*     GtkWidget *domain_entry; */
/*     GtkWidget *password_entry; */
/*     GtkWidget *anonymous_toggle; */
/*     GList *user_widgets; */
/*  */
/*     GAskPasswordFlags ask_flags; */
/*     GPasswordSave     password_save; */
/*     gboolean          anonymous; */
/*  */
/*     #<{(| for the show-processes dialog |)}># */
/*     GtkWidget *process_tree_view; */
/*     GtkListStore *process_list_store; */
/* }; */
