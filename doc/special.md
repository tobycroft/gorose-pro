# GorosePro编程警告

## 正常使用中可能会出现的故障：
- 编程时没有加入限定参数（在GorosePro中修复）
  - WhereIn函数，需要[]any{}是有参数的情况下，原版框架在这里没有判断，这有助于提高"DejaVu性能"
  - 导致原因：
    - 可能数据库采用了直写编程思想，这有助于提高代码自释性，但是外部传入可能会传空
      - 会导致：[]any{}被传入,语句输出将变成select...where id in () ,这将导致数据库错误
- 分页错误：
  - 例如在大量数据查询时，第一页最后一条和第二页第一或者提二条重复的问题
    - 案例：如我需要做积分排行，我需要将一段时间内的新增学分累积起来并排行那么我正常将写一个带有join的读取方法
    - 这时如果排行仅使用积分排行，那么相同积分的情况下就会出现随机排行的情况，这就是为什么在跨页排行时会出现如此故障
    - 这个故障是MySQL的原理导致的，请各位开发者在编程时要特别注意
```go
func Api_paginator_byStudyTypeAndCoinIdAndSchoolIdAndYearAndClassAndDate(coin_id any, extra_like string, school_id, year, class, start_date, end_date any, groupby string, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(table + " a")
	db.Fields("a.*", "b.school_id", "b.year", "b.class", "sum(amount) as sum_amount", "b.name", "c.wx_name", "c.wx_img")
	db.LeftJoin(StudentModel.Table+" b", "a.student_id=b.id")
	db.LeftJoin(UserModel.Table+" c", "a.uid=c.id")
	if coin_id != nil {
		db.Where("coin_id", coin_id)
	}
	if extra_like != "" {
		db.Where("extra", "like", extra_like+".%")
	}
	if school_id != nil {
		db.Where("school_id", school_id)
	}
	if year != nil {
		db.Where("year", year)
	}
	if class != nil {
		db.Where("class", class)
	}
	db.Where("a.date", ">=", start_date)
	db.Where("a.date", "<", end_date)
	db.GroupBy(groupby)
	//错误这里需要在sum_amount的基础上加入id排行避免相同积分出现混乱排序
	//db.OrderBy("sum_amount desc")
	db.OrderBy("sum_amount desc,id desc")
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Paginator()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}
```
