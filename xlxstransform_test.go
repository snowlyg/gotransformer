package gotransformer

import (
	"reflect"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func TestNewXlxsTransform(t *testing.T) {
	type args struct {
		outObj     interface{}
		title      map[string]string
		row        []string
		excelName  string
		timeFormat string
		file       *excelize.File
	}
	var tests []struct {
		name string
		args args
		want *XlxsTransform
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewXlxsTransform(tt.args.outObj, tt.args.title, tt.args.row, tt.args.excelName, tt.args.timeFormat, tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXlxsTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_GetExcelCell(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	type args struct {
		axis string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   string
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &XlxsTransform{
				OutputObj:  tt.fields.OutputObj,
				Title:      tt.fields.Title,
				Row:        tt.fields.Row,
				ExcelName:  tt.fields.ExcelName,
				File:       tt.fields.File,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetExcelCell(tt.args.axis); got != tt.want {
				t1.Errorf("GetExcelCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_XlxsCellTransformer(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &XlxsTransform{
				OutputObj:  tt.fields.OutputObj,
				Title:      tt.fields.Title,
				Row:        tt.fields.Row,
				ExcelName:  tt.fields.ExcelName,
				File:       tt.fields.File,
				TimeFormat: tt.fields.TimeFormat,
			}
			if err := t.XlxsCellTransformer(); (err != nil) != tt.wantErr {
				t1.Errorf("XlxsCellTransformer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXlxsTransform_XlxsTransformer(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &XlxsTransform{
				OutputObj:  tt.fields.OutputObj,
				Title:      tt.fields.Title,
				Row:        tt.fields.Row,
				ExcelName:  tt.fields.ExcelName,
				File:       tt.fields.File,
				TimeFormat: tt.fields.TimeFormat,
			}
			if err := t.XlxsTransformer(); (err != nil) != tt.wantErr {
				t1.Errorf("XlxsTransformer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXlxsTransform_isExists(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	type args struct {
		s string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  int
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &XlxsTransform{
				OutputObj:  tt.fields.OutputObj,
				Title:      tt.fields.Title,
				Row:        tt.fields.Row,
				ExcelName:  tt.fields.ExcelName,
				File:       tt.fields.File,
				TimeFormat: tt.fields.TimeFormat,
			}
			got, got1 := t.isExists(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("isExists() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t1.Errorf("isExists() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestXlxsTransform_parseTime(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	type args struct {
		cell string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    time.Time
		wantErr bool
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &XlxsTransform{
				OutputObj:  tt.fields.OutputObj,
				Title:      tt.fields.Title,
				Row:        tt.fields.Row,
				ExcelName:  tt.fields.ExcelName,
				File:       tt.fields.File,
				TimeFormat: tt.fields.TimeFormat,
			}
			got, err := t.parseTime(tt.args.cell)
			if (err != nil) != tt.wantErr {
				t1.Errorf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("parseTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filipValueString(t *testing.T) {
	type args struct {
		o map[string]string
	}
	var tests []struct {
		name string
		args args
		want map[string]string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filipValueString(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filipValueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
