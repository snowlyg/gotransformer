package gotransformer

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
	tests := []struct {
		name string
		args args
		want *Transform
	}{
		// TODO: Add test cases.
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

func TestTransform_GetInsertMapKeys(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		want   []reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertMapKeys(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertMapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_GetInsertMapValue(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		key reflect.Value
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
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.GetInsertMapValue(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetInsertMapValue() = %v, want %v", got, tt.want)
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Type
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Kind
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Type
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   reflect.Kind
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
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

func TestTransform_setMapValue(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	type args struct {
		in  reflect.Value
		out reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
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
		inf        reflect.Value
		fieldName  string
		timeFormat string
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
			t := &Transform{
				OutputObj:  tt.fields.OutputObj,
				InsertObj:  tt.fields.InsertObj,
				TimeFormat: tt.fields.TimeFormat,
			}
			if got := t.setTime(tt.args.inf, tt.args.fieldName, tt.args.timeFormat); got != tt.want {
				t1.Errorf("setTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransform_transformerMap(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
		})
	}
}

func TestTransform_transformerPtr(t1 *testing.T) {
	type fields struct {
		OutputObj  interface{}
		InsertObj  interface{}
		TimeFormat string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
		})
	}
}

func Test_getMapValueB(t *testing.T) {
	type args struct {
		in reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMapValueB(tt.args.in); got != tt.want {
				t.Errorf("getMapValueB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMapValueF(t *testing.T) {
	type args struct {
		in reflect.Value
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMapValueF(tt.args.in); got != tt.want {
				t.Errorf("getMapValueF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMapValueI(t *testing.T) {
	type args struct {
		in reflect.Value
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMapValueI(tt.args.in); got != tt.want {
				t.Errorf("getMapValueI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMapValueS(t *testing.T) {
	type args struct {
		in reflect.Value
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMapValueS(tt.args.in); got != tt.want {
				t.Errorf("getMapValueS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMapValueU(t *testing.T) {
	type args struct {
		in reflect.Value
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMapValueU(tt.args.in); got != tt.want {
				t.Errorf("getMapValueU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTag(t *testing.T) {
	type args struct {
		otf reflect.StructField
	}
	tests := []struct {
		name string
		args args
		want *Tag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTag(tt.args.otf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setValue(t *testing.T) {
	type args struct {
		in  reflect.Value
		out reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
