package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "pen-1", "Price": 100, "Picture": "/static/image/pen-1.png"},
		{"Name": "pen-2", "Price": 110, "Picture": "/static/image/pen-2.png"},
		{"Name": "pen-3", "Price": 120, "Picture": "/static/image/pen-3.png"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}