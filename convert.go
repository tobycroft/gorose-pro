package gorose

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ============================================================
// 本地类型转换工具 (替代外部依赖 github.com/gohouse/t)
//
// 与 t 库保持相同的类型与方法签名, 但针对 ORM 热路径做了直连类型断言优化:
//   - 基础类型转换 (String/Int64/Float64/Bool) 直接 type switch, 无反射
//   - 容器转换 (Slice/MapStringT/MapStringInterface 等) 优先命中常见类型,
//     避免原库先经 Map()/Slice() 反射包装、再二次解包的多趟开销
// ============================================================

// Type 超级类型接口
type Type interface {
	Interface() interface{}
	String() string
	Int() int
	Int64() int64
	Float64() float64
	Bool() bool
	Slice() Slice
	MapStringT() MapStringT
	MapInterfaceT() MapInterfaceT
	MapStringInterface() map[string]interface{}
	SliceMapStringInterface() []map[string]interface{}
}

// Slice 超级类型切片
type Slice []Type

// MapStringT 超级类型 map (string key)
type MapStringT map[string]Type

// MapInterfaceT 超级类型 map (interface key)
type MapInterfaceT map[interface{}]Type

// TypeContext 值包装器
type TypeContext struct {
	val interface{}
}

var _ Type = (*TypeContext)(nil)

// New 初始化 Type 对象
func New(o interface{}) TypeContext {
	switch v := o.(type) {
	case Type:
		return TypeContext{val: v.Interface()}
	default:
		return TypeContext{val: o}
	}
}

// Interface 返回原始值
func (tc TypeContext) Interface() interface{} { return tc.val }

// String 任意对象转字符串
func (tc TypeContext) String() string { return parseString(tc.val) }

// Int 转 int
func (tc TypeContext) Int() int { return int(parseInt64(tc.val)) }

// Int64 转 int64
func (tc TypeContext) Int64() int64 { return parseInt64(tc.val) }

// Float64 转 float64
func (tc TypeContext) Float64() float64 { return parseFloat64(tc.val) }

// Bool 转 bool
func (tc TypeContext) Bool() bool { return parseBool(tc.val) }

// Slice 转切片
func (tc TypeContext) Slice() Slice { return toSlice(tc.val) }

// MapStringT 转 map[string]Type
func (tc TypeContext) MapStringT() MapStringT { return toMapStringT(tc.val) }

// MapInterfaceT 转 map[interface{}]Type
func (tc TypeContext) MapInterfaceT() MapInterfaceT { return toMapInterfaceT(tc.val) }

// MapStringInterface 转 map[string]interface{}
func (tc TypeContext) MapStringInterface() map[string]interface{} {
	return toMapStringInterface(tc.val)
}

// SliceMapStringInterface 转 []map[string]interface{}
func (tc TypeContext) SliceMapStringInterface() []map[string]interface{} {
	return toSliceMapStringInterface(tc.val)
}

// ============================================================
// 基础类型转换
// ============================================================

func parseString(o interface{}) string {
	if o == nil {
		return ""
	}
	switch v := o.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case *time.Time:
		if v == nil {
			return ""
		}
		return v.String()
	}

	rv := reflect.ValueOf(o)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return ""
		}
		return parseString(rv.Elem().Interface())
	}
	switch rv.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		if rv.IsNil() {
			return ""
		}
	}
	if js, err := json.Marshal(o); err == nil {
		return string(js)
	}
	return fmt.Sprint(o)
}

func parseFloat64(o interface{}) float64 {
	switch v := o.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	}
	str := strings.ToUpper(parseString(o))
	switch str {
	case "TRUE":
		str = "1"
	case "FALSE":
		str = "0"
	}
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func parseInt64(o interface{}) int64 {
	switch v := o.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int16:
		return int64(v)
	case int8:
		return int64(v)
	case uint64:
		return int64(v)
	case uint:
		return int64(v)
	case uint32:
		return int64(v)
	case uint16:
		return int64(v)
	case uint8:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	}
	return int64(parseFloat64(o))
}

func parseBool(o interface{}) bool {
	if v, ok := o.(bool); ok {
		return v
	}
	switch strings.ToUpper(parseString(o)) {
	case "", "0", "FALSE", "<NIL>":
		return false
	default:
		return true
	}
}

// ============================================================
// 容器转换
// ============================================================

