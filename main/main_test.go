package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"hash"
	"net"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"testing"
	"time"
)

type Cat struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type CustomInt int

type Usb interface {
	Start() error
	Stop() error
}

type Computer struct {
}

func (*Computer) Start() error {
	return nil
}

func (*Computer) Stop() error {
	return nil
}

type Phone struct {
}

func (*Phone) Start() error {
	return nil
}

func (*Phone) Stop() error {
	return nil
}

type Student struct {
	Name string
}

func TestB(t *testing.T) {
	computer1 := &Computer{}

	var ch chan struct{}
	fmt.Println(ch)

	var b interface {
		Start() error
		Stop() error
	}
	b = computer1
	fmt.Println(b)
	student1 := &struct {
		Name string
	}{
		Name: "gwh1",
	}
	student2 := &Student{
		Name: "gwh2",
	}
	student1 = (*struct {
		Name string
	})(student2)
	student2 = (*Student)(student1)
	fmt.Println(student2)
	fmt.Println(student1)

	var a interface{}
	cat1 := &Cat{
		Name: "gwh",
		Age:  1,
	}
	a = cat1
	var cat2 *Cat
	if cat, ok := a.(*Cat); ok { // 类型断言
		cat2 = cat
	}

	fmt.Println(cat2)
}

func TestC(t *testing.T) {
	computer := &Computer{}
	getType(computer)
}

func TestD(t *testing.T) {
	var a *string
	fmt.Println(a == nil)
	fmt.Println(IsNil(a))
}

func IsNil(i interface{}) bool {
	return i == nil
}

// 类型断言的另一种写法 用于判断某个 interface 变量中实际存储的变量类型
func getType(a Usb) {
	switch a.(type) { // 只能在switch中使用 左侧必须为接口类型
	case *Computer:
		fmt.Println("type of a is Computer")
	case *Phone:
		fmt.Println("type of a is Phone")
	default:
		fmt.Println("unknown type")
	}
}

func getType2(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("type of a is int")
	case string:
		fmt.Println("type of a is string")
	case *Computer:
		fmt.Println("type of a is Computer")
	default:
		fmt.Println("unknown type")
	}
}

func TestA(t *testing.T) {
	customInt := CustomInt(1)
	typeOfCusInt := reflect.TypeOf(customInt)
	fmt.Println(typeOfCusInt.Name(), typeOfCusInt.Kind())

	cat := &Cat{
		Name: "gwh",
		Age:  1,
	}
	typeOfCatPtr := reflect.TypeOf(cat)
	fmt.Println(typeOfCatPtr.Name(), typeOfCatPtr.Kind())
	typeOfCat := typeOfCatPtr.Elem() // Elem()对指针type取值
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	fmt.Println(typeOfCat.NumField())
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Println(fieldType.Name, fieldType.Tag, fieldType.Type.Name(), fieldType.Type.Kind())
	}

	if fieldType, ok := typeOfCat.FieldByName("Name"); ok {
		fmt.Println(fieldType.Tag.Get("json"))
	}

}

func TestAA(t *testing.T) {
	a := reflect.ValueOf(1).Interface().(int)
	fmt.Println(a)
}

func TestReflectValue(t *testing.T) {
	var a int64 = 10
	v := reflect.ValueOf(&a)
	v.Elem().SetInt(20)
	fmt.Println(a)
}

func TestReflectValu3(t *testing.T) {
	var a int64 = 10
	v := reflect.ValueOf(&a)
	v.Elem().SetInt(20)
	fmt.Println(a)
}

func TestReflectValue2(test *testing.T) {
	cat := &Cat{
		Name: "gwh",
		Age:  25,
	}
	v := reflect.ValueOf(cat)
	fmt.Println(v.Kind())
	if _, ok := v.Interface().(*Cat); ok {
		fmt.Println(ok)
	}
	v2 := v.Elem()
	fmt.Println(v2.Kind())
	if _, ok := v2.Interface().(Cat); ok {
		fmt.Println(ok)
	}
	nameField := v2.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("gwh123")
	}
	ageField := v2.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(30)
	}
	fmt.Println(cat)
}

type CustomError struct{}

func (*CustomError) Error() string {
	return fmt.Sprintf("custom error")
}

