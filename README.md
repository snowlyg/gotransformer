<h1 align="center">go-transform</h1>

<p align="center">
    <a href="https://travis-ci.org/snowlyg/gotransformer"><img src="https://travis-ci.org/snowlyg/gotransformer.svg?branch=master" alt="Build Status"></a>
    <a href="https://codecov.io/gh/snowlyg/gotransformer"><img src="https://codecov.io/gh/snowlyg/gotransformer/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/gotransformer"><img src="https://goreportcard.com/badge/github.com/snowlyg/gotransformer" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/gotransformer"><img src="https://godoc.org/github.com/snowlyg/gotransformer?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/gotransformer/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/gotransformer" alt="Licenses"></a>
</p>

#### 更新日志
[更新日志](UPDATE.MD)

#### Require
- go 1.13.x

#### 支持格式化方式
- 简单格式化
- 自定义方法格式化
- 关联数据格式化
- 时间数据格式化
- map数据格式化
- excel 导入数据格式化

## Installation

```
go get github.com/snowlyg/gotransformer

```


## Example
golang 简单实现接口数据格式化
```go
package main

import (
"fmt"
"time"

"github.com/snowlyg/gotransformer"
)

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
	Rmk          string    `orm:"column(rmk);size(255);null" description:""`
	DeletedAt    time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 格式化数据
type Response struct {
	Id           int64
	Name         string
    Rmk          string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}


func main()  {
    // struct
    response := Response{}
    baseModel := BaseModel{1,time.Now(),time.Now()}
    model := Model{baseModel,"name","remark",time.Now()}
    g := gotransform.NewTransform(&response, model, time.RFC3339)
    err := g.Transformer()
    if err != nil {
        _ = fmt.Sprintf("err:%v",err)
    }

    _ = fmt.Sprintf("response:%v",response)

    // slice
    models := []*Model{&model}
	var responses []*Response
	for _, m := range models {
		r := Response{}
		    g1 := gotransform.NewTransform(&r, m, time.RFC3339)
            err := g1.Transformer()
            if err != nil {
                _ = fmt.Sprintf("err:%v",err)
            }
           
		responses = append(responses, &r)
        _ = fmt.Sprintf("responses:%v",responses)
	}

}

```

## Simple Relation Example
简单关联关系格式化，使用 gtf 标识加`relationName.fieldName`
 
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
    Parent       *Parent   `orm:"null;rel(fk)"`      // RelForeignKey relation
	Rmk          string    `orm:"column(rmk);size(255);null" description:""`
	DeletedAt    time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

type Parent struct {
	BaseModel
	Name        string `orm:"column(Name);null" description:""`
}

// 格式化数据
type Response struct {
	Id           int64
	ParentName   string `gtf:"Parent.Name"`
    Rmk          string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}


```

## One To More Relation Example
一对多关联关系格式化，使用 gtf 标识加`Func.FormatTime(arg)`,参数为关联关系名称
 
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
    Parents      []*Parent   `orm:"null;rel(fk)"`      // RelForeignKey relation
	Rmk          string    `orm:"column(rmk);size(255);null" description:""`
	DeletedAt    time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 数据模型
type Parent struct {
	BaseModel
	Name        string `orm:"column(Name);null" description:""`
}

// 格式化数据
type Response struct {
	Id           int64
	ParentName   string `gtf:"Func.FormatTime(Parents)"`
    Rmk          string
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}

func (r *Response) GetAdminName(vs []*models.Parent) string {
	for _, v := range vs {
		//...
	}

	return ""
}

```

## 时间数据格式化
- 根据数据模型类型 time.Time 自动格式化，
- 默认使用 `gotransform.NewTransform(&response, model, time.RFC3339)` 定义的时间格式
- 自定义时间格式 使用标识 `gtf:"Time.2006-01-02 15:04:05"`
 
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
    Time         string `gtf:"Time.2006-01-02 15:04:05"`
    Time1        string 
	DeletedAt    string
	CreatedAt    string
	UpdatedAt    string
}
```


## Func Example
[Func Example](FUNC.md)

## Map Example
[Map Example](MAP.md)

 
## Excel 导入数据转换
[Excel 导入数据转换](XLSX.md)
