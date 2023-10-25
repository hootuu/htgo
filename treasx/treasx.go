package treasx

import (
	"github.com/hootuu/htgo/network/rest"
	"github.com/hootuu/htgoapi/treas"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/utils/errors"
)

var Client = &client{}

type client struct {
}

func (c client) CoinGet(req treas.CoinGetReq) (*uc.Coin, *errors.Error) {
	return rest.Go[uc.Coin]("/treas/coin/g", req)
}

func (c client) AccountGet(req treas.AccountGetReq) (*uc.Account, *errors.Error) {
	return rest.Go[uc.Account]("/treas/acc/g", req)
}

func (c client) AccountLoad(req treas.AccountLoadReq) (*treas.AccountLoadResp, *errors.Error) {
	return rest.Go[treas.AccountLoadResp]("/treas/acc/q", req)
}

func (c client) AlterLoad(req treas.AlterLoadReq) (*treas.AlterLoadResp, *errors.Error) {
	return rest.Go[treas.AlterLoadResp]("/treas/acc/alter/q", req)
}
