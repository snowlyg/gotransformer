package main

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	gft "github.com/snowlyg/gotransformer"
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
	FirstName string    `orm:"column(first_name);size(255);null" description:""`
	DeletedAt time.Time `form:"-" orm:"column(deleted_at);type(timestamp);null" `
}

func main() {

	titles := map[string]string{"0": "Name", "1": "FirstName"}

	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Excel 导入行数据转换
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows("Sheet1")
	// 提取 excel 数据
	ms := make([]*Model, 0)
	for roI, row := range rows {
		if roI > 0 {
			// 将数组  转成对应的 map
			m := Model{}
			x := gft.NewXlxsTransform(&m, titles, row, "", time.RFC3339, nil)
			err := x.XlxsTransformer()
			if err != nil {
				fmt.Println(err)
				return
			}

			ms = append(ms, &m)
		}
	}
	fmt.Println(ms)

	//Excel 导入单元格数据转换
	titlesCell := map[string]string{"Name": "B1", "FirstName": "B2"}
	m1 := Model{}
	x := gft.NewXlxsTransform(&m1, titlesCell, nil, "sheet1", time.RFC3339, f)
	err = x.XlxsCellTransformer()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m1)

}
