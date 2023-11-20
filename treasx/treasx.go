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

func (c *client) NftIssue(req treas.NftIssueReq) (*treas.NftIssueResp, *errors.Error) {
	return rest.Go[treas.NftIssueResp]("/treas/nft/issue", req)
}

func (c *client) NftAirdrop(req treas.NftAirdropReq) (*treas.NftAirdropResp, *errors.Error) {
	return rest.Go[treas.NftAirdropResp]("/treas/nft/airdrop", req)
}

func (c *client) NftPledge(req treas.NftPledgeReq) (*treas.NftPledgeResp, *errors.Error) {
	return rest.Go[treas.NftPledgeResp]("/treas/nft/pledge", req)
}

func (c *client) NftRent(req treas.NftRentReq) (*treas.NftRentResp, *errors.Error) {
	return rest.Go[treas.NftRentResp]("/treas/nft/rent", req)
}

func (c *client) NftTrans(req treas.NftTransReq) (*treas.NftTransResp, *errors.Error) {
	return rest.Go[treas.NftTransResp]("/treas/nft/mine", req)
}

func (c *client) NftMine(req treas.NftMineReq) (*treas.NftMineResp, *errors.Error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) CoinIssue(req treas.CoinIssueReq) (*treas.CoinIssueResp, *errors.Error) {
	return rest.Go[treas.CoinIssueResp]("/treas/coin/issue", req)
}

func (c *client) AccountCreate(req treas.AccountCreateReq) (*treas.AccountCreateResp, *errors.Error) {
	return rest.Go[treas.AccountCreateResp]("/treas/acc/g", req)
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

func (c *client) AlterLoad(req treas.AlterLoadReq) (*treas.AlterLoadResp, *errors.Error) {
	return rest.Go[treas.AlterLoadResp]("/treas/acc/alter/q", req)
}
