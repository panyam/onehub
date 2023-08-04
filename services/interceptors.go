package services

import (
	"context"

	"google.golang.org/grpc/metadata"
)

/**
 * Given a request context and metadata returns the UserID of the authenticated user if
 * request is authenticated.
 */
func GetAuthedUser(ctx context.Context) (userid string) {
	md, ok := metadata.FromIncomingContext(ctx)
	// log.Println("MD: ", md)
	var values []string
	if ok {
		values = md.Get("OneHubUsername")
	}

	if len(values) > 0 {
		userid = values[0]
	}
	return
}
