package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) String(i int) string {
	return fmt.Sprintf("<Name: %s, Age: %d, %d>", u.Name, u.Age, i)
}

func main() {
	author := User{"Ahala", 20}
	//author := 1
	t := reflect.TypeOf(author)
	v := reflect.ValueOf(author)
	fmt.Printf("TypeOf author:%T:%v\n", t, t)
	fmt.Printf("TypeOf author:%T:%v\n", v, v)

	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%v: %v\n", t.Field(i).Name, v.Field(i))
	}
	vs := make([]reflect.Value, v.Method(0).Type().NumIn())
	vs[0] = reflect.ValueOf(1)
	fmt.Printf("%v: %v\n", t.Method(0).Name, v.MethodByName("String").Call(vs))

	vchange := reflect.ValueOf(&author)
	vchange.Elem().Field(1).SetInt(60)
	fmt.Printf("After change value is %d", vchange.Elem().Field(1))

}
