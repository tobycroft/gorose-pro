package gorose

// OrmApi ...
type OrmApi struct {
	table    string
	fields   []string
	where    [][]interface{}
	order    string
	limit    int
	offset   int
	join     [][]interface{}
	distinct bool
	//union     string
	group     string
	having    string
	data      interface{}
	force     bool
	extraCols []string
	// 悲观锁
	pessimisticLock string
	// 向量支持: 延迟合并的向量字段(VecData 作用在 struct/slice 数据上时)
	vectorData map[string]VectorValue
	// 向量距离过滤条件
	vectorWheres []VectorWhereItem
	// 向量距离排序项(KNN)
	vectorOrders []VectorOrderItem
}

// GetTable ...
func (o *Orm) GetTable() string {
	return o.table
}

// GetFields ...
func (o *Orm) GetFields() []string {
	return o.fields
}

// SetWhere ...
func (o *Orm) SetWhere(arg [][]interface{}) {
	o.where = arg
}

// GetWhere ...
func (o *Orm) GetWhere() [][]interface{} {
	return o.where
}

// GetOrder ...
func (o *Orm) GetOrder() string {
	return o.order
}

// GetLimit ...
func (o *Orm) GetLimit() int {
	return o.limit
}

// GetOffset ...
func (o *Orm) GetOffset() int {
	return o.offset
}

// GetJoin ...
func (o *Orm) GetJoin() [][]interface{} {
	return o.join
}

// GetDistinct ...
func (o *Orm) GetDistinct() bool {
	return o.distinct
}

// GetGroup ...
func (o *Orm) GetGroup() string {
	return o.group
}

// GetHaving ...
func (o *Orm) GetHaving() string {
	return o.having
}

// GetData ...
func (o *Orm) GetData() interface{} {
	return o.data
}

// GetForce ...
func (o *Orm) GetForce() bool {
	return o.force
}

// GetExtraCols ...
func (o *Orm) GetExtraCols() []string {
	return o.extraCols
}

// GetPessimisticLock ...
func (o *Orm) GetPessimisticLock() string {
	return o.pessimisticLock
}

// GetVectorData 获取延迟合并的向量字段
func (o *Orm) GetVectorData() map[string]VectorValue {
	return o.vectorData
}

// GetVectorWheres 获取向量距离过滤条件
func (o *Orm) GetVectorWheres() []VectorWhereItem {
	return o.vectorWheres
}

// GetVectorOrders 获取向量距离排序项
func (o *Orm) GetVectorOrders() []VectorOrderItem {
	return o.vectorOrders
}
