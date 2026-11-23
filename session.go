package gorose

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const beginStatus = 0

// Session ...
type Session struct {
	IEngin
	IBinder
	master       *sql.DB
	tx           *sql.Tx
	slave        *sql.DB
	lastInsertId int64
	sqlLogs      []string
	lastSql      string
	union        interface{}
	transaction  bool
	tx_index     int64
	err          error
}

var _ ISession = (*Session)(nil)

// NewSession : 初始化 Session
func NewSession(e IEngin) *Session {

	var s = new(Session)
	s.IEngin = e
	// 初始化 IBinder
	s.SetIBinder(NewBinder())

	s.master = e.GetExecuteDB()
	s.slave = e.GetQueryDB()
	//初始化数据
	s.tx_index = beginStatus

	return s
}

func (s *Session) Close() {
	s.master.Close()
	s.slave.Close()
}

// GetIEngin 获取engin
func (s *Session) GetIEngin() IEngin {
	return s.IEngin
}

// GetDriver 获取驱动
func (s *Session) SetIEngin(ie IEngin) {
	s.IEngin = ie
}

// Bind : 传入绑定结果的对象, 参数一为对象, 可以是 struct, gorose.MapRow 或对应的切片
//
//	如果是做非query操作,第一个参数也可以仅仅指定为字符串表名
func (s *Session) Bind(tab interface{}) ISession {
	//fmt.Println(tab, NewBinder(tab))
	//s.SetIBinder(NewBinder(tab))
	s.GetIBinder().SetBindOrigin(tab)
	s.err = s.IBinder.BindParse(s.GetIEngin().GetPrefix())
	return s
}

// GetBinder 获取绑定对象
func (s *Session) GetErr() error {
	return s.err
}

// GetBinder 获取绑定对象
func (s *Session) SetIBinder(ib IBinder) {
	s.IBinder = ib
}

// GetBinder 获取绑定对象
func (s *Session) GetIBinder() IBinder {
	return s.IBinder
}

// GetBinder 获取绑定对象
func (s *Session) ResetBinderResult() {
	_ = s.IBinder.BindParse(s.GetIEngin().GetPrefix())
}

// GetTableName 获取解析后的名字, 提供给orm使用
// 为什么要在这里重复添加该方法, 而不是直接继承 IBinder 的方法呢?
// 是因为, 这里涉及到表前缀的问题, 只能通过session来传递, 所以IOrm就可以选择直接继承
func (s *Session) GetTableName() (string, error) {
	//err := s.IBinder.BindParse(s.GetIEngin().GetPrefix())
	//fmt.Println(s.GetIBinder())
	return s.GetIBinder().GetBindName(), s.err
}

// Begin ...
func (s *Session) Begin() (err error) {
	s.SetTransaction(true)
	if s.tx == nil {
		s.tx, err = s.master.Begin()
	} else {
		num := atomic.AddInt64(&s.tx_index, 1)
		err = s.SavePoint("sp_" + strconv.FormatInt(num, 10))
		if err != nil {
			s.GetIEngin().GetLogger().Error(err.Error())
			return
		} else {
		}
	}
	return
}

// Rollback ...
func (s *Session) Rollback() (err error) {
	if s.tx != nil && s.transaction == true {
		if atomic.LoadInt64(&s.tx_index) == beginStatus {
			s.SetTransaction(false)
			err = s.tx.Rollback()
			s.tx = nil
		} else {
			num := atomic.LoadInt64(&s.tx_index)
			err = s.RollbackTo("sp_" + strconv.FormatInt(num, 10))
			if err != nil {
				s.GetIEngin().GetLogger().Error(err.Error())
				return
			} else {
				atomic.AddInt64(&s.tx_index, -1)
			}
		}
	}
	return
}

// RollbackTo ...
func (s *Session) RollbackTo(savepoint string) (err error) {
	if s.tx != nil && s.transaction == true {
		_, err = s.tx.Exec("rollback to " + savepoint)
		if err != nil {
			s.GetIEngin().GetLogger().Error(err.Error())
			return
		}
	}
	return
}

func (s *Session) SavePoint(savepoint string) (err error) {
	if s.tx != nil && s.transaction == true {
		_, err = s.tx.Exec("savepoint " + savepoint)
		if err != nil {
			s.GetIEngin().GetLogger().Error(err.Error())
			return
		}
	}
	return
}

