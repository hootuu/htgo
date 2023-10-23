package yamax

import (
	"github.com/hootuu/htgo/network/rest"
	"github.com/hootuu/htgoapi/yama"
	"github.com/hootuu/utils/errors"
)

var Client = &client{}

type client struct {
}

func (c client) YinPlant(req yama.YinPlantReq) (*yama.YinPlantResp, *errors.Error) {
	return rest.Go[yama.YinPlantResp]("/yama/yn/plant", req)
}
