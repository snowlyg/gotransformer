package gotransformer

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Transform struct {
	OutputObj  interface{}
	InsertObj  interface{}
	TimeFormat string
}

type Tag struct {
	Key       string
	Value     string
	FiledName string
	Args      []string
}

func NewTransform(outObj, inObj interface{}, timeFormat string) *Transform {
	return &Transform{
		OutputObj:  outObj,
		InsertObj:  inObj,
		TimeFormat: timeFormat,
	}
}

// 输出数据的值
func (t *Transform) GetOutputValue() reflect.Value {
	return reflect.ValueOf(t.OutputObj)
}

// 输出数据的值类型
func (t *Transform) GetOutputValueKind() reflect.Kind {
	return t.GetOutputValue().Kind()
}

// 输出数据的值的成员
func (t *Transform) GetOutputValueElem() reflect.Value {
	return reflect.ValueOf(t.OutputObj).Elem()
}

// 输出数据的值的成员类型
func (t *Transform) GetOutputValueElemType() reflect.Type {
	return reflect.ValueOf(t.OutputObj).Elem().Type()
}

// 输出数据的值的成员的值
func (t *Transform) GetOutputValueElemField(i int) reflect.Value {
	return reflect.ValueOf(t.OutputObj).Elem().Field(i)
}

// 输出数据的值的成员类型的值
func (t *Transform) GetOutputValueElemTypeField(i int) reflect.StructField {
	return reflect.ValueOf(t.OutputObj).Elem().Type().Field(i)
}

// 输入数据的值
func (t *Transform) GetInsertValue() reflect.Value {
	return reflect.ValueOf(t.InsertObj)
}

// 输入 map类型数据的 keys
func (t *Transform) GetInsertMapKeys() []reflect.Value {
	return reflect.ValueOf(t.InsertObj).MapKeys()
}

// 输入map类型数据的 key 对应数数据
func (t *Transform) GetInsertMapValue(key reflect.Value) reflect.Value {
	return reflect.ValueOf(t.InsertObj).MapIndex(key)
}

// 输入数据的值类型
func (t *Transform) GetInsertValueKind() reflect.Kind {
	return t.GetInsertValue().Kind()
}

// 输入数据的值的成员
func (t *Transform) GetInsertValueElem() reflect.Value {
	return reflect.ValueOf(t.InsertObj).Elem()
}

// 输入数据的值的成员的类型
func (t *Transform) GetInsertValueElemType() reflect.Type {
	return reflect.ValueOf(t.InsertObj).Elem().Type()
}

// 输入数据的值的成员的值
func (t *Transform) GetInsertValueElemField(i int) reflect.Value {
	return reflect.ValueOf(t.InsertObj).Elem().Field(i)
}

// 输入数据的值的成员类型的值
func (t *Transform) GetInsertValueElemTypeField(i int) reflect.StructField {
	return reflect.ValueOf(t.InsertObj).Elem().Type().Field(i)
}

func (t *Transform) Transformer() error {

	if t.GetOutputValueKind() != reflect.Ptr {
		return errors.New("输出数据格式必须是指针")
	}

	if t.GetInsertValueKind() == reflect.Map {
		t.transformerMap()
	} else if t.GetInsertValueKind() == reflect.Ptr {
		t.transformerPtr()
	}

	return nil
}

