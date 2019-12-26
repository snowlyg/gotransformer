## Func Example
 
### Simple Func Example
简单使用自定义方法格式化，使用 gtf 标识加`Func.funcName`
 
```
// 基础数据模型  beego/orm 
type BaseModel struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);column(created_at);type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);column(updated_at);type(timestamp)"`
}

// 数据模型
type Model struct {
	BaseModel
	Name        string `orm:"column(Name);null" description:""`
    Parent       *Resource   `orm:"null;rel(fk)"`      // RelForeignKey relation
	Rmk          string    `orm:"column(rmk);size(255);null" description:""`
	DeletedAt    time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 格式化数据
type Response struct {
	Id           int64
    Value     string `gtf:"Func.GetValueEnd"`
    Rmk          string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}

//  自定义方法
func (r *Response) GetValueEnd(v string) string {
	value := html.UnescapeString(v)
	valueEnd := value[:len(value)-1]
	if len(value) > 30 {
		valueEnd = value[:30] + "..."
	}

	return valueEnd
}


```

 
### Func Example with Args
1.使用带参数方法格式化(一个参数)，使用 gtf 标识加`Func.funcName(arg)`
- 当只有一个参数时候，自定义方法需要两个参数。自定义方法的第一个参数为转换数据本身
 
```
// 基础数据模型  beego/orm 
type BaseModel struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);column(created_at);type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);column(updated_at);type(timestamp)"`
}

// 数据模型
type Model struct {
	BaseModel
	Value        string `orm:"column(Value);null" description:""`
	Rmk          string    `orm:"column(rmk);size(255);null" description:""`
	DeletedAt    time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 格式化数据
type Response struct {
	Id           int64
    Value     string `gtf:"Func.GetValueEnd(格式化数据)"`
    Rmk          string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}

//  自定义方法
func (r *Response) GetValueEnd(v,s string) string {
    // v 相当于 r.Value 的值
    // s 等于 “格式化数据”
	value := html.UnescapeString(v)
	valueEnd := value[:len(value)-1]
	if len(value) > 30 {
		valueEnd = value[:30] + s
	}

	return valueEnd
}

```

2.使用带参数方法格式化(多个参数)，使用 gtf 标识加`Func.funcName(arg1,arg2,arg3)`
- 当有多个参数时候，自定义方法的第一个参数为需要转换数据的键名称
 
```
// 基础数据模型  beego/orm 
type BaseModel struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);column(created_at);type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);column(updated_at);type(timestamp)"`
}

// 数据模型
type Model struct {
	BaseModel
	Value        string `orm:"column(Value);null" description:""`
	Rmk          string    `orm:"column(rmk);size(255);null" description:""`
	DeletedAt    time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 格式化数据
type Response struct {
	Id           int64
    Value        string 
    ValueEnd     string `gtf:"Func.GetValueEnd(Value,格式化数据1,格式化数据2)"`
    Rmk          string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}

//  自定义方法
func (r *Response) GetValueEnd(v,s1,s2 string) string {
    // v 相当于 r.Value 的值
    // s1 等于 “格式化数据1”
    // s2 等于 “格式化数据2”
	value := html.UnescapeString(v)
	valueEnd := value[:len(value)-1]
	if len(value) > 30 {
		valueEnd = value[:30] + s1 + s2
	}

	return valueEnd
}

```


3.时间数据格式，使用 gtf 标识加 `Func.FormatTime(arg)`
- 时间数据格式不需要写自定义方法，直接使用标识 `gtf:"Func.FormatTime()"`，`gtf:"Func.FormatTime("2006-01-02 15:04:05")"`
- 参数为可选，为空是使用 `gotransform.NewTransform(&response, model, time.RFC3339)` 定义的时间格式
 
```
// 基础数据模型  beego/orm 
type BaseModel struct {
	Id        int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);column(created_at);type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime);column(updated_at);type(timestamp)"`
}

// 数据模型
type Model struct {
	BaseModel
	Time        time.Time `form:"-" orm:"column(Time);type(timestamp);null" `
	Time1       time.Tim   `form:"-" orm:"column(Time1);type(timestamp);null" `
	DeletedAt   time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 格式化数据
type Response struct {
	Id           int64
    Time         string `gtf:"Func.FormatTime("2006-01-02 15:04:05")"`
    Time1        string `gtf:"Func.FormatTime()"`
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}
```

