package curd

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func ApiSearch[T any](fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		query := body.ToQuery()

		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		var datum []*T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//OK(ctx, cs)
		api.List(ctx, datum, cnt)
	}
}

func ApiSearchHook[T any](after func(datum []*T) error, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		query := body.ToQuery()

		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		var datum []*T
		cnt, err := query.FindAndCount(&datum)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		if after != nil {
			if err := after(datum); err != nil {
				api.Error(ctx, err)
				return
			}
		}

		//OK(ctx, cs)
		api.List(ctx, datum, cnt)
	}
}

func ApiSearchMapHook[T any](after func(datum []map[string]any) error, fields ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body ParamSearch
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		query := body.ToQuery()

		//查询字段
		fs := ctx.QueryArray("field")
		if len(fs) > 0 {
			query.Cols(fs...)
		} else if len(fields) > 0 {
			query.Cols(fields...)
		}

		var data T
		var datum []map[string]any
		cnt, err := query.Table(data).FindAndCount(&datum)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//后续处理
		if after != nil {
			err := after(datum)
			if err != nil {
				api.Error(ctx, err)
				return
			}
		}

		//OK(ctx, cs)
		api.List(ctx, datum, cnt)
	}
}
