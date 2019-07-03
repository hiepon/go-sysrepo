// -*- coding: utf-8 -*-

#include "helper.h"

bool _sr_data_bool_get(sr_data_t* p) {
  return p->bool_val;
}

void _sr_data_bool_set(sr_data_t* p, bool v) {
  p->bool_val = v;
}

double _sr_data_decimal64_get(sr_data_t* p) {
  return p->decimal64_val;
}

void _sr_data_decimal64_set(sr_data_t* p, double v) {
  p->decimal64_val = v;
}

int8_t _sr_data_int8_get(sr_data_t* p) {
  return p->int8_val;
}

void _sr_data_int8_set(sr_data_t* p, int8_t v) {
  p->int8_val = v;
}

int16_t _sr_data_int16_get(sr_data_t* p) {
  return p->int16_val;
}

void _sr_data_int16_set(sr_data_t* p, int16_t v) {
  p->int16_val = v;
}

int32_t _sr_data_int32_get(sr_data_t* p) {
  return p->int32_val;
}

void _sr_data_int32_set(sr_data_t* p, int32_t v) {
  p->int32_val = v;
}

int64_t _sr_data_int64_get(sr_data_t* p) {
  return p->int64_val;
}

void _sr_data_int64_set(sr_data_t* p, int64_t v) {
  p->int64_val = v;
}

uint8_t _sr_data_uint8_get(sr_data_t* p) {
  return p->uint8_val;
}

void  _sr_data_uint8_set(sr_data_t* p, uint8_t v) {
  p->uint8_val = v;
}

uint16_t _sr_data_uint16_get(sr_data_t* p) {
  return p->uint16_val;
}

void _sr_data_uint16_set(sr_data_t* p, uint16_t v) {
  p->uint16_val = v;
}

uint32_t _sr_data_uint32_get(sr_data_t* p) {
  return p->uint32_val;
}

void _sr_data_uint32_set(sr_data_t* p, uint32_t v) {
  p->uint32_val = v;
}

uint64_t _sr_data_uint64_get(sr_data_t* p) {
  return p->uint64_val;
}

void _sr_data_uint64_set(sr_data_t* p, uint64_t v) {
  p->uint64_val = v;
}

sr_val_t* _sr_val_array_get(sr_val_t* val, size_t index) {
  return val + index;
}

sr_node_t* _sr_node_array_get(sr_node_t* node, size_t index) {
  return node + index;
}

sr_val_t* _sr_node_to_val(sr_node_t* node) {
  return (sr_val_t*)(node);
}

extern int go_sr_module_change_cb(sr_session_ctx_t* session,
				   const char* module_name,
				   sr_notif_event_t event,
				   void *private_ctx);

int _sr_module_change_cb(sr_session_ctx_t* session,
			 const char* module_name,
			 sr_notif_event_t event,
			 void *private_ctx) {
  return go_sr_module_change_cb(session, module_name, event, private_ctx);
}

extern int go_sr_subtree_change_cb(sr_session_ctx_t* session,
				   const char* xpath,
				   sr_notif_event_t event,
				   void* private_ctx);

int _sr_subtree_change_cb(sr_session_ctx_t* session,
			  const char* xpath,
			  sr_notif_event_t event,
			  void* private_ctx) {
  return go_sr_subtree_change_cb(session, xpath, event, private_ctx);
}

extern void go_sr_module_install_cb(const char* module_name,
			   const char* revision,
			   sr_module_state_t state,
			   void* private_ctx);

void _sr_module_install_cb(const char* module_name,
			   const char* revision,
			   sr_module_state_t state,
			   void* private_ctx) {
  go_sr_module_install_cb(module_name, revision, state, private_ctx);
}
  

extern void go_sr_feature_enable_cb(const char* module_name,
				    const char* feature_name,
				    bool enabled,
				    void* private_ctx);

void _sr_feature_enable_cb(const char* module_name,
			   const char* feature_name,
			   bool enabled,
			   void* private_ctx) {
  go_sr_feature_enable_cb(module_name, feature_name, enabled, private_ctx);
}

extern int go_sr_dp_get_items_cb(const char *xpath,
			     sr_val_t **values,
			     size_t *values_cnt,
			     void *private_ctx);

int _sr_dp_get_items_cb(const char *xpath,
			sr_val_t **values,
			size_t *values_cnt,
			void *private_ctx) {
  return go_sr_dp_get_items_cb(xpath, values, values_cnt, private_ctx);
}
