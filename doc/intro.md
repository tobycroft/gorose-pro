# 先说思想

1.为什么开头都叫Api？叫别的行不行？

受到Golang的限制，首字母大写挎包能调用到，那么前面几个关键字就很重要了，
所以在设计数据库层的时候，这里使用的是单例模式，每个对数据的调用都是API的形式，
所以开头叫Api，你也可以改成别的

2.为什么又蛇形又大小写？

可恶心到很多人，其实蛇形的目的是为了把动作分开，

例如通过“username字段找1个用户”，Java或者PHP开发喜欢写GetUser，Py开发喜欢get_user这也没问题，
很多朋友是大神喜欢直接写目的，不过后续2开让人接手就很吐了，或者过1年看自己代码就难受了，我一开始开发也是这样的，
后来才开始使用“过程”代替“目的”的命名方式


给个示例：

例如:
~~~
Api_find_byQqandPassword()
~~~

~~~
1.这种命名方式请问他是数据库方法还是一个逻辑方法

2.他是增删改查中的哪一种？

3.他返回的是单条数据还是多条数据，我是使用map[string]interface{}来处理还是[]map[string]interface{}来处理？

4.这个方法主要需要哪两种数据类型才能完成查询？
~~~

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答

好的如上问题请在看下面的代码前回答



·

·

·

·


```go
func Api_find_byQqandPassword(qq, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq":       qq,
		"password": password,
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

好的，你在对照下Model代码，你看过程式命名方法是不是给后面接手你代码的人带来了很大的便利

写代码就像玩魂斗罗，要么一人跳太快给另一人给拖死，要么后面的太慢了把前面的拖死，如果你还没有决定项目规范，GorosePro的这个方案可以作为你的备选呢！


那么如果是展示呢？一般在读取数据的时候，find和select后面都应该跟着重点字段，例如这个数据输出，他是什么类型的
例如这里我需要把可展示的内容列出，那么这里的命名就变成了canShow，因为字段的问题，所以前小后大，请大家注意


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


如果这些案例你能看懂那么接下来这个案例就比较特别，需要使用群id，用户id，道具id来查询对应那条数据的num字段

命名是Api_value_num，你会发现在这里没有by了，因为数据太复杂了，请注意，虽然过程命名法可以让后来的人更方便的读懂你的项目，
知道你每个方法期待调用什么类型的数据，但是你可别把所有的查询限定项都写到方法名称中去，这么做虽然方便别人了，但是容易挨揍！

所以各位用户请掌握好度


```go
func (self *Interface) Api_value_num(group_id, user_id, dj_id interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	data, err := db.Value("num")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		if data == nil {
			return 0
		} else {
			return data.(int64)
		}
	}
}
```
同理还有

```go
func (self *Interface) Api_sum_byCid(cid interface{}) float64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"cid": cid,
	}
	db.Where(where)
	ret, err := db.Sum("amount")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		if ret != nil {
			return ret.(float64)
		} else {
			return 0
		}
	}
}
```

Tuuz.Db()方法是我的框架中创建数据库的方法，你可以根据你自己的情况来，后续不再赘述


以上是命名方法的解说

那么我们就要具体案例分析了：

0.[基础准备](./base.md)

1.[select方法](./select.md)

2.[find/first方法](./find.md)

3.[update方法](./update.md)

4.[delete方法](./delete.md)

5.[insert方法](./insert.md)

6.[安全相关](./security.md)

~~~
好的如果你能平心静气的看到这里你就会发现

下划线只是为了分割动作！！！命名方式依旧遵循了标准驼峰，蛇形只是表象而已

如果你心平气和的读到这里，那么恭喜你你是一个接受能力很强且不会刚愎自用的人

我真诚的邀请您加入我的群：94537310
~~~
