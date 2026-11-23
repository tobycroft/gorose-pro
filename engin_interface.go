package gorose

import "database/sql"

// IEngin ...
type IEngin interface {
	GetExecuteDB() *sql.DB
	GetQueryDB() *sql.DB
	GetExecuteDBStats() sql.DBStats
	GetQueryDBStats() sql.DBStats
	//EnableSqlLog(e ...bool)
	//IfEnableSqlLog() (e bool)
	//SetPrefix(pre string)
	GetPrefix() (pre string)
	//NewSession() ISession
	//NewOrm() IOrm
	SetLogger(lg ILogger)
	GetLogger() ILogger
	GetDriver() string
}
