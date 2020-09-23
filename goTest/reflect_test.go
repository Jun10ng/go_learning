package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

/*
	go反射笔记
*/

func TestReflect(t *testing.T) {
	act := &Account{
		"Testname",
		"Testpwd",
	}
	arms := []string{"test_arm1", "test_arm2"}
	u := &User{
		"ma yun",
		99,
		act,
		arms,
	}

	/*
		使用reflect.TypeOf()，返回参数类型
	*/
	fmt.Println(reflect.TypeOf(u))      // *utils.User
	fmt.Println(reflect.TypeOf(*u))     //utils.User
	fmt.Println(reflect.TypeOf(u.Name)) //string
	fmt.Println(reflect.TypeOf(u.Arms)) //[]string

	/*
		Name() 方法是Type的method，可返回类型的名称，结构体返回类型名称，
		某些类型（如slice或pointer）没有名称，此方法返回空字符串。
		和TypeOf返回值(Type类型)比起来，Name方法返回的类型字段比较"纯粹"（string类型)，不包括包名，仅含有类型。

		Kind() 可返回类型的种类，结构体返回struct,切片就返回slice，指针返回ptr，函数返回func等等
	*/
	fmt.Println(reflect.TypeOf(u).Name())      //空字符串
	fmt.Println(reflect.TypeOf(*u).Name())     //User
	fmt.Println(reflect.TypeOf(u.Name).Name()) //string
	fmt.Println(reflect.TypeOf(u.Arms).Name()) //空字符串

	fmt.Println(reflect.TypeOf(u).Kind())      //prt
	fmt.Println(reflect.TypeOf(*u).Kind())     // struct
	fmt.Println(reflect.TypeOf(u.Name).Kind()) //string
	fmt.Println(reflect.TypeOf(u.Arms).Kind()) // slice

	/*
		对序列类型，如ptr，map，slice，channel等，可以使用Elem()找到包含的类型
	*/
	fmt.Println(reflect.TypeOf(u).Elem().Name())      //User :u是ptr 指向的元素类型
	fmt.Println(reflect.TypeOf(u.Arms).Elem().Name()) // string :arms切片的元素类型

	/*
		Type.Filed(n)方法可以返回struct的第n个字段,NumField()返回字段总数,结构体的方法可以使用Method()
		filed 可以有Tag，Type，Name等属性可获取
	*/
	fmt.Println(reflect.TypeOf(*u).NumField())    // 4
	fmt.Println(reflect.TypeOf(*u).Field(2))      // {Account  *utils.Account json:"account" 24 [2] false} ：name，pkgPath,tag，偏移量，index,
	fmt.Println(reflect.TypeOf(*u).Field(2).Name) // Account

	/*
		使用reflect.ValueOf()来读取，设置或创建值
		**注意** 需要使用传入指针参数才能进行修改，否则只是对值的一份拷贝

		使用reflect.New(varType)创建新值，返回一个可以修改的结构体指针值,所以记得加Elem()
	*/

	user := reflect.ValueOf(u).Elem()
	user.FieldByName("Age").SetInt(100)
	fmt.Println(u.Age) //100

	u2 := reflect.New(reflect.TypeOf(*u))
	u2.Elem().FieldByName("Name").SetString("mayun2")
	fmt.Println(u2)        //&{mayun2 0 <nil> []}
	fmt.Println(u2.Elem()) //{mayun2 0 <nil> []}

	/*
		u2.Elem只是一个Value变量，其中丢失了变量的原始类型，所以我们需要将空接口转换为实际类型才能使用它
		通过调用Interface()方法返回到正常变量
	*/
	u22 := u2.Elem().Interface().(User)
	fmt.Println(u22)
	/*
		还可以使用反射来创建通常需要make函数的实例
		reflect.MakeSlice()
		reflect.MakeMap()
	*/

	/*
		使用reflect.MakeFunc创建函数
		计算一个函数的运行时间
	*/
	fncType := reflect.TypeOf(u.Say)
	fncValue := reflect.ValueOf(u.Say)

	calTimeFunc := reflect.MakeFunc(fncType, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		out := fncValue.Call(args)
		end := time.Now()
		fmt.Println(end.Sub(start)) //6.828µs
		return out
	})
	fncReturn := calTimeFunc.Call([]reflect.Value{reflect.ValueOf("Hi")}) //Hi,my name is ma yun
	fmt.Println(fncReturn[0])                                             // done

	/*
		使用reflect.StructOf
		在运行时，根据某些数据来动态的生成struct

		可以看下json库Unmarshal的实现
	*/
	userType := reflect.StructOf([]reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag:  `json:"name"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
			Tag:  `json:"age"`,
		},
	})
	u3 := reflect.New(userType)
	u3.Elem().FieldByName("Name").SetString("mayun3")
	u3.Elem().FieldByName("Age").SetInt(100)
	fmt.Printf("%+v\n", u3) //&{Name:mayun3 Age:100}
	//u33 := u3.Elem().Interface().(User)
	//u33 := u3.Elem().Interface().(User)
	u33 := &User{}
	u44 := &User{
		Name: "ABC",
		Age:  123,
	}
	jdata, _ := json.Marshal(u44)
	json.Unmarshal(jdata, u33)
	fmt.Printf("%+v\n", u33) //
}

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Account *Account `json:"account"`
	Arms    []string
}
type Account struct {
	Username string
	Password string
}

func (u *User) Say(prefix string) string {
	fmt.Printf("%s,my name is %s \n", prefix, u.Name)
	return "done"
}
func NewUser(name string, age int) *User {
	return &User{
		name,
		age,
		nil,
		nil,
	}
}
