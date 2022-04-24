package processor

type RecognitionResponse struct {
	CameraID       string              `json:"camera_id"`
	Filename       string              `json:"filename"`
	ProcessingTime float64             `json:"processing_time"`
	Timestamp      string              `json:"timestamp"`
	Version        int                 `json:"version"`
	Results        []RecognitionResult `json:"results"`
}

type RecognitionResult struct {
	Box        RecognitionBox         `json:"box"`
	Region     RecognitionRegion      `json:"region"`
	Vehicle    RecognitionVehicle     `json:"vehicle"`
	DScore     float64                `json:"dscore"`
	Plate      string                 `json:"plate"`
	Score      float64                `json:"score"`
	Candidates []RecognitionCandidate `json:"candidates"`
}

type RecognitionBox struct {
	XMax int `json:"xmax"`
	XMin int `json:"xmin"`
	YMax int `json:"ymax"`
	YMin int `json:"ymin"`
}

type RecognitionCandidate struct {
	Plate string  `json:"plate"`
	Score float64 `json:"score"`
}

type RecognitionRegion struct {
	Code  string  `json:"code"`
	Score float64 `json:"score"`
}

type RecognitionVehicle struct {
	Box   RecognitionBox `json:"box"`
	Score float64        `json:"score"`
	Type  string         `json:"type"`
}

func (rr RecognitionResponse) GetPlates() []string {
	var plates []string
	for _, result := range rr.Results {
		plates = append(plates, result.Plate)
		for _, candidate := range result.Candidates {
			plates = append(plates, candidate.Plate)
		}
	}
	return plates
}
