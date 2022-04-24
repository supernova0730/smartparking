package processor

import (
	"bytes"
	"smartparking/internal/processor"
)

type PlateRecognizer interface {
	Execute(filename string, content *bytes.Buffer) (result processor.RecognitionResponse, err error)
}