// Commit ...
func (s *Session) Commit() (err error) {
	if s.tx != nil {
		if atomic.LoadInt64(&s.tx_index) == beginStatus {
			s.SetTransaction(false)
			err = s.tx.Commit()
			s.tx = nil
		} else {
			atomic.AddInt64(&s.tx_index, -1)
		}
	} else {
		s.SetTransaction(false)
	}
	return
}

// Transaction ...
func (s *Session) Transaction(closers ...func(ses ISession) error) (err error) {
	err = s.Begin()
	if err != nil {
		s.GetIEngin().GetLogger().Error(err.Error())
		return err
	}

	for _, closer := range closers {
		err = closer(s)
		if err != nil {
			s.GetIEngin().GetLogger().Error(err.Error())
			_ = s.Rollback()
			return
		}
	}
	return s.Commit()
}

// Query ...
func (s *Session) Query(sqlstring string, args ...interface{}) (result []Data, err error) {
	// 记录开始时间
	start := time.Now()
	//withRunTimeContext(func() {
	if s.err != nil {
		err = s.err
		s.GetIEngin().GetLogger().Error(err.Error())
		// 之前的绑定解析出错, 直接返回, 不再执行查询
		return
	}
	// 记录sqlLog: 惰性格式化, 仅在启用 SQL 日志或慢查询日志时才格式化, 避免高频小查询的无效开销
	if s.GetIEngin().GetLogger().EnableSqlLog() || s.GetIEngin().GetLogger().EnableSlowLog() > 0 {
		s.lastSql = fmt.Sprint(sqlstring, ", ", args)
	} else {
		s.lastSql = ""
	}

	// 直接执行查询, 避免 Prepare+Close 带来的额外网络往返
	var rows *sql.Rows
	// 如果是事务, 则从主库中读写
	if s.tx == nil {
		rows, err = s.slave.Query(sqlstring, args...)
	} else {
		rows, err = s.tx.Query(sqlstring, args...)
	}

	if err != nil {
		s.GetIEngin().GetLogger().Error(err.Error())
		return
	}

	// make sure we always close rows
	defer rows.Close()

	err = s.scan(rows)
	if err != nil {
		s.GetIEngin().GetLogger().Error(err.Error())
		return
	}
	//}, func(duration time.Duration) {
	//	//if duration.Seconds() > 1 {
	//	//	s.GetIEngin().GetLogger().Slow(s.LastSql(), duration)
	//	//} else {
	//	//	s.GetIEngin().GetLogger().Sql(s.LastSql(), duration)
	//	//}
	//})

	timeduration := time.Since(start)
	//if timeduration.Seconds() > 1 {
	s.GetIEngin().GetLogger().Slow(s.LastSql(), timeduration)
	//} else {
	s.GetIEngin().GetLogger().Sql(s.LastSql(), timeduration)
	//}

	result = s.GetIBinder().GetBindAll()
	return
}

// Execute ...
func (s *Session) Execute(sqlstring string, args ...interface{}) (rowsAffected int64, err error) {
	// 记录开始时间
	start := time.Now()
	//withRunTimeContext(func() {
	//	err = s.GetIBinder().BindParse(s.GetIEngin().GetPrefix())
	if s.err != nil {
		// 之前的绑定解析出错, 直接返回, 避免对 nil err 调用 Error() 引发 panic
		err = s.err
		s.GetIEngin().GetLogger().Error(err.Error())
		return
	}
	// 记录sqlLog: 惰性格式化, 仅在启用 SQL 日志或慢查询日志时才格式化, 避免高频小查询的无效开销
	if s.GetIEngin().GetLogger().EnableSqlLog() || s.GetIEngin().GetLogger().EnableSlowLog() > 0 {
		s.lastSql = fmt.Sprint(sqlstring, ", ", args)
	} else {
		s.lastSql = ""
	}

	var operType = strings.ToLower(sqlstring[0:6])
	if operType == "select" {
		err = errors.New("Execute does not allow select operations, please use Query")
		s.GetIEngin().GetLogger().Error(err.Error())
		return
	}

	// 直接执行, 避免 Prepare+Close 带来的额外网络往返
	var result sql.Result
	if s.tx == nil {
		result, err = s.master.Exec(sqlstring, args...)
	} else {
		result, err = s.tx.Exec(sqlstring, args...)
	}

	if err != nil {
		s.GetIEngin().GetLogger().Error(err.Error())
		return
	}

	if operType == "insert" || operType == "replace" {
		// get last insert id
		lastInsertId, err := result.LastInsertId()
		if err == nil {
			s.lastInsertId = lastInsertId
		} else {
			s.GetIEngin().GetLogger().Error(err.Error())
		}
	}
	// get rows affected
	rowsAffected, err = result.RowsAffected()
	timeduration := time.Since(start)
	//}, func(duration time.Duration) {
	if timeduration.Seconds() > 1 {
		s.GetIEngin().GetLogger().Slow(s.LastSql(), timeduration)
	} else {
		s.GetIEngin().GetLogger().Sql(s.LastSql(), timeduration)
	}
	//})
	return
}

