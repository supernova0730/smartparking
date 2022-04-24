package processor

import (
	"bytes"
	"go.uber.org/zap"
	"smartparking/pkg/client"
	"smartparking/pkg/logger"
)

type plateRecognizerProcessor struct {
	url   string
	token string
}

func NewPlateRecognizerProcessor(url, token string) *plateRecognizerProcessor {
	return &plateRecognizerProcessor{url: url, token: token}
}

func (prc *plateRecognizerProcessor) Execute(filename string, content *bytes.Buffer) (result RecognitionResponse, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("plateRecognizerProcessor.Execute failed", zap.Error(err))
		} else {
			logger.Log.Info("plateRecognizerProcessor.Execute success", zap.Any("result", result))
		}
	}()

	err = client.NewMultipartFormClient(prc.url+"/v1/plate-reader").
		SetFile("upload", filename, content).
		SetHeader("Authorization", "Token "+prc.token).
		Do(&result)
	return
}
