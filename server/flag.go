package server

import (
	"flag"
	"fmt"
	"os"
)

var (
	TokenUrl string = ""
	h        bool
)

func init() {
	flag.StringVar(&TokenUrl, "u", "", `if you set the value, app will check token with it
    this api must be:    https: / http:// {{host}}:{{port}}/XXXXX

    app will request  https: / http:// {{host}}:{{port}}/XXXXX?token=XXXXXXX


    you must response 

    {
      status: int,    ## 0:ok 1:error
      gid:    string,
      name:   string,
      uid:    string,
    }
 
	`)
	flag.BoolVar(&h, "h", false, "this help")
	flag.Parse()

	if h {
		flag.Usage()
		os.Exit(0)
	}

}

func usage() {

	fmt.Fprintf(os.Stderr, `gows version: gows/1.0.1
Usage: gows [-h]  [-u TokenCheckUrl] 

Options:

`)

	flag.PrintDefaults()

}
