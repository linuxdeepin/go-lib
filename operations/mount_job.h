// #ifndef __D_MOUNT_OPERATION_H__
// #define __D_MOUNT_OPERATION_H__
//
// #include <gio/giotypes.h>
//
// G_BEGIN_DECLS
//
// #define D_TYPE_MOUNT_OPERATION         (d_mount_operation_get_type ())
// #define D_MOUNT_OPERATION(o)           (G_TYPE_CHECK_INSTANCE_CAST ((o), D_TYPE_MOUNT_OPERATION, DMountOperation))
// #define D_MOUNT_OPERATION_CLASS(k)     (G_TYPE_CHECK_CLASS_CAST((k), D_TYPE_MOUNT_OPERATION, DMountOperationClass))
// #define D_IS_MOUNT_OPERATION(o)        (G_TYPE_CHECK_INSTANCE_TYPE ((o), D_TYPE_MOUNT_OPERATION))
// #define D_IS_MOUNT_OPERATION_CLASS(k)  (G_TYPE_CHECK_CLASS_TYPE ((k), D_TYPE_MOUNT_OPERATION))
// #define D_MOUNT_OPERATION_GET_CLASS(o) (G_TYPE_INSTANCE_GET_CLASS ((o), D_TYPE_MOUNT_OPERATION, DMountOperationClass))
//
// typedef struct _DMountOperation         DMountOperation;
// typedef struct _DMountOperationClass    DMountOperationClass;
// typedef struct _DMountOperationPrivate  DMountOperationPrivate;
//
// #<{(|*
//  * DMountOperation:
//  *
//  * This should not be accessed directly. Use the accessor functions below.
//  |)}>#
// struct _DMountOperation
// {
//     GMountOperation parent_instance;
//
//     DMountOperationPrivate *priv;
// };
//
// struct _DMountOperationClass
// {
//     GMountOperationClass parent_class;
//
//     #<{(| Padding for future expansion |)}>#
//     void (*_d_reserved1) (void);
//     void (*_d_reserved2) (void);
//     void (*_d_reserved3) (void);
//     void (*_d_reserved4) (void);
// };
//
//
// GDK_AVAILABLE_IN_ALL
// GType            d_mount_operation_get_type   (void);
// GDK_AVAILABLE_IN_ALL
// DMountOperation *d_mount_operation_new        ();
// GDK_AVAILABLE_IN_ALL
// gboolean         d_mount_operation_is_showing (DMountOperation *op);
// GDK_AVAILABLE_IN_ALL
// void             d_mount_operation_set_parent (DMountOperation *op,
//                                                GtkWindow         *parent);
// GDK_AVAILABLE_IN_ALL
// GtkWindow *      d_mount_operation_get_parent (DMountOperation *op);
// GDK_AVAILABLE_IN_ALL
// void             d_mount_operation_set_screen (DMountOperation *op,
//                                                GdkScreen         *screen);
// GDK_AVAILABLE_IN_ALL
// GdkScreen       *d_mount_operation_get_screen (DMountOperation *op);
//
// G_END_DECLS
//
// #endif #<{(| __D_MOUNT_OPERATION_H__ |)}>#
