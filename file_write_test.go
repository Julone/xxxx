package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestMmain(t *testing.T) {
	persons := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	// 获取persons切片的反射对象
	v := reflect.ValueOf(persons)

	// 检查v是否为切片类型
	if v.Kind() == reflect.Slice {
		fmt.Println("persons is a slice")
	} else {
		fmt.Println("persons is not a slice")
	}

	// 遍历切片中的元素
	for i := 0; i < v.Len(); i++ {
		person := v.Index(i) // 获取切片中索引为i的元素

		// 获取Name字段的值
		nameField := person.FieldByName("Name")
		if nameField.IsValid() {
			fmt.Println("Name:", nameField.String())
		} else {
			fmt.Println("Name field not found")
		}

		// 获取Age字段的值
		ageField := person.FieldByName("Age")
		if ageField.IsValid() {
			fmt.Println("Age:", ageField.Int())
		} else {
			fmt.Println("Age field not found")
		}
	}
}
