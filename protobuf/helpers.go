package protobuf

import (
	"context"

	"strings"

	"google.golang.org/grpc/metadata"
)

func GetFromMetadata(ctx context.Context, key string) (string, bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	if len(headers[key]) == 0 {
		return "", false
	}

	return strings.Join(headers[key], ""), true
}

func CreateMetaDataContext(headers map[string]string) context.Context {
	h := metadata.New(headers)
	return metadata.NewOutgoingContext(context.Background(), h)
}

func CreateMetaDataContextWithContext(ctx context.Context, headers map[string]string) context.Context {
	h := metadata.New(headers)
	return metadata.NewOutgoingContext(ctx, h)
}
