package manager

import (
	"smartparking/config"
	"smartparking/internal/interfaces/processor"
	processor2 "smartparking/internal/processor"
	"sync"
)

type processorImpl struct {
	plateRecognizerProcessorInit sync.Once
	plateRecognizerProcessor     processor.PlateRecognizer
}

func (pm *processorImpl) PlateRecognizer() processor.PlateRecognizer {
	pm.plateRecognizerProcessorInit.Do(func() {
		if pm.plateRecognizerProcessor == nil {
			pm.plateRecognizerProcessor = processor2.NewPlateRecognizerProcessor(config.GlobalConfig.Recognizer.URL, config.GlobalConfig.Recognizer.Token)
		}
	})
	return pm.plateRecognizerProcessor
}
