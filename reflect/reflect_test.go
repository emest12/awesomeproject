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

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 结构体 反射 FieldByName 指针 Elem()
func TestReflect2(test *testing.T) {
	// 结构体变量
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

	test.Log(v.CanSet())        // false
	test.Log(v.Elem().CanSet()) // true

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

// 反射 指针 elem()
func TestReflect3(t *testing.T) {
	i1 := 1
	i2 := &i1
	v := reflect.ValueOf(&i2)
	v.Elem().Elem().SetInt(10)
	t.Log(i1)
}

// reflect.New 创建该类型的新指针对象
func TestReflect4(t *testing.T) {
	gwh := Person{
		Name: "gwh",
		Age:  18,
	}

	v := reflect.New(reflect.TypeOf(gwh))
	t.Log(v.Type()) // *reflect.Person
	t.Log(v.Kind()) // ptr

	newGwh := v.Elem().Interface().(Person)
	t.Log(newGwh)
}

// 使用反射调用方法
func Add(a, b int) int {
	return a + b
}

var add = func(a, b int) int {
	return a + b
}

func TestReflect5(t *testing.T) {
	v := reflect.ValueOf(Add)
	t.Log(v.Type()) // func(int, int) int
	t.Log(v.Kind()) // func
	if v.Kind() != reflect.Func {
		t.Log("not func, can not call")
	}
	argv := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	result := v.Call(argv) // 前提是v.kind() == reflect.Func
	t.Log("result:", result[0].Interface().(int))
}
