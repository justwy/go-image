package cognitiveservice

import (
	"net/http"
	"strings"
)

// based on microsoft cognitive service
// https://dev.projectoxford.ai/docs/services/563879b61984550e40cbbe8d/operations/563879b61984550f30395244

const faceBaseURL = "https://api.projectoxford.ai/face/v1.0/detect?returnFaceId=true&returnFaceLandmarks=true&returnFaceAttributes=age,gender,headPose,smile,facialHair,glasses"

const (
	AgeAttr string = "age"        // an age number in years
	GenderAttr string = "gender"     // male or female
	HeadPoseAttr string = "headPose"   // 3-D roll/yew/pitch angles for face direction. Pitch value is reserved to 0
	SmileAttr string = "smile"      // smile intensity, a number between [0,1]
	FacialHairAttr string = "facialHair" // consists of lengths of three facial hair areas: moustache, beard and sideburns
	GlassesAttr string = "glasses"    // glasses type. Possible values are 'noGlasses', 'readingGlasses', 'sunglasses', 'swimmingGoggles'
)

const (
	Male string = "male"   // Male in string
	Female string = "female" // Female in string
)

// GlassType in string
type GlassType string

const (
	NoGlasses GlassType = "noGlasses"       // no glasses
	ReadingGlasses GlassType = "readingGlasses"  // reading glasses
	Sunglasses GlassType = "sunglasses"      // sun glasses
	SwimmingGoggles GlassType = "swimmingGoggles" // swimming goggles
)

// Point represents a position
type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type FaceLandmarks struct {
	PupilLeft           Point `json:"pupilLeft"`
	PupilRight          Point `json:"pupilRight"`
	NoseTip             Point `json:"noseTip"`
	MouthLeft           Point `json:"mouthLeft"`
	MouthRight          Point `json:"mouthRight"`
	EyebrowLeftOuter    Point `json:"eyebrowLeftOuter"`
	EyebrowLeftInner    Point `json:"eyebrowLeftInner"`
	EyeLeftOuter        Point `json:"eyeLeftOuter"`
	EyeLeftTop          Point `json:"eyeLeftTop"`
	EyeLeftBottom       Point `json:"eyeLeftBotom"`
	EyeLeftInner        Point `json:"eyeLeftInner"`
	EyebrowRightInner   Point `json:"eyebrowRightInner"`
	EyebrowRightOuter   Point `json:"eyebrowRightOuter"`
	EyeRightInner       Point `json:"eyeRightInner"`
	EyeRightTop         Point `json:"eyeRightTop"`
	EyeRightBottom      Point `json:"eyeRightBottom"`
	EyeRightOuter       Point `json:"eyeRightOuter"`
	NoseRootLeft        Point `json:"noseRootLeft"`
	NoseRootRight       Point `json:"noseRootRight"`
	NoseLeftAlarTop     Point `json:"noseLeftAlarTop"`
	NoseRightAlarTop    Point `json:"noseRightAlarTop"`
	NoseLeftAlarOutTip  Point `json:"noseLeftAlarOutTip"`
	NoseRightAlarOutTip Point `json:"noseRightAlarOutTip"`
	UpperLipTop         Point `json:"upperLipTop"`
	UpperLipBottom      Point `json:"upperLipBottom"`
	UnderLipTop         Point `json:"underLipTop"`
	UnderLipBottom      Point `json:"underLipBottom"`
}

// FacialHair consists of lengths of three facial hair areas: moustache, beard and sideburns.
type FacialHair struct {
	Mustache  float32 `json:"mustache"`
	Beard     float32 `json:"beard"`
	Sideburns float32 `json:"sideburns"`
}

// HeadPose is 3-D roll/yew/pitch angles for face direction. Pitch value is reserved to 0
type HeadPose struct {
	Roll  float32 `json:"roll"`
	Yaw   float32 `json:"yaw"`
	Pitch float32 `json:"pitch"`
}

