package cognitiveservice

import (
	//"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"encoding/json"
)

const apiKeyName = "Ocp-Apim-Subscription-Key";

func commonHTTPRequest(
	httpMethod string, url string, apiKey string, requestBody io.Reader, responseObject interface{}) error {

	client := &http.Client{}

	req, _ := http.NewRequest(httpMethod, url, requestBody)

	req.Header.Add(apiKeyName, apiKey)

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return errors.New("commonHTTPRequest: " + resp.Status + " " + string(bodyData))
	}

	// No need to look at the body for requests like update/new
	if responseObject == nil {
		return nil
	}

	err = json.Unmarshal(bodyData, responseObject)
	return err
}
