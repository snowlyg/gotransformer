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
	tests := []struct {
		name string
		args args
		want *XlxsTransform
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
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

func TestXlxsTransform_GetOutputValue(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
			if got := t.GetOutputValue(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_GetOutputValueElem(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
			if got := t.GetOutputValueElem(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_GetOutputValueElemField(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
			if got := t.GetOutputValueElemField(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElemField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_GetOutputValueElemType(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Type
	}{
		// TODO: Add test cases.
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
			if got := t.GetOutputValueElemType(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElemType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_GetOutputValueElemTypeField(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   reflect.StructField
	}{
		// TODO: Add test cases.
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
			if got := t.GetOutputValueElemTypeField(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElemTypeField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlxsTransform_GetOutputValueKind(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		Title      map[string]string
		Row        []string
		ExcelName  string
		File       *excelize.File
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Kind
	}{
		// TODO: Add test cases.
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
			if got := t.GetOutputValueKind(); got != tt.want {
				t1.Errorf("GetOutputValueKind() = %v, want %v", got, tt.want)
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  int
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filipValueString(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filipValueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