// Face Attributes:
//   age: an age number in years.
//   gender: male or female.
//   smile: smile intensity, a number between [0,1]
//   facialHair: consists of lengths of three facial hair areas: moustache, beard and sideburns.
//   headPose: 3-D roll/yew/pitch angles for face direction. Pitch value is reserved to 0.
//   glasses: glasses type. Possible values are 'noGlasses', 'readingGlasses', 'sunglasses', 'swimmingGoggles'.
type FaceAttributes struct {
	Age        float32 `json:"age"`
	Gender     string `json:"gender"`
	Smile      float32 `json:"smile"`
	FacialHair FacialHair `json:"facialHair"`
	Glasses    GlassType `json:"glasses"`
	HeadPose   HeadPose `json:"headPose"`
}

type FaceRectangle struct {
	Width float32 `json:"width"`
	Height float32 `json:"height"`
	Left float32 `json:"left"`
	Top float32 `json:"top"`
}

type DetectResponse struct {
	// Id of the detected face, created by detection API.
	// To return this, it requires "returnFaceId" parameter to be true.
	FaceId         string        `json:"faceID"`

	FaceRectangle FaceRectangle `json:"faceRectangle"`

	// An array of 27-point face landmarks pointing to the important positions of face components.
	// To return this, it requires "returnFaceLandmarks" parameter to be true.
	FaceLandmarks  FaceLandmarks `json:"faceLandmarks"`

	// Face Attributes:
	//   age: an age number in years.
	//   gender: male or female.
	//   smile: smile intensity, a number between [0,1]
	//   facialHair: consists of lengths of three facial hair areas: moustache, beard and sideburns.
	//   headPose: 3-D roll/yew/pitch angles for face direction. Pitch value is reserved to 0.
	//   glasses: glasses type. Possible values are 'noGlasses', 'readingGlasses', 'sunglasses', 'swimmingGoggles'.
	FaceAttributes FaceAttributes `json:"faceAttributes"`
}

type FindSimilarRequest struct {
	// Id of the detected face, created by detection API.
	// To return this, it requires "returnFaceId" parameter to be true.
	FaceId                     string

	// An array of 27-point face landmarks pointing to the important positions of face components.
	// To return this, it requires "returnFaceLandmarks" parameter to be true.
	FaceListId                 string

	MaxNumOfCandidatesReturned int
}

type FindSimilarResponse struct {
	PersistedFaceId string
	FaceId          string
	Confidence      float32
}

type FaceAPI interface {
	// Detect human faces in an image and returns face locations,
	// and optionally with face ID, landmarks, and attributes.
	Detect(url string) ([]DetectResponse, error)

	// looking faces for a query face from a list of candidate faces (given by a face list or a face ID array)
	// and return similar face IDs ranked by similarity. The candidate face list has a limitation of 1000 faces.
	// The first return argument is persistedFaceId.
	// FindSimilarByFaceListId(faceId string, faceListId string, maxNumOfCandidatesReturned int) (string, float32, error)

	// Similar to FindSimilarByFaceListId. The difference is that this interface takes an array of faceId instead
	// of a string of faceListId. Also, the first return in this interface is the faecId instead of persistedFaceId.
	// FindSimilarByFaceIds(faceId string, faceIds []string, maxNumOfCandidatesReturned int) (string, float32, error)
}

type MicrosoftFaceAPI struct {
	// API base url
	BaseURL string
	// API key
	APIKey  string
}

func (faceAPI MicrosoftFaceAPI) Detect(url string) ([]DetectResponse, error) {

	queryURL := faceAPI.BaseURL

	detectResponse := []DetectResponse{}

	err := commonHTTPRequest(http.MethodPost, queryURL, faceAPI.APIKey, strings.NewReader(`{"url": "` + url  + `"}`), &detectResponse)

	return detectResponse, err
}

// NewMicrosoftFaceAPI creates an instance of MicrosoftFaceApi
func NewMicrosoftFaceAPI(apiKey string) MicrosoftFaceAPI {
	return MicrosoftFaceAPI{faceBaseURL, apiKey}
}

// for Test
func NewMicrosoftFaceAPIWithURL(baseUrl string, apiKey string) MicrosoftFaceAPI {
	return MicrosoftFaceAPI{baseUrl, apiKey}
}
