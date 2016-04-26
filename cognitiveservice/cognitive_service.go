package cognitiveservice

const (
	personGroupBaseURL string = "https://api.projectoxford.ai/face/v1.0/persongroups/"
)

// CognitiveService exposes all interfaces
type CognitiveService struct {
	APIKey      string
	contentType string
	PersonGroupAPI
	FaceAPI
}

// NewMicrosoftCognitiveService creates a new CognitiveService object
func NewMicrosoftCognitiveService(apiKey string) CognitiveService {
	return CognitiveService{
		APIKey:         apiKey,
		contentType:    "application/json",
		PersonGroupAPI: NewMicrosoftPersonGroupAPI(personGroupBaseURL, apiKey),
		FaceAPI: NewMicrosoftFaceAPI(apiKey),
	}
}
