package nvwax

import (
	"github.com/hootuu/htgo/network/rest"
	"github.com/hootuu/htgoapi/nvwa"
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
