package gorose

import (
	"fmt"
	"regexp"
)

// IBuilder ...
type IBuilder interface {
	IFieldQuotes
	BuildQuery(orm IOrm) (sqlStr string, args []interface{}, err error)
	BuildExecute(orm IOrm, operType string) (sqlStr string, args []interface{}, err error)
	Clone() IBuilder
	//GetIOrm() IOrm
}

// IFieldQuotes 给系统关键词冲突的字段加引号,如: mysql是反引号, pg是双引号
type IFieldQuotes interface {
	AddFieldQuotes(field string) string
}

// 全局预编译正则: 只初始化一次, 性能提升巨大!!!
var fieldReg = regexp.MustCompile(`^[\w-]+$`)

type FieldQuotesDefault struct {
}

func (FieldQuotesDefault) AddFieldQuotes(field string) string {
	// 直接使用全局编译好的正则, 不再重复编译
	if fieldReg.MatchString(field) {
		return fmt.Sprintf("`%s`", field)
	}
	return field
}

func (FieldQuotesDefault) AddFieldQuotesOracle(field string) string {
	// 共用同一个正则规则
	if fieldReg.MatchString(field) {
		return fmt.Sprintf("\"%s\"", field)
	}
	return field
}
