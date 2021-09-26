# Find/First查找单条数据的方法

基础方法已经在base里面展示过了，这里说个特殊的

如果大多数时候希望直接创建法，在做资金查询的时候使用注入法，要怎么做呢？

很简单，先把注入法写好，再写一个创建法，创建后注入Interface里面即可

```go
type Interface struct {
Db gorose.IOrm
}

func Api_find(group_id, user_id, dj_id interface{}) gorose.Data {
	var self Interface
	self.Db = tuuz.Db()
	return Api_find(group_id, user_id, dj_id)
}

func (self *Interface) Api_find(group_id, user_id, dj_id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	data, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}
```


find方法十分的简单，没有难度，相信大家会很快上手

find方法中没有太多骚操作，除此之外还有value拉sum拉之类的方法，大同小异

```go
func Api_value_balance(group_id, user_id interface{}) interface{} {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	ret, err := db.Value("balance")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
```



0.[基础准备](./base.md)

1.[select方法](./select.md)

3.[update方法](./update.md)

4.[delete方法](./delete.md)

5.[insert方法](./insert.md)

6.[安全相关](./security.md)
