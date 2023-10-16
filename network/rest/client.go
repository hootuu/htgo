package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hootuu/htgoapi/htocol"
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/strs"
	"github.com/hootuu/utils/sys"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type H = map[string]string

const (
	DefaultHOTUGateway = "http://localhost:9090"
)

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
	request.Header.Add(htocol.HeaderNonce, xid.New().String())
	request.Header.Add(htocol.HeaderTimestamp, fmt.Sprintf("%d", time.Now().UnixMilli()))
	request.Header.Add(htocol.HeaderSignature, strs.MD5(request.RequestURI))
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
	gBaseUrl = configure.GetString("hotu.gateway", DefaultHOTUGateway)
	sys.Info("HOTU Gateway [hotu.gateway]: ", gBaseUrl)
}
