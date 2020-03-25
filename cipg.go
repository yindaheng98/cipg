package cipg

import (
	"flag"
	"fmt"
	"reflect"
	"time"
)

//Generate options from command line.
func Generate(opt interface{}, logger func(i ...interface{})) {
	generateValue(reflect.ValueOf(opt), "", "")
	flag.Parse()
	printValue(reflect.ValueOf(opt), "", logger)
	return
}

//通过反射获取设置信息
func generateValue(Value reflect.Value, ValueName, ValueUsage string) {
	operateValue(Value, ValueName, ValueUsage,
		func(Value reflect.Value, ValueName, ValueUsage string) {
			if Value.Type() == reflect.TypeOf(time.Duration(0)) {
				pointer := Value.Addr().Interface().(*time.Duration)
				flag.DurationVar(pointer, ValueName, *pointer, ValueUsage)
				return
			}
			switch Value.Kind() {
			case reflect.Bool:
				pointer := Value.Addr().Interface().(*bool)
				flag.BoolVar(pointer, ValueName, *pointer, ValueUsage)
			case reflect.Int:
				pointer := Value.Addr().Interface().(*int)
				flag.IntVar(pointer, ValueName, *pointer, ValueUsage)
			case reflect.Int64:
				pointer := Value.Addr().Interface().(*int64)
				flag.Int64Var(pointer, ValueName, *pointer, ValueUsage)
			case reflect.Uint:
				pointer := Value.Addr().Interface().(*uint)
				flag.UintVar(pointer, ValueName, *pointer, ValueUsage)
			case reflect.Uint64:
				pointer := Value.Addr().Interface().(*uint64)
				flag.Uint64Var(pointer, ValueName, *pointer, ValueUsage)
			case reflect.String:
				pointer := Value.Addr().Interface().(*string)
				flag.StringVar(pointer, ValueName, *pointer, ValueUsage)
			case reflect.Float64:
				pointer := Value.Addr().Interface().(*float64)
				flag.Float64Var(pointer, ValueName, *pointer, ValueUsage)
			}
		})
}

func printValue(Value reflect.Value, ValueName string, logger func(...interface{})) {
	operateValue(Value, ValueName, "",
		func(Value reflect.Value, ValueName, ValueUsage string) {
			p := func(TypeStr string) {
				logger(ValueName, fmt.Sprintf("(%s) =", TypeStr), Value.Interface())
			}
			if Value.Type() == reflect.TypeOf(time.Duration(0)) {
				p("time.Duration")
				return
			}
			switch Value.Type().Kind() {
			case reflect.Bool:
				p("bool")
			case reflect.Int:
				p("int")
			case reflect.Int64:
				p("int64")
			case reflect.Uint:
				p("uint")
			case reflect.Uint64:
				p("uint64")
			case reflect.String:
				p("string")
			case reflect.Float64:
				p("float64")
			}
		})
}

func operateValue(Value reflect.Value, ValueName, ValueUsage string,
	operation func(Value reflect.Value, ValueName, ValueUsage string)) {
	Value = Value.Elem()
	Type := Value.Type()
	for i := 0; i < Value.NumField(); i++ { // 遍历结构体所有成员
		field := Type.Field(i) // 获取每个成员的结构体字段

		fieldName := field.Name //获取字段名称
		if ValueName != "" {
			fieldName = ValueName + "." + field.Name
		}
		fieldTag := field.Tag.Get("usage")
		if ValueUsage != "" {
			fieldTag = ValueUsage + " " + fieldTag
		}

		fieldValue := Value.Field(i)             //获取每个成员的结构体字段值
		if field.Type.Kind() == reflect.Struct { //如果还是一个结构体
			operateValue(fieldValue.Addr(), fieldName, fieldTag, operation) //就递归
		} else { //不是结构体，那么执行操作
			operation(fieldValue, fieldName, fieldTag)
		}
	}
}