func TestCheckInterface(t *testing.T) {
	// 获取接口的reflect.Type
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	// typeOfError := reflect.TypeOf((error)(nil)) 这种方式错误
	fmt.Println(typeOfError)

	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError))
	fmt.Println(customError.Implements(typeOfError))
}

func Add(a, b int) int {
	return a + b
}

func Test1(t *testing.T) {
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(1)}
	for i := range args {
		fmt.Println(i)
	}
}

func TestReflectFunc(t *testing.T) {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	addType := v.Type()
	args := []reflect.Value{}
	for i := 0; i < addType.NumIn(); i++ {
		if addType.In(i).Kind() != reflect.Int {
			return
		}
		args = append(args, reflect.ValueOf(i))
	}

	result := v.Call(args)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int())
}

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}

type Person struct {
	Name string
	Age  int
}

type Horse struct {
	Price int
	Inch  int
}

func TestMap(t *testing.T) {
	data :=
		`{
		"Name": "gwh",
		"Age": 28
	}`

	var target interface{}
	json.Unmarshal([]byte(data), &target)

	targetValue := reflect.ValueOf(target)
	fmt.Println(targetValue.Kind())
	switch targetValue.Kind() {
	case reflect.Map:
		for _, key := range targetValue.MapKeys() {
			fmt.Println(key, targetValue.MapIndex(key))
		}
	case reflect.Struct:
		for i := 0; i < targetValue.NumField(); i++ {
			fmt.Println(targetValue.Type().Field(i).Tag)
			fmt.Println(targetValue.Field(i))
		}
	case reflect.Ptr:
		if targetValue.IsNil() {
		} else {
			targetValue.Elem().Type()
		}
	}
	var i interface{} = 1
	fmt.Println(display("", reflect.ValueOf(&i)))

}

func display(path string, v reflect.Value) string {
	return v.Kind().String()
}

const (
	n1 = iota
	n2
	n3
	n4
)

func TestChannel(t *testing.T) {
	str := "hello world郭"
	fmt.Println(len([]byte(str)))
}

func TestMap1(t *testing.T) {
	ctx, cancelf := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelf()

	time.Sleep(2 * time.Second)
	useCtx(ctx)
}

func useCtx(ctx context.Context) {
	fmt.Println(ctx)
}

var (
	client_token  string
	client_params string

	server_token  string // 表示redis中所有token
	server_params string
	server_status string
	server_result string
)

func judge() {
	if client_token != server_token {
		// 执行，入库token，状态执行中(server_token,server_params,server_status = "running")
		// ... 业务逻辑
		// 执行完成，入库结果，状态执行完毕(server_result, server_status = "done")
	} else if client_params != server_params {
		// 报错，请求不符合幂等参数校验
	} else { // token和params都和server中对应
		if server_status == "running" {
			// 方案1：wait until server_status is 'done'，读取server_result，直接返回结果即可
			// 方案2：直接报错，提示上次幂等请求仍在执行中，勿重复操作
		} else if server_status == "done" {
			// 读取server_result，直接返回结果即可
		}
	}
}

func Test_Map(t *testing.T) {
	eniiipMap := map[string][]string{}
	eniiipMap["gwh"] = []string{"gwh1", "gwh2"}
	for _, s := range eniiipMap["wll"] {
		fmt.Println(s)
	}
}

type ParamsReq struct {
	Name          *string `json:"name,omitempty"`
	EnName        *string `json:"en_name,omitempty"`
	Status        *bool   `json:"status,omitempty"`
	IsAutoApporve *bool   `json:"is_auto_apporve,omitempty"`
}

func TestParamsRes(t *testing.T) {
	name := "郭伟豪"
	status := true
	isAutoApprove := false
	reqBody := &ParamsReq{
		Name:          &name,
		Status:        &status,
		IsAutoApporve: &isAutoApprove,
	}
	reqBodyBs, err := json.Marshal(reqBody)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(string(reqBodyBs))

	respBody := &ParamsReq{}
	if err := json.Unmarshal(reqBodyBs, &respBody); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(respBody)

}

type CreateCloudServerReq struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Sex  bool   `json:"sex,omitempty"`
}

var Hasher hash.Hash

func init() {
	Hasher = md5.New()
}

func TestHash(t *testing.T) {

	s := `{
    "instance_identity": "dev-veen90944523324202357592",
    "instance_name": "gwh测试ClientToken123"
}`

	res, _ := tranStringToHash(s)
	fmt.Println(res)

}

