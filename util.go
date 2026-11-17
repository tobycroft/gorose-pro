package gorose

import (
	"log"
	"math/rand"
	"os"
	"path"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gohouse/t"
)

// tagNameCache 按 (类型, tag名) 缓存 getTagName 的反射解析结果
var tagNameCache sync.Map

func getRandomInt(num int) int {
	// Go 1.20+ 全局 rand 源自动随机播种, 每次手动 Seed 反而引入锁争用与相同种子问题
	return rand.Intn(num)
}

func structForScan(u interface{}) []interface{} {
	val := reflect.Indirect(reflect.ValueOf(u))
	v := make([]interface{}, 0)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if val.Type().Field(i).Tag.Get(getTagname()) != getIgnorename() {
			if valueField.CanAddr() {
				v = append(v, valueField.Addr().Interface())
			} else {
				//v[i] = valueField
				v = append(v, valueField)
			}
		}
	}
	return v
}

// StructToMap this will be deprecated after v1.15.0 please use StructToMapV2
func StructToMap(obj interface{}) map[string]interface{} {
	ty := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < ty.NumField(); i++ {
		data[ty.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// getTagName 获取结构体中Tag的值，如果没有tag则返回字段值
// 结果按 (类型, tag名) 缓存, 同一结构体类型只解析一次反射
func getTagName(structName interface{}, tagstr string) []string {
	// 获取type
	tag := reflect.TypeOf(structName)
	// 如果是反射Ptr类型, 就获取他的 element type
	if tag.Kind() == reflect.Ptr {
		tag = tag.Elem()
	}

	// 判断是否是struct
	if tag.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}

	var cacheKey = struct {
		typ    reflect.Type
		tag    string
		ignore string
	}{tag, tagstr, getIgnorename()}
	if cached, ok := tagNameCache.Load(cacheKey); ok {
		// 返回副本, 防止调用方修改缓存
		src := cached.([]string)
		dst := make([]string, len(src))
		copy(dst, src)
		return dst
	}

	// 获取字段数量
	fieldNum := tag.NumField()
	ignoreName := getIgnorename()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		// tag 名字
		tagName := tag.Field(i).Tag.Get(tagstr)
		if tagName != ignoreName {
			// tag为-时, 不解析
			if tagName == "-" || tagName == "" {
				// 字段名字
				tagName = tag.Field(i).Name
			}
			result = append(result, tagName)
		}
	}
	tagNameCache.Store(cacheKey, result)

	// 返回副本, 防止调用方修改缓存
	dst := make([]string, len(result))
	copy(dst, result)
	return dst
}

// If : ternary operator (三元运算)
// condition:比较运算
// trueVal:运算结果为真时的值
// falseVal:运算结果为假时的值
// return: 由于不知道传入值的类型, 所有, 必须在接收结果时, 指定对应的值类型
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// InArray :给定元素值 是否在 指定的数组中
func inArray(needle, hystack interface{}) bool {
	nt := t.New(needle)
	for _, item := range t.New(hystack).Slice() {
		if strings.EqualFold(nt.String(), item.String()) {
			return true
		}
	}
	return false
}

func withRunTimeContext(closer func(), callback func(time.Duration)) {
	// 记录开始时间
	start := time.Now()
	closer()
	timeduration := time.Since(start)
	//log.Println("执行完毕,用时:", timeduration.Seconds(),timeduration.Seconds()>1.1)
	callback(timeduration)
}

func readFile(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		// 目录不存在时先创建目录再重试
		_ = os.MkdirAll(path.Dir(filepath), os.ModePerm)
		file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return nil
		}
	}
	return file
}
