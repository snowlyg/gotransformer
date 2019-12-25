package gotransform

import (
	"fmt"
	"reflect"
	"strconv"
)

type XlxsTransform struct {
	OutputObj interface{}
	Title     map[string]string
	Row       []string
}

func NewXlxsTransform(outObj interface{}, title map[string]string, row []string) *XlxsTransform {
	return &XlxsTransform{
		OutputObj: outObj,
		Title:     title,
		Row:       row,
	}
}

func (t *XlxsTransform) GetOutputValue() reflect.Value {
	return reflect.ValueOf(t.OutputObj)
}

func (t *XlxsTransform) GetOutputValueKind() reflect.Kind {
	return t.GetOutputValue().Kind()
}

func (t *XlxsTransform) GetOutputValueElem() reflect.Value {
	return reflect.ValueOf(t.OutputObj).Elem()
}

func (t *XlxsTransform) GetOutputValueElemType() reflect.Type {
	return reflect.ValueOf(t.OutputObj).Elem().Type()
}

func (t *XlxsTransform) GetOutputValueElemField(i int) reflect.Value {
	return reflect.ValueOf(t.OutputObj).Elem().Field(i)
}

func (t *XlxsTransform) GetOutputValueElemTypeField(i int) reflect.StructField {
	return reflect.ValueOf(t.OutputObj).Elem().Type().Field(i)
}

func (t *XlxsTransform) XlxsTransformer() error {
	for i := 0; i < t.GetOutputValueElem().NumField(); i++ {
		of := t.GetOutputValueElemField(i)
		otf := t.GetOutputValueElemTypeField(i)
		for _, iw := range t.Title {
			if iw == otf.Name {
				err, rI := t.isExists(iw)
				if err != nil {
					return err
				}

				if rI != -1 && rI <= len(t.Row) {
					switch of.Kind() {
					case reflect.String:
						of.SetString(t.Row[rI])
					case reflect.Slice:
						// reflect.Copy(cf, t.Row[rI])
					case reflect.Bool:
						// cf.SetBool(t.Row[rI])
					case reflect.Float64, reflect.Float32:
						float, _ := strconv.ParseFloat(t.Row[rI], 64)
						of.SetFloat(float)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						atob, _ := strconv.Atoi(t.Row[rI])
						of.SetInt(int64(atob))
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						// cf.SetUint()
					default:
						fmt.Printf("数据类型错误:%v,%v", of.Kind(), of)
					}
				} else {
					continue
				}
			}
		}
	}

	return nil
}

// IsExists
func (t *XlxsTransform) isExists(s string) (error, int) {
	fv := filipValueString(t.Title)
	if _, ok := fv[s]; ok {
		i, err := strconv.Atoi(t.Title[s])
		if err != nil {
			return err, 0
		}
		return nil, i
	} else {
		return nil, -1
	}
}

// filpValue
func filipValueString(o map[string]string) map[string]string {
	for i, v := range o {
		o[v] = i
	}
	return o
}
