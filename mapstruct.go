package property

import (
	"errors"
	"reflect"
	"strconv"
)

type Decoder struct {
	Kind    reflect.Kind
	IsValid func(string) bool
	Set     func(value reflect.Value, valueStr string) error
}

var methodMap map[reflect.Kind]Decoder

func init() {
	methodMap = make(map[reflect.Kind]Decoder)
	methodMap[reflect.Bool] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			v.SetBool(valueStr == "true")
			return nil
		},
	}
	methodMap[reflect.Int] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			i, err := strconv.Atoi(valueStr)
			if err != nil {
				return err
			}
			v.SetInt(int64(i))
			return nil
		},
	}
	methodMap[reflect.Int8] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			i, err := strconv.Atoi(valueStr)
			if err != nil {
				return err
			}
			v.SetInt(int64(i))
			return nil
		},
	}
	methodMap[reflect.Int16] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			i, err := strconv.Atoi(valueStr)
			if err != nil {
				return err
			}
			v.SetInt(int64(i))
			return nil
		},
	}
	methodMap[reflect.Int32] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			i, err := strconv.Atoi(valueStr)
			if err != nil {
				return err
			}
			v.SetInt(int64(i))
			return nil
		},
	}
	methodMap[reflect.Int64] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			i, err := strconv.Atoi(valueStr)
			if err != nil {
				return err
			}
			v.SetInt(int64(i))
			return nil
		},
	}
	methodMap[reflect.Float32] = Decoder{
		Set: func(value reflect.Value, valueStr string) error {
			floatValue, err := strconv.ParseFloat(valueStr, 32)
			if err != nil {
				return err
			}
			value.SetFloat(floatValue)
			return nil
		},
	}
	methodMap[reflect.Float64] = Decoder{
		Set: func(value reflect.Value, valueStr string) error {
			floatValue, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return err
			}
			value.SetFloat(floatValue)
			return nil
		},
	}
	methodMap[reflect.String] = Decoder{
		Set: func(v reflect.Value, valueStr string) error {
			v.SetString(valueStr)
			return nil
		},
	}
}

func Unmarshal(structurePointer interface{}, valueMap map[string]string) error {
	t := reflect.TypeOf(structurePointer)
	v := reflect.ValueOf(structurePointer)
	for i := 0; i < v.Elem().NumField(); i++ {
		kind := v.Elem().Field(i).Kind()
		var value string
		var key string
		if tagName := t.Elem().Field(i).Tag.Get("propName"); tagName != "" {
			key = tagName
		} else {
			key = t.Elem().Field(i).Name
		}
		value, ok := valueMap[key]
		if !ok {
			return errors.New("no such key " + key)
		}
		if decoder, ok := methodMap[kind]; !ok {
			return errors.New("no decoder for variable of kind " + v.Elem().Field(i).Kind().String())
		} else {
			if err := decoder.Set(v.Elem().Field(i), value); err != nil {
				return err
			}
		}
	}
	return nil
}
