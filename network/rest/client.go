package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hootuu/htgo/network/htpeer"
	"github.com/hootuu/htgoapi/htocol"
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"github.com/hootuu/utils/sys"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type H = map[string]string

func NewClient() *resty.Client {
	cli := resty.New().
		SetBaseURL(gBaseUrl).
		SetPreRequestHook(doSignature).
		OnAfterResponse(doWrapResponse)
	if sys.RunMode.IsRd() {
		cli.EnableTrace()
	}
	return cli
}

func doSignature(_ *resty.Client, request *http.Request) error {
	nonce := xid.New().String()
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	vnStr := htpeer.Here().VN.S()
	spStr := htpeer.Here().SP.S()
	signBuilder := htocol.SignBuilder{}
	signStr, err := signBuilder.
		Add(htocol.HeaderNonce, nonce).
		Add(htocol.HeaderTimestamp, timestamp).
		Add(htocol.HeaderVN, vnStr).
		Add(htocol.HeaderSP, spStr).
		Sign(htpeer.Here().PrivateKey)
	if err != nil {
		logger.Logger.Error("sign failed:", zap.Error(err))
		return err
	}
	request.Header.Add(htocol.HeaderNonce, nonce)
	request.Header.Add(htocol.HeaderTimestamp, timestamp)
	request.Header.Add(htocol.HeaderVN, vnStr)
	request.Header.Add(htocol.HeaderSP, spStr)
	request.Header.Add(htocol.HeaderSignature, signStr)
	return nil
}

func doWrapResponse(_ *resty.Client, response *resty.Response) error {
	if response == nil {
		return nil
	}
	if response.IsSuccess() {
		return nil
	}

	var respErr errors.Error
	err := json.Unmarshal(response.Body(), &respErr)
	if err != nil {
		return errors.Verify("[invalid - data] invalid response")
	}
	return &respErr
}

var gBaseUrl string

func init() {
	gBaseUrl = configure.GetString("hotu.gateway", "http://localhost:9090")
	sys.Info("HOTU Gateway [hotu.gateway]: ", gBaseUrl)
}
