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

func (c *client) CoinIssue(req treas.CoinIssueReq) (*treas.CoinIssueResp, *errors.Error) {
	return rest.Go[treas.CoinIssueResp]("/treas/coin/issue", req)
}

func (c *client) AccountCreate(req treas.AccountCreateReq) (*treas.AccountCreateResp, *errors.Error) {
	return rest.Go[treas.AccountCreateResp]("/treas/acc/g", req)
}

func (c *client) AlterLoad(req treas.AlterLoadReq) (*treas.AlterLoadResp, *errors.Error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) CoinGet(req treas.CoinGetReq) (*uc.Coin, *errors.Error) {
	return rest.Go[uc.Coin]("/treas/coin/g", req)
}

func (c *client) AccountGet(req treas.AccountGetReq) (*uc.Account, *errors.Error) {
	return rest.Go[uc.Account]("/treas/acc/g", req)
}

func (c *client) AccountLoad(req treas.AccountLoadReq) (*treas.AccountLoadResp, *errors.Error) {
	return rest.Go[treas.AccountLoadResp]("/treas/acc/q", req)
}
