package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	wc "github.com/xh-dev/webclip-client"
	"os"
	"strings"
)

//go:embed version.txt
var version string

var action string
var host string
var outUrl bool
var toClip bool
var fromClip bool

func main() {
	flag.StringVar(&action, "action", "none", "send - send a message\nretrive - retrive a message")
	flag.StringVar(&host, "host", "https://webclip.webutils.space", "url posting \n\tsend - /msg/create\n\tretrive - /msg/retrive")
	flag.BoolVar(&outUrl, "asUrl", false, "[mode==send] export result as the retrive url")
	flag.BoolVar(&fromClip, "ic", false, "[mode==send] input msg from clipboard")
	flag.BoolVar(&toClip, "oc", false, "also output result to clipboard")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "webclip version %s\n\n", strings.TrimSpace(version))
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s -action send [-ic] [-asUrl] [-oc] [-host <url>] <message...>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "      send a message (or read it from the clipboard with -ic);\n")
		fmt.Fprintf(os.Stderr, "      prints \"<code> <secret>\", or the retrieve url with -asUrl\n")
		fmt.Fprintf(os.Stderr, "  %s -action retrive [-oc] [-host <url>] <code> <secret>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "      retrieve a message using the <code> and <secret> returned by send\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if action == "none" {
		flag.Usage()
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
		var secret string
		code, secret, err = wc.SendMessage(msg, host+"/api")
		if err != nil {
			panic(err)
		}
		
		var o string
		if outUrl{
			o = host+"/#/get?id="+code+"&secret="+secret
			print(o)
			if toClip {
				clipboard.WriteAll(o)
			}
		}else{
			o = code + " " + secret
			print(o)
			if toClip {
				clipboard.WriteAll(o)
			}
		}
	case "retrive":
		if flag.NArg() != 2 {
			panic(errors.New("usage: -action retrive <code> <secret>"))
		}
		msg, err := wc.RetriveMessage(flag.Arg(0), flag.Arg(1), host+"/api")
		if err != nil {
			panic(err)
		}
		print(msg)
		if toClip {
			clipboard.WriteAll(msg)
		}
	}

}
