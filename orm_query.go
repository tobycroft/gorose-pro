package gorose

import (
	"errors"
	"github.com/gohouse/t"
	"math"
	"reflect"
	"strings"
	"sync"
)

// Select : select one or more rows , relation limit set
func (dba *Orm) Select() error {
	switch dba.GetIBinder().GetBindType() {
	case OBJECT_STRUCT, OBJECT_MAP, OBJECT_MAP_T:
		dba.Limit(1)
	}
	// 构建sql
	sqlStr, args, err := dba.BuildSql()
	if err != nil {
		return err
	}
	// 执行查询
	_, err = dba.GetISession().Query(sqlStr, args...)
	return err
}

func (dba *Orm) Scan(scan_to_struct interface{}) error {
	dstVal := reflect.ValueOf(scan_to_struct)
	sliceVal := reflect.Indirect(dstVal)
	switch sliceVal.Kind() {
	case reflect.Struct: // struct
		dba.Limit(1)
		sqlStr, args, err := dba.BuildSql()
		if err != nil {
			return err
		}
		dba.GetIBinder().SetBindType(OBJECT_STRUCT)
		dba.GetIBinder().SetBindResult(scan_to_struct)
		if len(dba.GetIBinder().GetBindFields()) == 0 {
			dba.GetIBinder().SetBindFields(getTagName(dba.GetIBinder().GetBindResult(), TAGNAME))
		}
		switch dstVal.Kind() {
		case reflect.Ptr, reflect.Struct:
			break
		default:
			return errors.New("传入的对象有误,示例:var user User,传入 &user{}")
		}
		_, err = dba.GetISession().Query(sqlStr, args...)
		return err

	case reflect.Slice:
		eltType := sliceVal.Type().Elem()
		switch eltType.Kind() {
		case reflect.Struct:
			sqlStr, args, err := dba.BuildSql()
			if err != nil {
				return err
			}
			dba.GetIBinder().SetBindType(OBJECT_STRUCT_SLICE)
			br := reflect.New(eltType)
			dba.GetIBinder().SetBindResult(br.Interface())
			dba.GetIBinder().SetBindResultSlice(sliceVal)
			if len(dba.GetIBinder().GetBindFields()) == 0 {
				dba.GetIBinder().SetBindFields(getTagName(dba.GetIBinder().GetBindResult(), TAGNAME))
			}
			switch dstVal.Kind() {
			case reflect.Ptr, reflect.Struct:
				break
			default:
				return errors.New("传入的对象有误,示例:var user User,传入 &user{}")
			}
			_, err = dba.GetISession().Query(sqlStr, args...)
			return err

		default:
			return errors.New("传入[]struct{}将会解析成多条，类似Get方法，注意需要传入指针值，例如传入：&User{},而不是：User{}，不要用这个方法传入Map")

		}

	default:
		return errors.New("传入struct{}可以解析单条，类似Find方法，传入[]struct{}将会解析成多条，类似Get方法，注意需要传入指针值，例如传入：&User{},而不是：User{}，不要用这个方法传入Map")
	}
}

// First : select one row , relation limit set
func (dba *Orm) First() (result Data, err error) {
	dba.GetIBinder().SetBindType(OBJECT_STRING)
	err = dba.Limit(1).Select()
	if err != nil {
		return
	}
	res := dba.GetISession().GetBindAll()
	if len(res) > 0 {
		result = res[0]
	}
	return
}

func (dba *Orm) Find() (result Data, err error) {
	return dba.First()
}

// Get : select more rows , relation limit set
func (dba *Orm) Get() (result []Data, err error) {
	dba.GetIBinder().SetBindType(OBJECT_STRING)
	tabname := dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)
	dba.ResetTable()
	dba.Table(tabname2)
	err = dba.Select()
	result = dba.GetISession().GetBindAll()
	return
}

// Count : select count rows
func (dba *Orm) Count(args ...string) (int64, error) {
	fields := "*"
	if len(args) > 0 {
		fields = args[0]
	}
	count, err := dba._unionBuild("COUNT", fields)
	if count == nil {
		return 0, err
	}
	return t.New(count).Int64(), err
}

