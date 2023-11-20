package rest

import (
	"github.com/go-resty/resty/v2"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"github.com/hootuu/utils/sys"
	"go.uber.org/zap"
	"time"
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
	s := time.Now().UnixMilli()
	defer func() {
		sys.Info("Elapse ", time.Now().UnixMilli()-s, " ms")
	}()
	logger.Logger.Info("Call ", zap.String("url", url))
	var toResp T
	_, nErr := req.SetResult(&toResp).Post(url)
	logger.Logger.Info("Call Return", zap.String("url", url))
	if nErr != nil {
		logger.Logger.Error("resty.post failed", zap.Error(nErr))
		return nil, errors.Sys("go [" + url + " ] failed: " + nErr.Error())
	}
	return &toResp, nil
}
