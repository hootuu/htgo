package main

import (
	"fmt"
	"github.com/hootuu/htgo/network/htpeer"
	"github.com/hootuu/htgo/nvwax"
	"github.com/hootuu/htgo/treasx"
	"github.com/hootuu/htgo/yamax"
	"github.com/hootuu/htgoapi/nvwa"
	"github.com/hootuu/htgoapi/treas"
	"github.com/hootuu/htgoapi/yama"
	"github.com/hootuu/tome/fq"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/yn"
	"github.com/hootuu/utils/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	htpeer.Initialize(&htpeer.Peer{
		VN:         "MOJU",
		SP:         sp.NilID,
		PrivateKey: "0xab3a2cc7edbcfb004be08ef6b61fa9ace77183ce6546ab7f2a969c7daefda55d",
	})
	//testCoinIssue()
	//testSpCreate()
	testYinPlant()
}

func testCoinIssue() {
	req := treas.CoinIssueReq{
		Token:    "37d01ecdf84ae69049c118ce2af8b3f9",
		Issuer:   "MOJU",
		Coin:     "MJC",
		Wei:      10,
		Issuance: 2000000000,
	}
	resp, err := treasx.Client.CoinIssue(req)
	if err != nil {
		return
	}
	logger.Logger.Info("testCoinIssue.resp", zap.Any("resp", resp))
}

func testSpCreate() {
	req := nvwa.SPCreateReq{
		Token: "ac4e8cd0d7c13dbe3877eb94d5693fcd",
		VN:    "MOJU",
		Scope: "",
		ID:    "MJNBMT",
		FQ:    []*fq.FQ{{Fq: fq.DIAMOND}},
	}
	resp, err := nvwax.Client.SPCreate(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}

func testYinPlant() {
	yin, _ := yn.NewYin(
		htpeer.Here().VN.S(),
		"MJNBMT",
	)
	_ = yin.WithTitle(xid.New().String())
	_ = yin.WithWho(ki.UnkADR.S(), "MOJU.User", fmt.Sprintf("MJUSER_%d", time.Now().Unix()))
	_ = yin.WithAct("ORDER")
	_ = yin.WithWhat("MOJU.Commodity", fmt.Sprintf("MJCOMM_%d", time.Now().Unix()))
	_ = yin.WithExpense("CNY", int64(rand.Intn(1000)))
	tags := []string{
		"SPORT", "FRUIT", "EDU.", "ORG", "SOFT",
	}
	_ = yin.WithTag(tags[rand.Intn(len(tags))])

	req := yama.YinPlantReq{Yin: yin}
	logger.Logger.Info("testYinPlant.req", zap.Any("req", req))
	resp, err := yamax.Client.YinPlant(req)
	if err != nil {
		return
	}
	logger.Logger.Info("testYinPlant.resp", zap.Any("resp", resp))
}