// LastInsertId ...
func (s *Session) LastInsertId() int64 {
	return s.lastInsertId
}

// LastSql ...
func (s *Session) LastSql() string {
	return s.lastSql
}

func (s *Session) scan(rows *sql.Rows) (err error) {
	// 如果不需要绑定, 则需要初始化一下binder
	if s.GetIBinder() == nil {
		s.SetIBinder(NewBinder())
	}
	// 检查实多维数组还是一维数组
	switch s.GetBindType() {
	case OBJECT_STRING:
		err = s.scanAll(rows)
	case OBJECT_STRUCT, OBJECT_STRUCT_SLICE:
		err = s.scanStructAll(rows)
	//case OBJECT_MAP, OBJECT_MAP_T:
	//	err = s.scanMap(rows, s.GetBindResult())
	case OBJECT_MAP, OBJECT_MAP_T, OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T:
		err = s.scanMapAll(rows)
	case OBJECT_NIL:
		err = s.scanAll(rows)
	default:
		err = errors.New("Bind value error")
	}
	return
}

//func (s *Session) scanMap(rows *sql.Rows, dst interface{}) (err error) {
//	return s.scanMapAll(rows, dst)
//}

func (s *Session) scanMapAll(rows *sql.Rows) (err error) {
	// 获取查询的所有字段
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	count := len(columns)

	// 反射操作提升到循环外, 避免在 行×列 的内层循环中重复执行
	br := reflect.Indirect(reflect.ValueOf(s.GetBindResult()))
	bindResultMapType := br.Type()
	bindType := s.GetBindType()

	// 复用扫描缓冲区, 避免每行重复分配
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// 预计算列名反射值, 避免内层循环重复创建反射对象
	colVals := make([]reflect.Value, count)
	for i, col := range columns {
		colVals[i] = reflect.ValueOf(col)
	}

	// 普通 map[string]interface{} 的元素类型是空接口, 走无反射快速路径
	isPlainMap := bindResultMapType.Elem().Kind() == reflect.Interface &&
		bindResultMapType.Elem().NumMethod() == 0
	// 记录最后一行, 循环结束后一次性更新最新一条结果对象
	var lastItem map[string]interface{}

	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return
		}

		// union 操作: 直接取第一列的值返回
		if s.GetUnion() != nil {
			if count > 0 {
				if b, ok := values[0].([]byte); ok {
					s.union = string(b)
				} else {
					s.union = values[0]
				}
			}
			return
		}

		if isPlainMap {
			// 快速路径: 普通 map[string]interface{}, 不做反射式赋值
			item := make(map[string]interface{}, count)
			for i, col := range columns {
				val := values[i]
				if b, ok := val.([]byte); ok {
					item[col] = string(b)
				} else {
					item[col] = val
				}
			}
			if bindType == OBJECT_MAP_SLICE {
				mv := reflect.ValueOf(item).Convert(bindResultMapType)
				s.GetBindResultSlice().Set(reflect.Append(s.GetBindResultSlice(), mv))
			}
			lastItem = item
			continue
		}

		// Type 或其他类型: 保留反射赋值
		var bindResultTmp = reflect.MakeMap(bindResultMapType)
		for i := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			switch bindType {
			case OBJECT_MAP_T, OBJECT_MAP_SLICE_T: // Type 类型
				tv := reflect.ValueOf(New(v))
				// 绑定到一条数据结果对象上,方便其他地方的调用,永远存储最新一条
				br.SetMapIndex(colVals[i], tv)
				if bindType == OBJECT_MAP_SLICE_T {
					bindResultTmp.SetMapIndex(colVals[i], tv)
				}
			default: // 其他类型
				br.SetMapIndex(colVals[i], reflect.ValueOf(v))
				if bindType == OBJECT_MAP_SLICE {
					bindResultTmp.SetMapIndex(colVals[i], reflect.ValueOf(v))
				}
			}
		}
		// 如果是多条数据集, 就插入到对应的结果集slice上
		if bindType == OBJECT_MAP_SLICE || bindType == OBJECT_MAP_SLICE_T {
			s.GetBindResultSlice().Set(reflect.Append(s.GetBindResultSlice(), bindResultTmp))
		}
	}

	// 更新最新一条结果对象, 便于 GetBindResult() 返回最后一条数据
	if isPlainMap && lastItem != nil {
		if bindType == OBJECT_MAP {
			br.Set(reflect.ValueOf(lastItem).Convert(bindResultMapType))
		} else {
			for i, col := range columns {
				br.SetMapIndex(colVals[i], reflect.ValueOf(lastItem[col]))
			}
		}
	}
	return rows.Err()
}

