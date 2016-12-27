package property

import (
	"reflect"
	"strconv"
	"errors"
)
type Decoder struct {
	Kind reflect.Kind
	IsValid func(string) bool
	Set func(value reflect.Value,valueStr string)
}

var methodMap map[reflect.Kind]Decoder
func init() {
	methodMap = make(map[reflect.Kind]Decoder)
	methodMap[reflect.Bool] = Decoder{
		Set:func(v reflect.Value, valueStr string) {
			v.SetBool(valueStr=="true")
		},
	}
	methodMap[reflect.Int] = Decoder{
		Set:func(v reflect.Value, valueStr string) {
			i, err:= strconv.Atoi(valueStr)
			if err!= nil {
				panic(err)
			}
			v.SetInt(int64(i))
		},
	}
	methodMap[reflect.Int8] = Decoder{
		Set:func(v reflect.Value, valueStr string) {
			i, err:= strconv.Atoi(valueStr)
			if err!= nil {
				panic(err)
			}
			v.SetInt(int64(i))
		},
	}
	methodMap[reflect.Int16] = Decoder{
		Set:func(v reflect.Value, valueStr string) {
			i, err:= strconv.Atoi(valueStr)
			if err!= nil {
				panic(err)
			}
			v.SetInt(int64(i))
		},
	}
	methodMap[reflect.Int32] = Decoder{
		Set:func(v reflect.Value, valueStr string) {
			i, err:= strconv.Atoi(valueStr)
			if err!= nil {
				panic(err)
			}
			v.SetInt(int64(i))
		},
	}
	methodMap[reflect.Int64] = Decoder{
		Set:func(v reflect.Value, valueStr string) {
			i, err:= strconv.Atoi(valueStr)
			if err!= nil {
				panic(err)
			}
			v.SetInt(int64(i))
		},
	}
	methodMap[reflect.Float32]=Decoder{
		Set:func(value reflect.Value, valueStr string) {
			floatValue,err := strconv.ParseFloat(valueStr,32)
			if err!=nil {
				panic(err)
			}
			value.SetFloat(floatValue)
		},
	}
	methodMap[reflect.Float64]=Decoder{
		Set:func(value reflect.Value, valueStr string) {
			floatValue,err := strconv.ParseFloat(valueStr,64)
			if err!=nil {
				panic(err)
			}
			value.SetFloat(floatValue)
		},
	}
	methodMap[reflect.String]=Decoder{
		Set:func(v reflect.Value, valueStr string) {
			v.SetString(valueStr)
		},
	}
}

func Unmarshal(structurePointer interface{}, valueMap map[string]string) error {
	t := reflect.TypeOf(structurePointer)
	v := reflect.ValueOf(structurePointer)
	for i:=0;i<v.Elem().NumField();i++ {
		kind := v.Elem().Field(i).Kind()
		var value string
		var key string
		if tagName := t.Elem().Field(i).Tag.Get("propName"); tagName!="" {
			key = tagName
		}else {
			key = t.Elem().Field(i).Name
		}
		value, ok := valueMap[key]
		if !ok {
			return errors.New("no such key "+key)
		}
		if decoder, ok := methodMap[kind]; !ok{
			return errors.New("no decoder for variable of kind "+ v.Elem().Field(i).Kind().String())
		}else{
			decoder.Set(v.Elem().Field(i), value)
		}
	}
	return nil
}
