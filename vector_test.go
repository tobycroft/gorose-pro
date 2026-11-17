package gorose

import (
	"strings"
	"testing"
)

func newVectorOrm(t *testing.T) *Orm {
	t.Helper()
	e, err := Open(&Config{Driver: "sqlite3", Dsn: ":memory:"})
	if err != nil {
		t.Fatalf("初始化engin失败: %v", err)
	}
	e.TagName("orm")
	e.IgnoreName("ignore")
	o := NewOrm(e)
	o.table = "docs"
	return o
}

// mysql 插入: 向量需要 STRING_TO_VECTOR 包裹
func TestVector_InsertMysql(t *testing.T) {
	o := newVectorOrm(t)
	o.Data(Data{"title": "hello"}).VecData("embedding", Vector{0.1, 0.2, 0.3})

	sqlStr, args, err := NewBuilderMysql().BuildExecute(o, "insert")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if !strings.Contains(sqlStr, "INSERT INTO `docs`") {
		t.Fatalf("sql错误: %s", sqlStr)
	}
	if !strings.Contains(sqlStr, "STRING_TO_VECTOR(?)") {
		t.Fatalf("mysql向量应包裹STRING_TO_VECTOR: %s", sqlStr)
	}
	if len(args) != 2 {
		t.Fatalf("绑定参数错误: %v", args)
	}
	if !containsValue(args, "hello") || !containsValue(args, "[0.1,0.2,0.3]") {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

func containsValue(args []interface{}, want interface{}) bool {
	for _, a := range args {
		if a == want {
			return true
		}
	}
	return false
}

// postgres 插入: 直接绑定文本形式
func TestVector_InsertPostgres(t *testing.T) {
	o := newVectorOrm(t)
	o.Data(Data{"title": "hello"}).VecData("embedding", Vector{0.1, 0.2, 0.3})

	sqlStr, args, err := NewBuilderPostgres().BuildExecute(o, "insert")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if strings.Contains(sqlStr, "STRING_TO_VECTOR") {
		t.Fatalf("pg不应包裹STRING_TO_VECTOR: %s", sqlStr)
	}
	if !strings.Contains(sqlStr, "$2") {
		t.Fatalf("pg应使用$N占位符: %s", sqlStr)
	}
	if len(args) != 2 || !containsValue(args, "[0.1,0.2,0.3]") {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// clickhouse 插入: 直接绑定 []float32
func TestVector_InsertClickhouse(t *testing.T) {
	o := newVectorOrm(t)
	o.VecData("embedding", Vector{0.1, 0.2, 0.3})

	sqlStr, args, err := NewBuilderClickhouse().BuildExecute(o, "insert")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if len(args) != 1 {
		t.Fatalf("绑定参数数量错误: %v", args)
	}
	f32, ok := args[0].([]float32)
	if !ok || len(f32) != 3 || f32[2] != 0.3 {
		t.Fatalf("clickhouse应绑定[]float32: %v(%T)", args[0], args[0])
	}
}

// oracle 插入: 绑定文本形式, 参数走 orm.bindValues
func TestVector_InsertOracle(t *testing.T) {
	o := newVectorOrm(t)
	o.VecData("embedding", Vector{0.1, 0.2, 0.3})

	sqlStr, args, err := NewBuilderOracle(o).BuildExecute(o, "insert")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if len(args) != 1 || args[0] != "[0.1,0.2,0.3]" {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// postgres KNN 排序: "embedding" <=> $1
func TestVector_OrderPostgres(t *testing.T) {
	o := newVectorOrm(t)
	o.OrderVector("embedding", Vector{1, 2, 3}, VectorCosine).Limit(10)

	sqlStr, args, err := NewBuilderPostgres().BuildQuery(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if !strings.Contains(sqlStr, `ORDER BY "embedding" <=> $1`) {
		t.Fatalf("sql错误: %s", sqlStr)
	}
	if !strings.Contains(sqlStr, "LIMIT 10") {
		t.Fatalf("应包含LIMIT: %s", sqlStr)
	}
	if len(args) != 1 || args[0] != "[1,2,3]" {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// mysql: 常规where + 向量距离过滤 + 向量排序, 绑定参数顺序要对
func TestVector_WhereAndOrderMysql(t *testing.T) {
	o := newVectorOrm(t)
	o.Where("uid", ">", 1).
		WhereVectorDistance("embedding", VectorString("[1,2,3]"), VectorCosine, VectorOpLt, 0.5).
		OrderVector("embedding", VectorString("[1,2,3]"), VectorCosine).
		Limit(5)

	sqlStr, args, err := NewBuilderMysql().BuildQuery(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if !strings.Contains(sqlStr, "WHERE `uid` > ? AND DISTANCE(`embedding`, ?, 'COSINE') < ?") {
		t.Fatalf("sql错误: %s", sqlStr)
	}
	if !strings.Contains(sqlStr, "ORDER BY DISTANCE(`embedding`, ?, 'COSINE') LIMIT 5") {
		t.Fatalf("sql错误: %s", sqlStr)
	}
	if len(args) != 4 || args[0] != 1 || args[1] != "[1,2,3]" || args[2] != 0.5 || args[3] != "[1,2,3]" {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// clickhouse 排序: L2Distance
func TestVector_OrderClickhouse(t *testing.T) {
	o := newVectorOrm(t)
	o.OrderVector("embedding", Vector{1, 2}, VectorL2).Limit(3)

	sqlStr, args, err := NewBuilderClickhouse().BuildQuery(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if !strings.Contains(sqlStr, "L2Distance(`embedding`, ?)") {
		t.Fatalf("sql错误: %s", sqlStr)
	}
	if len(args) != 1 {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// oracle: VECTOR_DISTANCE + :N 占位符
func TestVector_OrderOracle(t *testing.T) {
	o := newVectorOrm(t)
	o.OrderVector("embedding", Vector{1, 2, 3}, VectorDot).Limit(5)

	sqlStr, args, err := NewBuilderOracle(o).BuildQuery(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if !strings.Contains(sqlStr, `VECTOR_DISTANCE('DOT', "embedding", :1)`) {
		t.Fatalf("sql错误: %s", sqlStr)
	}
	if len(args) != 1 || args[0] != "[1,2,3]" {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// mssql: VECTOR_DISTANCE('cosine', ...)
func TestVector_OrderMssql(t *testing.T) {
	o := newVectorOrm(t)
	o.OrderVector("embedding", Vector{1, 2, 3}, VectorCosine).Limit(5)

	sqlStr, args, err := NewBuilderMsSql().BuildQuery(o)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if !strings.Contains(sqlStr, "VECTOR_DISTANCE('cosine', `embedding`, ?)") {
		t.Fatalf("sql错误: %s", sqlStr)
	}
}

// sqlite 无原生向量检索, 应报错
func TestVector_UnsupportedSqlite(t *testing.T) {
	o := newVectorOrm(t)
	o.OrderVector("embedding", Vector{1, 2, 3}, VectorCosine)

	_, _, err := NewBuilderSqlite3().BuildQuery(o)
	if err == nil || !strings.Contains(err.Error(), "不支持向量检索") {
		t.Fatalf("sqlite应报不支持错误: %v", err)
	}
}

// 非法度量与空向量应报错
func TestVector_InvalidInput(t *testing.T) {
	o := newVectorOrm(t)
	o.OrderVector("embedding", Vector{1}, VectorMetric("FOO"))
	if _, _, err := NewBuilderMysql().BuildQuery(o); err == nil {
		t.Fatal("非法度量应报错")
	}

	o2 := newVectorOrm(t)
	o2.OrderVector("embedding", Vector{}, VectorCosine)
	if _, _, err := NewBuilderMysql().BuildQuery(o2); err == nil {
		t.Fatal("空向量应报错")
	}
}

// DataStruct([]struct) + VecData: 向量字段延迟注入到每行
func TestVector_DeferredDataInject(t *testing.T) {
	o := newVectorOrm(t)
	o.DataStruct([]Users{
		{Name: "a", Age: 1},
		{Name: "b", Age: 2},
	}).VecData("embedding", Vector{0.1, 0.2})

	sqlStr, args, err := NewBuilderMysql().BuildExecute(o, "insert")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sqlStr, args)
	if strings.Count(sqlStr, "STRING_TO_VECTOR(?)") != 2 {
		t.Fatalf("每行都应注入向量字段: %s", sqlStr)
	}
	if len(args) != 8 {
		t.Fatalf("绑定参数错误: %v", args)
	}
}

// ParseVector 往返解析
func TestVector_ParseVector(t *testing.T) {
	origin := Vector{0.1, -2.5, 3e10, 0}
	str := origin.String()
	t.Log(str)

	parsed, err := ParseVector(str)
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 4 || parsed[0] != 0.1 || parsed[1] != -2.5 || parsed[3] != 0 {
		t.Fatalf("解析结果错误: %v", parsed)
	}

	if _, err := ParseVector("[abc]"); err == nil {
		t.Fatal("非法向量应报错")
	}
	if v, err := ParseVector("[]"); err != nil || v != nil {
		t.Fatalf("空向量: %v, %v", v, err)
	}
}
