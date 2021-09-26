# Select方法

```go
func Api_select(group_id, user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
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


如果你看完了上一个技术分享，那么select就很好理解了，select没啥复杂的，就是where然后get，没了

说个稍微复杂点的查询

```go
func Api_select_have(group_id, user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Where("num", ">", 0)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}

```
这里就是在where的基础上，再次where，你会发现不是所有的数据都是可以对应上的，很多时候需要做范围查询，
那么这个时候使用单一的where的map就无法正常的将代码写清楚了，那么这个时候，你就可以再次使用where的指定模式
来再次对范围内容进行限定，如上

好的，接下来来一个升级的

```go
func (self *Interface) Api_join_select(group_id, user_id interface{}) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Join("coin on coin.id=cid")
	db.Where("amount", ">", 0)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
```

这里来一个join方法，join方法是select方法中的骚操作，在Thinkphp中，有model方法来帮助我们做一对多多对一

那么用Gorose，we are on our own!难度不大，大家参考下就行了

使用数据库ORM很多时候需要我们自己处理数据，不过这也让我们对数据优化可以有更多的认识和提升


0.[基础准备](./base.md)

2.[find/first方法](./find.md)

3.[update方法](./update.md)

4.[delete方法](./delete.md)

5.[insert方法](./insert.md)

6.[安全相关](./security.md)
