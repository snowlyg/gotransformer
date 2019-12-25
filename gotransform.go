package gotransform

import (
	"errors"
	"fmt"
	"reflect"
)

type Transform struct {
	OutputObj  interface{}
	InsertObj  interface{}
	TimeFormat string
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
		// fmt.Sprintf("OutputType =》 %v , OutputValue => %v", otf, of)
		if !t.GetOutputValueElem().CanSet() {
			fmt.Sprintf("OutputType =》 %v , OutputValue => %v", otf, of)
			continue
		}

		for iI := 0; iI < t.GetInsertValueElem().NumField(); iI++ {
			inf := t.GetInsertValueElemField(iI)
			into := t.GetInsertValueElemTypeField(iI)
			// fmt.Sprintf("InsertType =》 %v , InsertValue => %v", into, inf)
			if otf.Name == into.Name {
				if inf.Type() == of.Type() {
					switch inf.Kind() {
					case reflect.String:
						of.SetString(inf.String())
					case reflect.Bool:
						of.SetBool(inf.Bool())
					case reflect.Float64, reflect.Float32:
						of.SetFloat(inf.Float())
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						of.SetInt(inf.Int())
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						of.SetUint(inf.Uint())
					default:
						fmt.Sprintf("数据类型错误:%v,%v", inf.Kind(), inf)
					}

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

func (t *Transform) setTime(inf reflect.Value, fieldName string) string {
	format := inf.FieldByName(fieldName).MethodByName("Format")
	m := []reflect.Value{reflect.ValueOf(t.TimeFormat)}
	createdAt := format.Call(m)

	return createdAt[0].Interface().(string)
}