func tranStringToHash(s string) (string, error) {
	_, err := Hasher.Write([]byte(s))
	if err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%x", Hasher.Sum(nil))
	return hash, nil
}

func TestHash2(t *testing.T) {
	req := &CreateCloudServerReq{
		Age:  18,
		Name: "gwh",
	}
	bs, _ := json.Marshal(req)
	fmt.Println(string(bs))

	hasher := md5.New()
	_, err := hasher.Write(bs)
	if err != nil {
		t.Errorf(err.Error())
	}
	result := hasher.Sum(nil)
	res := fmt.Sprintf("%x", result)
	fmt.Println(res)
}

func TestHash3(t *testing.T) {
	bs := []byte(`{"name":"gwh","age":18}`)
	fmt.Println(string(bs))

	hasher := md5.New()
	_, err := hasher.Write(bs)
	if err != nil {
		t.Errorf(err.Error())
	}
	result := hasher.Sum(nil)
	res := fmt.Sprintf("%x", result)
	fmt.Println(res)
}

func TestHash4(t *testing.T) {
	bs := []byte(`{"age":18,"name":"gwh"}`)
	fmt.Println(string(bs))

	hasher := md5.New()
	_, err := hasher.Write(bs)
	if err != nil {
		t.Errorf(err.Error())
	}
	result := hasher.Sum(nil)
	res := fmt.Sprintf("%x", result)
	fmt.Println(res)
}

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Add(-24 * time.Hour))
	fmt.Println(now.Unix())
	fmt.Println(now.Add(-24 * time.Hour).Unix())
}

func TestNew(t *testing.T) {
	cat := *new(Cat)
	fmt.Println(cat)
}

func TestInferface(t *testing.T) {
	var result interface{}
	if _, ok := result.(Usb); ok {
		fmt.Println(ok)
	} else {
		fmt.Println(ok)
	}
}

func TestSetValue(t *testing.T) {
	mapReqBody := map[string]interface{}{}
	if err := FromJson("{}", &mapReqBody); err != nil {
		fmt.Printf("err is %s", err.Error())
	}
	fmt.Println(mapReqBody)
}

func TestSetValue2(t *testing.T) {
	key := "client_token_gwh"
	expiration := 24 * time.Hour
	fmt.Printf("add client_token [%s] expiration is %v", key, expiration)
}

func ip2int(s string) uint32 {
	ip := net.ParseIP(s)
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func TestIp2Int(t *testing.T) {
	fmt.Println(ip2int("0.0.0.1"))
	fmt.Println(ip2int("0.0.1.1"))
}

func TestBool2String(t *testing.T) {
	// 定义网段A和网段B
	netA := "192.168.0.0/16"
	netB := "192.168.0.0/16"

	// 解析网段A和网段B
	_, subnetA, _ := net.ParseCIDR(netA)
	_, subnetB, _ := net.ParseCIDR(netB)

	// 判断网段A是否包含网段B
	if ContainsCIDR(subnetA, subnetB) {
		fmt.Println(netA, "包含", netB)
	} else {
		fmt.Println(netA, "不包含", netB)
	}
}

func TestBool2String2(t *testing.T) {
	// 定义一批网段
	nets := []string{
		"192.168.10.0/24",
		"192.168.0.0/16",
		"10.0.10.0/24",
		"10.0.0.0/8",
		"192.168.0.0/16",
		"10.0.0.0/16",
		"172.16.0.0/12",
	}

	// 解析网段并存储到IP网段对象的切片中
	ipNets := make([]*net.IPNet, len(nets))
	for i, netStr := range nets {
		_, ipNet, _ := net.ParseCIDR(netStr)
		ipNets[i] = ipNet
	}

	// 按照网段长度（子网掩码位数）进行排序
	//sort.Slice(ipNets, func(i, j int) bool {
	//	iSize, _ := ipNets[i].Mask.Size()
	//	jSize, _ := ipNets[j].Mask.Size()
	//	return iSize > jSize
	//})

	// 判断网段是否相互包含，并只保留最大的网段
	for i := 0; i < len(ipNets)-1; i++ {
		for j := i + 1; j < len(ipNets); j++ {
			if ContainsCIDR(ipNets[i], ipNets[j]) {
				ipNets = append(ipNets[:j], ipNets[j+1:]...)
				j--
			} else if ContainsCIDR(ipNets[j], ipNets[i]) {
				ipNets = append(ipNets[:i], ipNets[i+1:]...)
				i--
				break
			}
		}
	}

	// 输出最大的网段
	for _, ipNet := range ipNets {
		fmt.Println(ipNet.String())
	}
}

// 判断a网段是否包含b网段
func ContainsCIDR(a, b *net.IPNet) bool {
	ones1, _ := a.Mask.Size()
	ones2, _ := b.Mask.Size()
	return ones1 <= ones2 && a.Contains(b.IP)
}

func TestJson(t *testing.T) {
	cidr := "1.0.0.0/0"
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic("check1")
	}
	// cidr必须为网段
	if ip.String() != ipNet.IP.String() {
		panic("check2")
	}
}

