package manager

import "smartparking/internal/interfaces/processor"

type Processor interface {
	PlateRecognizer() processor.PlateRecognizer
}
