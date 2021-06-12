package weblcip

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xh-dev/webclip-client/model"
	"io/ioutil"
	"net/http"
)

func SendMessage(msg string, host string) (string, error) {
	url := host + "/msg/create"

	jString, err := json.Marshal(model.CreateMsg{
		Msg: msg,
	})

	if err != nil {
		return "", err
	}
	// println(string(jString))

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jString))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", errors.New("Status code: " + response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var r model.CreateMsgResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}

func RetriveMessage(code int, host string) (string, error) {
	url := host + "/msg/retrieve"
	jString, err := json.Marshal(model.RetriveMsg{
		Code: code,
	})

	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jString))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", errors.New("Status code: " + response.Status)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var r model.RetriveMsgResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	return r.Msg, nil
}
