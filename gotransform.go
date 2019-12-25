package gotransform

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Transform struct {
	OutputObj  interface{}
	InsertObj  interface{}
	TimeFormat string
}

type Tag struct {
	Key   string
	Value string
}

func NewTransform(outObj, inObj interface{}, timeFormat string) *Transform {
	return &Transform{
		OutputObj:  outObj,
		InsertObj:  inObj,
		TimeFormat: timeFormat,
	}
}

func (t *Transform) GetOutputValue() reflect.Value {
	return reflect.ValueOf(t.OutputObj)
}

func (t *Transform) GetOutputValueKind() reflect.Kind {
	return t.GetOutputValue().Kind()
}

func (t *Transform) GetOutputValueElem() reflect.Value {
	return reflect.ValueOf(t.OutputObj).Elem()
}

func (t *Transform) GetOutputValueElemType() reflect.Type {
	return reflect.ValueOf(t.OutputObj).Elem().Type()
}

func (t *Transform) GetOutputValueElemField(i int) reflect.Value {
	return reflect.ValueOf(t.OutputObj).Elem().Field(i)
}

func (t *Transform) GetOutputValueElemTypeField(i int) reflect.StructField {
	return reflect.ValueOf(t.OutputObj).Elem().Type().Field(i)
}

func (t *Transform) GetInsertValue() reflect.Value {
	return reflect.ValueOf(t.InsertObj)
}

func (t *Transform) GetInsertValueKind() reflect.Kind {
	return t.GetInsertValue().Kind()
}

func (t *Transform) GetInsertValueElem() reflect.Value {
	return reflect.ValueOf(t.InsertObj).Elem()
}

func (t *Transform) GetInsertValueElemType() reflect.Type {
	return reflect.ValueOf(t.InsertObj).Elem().Type()
}

func (t *Transform) GetInsertValueElemField(i int) reflect.Value {
	return reflect.ValueOf(t.InsertObj).Elem().Field(i)
}

func (t *Transform) GetInsertValueElemTypeField(i int) reflect.StructField {
	return reflect.ValueOf(t.InsertObj).Elem().Type().Field(i)
}

func (t *Transform) Transformer() error {

	if t.GetOutputValueKind() != reflect.Ptr {
		return errors.New("输出数据格式必须是指针")
	}

	for i := 0; i < t.GetOutputValueElem().NumField(); i++ {
		of := t.GetOutputValueElemField(i)
		otf := t.GetOutputValueElemTypeField(i)
		// fmt.Printf("OutputType =》 %v , OutputValue => %v", otf, of)
		if !t.GetOutputValueElem().CanSet() {
			fmt.Printf("OutputType =》 %v , OutputValue => %v", otf, of)
			continue
		}

		tag := t.getTag(otf)
		for iI := 0; iI < t.GetInsertValueElem().NumField(); iI++ {
			inf := t.GetInsertValueElemField(iI)
			into := t.GetInsertValueElemTypeField(iI)
			if tag != nil {
				if tag.Key == "Func" && into.Name == otf.Name {
					args := []reflect.Value{reflect.ValueOf(inf.String())}
					rs := t.CallOutFunc(tag).Call(args)
					of.SetString(rs[0].Interface().(string))
				} else if inf.Kind() == reflect.Ptr {
					if into.Name == tag.Key {
						relation := inf.Elem().FieldByName(tag.Value)
						t.setValue(relation, of)
					}
				}
			} else if into.Name == otf.Name {
				if inf.Type() == of.Type() {
					t.setValue(inf, of)
				}
			} else if into.Name == "BaseModel" {
				if otf.Name == "Id" {
					of.SetInt(inf.FieldByName("Id").Interface().(int64))
				} else if otf.Name == "CreatedAt" {
					createdAt := t.setTime(inf, "CreatedAt")
					of.SetString(createdAt)
				} else if otf.Name == "UpdatedAt" {
					createdAt := t.setTime(inf, "UpdatedAt")
					of.SetString(createdAt)
				}
			}

		}
	}

	return nil
}

// call out func
func (t *Transform) CallOutFunc(tag *Tag) reflect.Value {
	return t.GetOutputValue().MethodByName(tag.Value)
}

// set out value
func (t *Transform) setValue(in reflect.Value, out reflect.Value) {
	switch in.Kind() {
	case reflect.String:
		out.SetString(in.String())
	case reflect.Slice:
		reflect.Copy(out, in)
	case reflect.Bool:
		out.SetBool(in.Bool())
	case reflect.Float64, reflect.Float32:
		out.SetFloat(in.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out.SetInt(in.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		out.SetUint(in.Uint())
	default:
		fmt.Printf("数据类型错误:%v,%v", in.Kind(), in)
	}
}

// get tag
func (t *Transform) getTag(otf reflect.StructField) *Tag {
	tag := otf.Tag.Get("gtf")
	names := strings.Split(tag, ".")
	if len(names) > 1 {
		return &Tag{Key: names[0], Value: names[1]}
	}

	return nil
}

// time format
func (t *Transform) setTime(inf reflect.Value, fieldName string) string {
	args := []reflect.Value{reflect.ValueOf(t.TimeFormat)}
	f := inf.FieldByName(fieldName).MethodByName("Format")
	rs := f.Call(args)
	return rs[0].Interface().(string)
}
