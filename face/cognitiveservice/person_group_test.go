package cognitiveservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/justwy/treqme/face/cognitiveservice"
)

var personGroupAPI cognitiveservice.MicrosoftPersonGroupAPI

func TestMain(m *testing.M) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			if r.RequestURI == "/test_group" {
				fmt.Fprint(w, "done")
			} else {
				http.Error(w, "test error", http.StatusInternalServerError)
			}
		case http.MethodGet:
			fmt.Fprint(w, `
{
    "personGroupId": "sample_group",
    "name": "testgroup1",
    "userData":"User-provided data attached to the person group"
}
			`)
		case http.MethodDelete:
			fmt.Fprint(w, "done")
		}
	}))
	defer ts.Close()

	personGroupAPI = cognitiveservice.NewMicrosoftPersonGroupAPI(ts.URL, "test_api_key")

	returnCode := m.Run()

	os.Exit(returnCode)
}

func TestMicrosoftPersonGroupAPI_Create(t *testing.T) {
	err := personGroupAPI.Create(cognitiveservice.PersonGroupObject{
		PersonGroupID: "test_group",
		Name:          "test_name",
		UserData:      "test_user_data",
	})

	if err != nil {
		t.Errorf("expect error is nil but was %v", err)
	}
}

func TestMicrosoftPersonGroupAPI_Create_Error_Code(t *testing.T) {
	err := personGroupAPI.Create(cognitiveservice.PersonGroupObject{
		PersonGroupID: "error_test_group",
		Name:          "test_name",
		UserData:      "test_user_data",
	})

	if err == nil {
		t.Errorf("Expected error but not error")
	}
}

func TestMicrosoftPersonGroupAPI_Query(t *testing.T) {
	personGroupObj, err := personGroupAPI.Query("sample_group")
	if err != nil {
		t.Errorf("expect error is nil but was %v", err)
	}

	if personGroupObj.Name != "testgroup1" {
		t.Errorf("Expected testgroup1 but was %s", personGroupObj.Name)
	}
}

func TestNewMicrosoftPersonGroupAPI_Delete(t *testing.T) {
	err := personGroupAPI.Delete("testgroup")

	if err != nil {
		t.Errorf("expect error is nil but was %v", err)
	}
}

func ExampleMicrosoftPersonGroupAPI_Create() {
	cognitiveService := cognitiveservice.NewMicrosoftCognitiveService("test_key")

	groupID := "test_group"

	err := cognitiveService.PersonGroupAPI.Create(cognitiveservice.PersonGroupObject{PersonGroupID: groupID, Name: "test", UserData: "this is a test"})

	fmt.Println(err)
}

func ExampleMicrosoftPersonGroupAPI_Query() {
	cognitiveService := cognitiveservice.NewMicrosoftCognitiveService("131b5264f0954b608d41daac603276cd")
	groupID := "test_group"
	personGroupObj, err := cognitiveService.PersonGroupAPI.Query(groupID)
	fmt.Println("Get user group: ", personGroupObj, err)
}

func ExampleMicrosoftPersonGroupAPI_Delete() {
	cognitiveService := cognitiveservice.NewMicrosoftCognitiveService("131b5264f0954b608d41daac603276cd")

	groupID := "test_group_id"
	err := cognitiveService.PersonGroupAPI.Delete(groupID)

	fmt.Println(err)
}
