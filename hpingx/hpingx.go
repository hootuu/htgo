package hpingx

import (
	"github.com/hootuu/htgo/network/rest"
	"github.com/hootuu/htgoapi/hping"
	"github.com/hootuu/utils/errors"
)

var Client = &client{}

type client struct {
}

func (c client) Ping(req hping.PingReq) (*hping.PingResp, *errors.Error) {
	return rest.Go[hping.PingResp]("/net/ping", req)
}

func (c client) Load(req hping.NodeLoadReq) (*hping.NodeLoadResp, *errors.Error) {
	return rest.Go[hping.NodeLoadResp]("/net/load", req)
}
