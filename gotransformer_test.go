package gotransformer

import (
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	type args struct {
		o interface{}
	}
	var tests []struct {
		name string
		args args
		want reflect.Value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueElem(t *testing.T) {
	type args struct {
		o interface{}
	}
	var tests []struct {
		name string
		args args
		want reflect.Value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueElem(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueElemField(t *testing.T) {
	type args struct {
		o interface{}
		i int
	}
	var tests []struct {
		name string
		args args
		want reflect.Value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueElemField(tt.args.o, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueElemField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueElemType(t *testing.T) {
	type args struct {
		o interface{}
	}
	var tests []struct {
		name string
		args args
		want reflect.Type
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueElemType(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueElemType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueElemTypeField(t *testing.T) {
	type args struct {
		o interface{}
		i int
	}
	var tests []struct {
		name string
		args args
		want reflect.StructField
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueElemTypeField(tt.args.o, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueElemTypeField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueKind(t *testing.T) {
	type args struct {
		o interface{}
	}
	var tests []struct {
		name string
		args args
		want reflect.Kind
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueKind(tt.args.o); got != tt.want {
				t.Errorf("GetValueKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	var tests []struct {
		name   string
		fields fields
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
	var tests []struct {
		name   string
		fields fields
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
	var tests []struct {
		name string
		args args
		want bool
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
	var tests []struct {
		name string
		args args
		want float64
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
	var tests []struct {
		name string
		args args
		want int64
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
	var tests []struct {
		name string
		args args
		want string
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
	var tests []struct {
		name string
		args args
		want uint64
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
	var tests []struct {
		name string
		args args
		want *Tag
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
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
