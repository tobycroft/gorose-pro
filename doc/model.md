## 示例model

```go
func Api_update_password(qq, password interface{}) bool {
	//初始化db并且填写table名称
	db := tuuz.Db().Table(table)
	//创建一个where，使用类型是map[string]interface，这个和TP的Array是一样的
	where := map[string]interface{}{
		"qq": qq,
	}
	//导入where
	db.Where(where)
	//这里如果是查询就db.dFind()或者db.Get()结束掉了，这里是修改，所以继续
	//使用同样的方法创建同类型数据map
	data := map[string]interface{}{
		"password": password,
	}
	//导入map，使用Data方法
	db.Data(data)
	//这里使用db.Update来执行修改，会输出数据库数字和错误
	_, err := db.Update()
	//判断错误
	if err != nil {
		//记录错误(如果你使用TuuzGoWeb就请记录，你也可使用自动记录，我更偏爱手动)
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
```


如上方法是一个修改方法，那么如果是输出单条或者多条怎么用呢？


单条
```go
func Api_find(qq interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
```
多条
```go
func Api_select_canShow() []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"can_show": 1,
	}
	db.Where(where)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}
```
