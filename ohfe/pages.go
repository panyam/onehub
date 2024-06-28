package ohfe

import (
	"log"
	"net/http"

	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	s3 "github.com/panyam/s3gen"
)

type SiteView = s3.View[*Web]
type BaseView = s3.BaseView[*Web]

type LoginPage struct {
	BaseView
}

type HomePage struct {
	BaseView
	TopicPanel     TopicPanel
	TopicListPanel TopicListPanel
}

func (v *HomePage) InitView(w *Web, parentView SiteView) {
	v.BaseView.AddChildViews(&v.TopicListPanel, &v.TopicPanel)
	v.BaseView.InitView(w, parentView)
}

type TopicListPanel struct {
	BaseView
	TopicViews []*TopicView
}

func (v *TopicListPanel) InitView(w *Web, parentView SiteView) {
	for _, tv := range v.TopicViews {
		v.BaseView.AddChildViews(tv)
	}
	v.BaseView.InitView(w, parentView)
}

func (v *TopicListPanel) ValidateRequest(w http.ResponseWriter, r *http.Request) (err error) {
	// we load the actual topics here
	client, _ := v.Context.clientMgr.GetTopicSvcClient()

	req := protos.ListTopicsRequest{}
	username := v.Context.GetLoggedInUser(r)
	ctx := v.Context.UserContext(username)
	resp, err := client.ListTopics(ctx, &req)
	log.Println("Resp, err: ", resp)
	return nil
}

type TopicPanel struct {
	BaseView
}

type TopicView struct {
	BaseView
}
