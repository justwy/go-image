package cognitiveservice_test

import (
	"testing"
	"github.com/justwy/treqme/face/cognitiveservice"
	"fmt"
	"net/http/httptest"
	"net/http"
)

const testAPIKey = "abc"

func TestMicrosoftFaceAPI_Detect(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			fmt.Fprint(w, `
[
    {
        "faceId": "c5c24a82-6845-4031-9d5d-978df9175426",
        "faceRectangle": {
            "width": 78,
            "height": 78,
            "left": 394,
            "top": 54
        },
        "faceLandmarks": {
            "pupilLeft": {
                "x": 412.7,
                "y": 78.4
            },
            "pupilRight": {
                "x": 446.8,
                "y": 74.2
            },
            "noseTip": {
                "x": 437.7,
                "y": 92.4
            },
            "mouthLeft": {
                "x": 417.8,
                "y": 114.4
            },
            "mouthRight": {
                "x": 451.3,
                "y": 109.3
            },
            "eyebrowLeftOuter": {
                "x": 397.9,
                "y": 78.5
            },
            "eyebrowLeftInner": {
                "x": 425.4,
                "y": 70.5
            },
            "eyeLeftOuter": {
                "x": 406.7,
                "y": 80.6
            },
            "eyeLeftTop": {
                "x": 412.2,
                "y": 76.2
            },
            "eyeLeftBottom": {
                "x": 413.0,
                "y": 80.1
            },
            "eyeLeftInner": {
                "x": 418.9,
                "y": 78.0
            },
            "eyebrowRightInner": {
                "x": 4.8,
                "y": 69.7
            },
            "eyebrowRightOuter": {
                "x": 5.5,
                "y": 68.5
            },
            "eyeRightInner": {
                "x": 441.5,
                "y": 75.0
            },
            "eyeRightTop": {
                "x": 446.4,
                "y": 71.7
            },
            "eyeRightBottom": {
                "x": 447.0,
                "y": 75.3
            },
            "eyeRightOuter": {
                "x": 451.7,
                "y": 73.4
            },
            "noseRootLeft": {
                "x": 428.0,
                "y": 77.1
            },
            "noseRootRight": {
                "x": 435.8,
                "y": 75.6
            },
            "noseLeftAlarTop": {
                "x": 428.3,
                "y": 89.7
            },
            "noseRightAlarTop": {
                "x": 442.2,
                "y": 87.0
            },
            "noseLeftAlarOutTip": {
                "x": 424.3,
                "y": 96.4
            },
            "noseRightAlarOutTip": {
                "x": 446.6,
                "y": 92.5
            },
            "upperLipTop": {
                "x": 437.6,
                "y": 105.9
            },
            "upperLipBottom": {
                "x": 437.6,
                "y": 108.2
            },
            "underLipTop": {
                "x": 436.8,
                "y": 111.4
            },
            "underLipBottom": {
                "x": 437.3,
                "y": 114.5
            }
        },
        "faceAttributes": {
            "age": 71.0,
            "gender": "male",
            "smile": 0.88,
            "facialHair": {
                "mustache": 0.8,
                "beard": 0.1,
                "sideburns": 0.02
	    },
            "glasses": "sunglasses",
            "headPose": {
                "roll": 2.1,
                "yaw": 3,
                "pitch": 0
            }
        }
    }
]
			`)
		default:
			http.Error(w, "test error", http.StatusInternalServerError)
		}
	}))

	faceAPI := cognitiveservice.NewMicrosoftFaceAPIWithURL(ts.URL, testAPIKey)

	detectResults, err := faceAPI.Detect("http://treq.me/image1.png")
	if err != nil  {
		t.Error("expect no err but was ", err);
	}

	if detectResults == nil || len(detectResults) != 1 {
		t.Error("expect result with 1 result but was nil or 0: ", detectResults);
	}

	if detectResults[0].FaceID != "c5c24a82-6845-4031-9d5d-978df9175426" {
		t.Errorf("Expected c5c24a82-6845-4031-9d5d-978df9175426 but was %s", detectResults[0].FaceID)
	}

}

func TestNewMicrosoftFaceAPI(t *testing.T) {
	faceAPI := cognitiveservice.NewMicrosoftFaceAPI("abc")
	if faceAPI.APIKey != "abc" {
		t.Errorf("Expected abc but was %s", faceAPI.APIKey)
	}
}

func ExampleMicrosoftFaceAPI_Detect() {
	cognitiveService := cognitiveservice.NewMicrosoftCognitiveService("abcdefg")
	img := "https://dl.dropboxusercontent.com/u/test.jpg"
	detectedObjs, _ := cognitiveService.FaceAPI.Detect(img)
	for _, obj := range detectedObjs {
		// Print Face Rectangle information
		fmt.Printf("%+v", obj.FaceRectangle)
	}
}