package services

import (
	"context"
	"testing"

	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
)

func TestCreateTopicAndFetch(t *testing.T) {
	svc := NewTopicService(nil)
	RunTest(t,
		func(server *grpc.Server) {
			protos.RegisterTopicServiceServer(server, svc)
		},
		func(ctx context.Context, conn *grpc.ClientConn) {
			client := protos.NewTopicServiceClient(conn)
			resp, err := client.CreateTopic(ctx, &protos.CreateTopicRequest{
				Topic: &protos.Topic{
					Name:  "First Topic",
					Users: map[string]bool{"1": true, "2": true, "3": true},
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Topic.Id, "1")
			assert.Equal(t, resp.Topic.Name, "First Topic")
			assert.Equal(t, resp.Topic.Users, map[string]bool{"1": true, "2": true, "3": true})

			// Create another
			resp, err = client.CreateTopic(ctx, &protos.CreateTopicRequest{
				Topic: &protos.Topic{
					Name:  "An awesome second song",
					Users: map[string]bool{"4": true, "2": true, "3": true},
				},
			})
			assert.Equal(t, err, nil, "Error should be nil")
			assert.Equal(t, resp.Topic.Id, "2")
			assert.Equal(t, resp.Topic.Name, "An awesome second song")
			assert.Equal(t, resp.Topic.Users, []string{"4", "2", "5"})

			resp2, err := client.GetTopics(ctx, &protos.GetTopicsRequest{
				Ids: []string{"1", "2"},
			})
			assert.Equal(t, len(resp2.Topics), 2)
			assert.Equal(t, err, nil)
		})
}
