package rpc

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
)

func Test_iniUserClient(t *testing.T) {
	initUserClient()
	resp, err := UserClient.Login(context.Background(), &user.LoginReq{
		Email:    "1demo@damin.com",
		Password: "jfoajsfoji",
	})
	if err != nil {
		t.Errorf("err: %v", err)
		return
	}
	t.Logf("resp: %v", resp)
}
