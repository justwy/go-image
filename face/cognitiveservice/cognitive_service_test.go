package cognitiveservice_test

import (
	"testing"

	"github.com/justwy/treqme/face/cognitiveservice"
)

func TestNewMicrosoftCognitiveService(t *testing.T) {
	cognitiveservice.NewMicrosoftCognitiveService("testKey")
}
