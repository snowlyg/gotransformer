package gotransform

import (
	"reflect"
	"testing"
)

func TestNewTransform(t *testing.T) {
	type args struct {
		outObj     interface{}
		inObj      interface{}
		timeFormat string
	}
	var tests []struct {
		name string
		args args
		want *Transform
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransform(tt.args.outObj, tt.args.inObj, tt.args.timeFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_CallOutFunc(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		tag *Tag
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.CallOutFunc(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CallOutFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertValue(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertValue(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertValueElem(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertValueElem(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertValueElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertValueElemField(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		i int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertValueElemField(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertValueElemField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertValueElemType(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Type
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertValueElemType(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertValueElemType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertValueElemTypeField(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		i int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   reflect.StructField
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertValueElemTypeField(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertValueElemTypeField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertValueKind(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Kind
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertValueKind(); got != tt.want {
				t1.Errorf("GetInsertValueKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetOutputValue(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetOutputValue(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetOutputValueElem(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetOutputValueElem(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetOutputValueElemField(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		i int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   reflect.Value
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetOutputValueElemField(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElemField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetOutputValueElemType(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Type
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetOutputValueElemType(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElemType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetOutputValueElemTypeField(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		i int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   reflect.StructField
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetOutputValueElemTypeField(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetOutputValueElemTypeField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetOutputValueKind(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name   string
		fields fields
		want   reflect.Kind
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetOutputValueKind(); got != tt.want {
				t1.Errorf("GetOutputValueKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_Transformer(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if err := t.Transformer(); (err != nil) != tt.wantErr {
				t1.Errorf("Transformer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransform_getTag(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		otf reflect.StructField
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *Tag
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.getTag(tt.args.otf); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("getTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_setTime(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		inf       reflect.Value
		fieldName string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   string
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.setTime(tt.args.inf, tt.args.fieldName); got != tt.want {
				t1.Errorf("setTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_setValue(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		in  reflect.Value
		out reflect.Value
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
		})
	}
}
