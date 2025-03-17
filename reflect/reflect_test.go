package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect1(test *testing.T) {
	str := "hello world"
	// 变量 -> interface{} -> 反射变量
	v := reflect.ValueOf(str)
	_ = reflect.TypeOf(str)

	// 反射变量 -> interface{} -> 变量
	newStr := v.Interface().(string)
	fmt.Println(newStr)

	// 改变反射变量的值
	reflect.ValueOf(&str).Elem().Set(reflect.ValueOf("hello world2"))
	fmt.Println(str)
}

func TestReflect2(test *testing.T) {
	// 结构体变量
	type Person struct {
		Name string
		Age  int
	}
	person1 := Person{
		Name: "john",
		Age:  20,
	}

	v := reflect.ValueOf(person1)
	t := reflect.TypeOf(person1) // t = v.Type()

	fmt.Printf("person1 type: %v\n", v.Type()) // reflect.Person
	fmt.Printf("person1 kind: %v\n", t.Kind()) // struct

	if t.Kind() == reflect.Struct {
		fmt.Println
		for i := 0; i < t.NumField(); i++ {

		}
	}
}
