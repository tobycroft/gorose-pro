package gorose

import (
	"reflect"
)

// ========================
// 版本1：返回 []map[string]interface{}
// 支持：struct/*struct/[]struct/[]*struct
// ========================
func StructToMapSlices(data interface{}) []map[string]interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)
	return autoParse(v)
}

// ========================
// 版本2：返回 map[string]interface{}
// 支持：struct/*struct
// 就是原版
// ========================
func StructToMapV2(data interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)

	// 解指针
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	return parseStruct(v)
}

// 自动解析
func autoParse(v reflect.Value) []map[string]interface{} {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	var res []map[string]interface{}

	switch v.Kind() {
	case reflect.Struct:
		res = append(res, parseStruct(v))

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			for item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			if item.Kind() == reflect.Struct {
				res = append(res, parseStruct(item))
			}
		}
	}

	return res
}

// 解析结构体核心方法
func parseStruct(val reflect.Value) map[string]interface{} {
	m := make(map[string]interface{})
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fv := val.Field(i)

		// 跳过私有字段
		if field.PkgPath != "" {
			continue
		}

		// 跳过嵌套结构体
		if fv.Kind() == reflect.Struct {
			continue
		}

		// 读取 gorose 标签
		tag := field.Tag.Get(TAGNAME)

		// 忽略标记
		if tag == "ignore" {
			continue
		}

		// ======================
		// ✅ 关键：无标签 → 使用结构体原生字段名
		// ======================
		if tag == "" {
			tag = field.Name // 这里就是自动使用原生名字
		}

		// 保留所有值：0 / "" / false 都不会丢
		m[tag] = fv.Interface()
	}

	return m
}
