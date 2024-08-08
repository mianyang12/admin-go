package dao

import (
	"database/sql"
	"fmt"
	"naive-admin-go/model"
	"reflect"
	"strconv"
	"time"
)

// fillPostsFromRows 填充 posts 切片
func fillPostsFromRows(rows *sql.Rows, posts *[]model.Post) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	vals := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for k := range vals {
		scans[k] = &vals[k]
	}

	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			return err
		}

		post := new(model.Post)
		result := make(map[string]interface{})
		for index, val := range columns {
			result[val] = string(vals[index])
		}

		elem := reflect.ValueOf(post).Elem()
		for i := 0; i < elem.NumField(); i++ {
			structField := elem.Type().Field(i)
			fieldInfo := structField.Tag.Get("orm")
			v := result[fieldInfo]
			t := structField.Type
			switch t.String() {
			case "int":
				s := v.(string)
				vInt, _ := strconv.Atoi(s)
				elem.Field(i).Set(reflect.ValueOf(vInt))
			case "string":
				elem.Field(i).Set(reflect.ValueOf(v.(string)))
			case "int64":
				s := v.(string)
				vInt64, _ := strconv.ParseInt(s, 10, 64)
				elem.Field(i).Set(reflect.ValueOf(vInt64))
			case "int32":
				s := v.(string)
				vInt32, _ := strconv.ParseInt(s, 10, 32)
				elem.Field(i).Set(reflect.ValueOf(vInt32))
			case "time.Time":
				s := v.(string)
				t, _ := time.Parse(time.RFC3339, s)
				elem.Field(i).Set(reflect.ValueOf(t))
			default:
				// 如果有未处理的类型，这里可以添加错误处理或者日志记录
				continue
			}
		}

		*posts = append(*posts, *post)
	}

	return rows.Err()
}

// QueryPosts 是一个通用的函数，用于执行 SQL 查询并返回满足条件的 posts 切片
func QueryPosts(query string, args ...interface{}) ([]model.Post, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {

			fmt.Println("查询错误", err)
		}
	}(rows)

	var posts []model.Post
	err = fillPostsFromRows(rows, &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
