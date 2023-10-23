package nvwax

import (
	"github.com/hootuu/htgo/network/rest"
	"github.com/hootuu/htgoapi/nvwa"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

var Client = &client{}

type client struct {
}

func (c client) VNCreate(req nvwa.VNCreateReq) (*nvwa.VNCreateResp, *errors.Error) {
	return rest.Go[nvwa.VNCreateResp]("/nvwa/vn/birth", req)
}

func (c client) SPCreate(req nvwa.SPCreateReq) (*nvwa.SPCreateResp, *errors.Error) {
	return rest.Go[nvwa.SPCreateResp]("/nvwa/sp/birth", req)
}

func (c client) VNGet(req nvwa.VNGetReq) (*vn.VN, *errors.Error) {
	return rest.Go[vn.VN]("/nvwa/vn/g", req)
}

func (c client) VNLoad(req nvwa.VNLoadReq) (*nvwa.VNLoadResp, *errors.Error) {
	return rest.Go[nvwa.VNLoadResp]("/nvwa/vn/q", req)
}

func (c client) SPGet(req nvwa.SPGetReq) (*sp.SP, *errors.Error) {
	return rest.Go[sp.SP]("/nvwa/sp/g", req)
}

func (c client) SPLoad(req nvwa.SPLoadReq) (*nvwa.SPLoadResp, *errors.Error) {
	return rest.Go[nvwa.SPLoadResp]("/nvwa/sp/q", req)
}
