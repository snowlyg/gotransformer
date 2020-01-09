package main

import (
	"fmt"
	"html"
	"time"

	gtf "github.com/snowlyg/gotransformer"
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
	Name      string    `orm:"column(Name);null" description:""`
	LongText  string    `orm:"column(LongText);null" description:""`
	Status    string    `orm:"column(Status);null" description:""`
	Rmk       string    `orm:"column(rmk);size(255);null" description:""`
	Parent    *Parent   `orm:"null;rel(fk)"` // RelForeignKey relation
	Sons      []*Son    `orm:"null;rel(fk)"` // RelForeignKey relation
	Time      time.Time `form:"-" orm:"column(Time);type(timestamp);null" `
	Time1     time.Time `form:"-" orm:"column(Time1);type(timestamp);null" `
	DeletedAt time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

// 格式化数据
type Response struct {
	Id         int64
	Name       string
	LongText   string          `gtf:"Func.GetValueEnd"`               // 简单使用自定义方法格式化，使用 gtf 标识加`Func.funcName()`，默认第一个参数为转换数据本身
	Status     string          `gtf:"Func.GetUnitCode(LongText,1,2)"` // 使用带参数方法格式化(多个参数)，使用 gtf 标识加`Func.funcName(arg1,arg2,arg3)` 当有多个参数时候，自定义方法的第一个参数为需要转换数据的键名称
	ParentName string          `gtf:"Parent.Name"`                    // realtion 获取关联 Parent 的 Name
	SonNames   string          `gtf:"Func.FormatTime(Sons)"`          // realtion 获取关联 Sons 的 Names
	Sons       []*ResponseSon  // realtion 获取关联 Sons
	Parent     *ResponseParent // realtion 获取关联 Parent
	Time       string          `gtf:"Time.2006-01-02 15:04:05"` // 使用  2006-01-02 15:04:05 自定义格式化时间
	Time1      string          // 使用 `gotransform.NewTransform(&response, model, time.RFC3339)` 定义的时间格式
	Rmk        string
	DeletedAt  string
	CreatedAt  string
	UpdatedAt  string
}

// 格式化数据
type ResponseSon struct {
	Id        int64
	Name      string
	DeletedAt string
	CreatedAt string
	UpdatedAt string
}

// 格式化数据
type ResponseParent struct {
	Id        int64
	Name      string
	DeletedAt string
	CreatedAt string
	UpdatedAt string
}

type Parent struct {
	BaseModel
	Name string `orm:"column(Name);null" description:""`
}

type Son struct {
	BaseModel
	Name string `orm:"column(Name);null" description:""`
}

// realtion 获取关联 Sons 的 Names
func (r *Response) GetSonNames(vs []*Son) string {
	var ss string
	for _, v := range vs {
		ss += v.Name
	}

	return ss
}

/*
 * 自定义方法
 * 简单使用自定义方法格式化，使用 gtf 标识加`Func.funcName()`，
 * 默认第一个参数为转换数据本身 ,提示：反射函数调用会损耗性能不建议大量使用
 */
func (r *Response) GetValueEnd(v string) string {
	value := html.UnescapeString(v)
	valueEnd := value[:len(value)-1]
	if len(value) > 30 {
		valueEnd = value[:30] + "..."
	}

	return valueEnd
}

//使用带参数方法格式化(多个参数)，使用 gtf 标识加`Func.funcName(arg1,arg2,arg3)`
//当有多个参数时候，自定义方法的第一个参数为需要转换数据的键名称
//  `gtf:"Func.GetValueEnd"`  等同于 `gtf:"Func.GetValueEnd(LongText)"`
func (r *Response) GetUnitCode(v, t, s string) string {
	if v == t {
		return "激活"
	} else if v == s {
		return "关闭"
	} else {
		return "未知状态"
	}
}

func main() {

	baseModel := BaseModel{
		1,
		time.Now(),
		time.Now(),
	}
	parent := &Parent{
		BaseModel: baseModel,
		Name:      "name",
	}
	sons := []*Son{
		{
			BaseModel: baseModel,
			Name:      "name",
		},
	}
	longtext := "使用带参数方法格式化(多个参数)，使用 gtf 标识加`Func.funcName(arg1,arg2,arg3)"
	model := Model{
		BaseModel: baseModel,
		Name:      "name",
		Rmk:       "remark",
		LongText:  longtext,
		Status:    "1",
		Parent:    parent,
		Sons:      sons,
		DeletedAt: time.Now(),
	}

	// struct
	response := &Response{}
	g := gtf.NewTransform(response, model, time.RFC3339)
	err := g.Transformer()
	if err != nil {
		_ = fmt.Sprintf("err:%v", err)
	}

	// Parent 关联关系
	responseP := &ResponseParent{}
	gp := gtf.NewTransform(responseP, model.Parent, time.RFC3339)
	err = gp.Transformer()
	if err != nil {
		_ = fmt.Sprintf("err:%v", err)
	}
	response.Parent = responseP

	// Sons  关联关系
	var responseS []*ResponseSon
	for _, s := range model.Sons {
		r := &ResponseSon{}
		g1 := gtf.NewTransform(r, s, time.RFC3339)
		err := g1.Transformer()
		if err != nil {
			_ = fmt.Sprintf("err:%v", err)
		}

		responseS = append(responseS, r)
		_ = fmt.Sprintf("responses:%v", responseS)
	}
	response.Sons = responseS

	_ = fmt.Sprintf("response:%v", response)

	// slice
	models := []*Model{&model}
	var responses []*Response
	for _, m := range models {
		r := Response{}
		g1 := gtf.NewTransform(&r, m, time.RFC3339)
		err := g1.Transformer()
		if err != nil {
			_ = fmt.Sprintf("err:%v", err)
		}

		responses = append(responses, &r)
		_ = fmt.Sprintf("responses:%v", responses)
	}

}
