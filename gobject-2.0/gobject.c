#include <glib-object.h>
#include <stdint.h>
#include <string.h>
#include "gobject.h"

struct _GGoClosure {
	GClosure closure;
	void *func[2];
	void *recv[2];
};

GType _g_param_spec_type(GParamSpec *pspec) {
	return G_PARAM_SPEC_TYPE(pspec);
}

GType _g_param_spec_value_type(GParamSpec *pspec) {
	return G_PARAM_SPEC_VALUE_TYPE(pspec);
}

GType _g_object_type(GObject *obj) {
	return G_OBJECT_TYPE(obj);
}

GType _g_value_type(GValue *val) {
	return G_VALUE_TYPE(val);
}

GType _g_type_interface()		{ return G_TYPE_INTERFACE; }
GType _g_type_char()			{ return G_TYPE_CHAR; }
GType _g_type_uchar()			{ return G_TYPE_UCHAR; }
GType _g_type_boolean()			{ return G_TYPE_BOOLEAN; }
GType _g_type_int()			{ return G_TYPE_INT; }
GType _g_type_uint()			{ return G_TYPE_UINT; }
GType _g_type_long()			{ return G_TYPE_LONG; }
GType _g_type_ulong()			{ return G_TYPE_ULONG; }
GType _g_type_int64()			{ return G_TYPE_INT64; }
GType _g_type_uint64()			{ return G_TYPE_UINT64; }
GType _g_type_enum()			{ return G_TYPE_ENUM; }
GType _g_type_flags()			{ return G_TYPE_FLAGS; }
GType _g_type_float()			{ return G_TYPE_FLOAT; }
GType _g_type_double()			{ return G_TYPE_DOUBLE; }
GType _g_type_string()			{ return G_TYPE_STRING; }
GType _g_type_pointer()			{ return G_TYPE_POINTER; }
GType _g_type_boxed()			{ return G_TYPE_BOXED; }
GType _g_type_param()			{ return G_TYPE_PARAM; }
GType _g_type_object()			{ return G_TYPE_OBJECT; }
GType _g_type_gtype()			{ return G_TYPE_GTYPE; }
GType _g_type_variant()			{ return G_TYPE_VARIANT; }
GType _g_type_param_boolean()		{ return G_TYPE_PARAM_BOOLEAN; }
GType _g_type_param_boxed()		{ return G_TYPE_PARAM_BOXED; }
GType _g_type_param_char()		{ return G_TYPE_PARAM_CHAR; }
GType _g_type_param_double()		{ return G_TYPE_PARAM_DOUBLE; }
GType _g_type_param_enum()		{ return G_TYPE_PARAM_ENUM; }
GType _g_type_param_flags()		{ return G_TYPE_PARAM_FLAGS; }
GType _g_type_param_float()		{ return G_TYPE_PARAM_FLOAT; }
GType _g_type_param_gtype()		{ return G_TYPE_PARAM_GTYPE; }
GType _g_type_param_int()		{ return G_TYPE_PARAM_INT; }
GType _g_type_param_int64()		{ return G_TYPE_PARAM_INT64; }
GType _g_type_param_long()		{ return G_TYPE_PARAM_LONG; }
GType _g_type_param_object()		{ return G_TYPE_PARAM_OBJECT; }
GType _g_type_param_override()		{ return G_TYPE_PARAM_OVERRIDE; }
GType _g_type_param_param()		{ return G_TYPE_PARAM_PARAM; }
GType _g_type_param_pointer()		{ return G_TYPE_PARAM_POINTER; }
GType _g_type_param_string()		{ return G_TYPE_PARAM_STRING; }
GType _g_type_param_uchar()		{ return G_TYPE_PARAM_UCHAR; }
GType _g_type_param_uint()		{ return G_TYPE_PARAM_UINT; }
GType _g_type_param_uint64()		{ return G_TYPE_PARAM_UINT64; }
GType _g_type_param_ulong()		{ return G_TYPE_PARAM_ULONG; }
GType _g_type_param_unichar()		{ return G_TYPE_PARAM_UNICHAR; }
GType _g_type_param_value_array()	{ return G_TYPE_PARAM_VALUE_ARRAY; }
GType _g_type_param_variant()		{ return G_TYPE_PARAM_VARIANT; }

GParamSpec *_g_object_find_property(GObject *object, const char *name)
{
	GObjectClass *cls = G_OBJECT_GET_CLASS(object);
	return g_object_class_find_property(cls, name);
}

extern void *g_go_interface_copy_go(void *boxed);
extern void g_go_interface_free_go(void *boxed);

GType _g_type_go_interface()
{
	static GType go_interface_type = G_TYPE_NONE;
	if (go_interface_type == G_TYPE_NONE) {
		go_interface_type = g_boxed_type_register_static("gointerface",
								 g_go_interface_copy_go,
								 g_go_interface_free_go);
	}
	return go_interface_type;
}

//-----------------------------------------------------------------------------

extern void g_goclosure_marshal_go(GGoClosure*, GValue*, int32_t, GValue*);
extern void g_goclosure_finalize_go(GGoClosure*);

static void g_goclosure_finalize(void *notify_data, GClosure *closure)
{
	GGoClosure *goclosure = (GGoClosure*)closure;
	g_goclosure_finalize_go(goclosure);
}

static void g_goclosure_marshal(GClosure *closure, GValue *return_value,
				uint32_t n_param_values, const GValue *param_values,
				void *invocation_hint, void *data)
{
	g_goclosure_marshal_go((GGoClosure*)closure,
			       return_value,
			       n_param_values,
			       (GValue*)param_values);
}

GGoClosure *g_goclosure_new(void *func, void *recv)
{
	GClosure *closure;
	GGoClosure *goclosure;

	closure = g_closure_new_simple(sizeof(GGoClosure), 0);
	goclosure = (GGoClosure*)closure;
	memset(goclosure->func, 0, sizeof(void*)*2);
	memset(goclosure->recv, 0, sizeof(void*)*2);
	if (func)
		memcpy(goclosure->func, func, sizeof(void*)*2);
	if (recv)
		memcpy(goclosure->recv, recv, sizeof(void*)*2);

	g_closure_add_finalize_notifier(closure, 0, g_goclosure_finalize);
	g_closure_set_marshal(closure, g_goclosure_marshal);

	return goclosure;
}

void *g_goclosure_get_func(GGoClosure *clo) {
	return clo->func;
}

void *g_goclosure_get_recv(GGoClosure *clo) {
	return clo->recv;
}