// ScanAll scans all sql result rows into a slice of structs.
// It reads all rows and closes rows when finished.
// dst should be a pointer to a slice of the appropriate type.
// The new results will be appended to any existing data in dst.
func (s *Session) scanStructAll(rows *sql.Rows) error {
	// check if there is data waiting
	//if !rows.Next() {
	//	if err := rows.Err(); err != nil {
	//		s.GetIEngin().GetLogger().Error(err.Error())
	//		return err
	//	}
	//	return sql.ErrNoRows
	//}
	var sfs = structForScan(s.GetBindResult())
	for rows.Next() {
		if s.GetUnion() != nil {
			var union interface{}
			err := rows.Scan(&union)
			if err != nil {
				s.GetIEngin().GetLogger().Error(err.Error())
				return err
			}
			s.union = union
			return err
		}
		// scan it
		//fmt.Printf("%#v \n",structForScan(s.GetBindResult()))
		err := rows.Scan(sfs...)
		if err != nil {
			s.GetIEngin().GetLogger().Error(err.Error())
			return err
		}
		// 如果是union操作就不需要绑定数据直接返回, 否则就绑定数据
		if s.GetUnion() == nil {
			// 如果是多条数据集, 就插入到对应的结果集slice上
			if s.GetBindType() == OBJECT_STRUCT_SLICE {
				// add to the result slice
				s.GetBindResultSlice().Set(reflect.Append(s.GetBindResultSlice(),
					reflect.Indirect(reflect.ValueOf(s.GetBindResult()))))
			}
		}
	}

	return rows.Err()
}

func (s *Session) scanAll(rows *sql.Rows) (err error) {
	var columns []string
	// 获取查询的所有字段
	if columns, err = rows.Columns(); err != nil {
		return
	}
	count := len(columns)

	// 复用扫描缓冲区, 避免每行重复分配
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result = []Data{}
	for rows.Next() {
		// 获取结果
		if err = rows.Scan(scanArgs...); err != nil {
			return
		}

		// 定义预设的绑定对象, 按列数预估容量
		var resultTmp = make(Data, count)
		//// 定义union操作的map返回
		//var unionTmp = map[string]interface{}{}
		for i, col := range columns {
			var v interface{}
			val := values[i]
			if b, ok := val.([]byte); ok {
				v = string(b)
			} else {
				v = val
			}
			if s.GetUnion() != nil {
				s.union = v
				return
				// 以下上通用解决方法
				//unionTmp[col] = v
				//s.union = unionTmp
			}
			resultTmp[col] = v
		}
		result = append(result, resultTmp)
	}
	s.IBinder.SetBindAll(result)
	return rows.Err()
}

// SetUnion ...
func (s *Session) SetUnion(u interface{}) {
	s.union = u
}

// GetUnion ...
func (s *Session) GetUnion() interface{} {
	return s.union
}

// SetTransaction ...
func (s *Session) SetTransaction(b bool) {
	s.transaction = b
}

// GetTransaction 提供给 orm 使用的, 方便reset操作
func (s *Session) GetTransaction() bool {
	return s.transaction
}
