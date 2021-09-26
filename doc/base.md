# 基础方法


以下两种方法将会贯穿我们的教程始终，请务必区分，根据实际情况来区分

Tuuz.Db()方法是我的框架中创建数据库的方法，你可以根据你自己的情况来，后续不再赘述

~~~
看教程中，请仔细看我的文字！！！
看教程中，请仔细看我的文字！！！
看教程中，请仔细看我的文字！！！
看教程中，请仔细看我的文字！！！
看教程中，请仔细看我的文字！！！
看教程中，请仔细看我的文字！！！

~~~

注意观察DB的生成！

每次查询都创建新的数据库对象

```go
func Api_find(group_id, user_id, cid interface{}) gorose.Data {
	//Tuuz.Db()方法是我的框架中创建数据库的方法，你可以根据你自己的情况来，后续不再赘述
	db := Tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
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

每次查询使用外部注入的数据库对象，请注意这里的type

```go
type Interface struct {
    Db gorose.IOrm
}


func (self *Interface) Api_find(group_id, user_id, cid interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
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
这两种方法需要在限定场景和非限定场景中选择

请大家务必根据需要来进行选择，这里推荐使用新建数据的方法，能达到最高效率

这里选择使用注入方法和新生成方法的原因是如果在Transaction或者Nested Transaction方法下，事务因为是隔离的
所以不在一个数据流（对象）下的查询是无法查询到隔离内的事务修改的，举个栗子

数据库中有一个字段A，里面只有一条数据，数据为123

例如在完全无事务的条件下流程是这样的：

A->find[输出123]->update(A,456)->find[输出456]

在find使用创建法，update导入法

A->开始事务->(不在事务中find[输出123])->update(A,456)->(不在事务中find[输出123])->提交事务->(不在事务中find[输出456])

find导入法，update导入法

A->开始事务->find[输出123]->update(A,456)->find[输出456]->提交事务->find[输出456]


明白了吧？如果你做的是交易所或者资金类的程序，这个概念你必须要搞懂！


1.[select方法](./select.md)

2.[find/first方法](./find.md)

3.[update方法](./update.md)

4.[delete方法](./delete.md)

5.[insert方法](./insert.md)

6.[安全相关](./security.md)
