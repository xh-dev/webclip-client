package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/xh-dev/webclip-client/model"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

var action string
var host string
var outUrl bool
var toClip bool
var fromClip bool

func main() {
	flag.StringVar(&action, "action", "none", "send - send a message\nretrive - retrive a message")
	flag.StringVar(&host, "host", "https://webclip2.mytools.express", "url posting \n\tsend - /msg/create\n\tretrive - /msg/retrive")
	flag.BoolVar(&outUrl, "asUrl", false, "[mode==send] export result as the retrive url")
	flag.BoolVar(&fromClip, "ic", false, "[mode==send] input msg from clipboard")
	flag.BoolVar(&toClip, "oc", false, "also output result to clipboard")
	flag.Parse()

	if action == "none" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	var msg string
	var err error
	var code string
	switch action {
	case "send":
		if fromClip {
			msg, err = clipboard.ReadAll()
			if err != nil{
				panic(errors.New("Fail to read clipboard text"))
			}
		} else {
			msg = strings.Join(flag.Args(), " ")
		}
		code, err = SendMessage(msg, host+"/api")
		if err != nil {
			panic(err)
		}
		
		var o string
		if outUrl{
			o = host+"/#/get?id="+code
			print(o)
			if toClip {
				clipboard.WriteAll(o)
			}
		}else{
			o = code
			print(o)
			if toClip {
				clipboard.WriteAll(o)
			}
		}
	case "retrive":
		retriveCode, err := strconv.Atoi(strings.Join(flag.Args(), " "))
		if err != nil {
			panic(err)
		}
		msg, err := RetriveMessage(retriveCode, host+"/api")
		if err != nil {
			panic(err)
		}
		print(msg)
		if toClip {
			clipboard.WriteAll(msg)
		}
	}

}
