package gorose

// VectorWhereItem 向量距离过滤条件
type VectorWhereItem struct {
	Field     string
	Value     VectorValue
	Metric    VectorMetric
	Op        VectorOp
	Threshold interface{}
}

// VectorOrderItem 向量距离排序项(KNN)
type VectorOrderItem struct {
	Field  string
	Value  VectorValue
	Metric VectorMetric
}

// VecData 写入/更新向量字段, 输入项受编译期类型限制(gorose.Vector / gorose.VectorString)。
// data 为空时自动创建 map; data 已是 map/Data 时直接合并;
// 其他类型(struct/slice/string)时暂存, 构建阶段注入。
func (dba *Orm) VecData(field string, value VectorValue) IOrm {
	switch d := dba.data.(type) {
	case nil:
		dba.data = map[string]interface{}{field: value}
	case map[string]interface{}:
		if d == nil {
			dba.data = map[string]interface{}{field: value}
		} else {
			d[field] = value
		}
	case Data:
		if d == nil {
			dba.data = map[string]interface{}{field: value}
		} else {
			d[field] = value
		}
	default:
		if dba.vectorData == nil {
			dba.vectorData = make(map[string]VectorValue)
		}
		dba.vectorData[field] = value
	}
	return dba
}

// OrderVector 向量最近邻排序(KNN), 需配合 Limit 使用。
// 示例: dba.Table("docs").OrderVector("embedding", gorose.Vector{0.1, 0.2}, gorose.VectorCosine).Limit(10)
func (dba *Orm) OrderVector(field string, value VectorValue, metric VectorMetric) IOrm {
	dba.vectorOrders = append(dba.vectorOrders, VectorOrderItem{
		Field:  field,
		Value:  value,
		Metric: metric,
	})
	return dba
}

// WhereVectorDistance 向量距离过滤: DISTANCE(field, value, metric) op threshold。
// 示例: dba.Table("docs").WhereVectorDistance("embedding", vec, gorose.VectorCosine, gorose.VectorOpLt, 0.5)
func (dba *Orm) WhereVectorDistance(field string, value VectorValue, metric VectorMetric, op VectorOp, threshold interface{}) IOrm {
	dba.vectorWheres = append(dba.vectorWheres, VectorWhereItem{
		Field:     field,
		Value:     value,
		Metric:    metric,
		Op:        op,
		Threshold: threshold,
	})
	return dba
}
