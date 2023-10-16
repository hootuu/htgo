package rest

import (
	"github.com/go-resty/resty/v2"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"go.uber.org/zap"
)

func Go[T any](
	url string,
	data interface{},
	inject ...func(req *resty.Request) *errors.Error,
) (*T, *errors.Error) {
	req := NewClient().R().SetBody(data)
	if len(inject) > 0 {
		err := inject[0](req)
		if err != nil {
			return nil, err
		}
	}
	var toResp T
	_, nErr := req.SetResult(&toResp).Post(url)
	if nErr != nil {
		logger.Logger.Error("resty.post failed", zap.Error(nErr))
		return nil, errors.Sys("go [" + url + " ] failed: " + nErr.Error())
	}
	return &toResp, nil
}
