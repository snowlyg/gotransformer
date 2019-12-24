package GoTransform

import (
	"fmt"
	"reflect"

	"BeeCustom/utils"
)

// 设置值
func Transform(outObj interface{}, inObj interface{}) {
	outObjValue := reflect.ValueOf(outObj)
	if outObjValue.Kind() == reflect.Struct { // 不是指针获取指针
		//outObjValue = reflect.ValueOf(&outObj)
	} else {
		transformDataKindIsStructOrPtr(outObjValue.Elem(), inObj)
	}

}

func transformDataKindIsStructOrPtr(outObjE reflect.Value, inObj interface{}) {
	outObjET := outObjE.Type()
	vaInObj := reflect.ValueOf(inObj)
	for i := 0; i < outObjE.NumField(); i++ {
		outObjEFValue := outObjE.Field(i)
		if !outObjEFValue.CanSet() {
			utils.LogDebug(fmt.Sprintf("数据无法修改:%v", outObjEFValue))
		}
		outObjETField := outObjET.Field(i)
		switch vaInObj.Kind() {
		case reflect.String:
			outObjEFValue.SetString(vaInObj.String())
		case reflect.Bool:
			outObjEFValue.SetBool(vaInObj.Bool())
		case reflect.Float64, reflect.Float32:
			outObjEFValue.SetFloat(vaInObj.Float())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			outObjEFValue.SetInt(vaInObj.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			outObjEFValue.SetUint(vaInObj.Uint())
		case reflect.Struct:
			vaInObj := reflect.ValueOf(&inObj) // 不是指针获取指针
			transformInnerDataKindIsStructOrPtr(vaInObj.Elem(), outObjETField, outObjEFValue)
		case reflect.Ptr:
			transformInnerDataKindIsStructOrPtr(vaInObj.Elem(), outObjETField, outObjEFValue)
		default:
			utils.LogDebug(fmt.Sprintf("数据类型错误:%v,%v", vaInObj.Kind(), vaInObj))
		}

	}
}

func transformInnerDataKindIsStructOrPtr(inObjE reflect.Value, outObjETField reflect.StructField, outObjEFValue reflect.Value) {
	inObjEType := inObjE.Type()
	for iI := 0; iI < inObjE.NumField(); iI++ {
		inObjEFValue := inObjE.Field(iI)
		inObjETypeField := inObjEType.Field(iI)
		if outObjETField.Name == inObjETypeField.Name {
			if outObjEFValue.Type() == inObjEFValue.Type() {
				switch inObjEFValue.Kind() {
				case reflect.String:
					outObjEFValue.SetString(inObjEFValue.String())
				case reflect.Bool:
					outObjEFValue.SetBool(inObjEFValue.Bool())
				case reflect.Float64, reflect.Float32:
					outObjEFValue.SetFloat(inObjEFValue.Float())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					outObjEFValue.SetInt(inObjEFValue.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					outObjEFValue.SetUint(inObjEFValue.Uint())
				case reflect.Struct:
					inObjEFValue = reflect.ValueOf(inObjEFValue)
					outObjEFValue = reflect.ValueOf(outObjEFValue)
					Transform(outObjEFValue, inObjEFValue)
				case reflect.Ptr:
					Transform(outObjEFValue, inObjEFValue)
				default:
					utils.LogDebug(fmt.Sprintf("数据类型错误:%v,%v", inObjEFValue.Kind(), inObjEFValue))
				}

			}
		} else if outObjETField.Name == "BaseTransform" && inObjETypeField.Name == "BaseModel" {
			inObjEFValue = reflect.ValueOf(inObjEFValue)
			outObjEFValue = reflect.ValueOf(outObjEFValue)
			Transform(outObjEFValue, inObjEFValue)
		}
	}
}
