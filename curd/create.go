package curd

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"reflect"
)

func ApiCreate[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//默认值
		field := reflect.ValueOf(data).Elem().FieldByName("Id")
		if !field.IsZero() && field.Kind() == reflect.String {
			if field.Len() == 0 {
				key := xid.New().String()
				field.SetString(key)
			}
		}

		_, err = db.Engine().InsertOne(&data)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		api.OK(ctx, &data)
	}
}

func ApiCreateHook[T any](before, after func(m *T) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//默认值
		field := reflect.ValueOf(data).Elem().FieldByName("Id")
		if !field.IsZero() && field.Kind() == reflect.String {
			if field.Len() == 0 {
				key := xid.New().String()
				field.SetString(key)
			}
		}

		if before != nil {
			if err := before(&data); err != nil {
				api.Error(ctx, err)
				return
			}
		}

		_, err = db.Engine().InsertOne(&data)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		if after != nil {
			if err := after(&data); err != nil {
				api.Error(ctx, err)
				return
			}
		}

		api.OK(ctx, &data)
	}
}
