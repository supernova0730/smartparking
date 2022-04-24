package utils

import "context"

func GetClientID(ctx context.Context) int64 {
	clientID, ok := GetInt64Attribute(ctx, AttributeClientID)
	if !ok {
		return 0
	}
	return clientID
}
