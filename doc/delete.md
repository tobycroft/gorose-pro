# Delete方法


```go

func Api_delete(qq interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

```

很简单，没啥说的，没有where的时候delete会出错


0.[基础准备](./base.md)

1.[select方法](./select.md)

2.[find/first方法](./find.md)

3.[update方法](./update.md)

5.[insert方法](./insert.md)

6.[安全相关](./security.md)