// ptr 类型数据转换
func (t *Transform) transformerPtr() {
	for i := 0; i < t.GetOutputValueElem().NumField(); i++ {
		of := t.GetOutputValueElemField(i)
		otf := t.GetOutputValueElemTypeField(i)
		if !of.CanSet() {
			fmt.Printf("%v:不能被修改 \n", otf.Name)
			continue
		}

		tag := getTag(otf)
		timeFormat := ""
		for iI := 0; iI < t.GetInsertValueElem().NumField(); iI++ {
			inf := t.GetInsertValueElemField(iI)
			into := t.GetInsertValueElemTypeField(iI)
			if tag != nil {
				if tag.Key == "Time" {
					timeFormat = tag.Value
				}
				var args []reflect.Value
				startFunc := false
				// 执行自定义方法
				if len(tag.FiledName) < 1 { // 标签只定义了一个参数，则默认第一个参数为 inf.String()
					args = []reflect.Value{reflect.ValueOf(inf.String())} // append args,first arg is inf.string()
					startFunc = into.Name == otf.Name
				} else {
					startFunc = into.Name == tag.FiledName
					args = append(args, t.GetInsertValueElem().FieldByName(tag.FiledName))
					for _, vt := range tag.Args {
						args = append(args, reflect.ValueOf(vt))
					}
				}
				if tag.Key == "Func" && startFunc {
					rs := t.CallOutFunc(tag).Call(args)
					if rs[0].Interface() != nil {
						of.SetString(rs[0].Interface().(string))
						continue
					}
				}

				if inf.Kind() == reflect.Ptr { // for beego orm
					if into.Name == tag.Key {
						relation := inf.Elem().FieldByName(tag.Value)
						setValue(relation, of)
						continue
					}
				} else if inf.Kind() == reflect.Struct { // for gorm
					if into.Name == tag.Key {
						relation := inf.FieldByName(tag.Value)
						setValue(relation, of)
						continue
					}
				}
			}
			if into.Name == otf.Name {
				if into.Type.Name() == "Time" {
					of.SetString(t.setTime(inf, "", timeFormat))
					continue
				}

				if inf.Type() == of.Type() {
					setValue(inf, of)
					continue
				}
			}
			if into.Name == "BaseModel" { // BaseModel for beego orm
				if otf.Name == "Id" {
					of.SetInt(inf.FieldByName("Id").Interface().(int64))
					continue
				} else if otf.Name == "CreatedAt" {
					of.SetString(t.setTime(inf, "CreatedAt", timeFormat))
					continue
				} else if otf.Name == "UpdatedAt" {
					of.SetString(t.setTime(inf, "UpdatedAt", timeFormat))
					continue
				}
			} else if into.Name == "Model" { //Model for gorm
				if otf.Name == "Id" {
					of.SetInt(int64(inf.FieldByName("ID").Interface().(uint)))
					continue
				} else if otf.Name == "CreatedAt" {
					of.SetString(t.setTime(inf, "CreatedAt", timeFormat))
					continue
				} else if otf.Name == "UpdatedAt" {
					of.SetString(t.setTime(inf, "UpdatedAt", timeFormat))
					continue
				}
			}

		}
	}
}

// map 类型数据转换
func (t *Transform) transformerMap() {
	for i := 0; i < t.GetOutputValueElem().NumField(); i++ {
		of := t.GetOutputValueElemField(i)
		otf := t.GetOutputValueElemTypeField(i)
		if !of.CanSet() {
			fmt.Printf("%v:不能被修改 \n", otf.Name)
			continue
		}

		tag := getTag(otf)
		timeFormat := ""
		for _, k := range t.GetInsertMapKeys() {
			inf := t.GetInsertMapValue(k)
			keyName := k.String()

			if tag != nil {
				if tag.Key == "Time" {
					timeFormat = tag.Value
				}
				var args []reflect.Value
				startFunc := false
				// 执行自定义方法
				if len(tag.FiledName) < 1 { // 标签只定义了一个参数，则默认第一个参数为 inf.String()
					args = []reflect.Value{reflect.ValueOf(inf.Interface().(string))} // append args,first arg is inf.string()
					startFunc = keyName == otf.Name
				}

				if tag.Key == "Func" && startFunc {
					rs := t.CallOutFunc(tag).Call(args)
					if rs[0].Interface() != nil {
						of.SetString(rs[0].Interface().(string))
						continue
					}
				}

				if keyName == tag.Key {
					relation := inf.Elem().FieldByName(tag.Value)
					t.setMapValue(relation, of)
					continue
				}
			}

			if keyName == otf.Name {
				if otf.Type.Name() == "Time" {
					of.SetString(t.setTime(inf, "", timeFormat))
					continue
				}

				t.setMapValue(inf, of)
				continue
			}

		}
	}
}

// call out func
func (t *Transform) CallOutFunc(tag *Tag) reflect.Value {
	return t.GetOutputValue().MethodByName(tag.Value)
}

