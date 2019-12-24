package gotransform

import (
	"errors"
	"fmt"
	"reflect"

	"BeeCustom/utils"
)

type Transform struct {
	OutputObj interface{}
	InsertObj interface{}
}

func NewTransform(outObj, inObj interface{}) *Transform {
	return &Transform{
		OutputObj: outObj,
		InsertObj: inObj,
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

func (t *Transform) Transformer() error {
	if t.GetOutputValueKind() != reflect.Ptr {
		return errors.New("输出数据格式必须是指针")
	}

	for i := 0; i < t.GetOutputValueElem().NumField(); i++ {
		of := t.GetOutputValueElem().Field(i)
		otf := t.GetOutputValueElemType().Field(i)
		if !t.GetOutputValueElem().CanSet() {
			utils.LogDebug(fmt.Sprintf("数据无法修改：  Type =》 %v , Value => %v", otf, of))
			continue
		}

		if t.GetInsertValueKind() == reflect.Struct {
			utils.LogDebug(fmt.Sprintf("数据无法修改：  Type =》 %v , Value => %v", t.GetInsertValueElemType(), t.GetInsertValue))
		}

		// itf := t.GetInsertValueElemType().Field(i)
		for iI := 0; iI < t.GetInsertValueElem().NumField(); iI++ {
			inf := t.GetInsertValueElem().Field(iI)
			intf := t.GetInsertValueElemType().Field(iI)
			if otf.Name == intf.Name {
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
					// case reflect.Struct:
					// 	inf = reflect.ValueOf(inf)
					// 	of = reflect.ValueOf(of)
					// 	Transform(of, inf)
					// case reflect.Ptr:
					// 	Transform(of, inf)
					default:
						utils.LogDebug(fmt.Sprintf("数据类型错误:%v,%v", inf.Kind(), inf))
					}

				}
			} else if otf.Name == "Id" && intf.Name == "BaseModel" {
				// inf = reflect.ValueOf(&inf)
				// of = reflect.ValueOf(of)
				// id := of.Elem().FieldByName("Id").Elem().Int()
				utils.LogDebug(of.Elem().FieldByName("Id").Elem())
				utils.LogDebug(of.Elem().FieldByName("Id"))
				utils.LogDebug(of.Elem())
				utils.LogDebug(of)
				utils.LogDebug(inf)
				// utils.LogDebug(id)
				// of.SetInt(id)
			}
		}
	}

	return nil
}