func TestBase64(t *testing.T) {
	bs, err := os.ReadFile("/root/code/byteedge/oort_ben_tob/internal/utils/a.csv")
	if err != nil {
		t.Fatal(err)
	}

	resBs, err := base64.StdEncoding.DecodeString(string(bs))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("/root/code/byteedge/oort_ben_tob/internal/utils/b.csv", resBs, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestNum(t *testing.T) {
	s1, err := genNewPvcNameForLocalSysDisk("veen29224091114241222102")
	if err != nil {
		panic(err)
	}
	s2, err := genNewPvcNameForLocalSysDisk("veen29224091114241222102-disk1")
	if err != nil {
		panic(err)
	}
	fmt.Println(s1, s2)
}

func genNewPvcNameForLocalDataDisk(pvcName string, dataDiskLength int) (string, error) {

	r := regexp.MustCompile("disk([0-9]+)")

	matches := r.FindStringSubmatch(pvcName)
	if len(matches) > 1 {
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			return "", err
		}
		newNum := num + dataDiskLength
		newStr := r.ReplaceAllString(pvcName, fmt.Sprintf("disk%d", newNum))
		return newStr, nil
	} else {
		return "", fmt.Errorf("local data disk pvc name is illegal")
	}
}

func genNewPvcNameForLocalSysDisk(str string) (string, error) {
	r := regexp.MustCompile(`(.*-disk)(\d*)$`)
	matches := r.FindStringSubmatch(str)
	if len(matches) > 2 {
		num := 1
		if matches[2] != "" {
			var err error
			num, err = strconv.Atoi(matches[2])
			if err != nil {
				return "", fmt.Errorf("local sys disk pvc name is illegal")
			}
			num++
		} else {
			return "", fmt.Errorf("local sys disk pvc name is illegal")
		}
		newStr := fmt.Sprintf("%s%d", matches[1], num)
		return newStr, nil
	}
	return str + "-disk1", nil
}

func TestBase641(t *testing.T) {
	bs, err := os.ReadFile("./a.csv")
	if err != nil {
		t.Fatal(err)
	}

	resBs, err := base64.StdEncoding.DecodeString(string(bs))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("./b.csv", resBs, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestList(t *testing.T) {
	fmt.Println(compareExtInfo())
}

type Extra struct {
	ResourceClass  string
	ResourceNumber int64
	Capacity       int64
}

func compareExtInfo() bool {
	oldExtras := []*Extra{
		{
			ResourceClass:  "123",
			ResourceNumber: 123,
			Capacity:       123,
		},
		{
			ResourceClass:  "321",
			ResourceNumber: 321,
			Capacity:       321,
		},
	}
	newExtras := []*Extra{
		{
			ResourceClass:  "321",
			ResourceNumber: 321,
			Capacity:       321,
		},
		{
			ResourceClass:  "123",
			ResourceNumber: 123,
			Capacity:       123,
		},
		{
			ResourceClass:  "321",
			ResourceNumber: 321,
			Capacity:       321,
		},
	}

	SortExtras(oldExtras)
	SortExtras(newExtras)
	return reflect.DeepEqual(oldExtras, newExtras)
}

func SortExtras(extras []*Extra) {
	sort.Slice(extras, func(i, j int) bool {
		if extras[i].ResourceClass != extras[j].ResourceClass {
			return extras[i].ResourceClass < extras[j].ResourceClass
		}
		if extras[i].ResourceNumber != extras[j].ResourceNumber {
			return extras[i].ResourceNumber < extras[j].ResourceNumber
		}
		return extras[i].Capacity < extras[j].Capacity
	})
}
