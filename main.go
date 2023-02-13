package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type A struct {
	Name string
	Desc string
}

func (*A) call() {
	fmt.Println("call")
}

type B interface {
	sort(a, b int) int
}

type Enum int

const (
	Zero Enum = 0
)

func main() {
	//customInt := CustomInt(1)
	//typeOfCusInt := reflect.TypeOf(customInt)
	//fmt.Println(typeOfCusInt.Name(), typeOfCusInt.Kind())
	////fmt.Println(typeOfCusInt.NumField())
	//cat := Cat{
	//	Name: "gwh",
	//	Age:  1,
	//}
	//typeOfCat := reflect.TypeOf(cat)
	//fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	//fmt.Println(typeOfCat.NumField())
	//for i := 0; i < typeOfCat.NumField(); i++ {
	//	fieldType := typeOfCat.Field(i)
	//	fmt.Println(fieldType.Name, fieldType.Tag)
	//}
	//if fieldType, ok := typeOfCat.FieldByName("Name"); ok {
	//	fmt.Println(fieldType.Tag.Get("json"))
	//}

	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Minute)
		<-ch
	}()

	fmt.Println(`start`)
	ch <- "1"
	fmt.Println(`ch <- "1"`)
	ch <- "2"
	fmt.Println(`ch <- "2"`)
	ch <- "3"
	fmt.Println(`ch <- "3"`)
	ch <- "4"
	fmt.Println(`ch <- "4"`)

	time.Sleep(10 * time.Second)

}

type ipKey struct {
	ClusterName string
	Isp         string
	IpVersion   int
	IPType      string
}

var err1 = fmt.Errorf("gwh")

func testError() error {
	return err1
}

type PhysicalNetwork struct {
	Interfaces []*Interface `json:"interfaces,omitempty"`
}

type Interface struct {
	NetworkType    string `json:"networkType,omitempty"`
	PhysicalVlanID int32  `json:"physicalVlanID,omitempty"`
}

func gen(msgArgs ...string) string {
	return ToJson(msgArgs)
}

func FromJson(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

func ToJson(v interface{}) string {
	bs, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bs)
}

func appendSlice(s *[]int) {
	*s = append(*s, 4)
	fmt.Println("Inside function:", *s)
}

// 获取两数之和
func GetSum(a, b int) int {
	return a + b
}

// 获取两数之差
func GetCut(a, b int) int {
	return a - b
}

// 函数也有类型
// GetSum是一个函数，其类型是func(int, int) int，这是一种函数类型
// {}里面是函数的声明

// 因此函数可以作为形参，来调用

// 给类型起别名
type Caculate func(int, int) int

// 函数语意:使用caculate函数计算出i和j的"结果"
func getResult(caculate Caculate, i int, j int) int {
	return caculate(i, j)
}

type NatService interface {
	ListNats()
}

type ServiceImp struct {
}

func (*ServiceImp) ListNats() {

}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
}
