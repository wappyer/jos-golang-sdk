package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person interface {
	GetName() string
	GetAge() int
	GetDesc() map[string]interface{}
}

type Man struct {
	Desc map[string]interface{}
	Name string
	Age  int
	Tag  []string
}

func (m *Man) GetName() string {
	return m.Name
}

func (m *Man) SetName(name string) {
	m.Name = name
	m.Desc["name"] = m.Name
}
func (m *Man) GetAge() int {
	return m.Age
}

func (m *Man) SetAge(age int) {
	m.Age = age
	m.Desc["age"] = m.Age
}

func (m *Man) GetTag() []string {
	return m.Tag
}

func (m *Man) SetTag(tag []string) {
	m.Tag = tag
	m.Desc["tag"] = strings.Join(tag, ",")
}

func (m *Man) GetDesc() map[string]interface{} {
	return m.Desc
}

func main() {
	zxc := &Man{
		Desc: make(map[string]interface{}),
		Name: "zxc",
		Age:  18,
	}
	ReadPerson(zxc)
}

func ReadPerson(person Person) string {
	requestValuePtr := reflect.ValueOf(person)
	requestValue := requestValuePtr
	if requestValuePtr.Kind() == reflect.Ptr {
		requestValue = requestValue.Elem()
	}

	requestType := requestValue.Type()
	for i := 0; i < requestValue.NumField(); i++ {
		field := requestType.Field(i)
		setterMethodName := "Set" + field.Name
		fmt.Println(setterMethodName)
		mtd := requestValuePtr.MethodByName(setterMethodName)
		fmt.Println(mtd)
		if !mtd.IsValid() {
			fmt.Println("no method")
			continue
		}

		value := requestValue.Field(i)
		ret := mtd.Call([]reflect.Value{value})
		fmt.Println(ret)
	}

	fmt.Println(person.GetDesc())
	return ""
}