// Counts is a Nested-Count function in order to solve the wrong number output when using count and groupBy in the same time : select count (select count.....)
func (dba *Orm) Counts(count_fileds ...string) (int64, error) {
	if dba.group == "" {
		return dba.Count(count_fileds...)
	} else {
		//temporary remove order in case of the renamed order not found which cause error sql
		order := dba.order
		dba.order = ""
		dba.limit = 0
		if len(dba.fields) < 1 {
			dba.fields = []string{"1"}
		}
		//fmt.Println(dba.fields)
		//dba.fields = []string{"count(DISTINCT " + dba.group + ") as count"}
		// 构建sql
		sqls, args, err := dba.BuildSql()
		//fmt.Println(sqls)
		if err != nil {
			return 0, err
		}
		dba.order = order
		total_number, err := dba.Query(`SELECT COUNT(*) as COUNT from(`+sqls+`) as COUNTS`, args...)
		if err != nil {
			return 0, err
		}
		//fmt.Println(dba.LastSql())
		if len(total_number) < 1 {
			return 0, err
		}
		return t.New(total_number[0]["COUNT"]).Int64(), err
	}
}

// Sum : select sum field
func (dba *Orm) Sum(sum string) (interface{}, error) {
	return dba._unionBuild("sum", sum)
}

// Avg : select avg field
func (dba *Orm) Avg(avg string) (interface{}, error) {
	return dba._unionBuild("avg", avg)
}

// Max : select max field
func (dba *Orm) Max(max string) (interface{}, error) {
	return dba._unionBuild("max", max)
}

// Min : select min field
func (dba *Orm) Min(min string) (interface{}, error) {
	return dba._unionBuild("min", min)
}

// _unionBuild : build union select real
func (dba *Orm) _unionBuild(union, field string) (interface{}, error) {
	fields := union + "(" + field + ") as " + union
	dba.fields = []string{fields}

	res, err := dba.First()
	if r, ok := res[union]; ok {
		return r, err
	}
	return 0, err
}

//func (dba *Orm) _unionBuild_bak(union, field string) (interface{}, error) {
//	var tmp interface{}
//
//	dba.union = union + "(" + field + ") as " + union
//	// 缓存fields字段,暂时由union占用
//	fieldsTmp := dba.fields
//	dba.fields = []string{dba.union}
//	dba.GetISession().SetUnion(true)
//
//	// 构建sql
//	sqls, args, err := dba.BuildSql()
//	if err != nil {
//		return tmp, err
//	}
//
//	// 执行查询
//	_, err = dba.GetISession().Query(sqls, args...)
//	if err != nil {
//		return tmp, err
//	}
//
//	// 重置union, 防止复用的时候感染
//	dba.union = ""
//	// 返还fields
//	dba.fields = fieldsTmp
//
//	// 语法糖获取union值
//	if dba.GetISession().GetUnion() != nil {
//		tmp = dba.GetISession().GetUnion()
//		// 获取之后, 释放掉
//		dba.GetISession().SetUnion(nil)
//	}
//
//	return tmp, nil
//}

// Pluck 获取一列数据, 第二个字段可以指定另一个字段的值作为这一列数据的key
func (dba *Orm) Pluck(field string, fieldKey ...string) (v interface{}, err error) {
	var resMap = make(map[interface{}]interface{}, 0)
	var resSlice = make([]interface{}, 0)

	res, err := dba.Get()

	if err != nil {
		return
	}

	if len(res) > 0 {
		for _, val := range res {
			if len(fieldKey) > 0 {
				resMap[val[fieldKey[0]]] = val[field]
			} else {
				resSlice = append(resSlice, val[field])
			}
		}
	}
	if len(fieldKey) > 0 {
		v = resMap
	} else {
		v = resSlice
	}
	return
}

