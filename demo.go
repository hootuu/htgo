package main

import (
	"github.com/hootuu/htgo/network/htpeer"
	"github.com/hootuu/htgo/nvwax"
	"github.com/hootuu/htgoapi/nvwa"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/utils/logger"
	"go.uber.org/zap"
)

func main() {
	htpeer.Initialize(&htpeer.Peer{
		VN:         "TEST1023",
		SP:         sp.NilID,
		PrivateKey: "0x728967a591094509ac45444045fa6701378cb4ee0cd11e15034927a88c2acb1c",
	})
	req := nvwa.SPCreateReq{
		Token: "ae7756b5ea7bc78f797626f2d82ad253",
		VN:    "TEST1023",
		ID:    "TEST1023001",
	}
	resp, err := nvwax.Client.SPCreate(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}
