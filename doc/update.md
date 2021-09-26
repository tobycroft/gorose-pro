# Update方法


举两个例子，自己看，实在是很简单，通过XXX来查询XXX

这两个栗子，都是不需要事务处理的，因为应用场景中，只要if update成功后就可以直接显示成功了

因为场景非常简单，所以这里不需要使用“注入法”来使用数据库，直接使用创建法，便捷也快速


```go

func Api_update_uname(qq, uname interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]interface{}{
		"uname": uname,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
```

```go

func Api_update_all(qq, uname, password interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]interface{}{
		"uname":    uname,
		"password": password,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

```

接下来的这个栗子，是资金修改的，所以需要使用导入法

(self *Interface)这个没有特殊的意思，名字都可以自己取

请注意，一旦涉及到transaction必须使用导入法

如果涉及到嵌套事务，一定要使用导入法，否则事务和事务隔离均不能生效

```go
//Interface你自己取个名字
type Interface struct {
	//这里面的Db你也可以自己取名字
	Db gorose.IOrm
}
```

```go
//这里的self你也可以自己取名字
func (self *Interface) Api_update(group_id, user_id, balance interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"balance": balance,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
```

所以非常的简单的！如果你是PHP程序员相信你秒懂


这里说个特殊的Increasement方法，这个方法就很单纯，通过某个字段来递增，然后增量你可以自己定或者使用传入值

非常简单，这些值ORM都已经做成可接受interface数据了，所以只要你不作死去导入特殊值例如string之类的，
一般都没问题（就算传入特殊值ORM也会处理）

```go

func (self *Interface) Api_incr(group_id, user_id, cid, amount interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
	}
	db.Where(where)
	_, err := db.Increment("amount", amount)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

```
递减也很简单，你也可以使用负值传入递增实现递减的效果，记得ABS后在负数哦~不然给人家撸口小可爱干了你的系统

你别怪我框架的问题啊！

```go

func (self *Interface) Api_decr(group_id, user_id, dj_id interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	_, err := db.Decrement("num", 1)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
```


0.[基础准备](./base.md)

1.[select方法](./select.md)

2.[find/first方法](./find.md)

4.[delete方法](./delete.md)

5.[insert方法](./insert.md)

6.[安全相关](./security.md)
