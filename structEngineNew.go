package gorose

import (
	"reflect"
)

// ========================
// 【原有版本1】切片版
// ========================
func StructToMapSlices(data interface{}) []map[string]interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)
	return autoParse(v, parseStructWithDefault)
}

// ========================
// 【原有版本2】单个结构体
// ========================
func StructToMapV2(data interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)

	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	return parseStructWithDefault(v)
}

// ========================
// 【新增：严格模式 - 只有带tag才解析，无tag跳过】
// ========================
func StructToMapStrict(data interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)

	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	return parseStructStrict(v)
}

// ========================
// 【新增：严格模式切片版】
// ========================
func StructToMapStrictSlices(data interface{}) []map[string]interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)
	return autoParse(v, parseStructStrict)
}

// ------------------------------
// 内部通用解析
// ------------------------------
func autoParse(v reflect.Value, parseFunc func(val reflect.Value) map[string]interface{}) []map[string]interface{} {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	var res []map[string]interface{}

	switch v.Kind() {
	case reflect.Struct:
		res = append(res, parseFunc(v))

	case reflect.Slice, reflect.Array:
		// 🔥 这里修复了：val.Len() → v.Len()
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			for item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			if item.Kind() == reflect.Struct {
				res = append(res, parseFunc(item))
			}
		}
	default:
	}

	return res
}

// ------------------------------
// 普通模式：无tag用字段名
// ------------------------------
func parseStructWithDefault(val reflect.Value) map[string]interface{} {
	m := make(map[string]interface{})
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fv := val.Field(i)

		if field.PkgPath != "" {
			continue
		}
		if fv.Kind() == reflect.Struct {
			continue
		}

		tag := field.Tag.Get(TAGNAME)
		if tag == "ignore" {
			continue
		}
		if tag == "" {
			tag = field.Name
		}

		m[tag] = fv.Interface()
	}

	return m
}

// ------------------------------
// 严格模式：无tag 直接跳过不解析
// ------------------------------
func parseStructStrict(val reflect.Value) map[string]interface{} {
	m := make(map[string]interface{})
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fv := val.Field(i)

		if field.PkgPath != "" {
			continue
		}
		if fv.Kind() == reflect.Struct {
			continue
		}

		tag := field.Tag.Get(TAGNAME)

		// 🔥 无tag 或 ignore → 跳过
		if tag == "" || tag == "ignore" {
			continue
		}

		m[tag] = fv.Interface()
	}

	return m
}
