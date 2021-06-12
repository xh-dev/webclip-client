package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	wc "github.com/xh-dev/webclip-client"
	"os"
	"strconv"
	"strings"
)

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
		code, err = wc.SendMessage(msg, host+"/api")
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
		msg, err := wc.RetriveMessage(retriveCode, host+"/api")
		if err != nil {
			panic(err)
		}
		print(msg)
		if toClip {
			clipboard.WriteAll(msg)
		}
	}

}
