package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"reflect"
)

func GwIjMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if !IsCanInject(req) {
				return handler(ctx, req)
			}
			reflectV := reflect.ValueOf(req).Elem()
			if !reflectV.IsValid() {
				return handler(ctx, req)
			}

			//inject uid
			StructField := reflectV.FieldByName(GatewayUID_Field)
			if StructField.IsValid() && StructField.Kind() == reflect.String && StructField.CanSet() && StructField.String() == "" {
				StructField.SetString(GetHeaderKey(ctx, GatewayUID))
			}

			//inject orgID
			StructField = reflectV.FieldByName(GatewayOrgID_Field)
			if StructField.IsValid() && StructField.Kind() == reflect.String && StructField.CanSet() && StructField.String() == "" {
				StructField.SetString(GetHeaderKey(ctx, GatewayOrgID))
			}

			return handler(ctx, req)
		}
	}
}

func IsCanInject(req interface{}) bool {
	if req == nil {
		return false
	}
	reflectT := reflect.TypeOf(req)
	switch reflectT.Kind() {
	case reflect.Ptr:
		if reflect.TypeOf(req).Elem().Kind() == reflect.Struct {
			return true
		}
	case reflect.Struct:
		return true
	}
	return false
}

func GetHeaderKey(ctx context.Context, key string) string {
	hp, ok := http.FromServerContext(ctx)
	if !ok {
		return ""
	}
	return hp.Request.Header.Get(key)
}