// Pluck_bak ...
func (dba *Orm) Pluck_bak(field string, fieldKey ...string) (v interface{}, err error) {
	var binder = dba.GetISession().GetIBinder()
	var resMap = make(map[interface{}]interface{}, 0)
	var resSlice = make([]interface{}, 0)

	err = dba.Select()
	if err != nil {
		return
	}

	switch binder.GetBindType() {
	case OBJECT_MAP, OBJECT_MAP_T, OBJECT_STRUCT: // row
		var key, val interface{}
		if len(fieldKey) > 0 {
			key, err = dba.Value(fieldKey[0])
			if err != nil {
				return
			}
			val, err = dba.Value(field)
			if err != nil {
				return
			}
			resMap[key] = val
		} else {
			v, err = dba.Value(field)
			if err != nil {
				return
			}
		}
	case OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T:
		for _, item := range t.New(binder.GetBindResultSlice().Interface()).Slice() {
			val := item.MapInterfaceT()
			if len(fieldKey) > 0 {
				resMap[val[fieldKey[0]].Interface()] = val[field].Interface()
			} else {
				resSlice = append(resSlice, val[field].Interface())
			}
		}
	case OBJECT_STRUCT_SLICE: // rows
		var brs = binder.GetBindResultSlice()
		for i := 0; i < brs.Len(); i++ {
			val := reflect.Indirect(brs.Index(i))
			if len(fieldKey) > 0 {
				mapkey := dba._valueFromStruct(val, fieldKey[0])
				mapVal := dba._valueFromStruct(val, field)
				resMap[mapkey] = mapVal
			} else {
				resSlice = append(resSlice, dba._valueFromStruct(val, field))
			}
		}
	case OBJECT_STRING:
		res := dba.GetISession().GetBindAll()
		if len(res) > 0 {
			for _, val := range res {
				if len(fieldKey) > 0 {
					resMap[val[fieldKey[0]]] = val[field]
				} else {
					resSlice = append(resSlice, val[field])
				}
			}
		}
	}
	if len(fieldKey) > 0 {
		v = resMap
	} else {
		v = resSlice
	}
	return
}

// Type is get a row of a field value
func (dba *Orm) Value(field string) (v interface{}, err error) {
	res, err := dba.First()
	if v, ok := res[field]; ok {
		return v, err
	}
	return
}

// Type is get a row of a field value
func (dba *Orm) Column(field string) (v []interface{}, err error) {
	dba.fields = []string{field}
	res, err := dba.Get()
	if err != nil {
		return
	}
	v = []interface{}{}
	for _, re := range res {
		if vt, ok := re[field]; ok {
			v = append(v, vt)
		}
	}
	return
}

// Value_bak ...
func (dba *Orm) Value_bak(field string) (v interface{}, err error) {
	dba.Limit(1)
	err = dba.Select()
	if err != nil {
		return
	}
	var binder = dba.GetISession().GetIBinder()
	switch binder.GetBindType() {
	case OBJECT_MAP, OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T, OBJECT_MAP_T:
		v = reflect.ValueOf(binder.GetBindResult()).MapIndex(reflect.ValueOf(field)).Interface()
	case OBJECT_STRUCT, OBJECT_STRUCT_SLICE:
		bindResult := reflect.Indirect(reflect.ValueOf(binder.GetBindResult()))
		v = dba._valueFromStruct(bindResult, field)
	case OBJECT_STRING:
		res := dba.GetISession().GetBindAll()
		if len(res) > 0 {
			v = res[0][field]
		}
	}
	return
}
func (dba *Orm) _valueFromStruct(bindResult reflect.Value, field string) (v interface{}) {
	ostype := bindResult.Type()
	for i := 0; i < ostype.NumField(); i++ {
		tag := ostype.Field(i).Tag.Get(TAGNAME)
		if tag == field || ostype.Field(i).Name == field {
			v = bindResult.FieldByName(ostype.Field(i).Name).Interface()
		}
	}
	return
}

// Chunk : 分块处理数据,当要处理很多数据的时候, 我不需要知道具体是多少数据, 我只需要每次取limit条数据,
// 然后不断的增加offset去取更多数据, 从而达到分块处理更多数据的目的
func (dba *Orm) Chunk(limit int, callback func([]Data) error) (err error) {
	var page = 1
	var tabname = dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)
	where, fields, group := dba.where, dba.fields, dba.group

	// 先执行一条看看是否报错, 同时设置指定的limit, offset
	dba.Table(tabname2).Limit(limit).Page(page)
	if err = dba.Select(); err != nil {
		return
	}
	result := dba.GetBindAll()
	for len(result) > 0 {
		if err = callback(result); err != nil {
			break
		}
		page++
		// 清理绑定数据, 进行下一次操作, 因为绑定数据是每一次执行的时候都会解析并保存的
		// 而第二次以后执行的, 都会再次解析并保存, 数据结构是slice, 故会累积起来
		dba.ClearBindValues()
		dba.where, dba.fields, dba.group = where, fields, group
		dba.Table(tabname2).Limit(limit).Page(page)
		if err = dba.Select(); err != nil {
			break
		}
		result = dba.GetBindAll()
	}
	return
}

