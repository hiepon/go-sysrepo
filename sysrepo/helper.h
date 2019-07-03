// -*- coding: utf-8 -*-

#ifndef _GO_SYSREPO_HELPER_H
#define _GO_SYSREPO_HELPER_H

#include <sysrepo.h>

#ifdef __cplusplus
extern "C" {
#endif

  bool     _sr_data_bool_get(sr_data_t* p);
  void     _sr_data_bool_set(sr_data_t* p, bool v);
  double   _sr_data_decimal64_get(sr_data_t* p);
  void     _sr_data_decimal64_set(sr_data_t* p, double v);
  int8_t   _sr_data_int8_get(sr_data_t* p);
  void     _sr_data_int8_set(sr_data_t* p, int8_t v);
  int16_t  _sr_data_int16_get(sr_data_t* p);
  void     _sr_data_int16_set(sr_data_t* p, int16_t v);
  int32_t  _sr_data_int32_get(sr_data_t* p);
  void     _sr_data_int32_set(sr_data_t* p, int32_t v);
  int64_t  _sr_data_int64_get(sr_data_t* p);
  void     _sr_data_int64_set(sr_data_t* p, int64_t v);
  uint8_t  _sr_data_uint8_get(sr_data_t* p);
  void     _sr_data_uint8_set(sr_data_t* p, uint8_t v);
  uint16_t _sr_data_uint16_get(sr_data_t* p);
  void     _sr_data_uint16_set(sr_data_t* p, uint16_t v);
  uint32_t _sr_data_uint32_get(sr_data_t* p);
  void     _sr_data_uint32_set(sr_data_t* p, uint32_t v);
  uint64_t _sr_data_uint64_get(sr_data_t* p);
  void     _sr_data_uint64_set(sr_data_t* p, uint64_t v);

  sr_val_t* _sr_val_array_get(sr_val_t* val, size_t index);
  sr_node_t* _sr_node_array_get(sr_node_t* node, size_t index);
  sr_val_t * _sr_node_to_val(sr_node_t* node);

  int _sr_module_change_cb(sr_session_ctx_t* session,
			   const char* module_name,
			   sr_notif_event_t event,
			   void *private_ctx);

  int _sr_subtree_change_cb(sr_session_ctx_t* session,
			    const char* xpath,
			    sr_notif_event_t event,
			    void* private_ctx);

  void _sr_module_install_cb(const char* module_name,
			     const char* revision,
			     sr_module_state_t state,
			     void* private_ctx);

  void _sr_feature_enable_cb(const char* module_name,
			     const char* feature_name,
			     bool enabled,
			     void* private_ctx);

  int _sr_dp_get_items_cb(const char *xpath,
			  sr_val_t **values,
			  size_t *values_cnt,
			  void *private_ctx);

#ifdef __cplusplus
}
#endif

#endif // _GO_SYSREPO_HELPER_H
