package curd

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
)

func ApiCreate[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			api.Error(ctx, err)
			return
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