// ChunkWG : ChunkWG是保留Chunk的使用方法的基础上，新增多线程读取&多线程执行的方式，注意onetime_exec_thread不宜过多，推荐4，不宜过大因为采用的是盲读的方法，详情请参考github-wiki的介绍部分
// 原理与PaginatorWG方法类似，保留了事务隔离性并且加入了
func (dba *Orm) ChunkWG(onetime_exec_thread int, limit int, callback func([]Data) error) (err error) {
	if onetime_exec_thread <= 0 {
		onetime_exec_thread = 1
	}
	if onetime_exec_thread > 20 {
		onetime_exec_thread = 20
	}
	var page = 1
	tabname := dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)
	where, fields, group := dba.where, dba.fields, dba.group
	dba.Table(tabname2).Limit(limit).Page(page)

	if err = dba.Select(); err != nil {
		return
	}
	result := dba.GetBindAll()
	if len(result) < 1 {
		return
	}
	if err = callback(result); err != nil {
		return
	}
	continue_run := true
	for continue_run {
		var mp sync.Map
		for i := 0; i < onetime_exec_thread; i++ {
			page++
			dba.ClearBindValues()
			dba.Table(tabname2).
				Limit(limit).
				Page(page)
			dba.where, dba.fields, dba.group = where, fields, group
			sqlStr, args, err := dba.BuildSql()
			if err != nil {
				continue_run = false
				return err
			}
			mp.Store(sqlStr, args)
		}
		var wg sync.WaitGroup
		wg.Add(onetime_exec_thread)
		mp.Range(func(sqlStr, args interface{}) bool {
			go func(db *Orm, sql string, arg []interface{}) {
				_result, _err := db.Query(sql, arg...)
				if _err != nil {
					wg.Done()
					continue_run = false
					logger.Error(_err.Error())
					return
				}
				if len(_result) < 1 {
					wg.Done()
					continue_run = false
					return
				}
				if _err = callback(_result); _err != nil {
					wg.Done()
					continue_run = false
					return
				}
				wg.Done()
			}(dba, sqlStr.(string), args.([]interface{}))
			return true
		})
		wg.Wait()
	}
	return
}

// ChunkStruct : 同Chunk,只不过不用返回map, 而是绑定数据到传入的对象上
// 这里一定要传入绑定struct
func (dba *Orm) ChunkStruct(limit int, callback func() error) (err error) {
	var page = 0
	//var tableName = dba.GetISession().GetIBinder().GetBindName()
	// 先执行一条看看是否报错, 同时设置指定的limit, offset
	err = dba.Limit(limit).Offset(page * limit).Select()
	if err != nil {
		return
	}
	switch dba.GetIBinder().GetBindType() {
	case OBJECT_STRUCT, OBJECT_MAP, OBJECT_MAP_T:
		var ibinder = dba.GetIBinder()
		var result = ibinder.GetBindResult()
		for result != nil {
			if err = callback(); err != nil {
				break
			}
			page++
			// 清空结果
			//result = nil
			var rfRes = reflect.ValueOf(result)
			rfRes.Set(reflect.Zero(rfRes.Type()))
			// 清理绑定数据, 进行下一次操作, 因为绑定数据是每一次执行的时候都会解析并保存的
			// 而第二次以后执行的, 都会再次解析并保存, 数据结构是slice, 故会累积起来
			dba.ClearBindValues()
			_ = dba.Table(ibinder.GetBindOrigin()).Offset(page * limit).Select()
			result = dba.GetIBinder().GetBindResultSlice()
		}
	case OBJECT_STRUCT_SLICE, OBJECT_MAP_SLICE, OBJECT_MAP_SLICE_T:
		var ibinder = dba.GetIBinder()
		var result = ibinder.GetBindResultSlice()
		for result.Interface() != nil {
			if err = callback(); err != nil {
				break
			}
			page++
			// 清空结果
			result.Set(result.Slice(0, 0))
			// 清理绑定数据, 进行下一次操作, 因为绑定数据是每一次执行的时候都会解析并保存的
			// 而第二次以后执行的, 都会再次解析并保存, 数据结构是slice, 故会累积起来
			dba.ClearBindValues()
			_ = dba.Table(ibinder.GetBindOrigin()).Offset(page * limit).Select()
			result = dba.GetIBinder().GetBindResultSlice()
		}
	}
	return
}

