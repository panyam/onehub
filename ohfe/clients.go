package ohfe

import (
	"log"

	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	"google.golang.org/grpc"
)

type ClientMgr struct {
	svcAddr          string
	topicSvcClient   protos.TopicServiceClient
	messageSvcClient protos.MessageServiceClient
	userSvcClient    protos.UserServiceClient
}

func NewClientMgr(svc_addr string) *ClientMgr {
	return &ClientMgr{svcAddr: svc_addr}
}

func (c *ClientMgr) GetUserSvcClient() (out protos.UserServiceClient, err error) {
	if c.userSvcClient == nil {
		userSvcConn, err := grpc.Dial(c.svcAddr, grpc.WithInsecure())
		if err != nil {
			log.Printf("cannot connect with server %v", err)
			return nil, err
		}

		c.userSvcClient = protos.NewUserServiceClient(userSvcConn)
	}
	return c.userSvcClient, nil

}

func (c *ClientMgr) GetTopicSvcClient() (out protos.TopicServiceClient, err error) {
	if c.topicSvcClient == nil {
		topicSvcConn, err := grpc.Dial(c.svcAddr, grpc.WithInsecure())
		if err != nil {
			log.Printf("cannot connect with server %v", err)
			return nil, err
		}

		c.topicSvcClient = protos.NewTopicServiceClient(topicSvcConn)
	}
	return c.topicSvcClient, nil

}

func (c *ClientMgr) GetMessageSvcClient() (out protos.MessageServiceClient, err error) {
	if c.messageSvcClient == nil {
		messageSvcConn, err := grpc.Dial(c.svcAddr, grpc.WithInsecure())
		if err != nil {
			log.Printf("cannot connect with server %v", err)
			return nil, err
		}

		c.messageSvcClient = protos.NewMessageServiceClient(messageSvcConn)
	}
	return c.messageSvcClient, nil
}
