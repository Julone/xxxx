package middleware

import (
	"fmt"
	"reflect"
	"testing"
)

func Validate2[T any](value T) bool {
	v := reflect.ValueOf(value)
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		if len() != 0 {
			v = reflect.ValueOf(v.Elem())
			fmt.Printf("%v is ptr %v", v.Elem(), v.Kind())
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
		// 对int类型进行校验的逻辑
		if v.Int() == 1 {
			return true
		}
	case reflect.Slice:
		// 对string类型进行校验的逻辑
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).CanInt() && v.Index(i).Int() == 23 {
				return true
			}
		}

	case reflect.Ptr:
		fmt.Printf("\n %v is pter", v)
		return false
	default:
		// 处理其他类型的情况
		return false
	}
	return false
}

func TestValidate(t *testing.T) {
	r1 := Validate2([]int{1, 2, 23, 23, 32})
	r2 := Validate2([]string{"s"})
	a := []int{1, 23}
	r3 := Validate2(&a)
	r4 := Validate2(1)
	r5 := Validate2(&a)
	r6 := Validate2([]int{})
	r7 := Validate2([]string{})
	r8 := Validate2(0) // 这里可以替换为你想要的任何int值，因为0不等于1或23
	t.Logf("%v, %v, %v, %v, %v, %v, %v, %v", r1, r2, r3, r4, r5, r6, r7, r8)
}
