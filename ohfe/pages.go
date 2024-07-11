package ohfe

import (
	"net/http"

	gfn "github.com/panyam/goutils/fn"
	protos "github.com/panyam/onehub/gen/go/onehub/v1"
	s3views "github.com/panyam/s3gen/views"
)

type SiteView = s3views.View[*Web]
type BaseView = s3views.BaseView[*Web]

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
	TopicViews []*TopicDetailView
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
	v.TopicViews = gfn.Map(resp.Topics, func(t *protos.Topic) *TopicDetailView {
		out := &TopicDetailView{Topic: t}
		out.InitView(v.Context, v)
		return out
	})

	return err
}

type TopicPanel struct {
	BaseView
}

type TopicDetailView struct {
	BaseView
	Topic *protos.Topic
}
