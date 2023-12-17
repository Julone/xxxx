package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"time"
)

type SuccessResponse[T any] struct {
	Success  bool `json:"success"`
	Data     []T  `json:"data"`
	Duration int  `json:"duration"`
}

func Validate[T any](value T) bool {
	// 根据泛型类型进行不同的校验逻辑
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
		// 对int类型进行校验的逻辑
		return reflect.ValueOf(value).Int() == 1
	case reflect.Slice:
		// 对string类型进行校验的逻辑
		v := reflect.ValueOf(value)
		for i := 0; i < v.Len(); i++ {
			r := v.Index(i).Int()
			if r == 2 {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func DefaultTimer[T any](condFunc func(item T) bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		now := time.Now()
		c.Locals("basicRole", 1)
		c.Locals("extraRoles", []int{1, 2})

		c.Next()
		fmt.Printf("%v", Validate(c.Locals("basicRole")))
		fmt.Printf("%v", Validate(c.Locals("extraRoles")))

		latency := time.Since(now)

		status := c.Response().Body()
		var srsp SuccessResponse[T]
		if err := json.Unmarshal(status, &srsp); err != nil {
			return err
		}
		newOutput := make([]T, 0)

		isHrbp := true //c.Locals("basicRole").(int) == 2
		for _, item := range srsp.Data {
			if condFunc(item) == true && isHrbp {
				newOutput = append(newOutput, item)
			}
		}
		srsp.Duration = int(latency.Milliseconds()) // You may want milliseconds here.
		srsp.Data = newOutput
		return c.JSON(srsp)
	}
}