// Loop : 同chunk, 不过, 这个是循环的取前limit条数据, 为什么是循环取这些数据呢
// 因为, 我们考虑到一种情况, 那就是where条件如果刚好是要修改的值,
// 那么最后的修改结果因为offset的原因, 只会修改一半, 比如:
// DB().Where("age", 18) ===> DB().Data(gorose.Data{"age":19}).Where().Update()
func (dba *Orm) Loop(limit int, callback func([]Data) error) (err error) {
	var page = 0
	var tabname = dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	tabname2 := strings.TrimPrefix(tabname, prefix)
	// 先执行一条看看是否报错, 同时设置指定的limit
	result, err := dba.Table(tabname2).Limit(limit).Get()
	if err != nil {
		return
	}
	for len(result) > 0 {
		if err = callback(result); err != nil {
			break
		}
		page++
		// 同chunk
		dba.ClearBindValues()
		result, _ = dba.Get()
	}
	return
}

// Paginate 自动分页
// @param limit 每页展示数量
// @param current_page 当前第几页, 从1开始
// 以下是laravel的Paginate返回示例
//
//	{
//		"total": 50,
//		"per_page": 15,
//		"current_page": 1,
//		"lastPage": 4,
//		"first_page_url": "http://laravel.app?page=1",
//		"lastPage_url": "http://laravel.app?page=4",
//		"nextPage_url": "http://laravel.app?page=2",
//		"prevPage_url": null,
//		"path": "http://laravel.app",
//		"from": 1,
//		"to": 15,
//		"data":[
//			{
//			// Result Object
//			},
//			{
//			// Result Object
//			}
//		]
//	}
func (dba *Orm) Paginate(page ...int) (res Data, err error) {
	if len(page) > 0 {
		dba.Page(page[0])
	}
	var limit = dba.GetLimit()
	if limit == 0 {
		limit = 15
	}
	var offset = dba.GetOffset()
	var currentPage = int(math.Ceil(float64(offset+1) / float64(limit)))
	//dba.ResetUnion()
	// 统计总量
	tabname := dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	where := dba.where
	resData, err := dba.Get()
	//fmt.Println(dba.LastSql())
	if err != nil {
		return
	}

	if resData == nil {
		resData = []Data{}
	}
	dba.offset = 0
	dba.GetISession().GetIBinder().SetBindName(tabname)
	dba.GetISession().GetIBinder().SetBindPrefix(prefix)
	dba.where = where
	count, err := dba.Count()
	//fmt.Println(dba.LastSql())
	if err != nil {
		return
	}
	var lastPage = int(math.Ceil(float64(count) / float64(limit)))
	var nextPage = currentPage + 1
	var prevPage = currentPage - 1
	// 获取结果

	res = Data{
		"total":          count,
		"per_page":       limit,
		"current_page":   currentPage,
		"last_page":      lastPage,
		"first_page_url": 1,
		"last_page_url":  lastPage,
		"next_page_url":  If(nextPage > lastPage, nil, nextPage),
		"prev_page_url":  If(prevPage < 1, nil, prevPage),
		//"data":          dba.GetISession().GetBindAll(),
		"data": resData,
	}
	return
}

// Paginate 自动分页
// @param limit 每页展示数量
// @param current_page 当前第几页, 从1开始
// 以下是laravel的Paginate返回示例
//
//	{
//		"total": 50,
//		"per_page": 15,
//		"current_page": 1,
//		"lastPage": 4,
//		"first_page_url": "http://laravel.app?page=1",
//		"lastPage_url": "http://laravel.app?page=4",
//		"nextPage_url": "http://laravel.app?page=2",
//		"prevPage_url": null,
//		"path": "http://laravel.app",
//		"from": 1,
//		"to": 15,
//		"data":[
//			{
//			// Result Object
//			},
//			{
//			// Result Object
//			}
//		]
//	}
func (dba *Orm) Paginator(page ...int) (res Paginate, err error) {
	if len(page) > 0 {
		dba.Page(page[0])
	}
	var limit = dba.GetLimit()
	if limit == 0 {
		limit = 15
	}
	var offset = dba.GetOffset()
	var currentPage = int(math.Ceil(float64(offset+1) / float64(limit)))
	//dba.ResetUnion()
	// 统计总量
	tabname := dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	where := dba.where
	fields := dba.fields
	resData, err1 := dba.Get()
	//fmt.Println(dba.LastSql())
	if err1 != nil {
		err = err1
		return
	}
	dba.offset = 0
	dba.fields = fields
	dba.GetISession().GetIBinder().SetBindName(tabname)
	dba.GetISession().GetIBinder().SetBindPrefix(prefix)
	dba.where = where
	count, err2 := dba.Counts()
	if err2 != nil {
		err = err2
		return
	}
	//fmt.Println(dba.LastSql())

	var lastPage = int(math.Ceil(float64(count) / float64(limit)))
	var nextPage = currentPage + 1
	var prevPage = currentPage - 1
	// 获取结果

	res = Paginate{
		Total:        count,
		PerPage:      limit,
		CurrentPage:  currentPage,
		LastPage:     lastPage,
		FirstPageUrl: 1,
		LastPageUrl:  lastPage,
		NextPageUrl:  If(nextPage > lastPage, nil, nextPage),
		PrevPageUrl:  If(prevPage < 1, nil, prevPage),
		//"data":          dba.GetISession().GetBindAll(),
		Data: resData,
	}
	return
}

