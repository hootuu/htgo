package main

import (
	"github.com/hootuu/htgo/nvwax"
	"github.com/hootuu/htgoapi/nvwa"
	"github.com/hootuu/utils/logger"
	"go.uber.org/zap"
)

func main() {
	req := nvwa.VNCreateReq{
		Token: "2020eb494bebafe40c3aa1d36e93f719",
		ID:    "NBMOJU_666",
	}
	resp, err := nvwax.Client.VNCreate(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}
