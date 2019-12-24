# gotransform
golang 简单实现接口数据格式化

## Example

```go
package main

import (
"fmt"
"time"

"github.com/snowlyg/gotransform"
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