// PaginatorWG this is a waitgroup Paginator function might be 20-50% faster than the original one,975ms->756ms
func (dba *Orm) PaginatorWG(page ...int) (res Paginate, err error) {

	if len(page) > 0 {
		dba.Page(page[0])
	}
	var limit = dba.GetLimit()
	if limit == 0 {
		limit = 15
	}
	var offset = dba.GetOffset()
	var currentPage = int(math.Ceil(float64(offset+1) / float64(limit)))
	//dba.ResetUnion()
	// 统计总量
	tabname := dba.GetISession().GetIBinder().GetBindName()
	prefix := dba.GetISession().GetIBinder().GetBindPrefix()
	where := dba.where
	fields := dba.fields

	var wg sync.WaitGroup
	wg.Add(2)

	var resData []Data
	var err1 error

	dba.GetIBinder().SetBindType(OBJECT_STRING)
	dba.ResetTable()
	dba.Table(strings.TrimPrefix(tabname, prefix))
	sqlStr, args, err := dba.BuildSql()
	if err != nil {
		return
	}
	go func(db *Orm, data *[]Data, errs1 *error, sqls *string, ags *[]interface{}) {
		// 执行查询
		*data, *errs1 = db.Query(*sqls, *ags...)
		wg.Done()
	}(dba, &resData, &err1, &sqlStr, &args)

	var count int64
	var err2 error
	go func(db *Orm, c *int64, errs2 *error) {
		db.offset = 0
		new_fields := []string{}
		for _, field := range fields {
			field_low := strings.ToLower(field)
			if strings.Contains(field_low, "count(") {
				new_fields = append(new_fields, field)
				break
			}
			if strings.Contains(field_low, "sum(") {
				new_fields = append(new_fields, field)
				break
			}
			if strings.Contains(field_low, "max(") {
				new_fields = append(new_fields, field)
				break
			}
			if strings.Contains(field_low, "min(") {
				new_fields = append(new_fields, field)
				break
			}
			if strings.Contains(field_low, "avg(") {
				new_fields = append(new_fields, field)
				break
			}
		}
		db.fields = new_fields
		db.GetISession().GetIBinder().SetBindName(tabname)
		db.GetISession().GetIBinder().SetBindPrefix(prefix)
		db.where = where
		*c, *errs2 = db.Counts()
		if *errs2 != nil {
			wg.Done()
			return
		}
		wg.Done()
	}(dba, &count, &err2)

	wg.Wait()
	if err1 != nil {
		err = err1
		return
	}
	if err2 != nil {
		err = err2
		return
	}

	//fmt.Println(dba.LastSql())

	var lastPage = int(math.Ceil(float64(count) / float64(limit)))
	var nextPage = currentPage + 1
	var prevPage = currentPage - 1
	// 获取结果

	res = Paginate{
		Total:        count,
		PerPage:      limit,
		CurrentPage:  currentPage,
		LastPage:     lastPage,
		FirstPageUrl: 1,
		LastPageUrl:  lastPage,
		NextPageUrl:  If(nextPage > lastPage, nil, nextPage),
		PrevPageUrl:  If(prevPage < 1, nil, prevPage),
		//"data":          dba.GetISession().GetBindAll(),
		Data: resData,
	}
	return
}
