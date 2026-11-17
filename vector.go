package gorose

import (
	"errors"
	"strconv"
	"strings"
)

// Vector 单精度(float32)稠密向量。
// 定义类型(不同于 []float32): 普通 []float32 无法直接传入向量 API,
// 必须使用 gorose.Vector{...} 或显式转换 gorose.Vector([]float32{...}),
// 借助编译期检查把误用拦截在 IDE/编译阶段。
type Vector []float32

// VectorString 预序列化的向量字符串, 如 "[1,2,3]"。
// 适合需要用户自行控制类型的场景, 如 pgvector 的 halfvec/sparsevec/bit。
type VectorString string

// VectorMetric 向量距离度量
type VectorMetric string

const (
	// VectorL2 欧几里得距离(L2)
	VectorL2 VectorMetric = "L2"
	// VectorCosine 余弦距离
	VectorCosine VectorMetric = "COSINE"
	// VectorDot 内积(点积)
	// 统一为"距离"语义: 取负内积, 值越小越相似, 配合升序排序使用
	VectorDot VectorMetric = "DOT"
)

// VectorOp 向量距离比较操作符
type VectorOp string

const (
	VectorOpEq  VectorOp = "="
	VectorOpLt  VectorOp = "<"
	VectorOpLte VectorOp = "<="
	VectorOpGt  VectorOp = ">"
	VectorOpGte VectorOp = ">="
)

// VectorValue 向量取值。
// 密封接口(私有方法): 包外类型无法实现, 编译期只允许 Vector / VectorString 两种输入。
type VectorValue interface {
	isVectorValue()
}

func (Vector) isVectorValue()      {}
func (VectorString) isVectorValue() {}

// String 将向量序列化为数据库通用文本表示 "[1,2,3]"
func (v Vector) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	for i, f := range v {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.FormatFloat(float64(f), 'g', -1, 32))
	}
	sb.WriteByte(']')
	return sb.String()
}

// ParseVector 将数据库返回的向量文本 "[1,2,3]" 解析为 Vector
func ParseVector(s string) (Vector, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	if s == "" {
		return nil, nil
	}
	parts := strings.Split(s, ",")
	v := make(Vector, 0, len(parts))
	for _, p := range parts {
		f, err := strconv.ParseFloat(strings.TrimSpace(p), 32)
		if err != nil {
			return nil, errors.New("向量解析失败: " + err.Error())
		}
		v = append(v, float32(f))
	}
	return v, nil
}

// valid 校验度量是否受支持
func (m VectorMetric) valid() bool {
	switch m {
	case VectorL2, VectorCosine, VectorDot:
		return true
	}
	return false
}

// canonicalName 统一度量名映射为 mysql/mssql/oracle 的距离函数参数名
func (m VectorMetric) canonicalName() string {
	switch m {
	case VectorL2:
		return "EUCLIDEAN"
	case VectorCosine:
		return "COSINE"
	case VectorDot:
		return "DOT"
	}
	return string(m)
}

// vectorToBind 将向量值按驱动转换为绑定参数值:
// clickhouse 绑定 []float32; mysql/postgres/mssql/oracle/sqlite 绑定文本形式
func vectorToBind(v VectorValue, driver string) interface{} {
	switch vv := v.(type) {
	case Vector:
		if driver == DriverClickhouse {
			return []float32(vv)
		}
		return vv.String()
	case VectorString:
		return string(vv)
	}
	return v
}

// vectorBindIfVector 如果值是向量类型则按驱动转换为绑定值, 否则原样返回
func vectorBindIfVector(v interface{}, driver string) interface{} {
	if vi, ok := v.(VectorValue); ok {
		return vectorToBind(vi, driver)
	}
	return v
}
