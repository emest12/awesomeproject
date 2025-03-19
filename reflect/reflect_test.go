package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// Type和Value类型的变量，都被称为反射对象

/*
1. 从 interface{} 变量可以反射出反射对象；
2. 从反射对象可以获取 interface{} 变量；
3. 要修改反射对象，其值必须可设置；
*/

// 变量 interface{} 反射对象 三者相互转化
func TestReflect1(test *testing.T) {
	str := "hello world"
	// 变量 -> interface{} -> 反射对象
	v := reflect.ValueOf(str)
	_ = reflect.TypeOf(str)

	// 反射对象 -> interface{} -> 变量
	newStr := v.Interface().(string)
	fmt.Println(newStr)

	// 改变反射对象的值
	reflect.ValueOf(&str).Elem().Set(reflect.ValueOf("hello world2"))
	fmt.Println(str)
}

func TestReflect2(test *testing.T) {
	// 结构体变量
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person1 := &Person{
		Name: "john",
		Age:  20,
	}

	v := reflect.ValueOf(person1)
	t := reflect.TypeOf(person1) // t = v.Type()

	fmt.Printf("person1 type: %v\n", v.Type()) // *reflect.Person
	fmt.Printf("person1 kind: %v\n", t.Kind()) // ptr

	fmt.Printf("person1 type: %v\n", v.Elem().Type()) // reflect.Person
	fmt.Printf("person1 kind: %v\n", t.Elem().Kind()) // struct

	test.Log(v.CanSet())
	test.Log(v.Elem().CanSet())

	// value和type都有FieldByName方法
	nameField, ok := t.Elem().FieldByName("Name")
	if ok {
		test.Log(nameField.Name, nameField.Type, nameField.Tag)
	} else {
		test.Log("Name Field not found")
	}

	_, ok = v.Elem().FieldByName("Name").Interface().(string)
	if ok {
		v.Elem().FieldByName("Name").SetString("jack")
	}
	test.Log(person1)
}
