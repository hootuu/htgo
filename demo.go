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
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/tome/uc/nft"
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
		PrivateKey: "0xdbbab671ebc5c1a75c0d842c3fc6291600a5a599d5cb230172fd7a73d6f9fbf1",
	})
	//testNftIssue()
	testNftAirdrop("bafyreibzsfgr62xgsnkdkgts4r6ylgllnekrnjobebf54azmaw37gwd3iq",
		"0x50dc95e9a14b45e726325Fda95A0E81810f95174")
	//testValueLoad()
	//testCoinIssue()
	//testSpCreate()
	//for i := 0; i < 1; i++ {
	//	testYinPlant()
	//}
	//s := time.Now().UnixMilli()
	//var max int64 = 10000000
	//var i int64 = 0
	//for i = 0; i < max; i++ {
	//	testValueLoad()
	//	testAlterLoad()
	//
	//	e := time.Now().UnixMilli()
	//	fmt.Println("[", i, "] elapse : ", e-s, " ms; each", (e-s)/max, " ms")
	//}
	//e := time.Now().UnixMilli()
	//fmt.Println("elapse : ", e-s, " ms; each", (e-s)/max, " ms")

	//var totalElapse int64 = 0
	//for i := 0; i < 10000; i++ {
	//	s := time.Now().UnixMilli()
	//	testYinPlant()
	//	e := time.Now().UnixMilli()
	//	totalElapse += e - s
	//	fmt.Println("##### AVG Elapse UnixMilli: ", totalElapse/int64(i+1), " ms [", i+1, "]#####")
	//	waitSeconds := 5 + rand.Int63n(20)
	//	time.Sleep(time.Duration(waitSeconds) * time.Second)
	//}
	//testValueLoad()
	//testAlterLoad()
}

func testCoinIssue() {
	req := treas.CoinIssueReq{
		Token:    "d38d64f082d693b2bc2dc7cb6a425fcc",
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
		Token: "ce57725b9d25c870213dcb2bc17b6c2b",
		VN:    "MOJU",
		Scope: "",
		ID:    "MJGOLD",
		FQ:    []*fq.FQ{{Fq: fq.GOLD}},
	}
	resp, err := nvwax.Client.SPCreate(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}

func testValueLoad() {
	req := yama.ValueLoadReq{
		VN:   "MOJU",
		Page: nil,
	}
	resp, err := yamax.Client.ValueLoad(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}

func testAlterLoad() {
	req := treas.AlterLoadReq{
		Lead: &uc.AccountLead{
			Owner: "0xfcDd9761E6be3fBcA401ca687bb08C8E5a0266f6",
			Coin:  "MJC",
		},
		Page: nil,
	}
	resp, err := treasx.Client.AlterLoad(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}

func testYinPlant() {
	yin, _ := yn.NewYin(
		htpeer.Here().VN.S(),
		//"MJTEST",
		"MJGOLD",
		//"PIQUXD",
		//"HZPUFF",
	)
	_ = yin.WithTitle(xid.New().String())
	_ = yin.WithWho(ki.UnkADR.S(), "MOJU.User", fmt.Sprintf("MJUSER_%s", xid.New().String()))
	_ = yin.WithAct("ORDER")
	_ = yin.WithWhat("MOJU.Commodity", fmt.Sprintf("MJCOMM_%d", time.Now().Unix()))
	_ = yin.WithExpense("CNY", 300)
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

func testNftIssue() {
	req := treas.NftIssueReq{
		VN:          "MOJU",
		Category:    "PIQU",
		Token:       nft.Token(fmt.Sprintf("PIQU_%d", time.Now().Unix())),
		Tag:         []nft.Tag{},
		Link:        "",
		Title:       "",
		Description: "",
		Meta:        nft.NewMeta(),
	}
	resp, err := treasx.Client.NftIssue(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}

func testNftAirdrop(nid nft.NID, owner ki.ADR) {
	req := treas.NftAirdropReq{
		NID:   nid,
		Owner: owner,
	}
	resp, err := treasx.Client.NftAirdrop(req)
	if err != nil {
		return
	}
	logger.Logger.Info("resp", zap.Any("resp", resp))
}
