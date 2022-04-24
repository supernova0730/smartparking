package utils

import (
	"context"
	"sync"
)

func GetBoolAttribute(ctx context.Context, attribute string) (bool, bool) {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if result, ok := value.(bool); ok {
			return result, true
		}
	}
	return false, false
}

func GetStringAttribute(ctx context.Context, attribute string) (string, bool) {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if result, ok := value.(string); ok {
			return result, true
		}
	}
	return "", false
}

func GetInt64Attribute(ctx context.Context, attribute string) (int64, bool) {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if result, ok := value.(int64); ok {
			return result, true
		}
	}
	return 0, false
}

func SetAttributeBool(ctx context.Context, attribute string, value bool) {
	setAttribute(ctx, attribute, value)
}

func SetAttributeString(ctx context.Context, attribute string, value string) {
	setAttribute(ctx, attribute, value)
}

func SetAttributeInt64(ctx context.Context, attribute string, value int64) {
	setAttribute(ctx, attribute, value)
}

func getAttribute(ctx context.Context, attribute string) interface{} {
	if contextHolder, ok := ctx.Value(ContextHolderKey).(*sync.Map); ok {
		if value, ok := contextHolder.Load(attribute); ok {
			return value
		}
	}
	return nil
}

func setAttribute(ctx context.Context, attribute string, value interface{}) {
	if contextHolder, ok := ctx.Value(ContextHolderKey).(*sync.Map); ok {
		contextHolder.Store(attribute, value)
	}
}
