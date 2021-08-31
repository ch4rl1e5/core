package httphelper

import (
	"encoding/json"
	"github.com/ch4rl1e5/go-common/logger"
	"go.uber.org/zap"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, response interface{}) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.ZapLogger.Panic("error:", zap.Error(err))
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.ZapLogger.Panic("error:", zap.Error(err))
		return
	}
}
