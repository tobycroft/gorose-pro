package gorose

const (
	// DriverPostgres ...
	DriverPostgres = "postgres"
)

// BuilderPostgres ...
type BuilderPostgres struct {
	//IOrm
	driver string
}

// sqlstr := fmt.Sprintf("SELECT %s%s FROM %s%s%s%s%s%s%s%s",
//		distinct, fields, table, join, where, group, having, order, limit, offset)
// select {distinct} {fields} from {table} {join} {where} {group} {having} {order} {limit} {offset}
// {execute} {table} {data} {where}
func init() {
	var builder = &BuilderPostgres{driver: DriverPostgres}
	NewBuilderDriver().Register(DriverPostgres, builder)
}

// NewBuilderPostgres ...
func NewBuilderPostgres() *BuilderPostgres {
	return &BuilderPostgres{driver: DriverPostgres}
}

// Clone : a new obj
func (b *BuilderPostgres) Clone() IBuilder {
	return &BuilderPostgres{driver: DriverPostgres}
}

// BuildQuery : build query sql string
func (b *BuilderPostgres) BuildQuery(o IOrm) (sqlStr string, args []interface{}, err error) {
	return NewBuilderDefault(o, NewBuilderPostgres()).SetDriver(b.driver).BuildQuery()
}

// BuildExecut : build execute sql string
func (b *BuilderPostgres) BuildExecute(o IOrm, operType string) (sqlStr string, args []interface{}, err error) {
	return NewBuilderDefault(o, NewBuilderPostgres()).SetDriver(b.driver).BuildExecute(operType)
}

func (*BuilderPostgres) AddFieldQuotes(field string) string {
	// 复用全局预编译正则, 避免每次调用都重新编译正则表达式
	if fieldReg.MatchString(field) {
		return `"` + field + `"`
	}
	return field
}