func toSlice(v interface{}) Slice {
	if v == nil {
		return nil
	}
	switch val := v.(type) {
	case Slice:
		return val
	case []Type:
		return Slice(val)
	case []interface{}:
		res := make(Slice, 0, len(val))
		for _, item := range val {
			res = append(res, New(item))
		}
		return res
	case []string:
		res := make(Slice, 0, len(val))
		for _, item := range val {
			res = append(res, New(item))
		}
		return res
	case []map[string]interface{}:
		res := make(Slice, 0, len(val))
		for _, item := range val {
			res = append(res, New(item))
		}
		return res
	case []map[string]Type:
		res := make(Slice, 0, len(val))
		for _, item := range val {
			res = append(res, New(item))
		}
		return res
	case []Data:
		res := make(Slice, 0, len(val))
		for _, item := range val {
			res = append(res, New(item))
		}
		return res
	case []Map:
		res := make(Slice, 0, len(val))
		for _, item := range val {
			res = append(res, New(item))
		}
		return res
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		l := rv.Len()
		res := make(Slice, 0, l)
		for i := 0; i < l; i++ {
			res = append(res, New(rv.Index(i).Interface()))
		}
		return res
	default:
		return Slice{New(v)}
	}
}

func toMapStringT(v interface{}) MapStringT {
	if v == nil {
		return nil
	}
	switch m := v.(type) {
	case MapStringT:
		return m
	case Map:
		return MapStringT(m)
	case map[string]Type:
		return MapStringT(m)
	case map[string]interface{}:
		res := make(MapStringT, len(m))
		for k, val := range m {
			res[k] = New(val)
		}
		return res
	case Data:
		res := make(MapStringT, len(m))
		for k, val := range m {
			res[k] = New(val)
		}
		return res
	}
	return reflectToMapStringT(v)
}

func reflectToMapStringT(v interface{}) MapStringT {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map {
		return nil
	}
	res := make(MapStringT, rv.Len())
	for _, key := range rv.MapKeys() {
		res[parseString(key.Interface())] = New(rv.MapIndex(key).Interface())
	}
	return res
}

func toMapInterfaceT(v interface{}) MapInterfaceT {
	if v == nil {
		return nil
	}
	switch m := v.(type) {
	case MapInterfaceT:
		return m
	case map[interface{}]Type:
		return MapInterfaceT(m)
	case Map:
		res := make(MapInterfaceT, len(m))
		for k, val := range m {
			res[k] = val
		}
		return res
	case MapStringT:
		res := make(MapInterfaceT, len(m))
		for k, val := range m {
			res[k] = val
		}
		return res
	case map[string]Type:
		res := make(MapInterfaceT, len(m))
		for k, val := range m {
			res[k] = val
		}
		return res
	case map[string]interface{}:
		res := make(MapInterfaceT, len(m))
		for k, val := range m {
			res[k] = New(val)
		}
		return res
	case Data:
		res := make(MapInterfaceT, len(m))
		for k, val := range m {
			res[k] = New(val)
		}
		return res
	}
	return reflectToMapInterfaceT(v)
}

func reflectToMapInterfaceT(v interface{}) MapInterfaceT {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map {
		return nil
	}
	res := make(MapInterfaceT, rv.Len())
	for _, key := range rv.MapKeys() {
		res[key.Interface()] = New(rv.MapIndex(key).Interface())
	}
	return res
}

func toMapStringInterface(v interface{}) map[string]interface{} {
	if v == nil {
		return nil
	}
	switch m := v.(type) {
	case map[string]interface{}:
		return m
	case Data:
		return map[string]interface{}(m)
	case Map:
		res := make(map[string]interface{}, len(m))
		for k, val := range m {
			res[k] = val.Interface()
		}
		return res
	case MapStringT:
		res := make(map[string]interface{}, len(m))
		for k, val := range m {
			res[k] = val.Interface()
		}
		return res
	case map[string]Type:
		res := make(map[string]interface{}, len(m))
		for k, val := range m {
			res[k] = val.Interface()
		}
		return res
	}
	mt := toMapStringT(v)
	res := make(map[string]interface{}, len(mt))
	for k, val := range mt {
		res[k] = val.Interface()
	}
	return res
}

func toSliceMapStringInterface(v interface{}) []map[string]interface{} {
	if v == nil {
		return nil
	}
	switch s := v.(type) {
	case []map[string]interface{}:
		return s
	case []Data:
		res := make([]map[string]interface{}, len(s))
		for i, m := range s {
			res[i] = map[string]interface{}(m)
		}
		return res
	case []Map:
		res := make([]map[string]interface{}, len(s))
		for i, m := range s {
			res[i] = toMapStringInterface(m)
		}
		return res
	case []map[string]Type:
		res := make([]map[string]interface{}, len(s))
		for i, m := range s {
			res[i] = toMapStringInterface(m)
		}
		return res
	}
	sl := toSlice(v)
	res := make([]map[string]interface{}, 0, len(sl))
	for _, item := range sl {
		res = append(res, item.MapStringInterface())
	}
	return res
}