// set out map value
func (t *Transform) setMapValue(in, out reflect.Value) {
	switch out.Kind() {
	case reflect.String:
		out.SetString(getMapValueS(in))
	case reflect.Slice:
		reflect.Copy(out, in)
	case reflect.Bool:
		out.SetBool(getMapValueB(in))
	case reflect.Float64, reflect.Float32:
		out.SetFloat(getMapValueF(in))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out.SetInt(getMapValueI(in))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		out.SetUint(getMapValueU(in))
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
	}
}

// time format
func (t *Transform) setTime(inf reflect.Value, fieldName string, timeFormat string) string {

	if inf.IsZero() {
		return ""
	}
	if len(timeFormat) < 1 { // 自定义时间格式
		timeFormat = t.TimeFormat
	}
	args := []reflect.Value{reflect.ValueOf(timeFormat)}
	if len(fieldName) > 0 { // CreatedAt ,UpdatedAt in BaseModel
		inf = inf.FieldByName(fieldName)
	}
	f := inf.MethodByName("Format")
	rs := f.Call(args)

	return rs[0].Interface().(string)
}

// set out value
func setValue(in, out reflect.Value) {
	switch out.Kind() {
	case reflect.String:
		out.SetString(in.Interface().(string))
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
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
	}
}

// transform map data
func getMapValueS(in reflect.Value) string {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		return in.Interface().(string)
	case reflect.Bool:
		if in.Bool() {
			return "1"
		} else {
			return "0"
		}
	case reflect.Float64, reflect.Float32:
		return strconv.FormatFloat(in.Float(), 'e', 0, 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(in.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(in.Uint(), 10)
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return ""
	}

}

// transform map data
func getMapValueB(in reflect.Value) bool {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		if len(in.Interface().(string)) > 0 {
			return true
		} else {
			return false
		}
	case reflect.Bool:
		return in.Bool()
	case reflect.Float64, reflect.Float32:
		if in.Float() > 0 {
			return true
		} else {
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if in.Int() > 0 {
			return true
		} else {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if in.Uint() > 0 {
			return true
		} else {
			return false
		}
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return false
	}

}

// transform map data
func getMapValueF(in reflect.Value) float64 {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		f, _ := strconv.ParseFloat(in.Interface().(string), 64)
		return f
	case reflect.Float64, reflect.Float32:
		return in.Float()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(in.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(in.Uint())
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return 0
	}
}

// transform map data
func getMapValueI(in reflect.Value) int64 {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		i, _ := strconv.Atoi(in.Interface().(string))
		return int64(i)
	case reflect.Bool:
		if in.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.Float64, reflect.Float32:
		return int64(in.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return in.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(in.Uint())
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return 0
	}

}

// transform map data
func getMapValueU(in reflect.Value) uint64 {
	inTypeKind := in.Elem().Type().Kind()
	switch inTypeKind {
	case reflect.String:
		i, _ := strconv.ParseUint(in.Interface().(string), 10, 64)
		return i
	case reflect.Bool:
		if in.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.Float64, reflect.Float32:
		return uint64(in.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(in.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return in.Uint()
	default:
		fmt.Printf("数据类型错误:%v,%v  \n", in.Kind(), in)
		return 0
	}

}

// get tag
// first arg is insert object self
func getTag(otf reflect.StructField) *Tag {
	tag := otf.Tag.Get("gtf")
	rtag := strings.Replace(tag, ")", "", 1)
	names := strings.Split(rtag, ".")
	if len(names) <= 1 { // no func name
		return nil
	}

	if !strings.Contains(names[1], "(") { // no args
		return &Tag{Key: names[0], Value: names[1]}
	}

	arg := strings.Split(names[1], "(")
	if len(arg) == 1 { // no args
		return &Tag{Key: names[0], Value: arg[0]}
	}

	args := strings.Split(arg[1], ",") // all args
	if len(args) == 1 {
		return &Tag{Key: names[0], Value: arg[0], FiledName: args[0]}
	}

	return &Tag{Key: names[0], Value: arg[0], FiledName: args[0], Args: args[1:]}

}
