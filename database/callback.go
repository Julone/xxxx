package database

import (
	"context"
	"fmt"
	"gorm.io/gorm/schema"
	"reflect"

	"gorm.io/gorm"
)

func getReflectValueElem(i interface{}) reflect.Value {
	value := reflect.ValueOf(i)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}

var dbAesTableColumnMap = map[[2]string]bool{
	[2]string{"books", "area"}: true,
}

// AfterQuery 只支持查询返回的结果是单表model的结构体，map、自定义的其他结构体（联表、部分字段）暂不支持
func AfterQuery(db *gorm.DB) {
	defer func() {
		if err := recover(); err != nil {
			db.AddError(fmt.Errorf("recover panic:%v", err))
		}
	}()
	if db.Error == nil && db.Statement.Schema != nil && db.RowsAffected > 0 && !db.Statement.SkipHooks {
		destReflectValue := getReflectValueElem(db.Statement.Dest)
		for fieldIndex, field := range db.Statement.Schema.Fields {
			switch destReflectValue.Kind() {
			case reflect.Slice, reflect.Array: //[]struct
				for i := 0; i < destReflectValue.Len(); i++ {
					index := destReflectValue.Index(i)

					if index.Kind() != reflect.Struct {
						continue
					}

					if fieldValue, isZero := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue.Index(i)); !isZero { // 从字段中获取数值
						if s, ok := fieldValue.(string); ok {
							if s == "fz" {
								continue
							}
							_ = db.AddError(field.Set(context.Background(), index, s+"fasd"))
						}
					}
				}
			case reflect.Struct: //struct
				if destReflectValue.NumField() != len(db.Statement.Schema.Fields) {
					return
				}
				if fieldValue, isZero := field.ValueOf(db.Statement.Context, destReflectValue); !isZero { // 从字段中获取数值
					if destReflectValue.Type().Field(fieldIndex).Name != field.Name {
						return
					}
					if dbAesTableColumnMap[[2]string{db.Statement.Schema.Table, field.DBName}] {
						if s, ok := fieldValue.(string); ok {
							_ = db.AddError(field.Set(context.Background(), destReflectValue, s+"fasd"))
						}
					}
				}
			}
		}
		var newArr []*schema.Field
		for _, field := range db.Statement.Schema.Fields {
			if field != nil {
				newArr = append(newArr, field)
			}
		}
		db.Statement.Schema.Fields = (newArr)
	}
}
