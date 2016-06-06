package main

import (
	"./meizar"
	"./rule"
	"flag"
	"net/http"
	"runtime"
	"strings"
)

var dir *string = flag.String("dir", "./images/", "")
var startPage *int = flag.Int("start", 2009, "")
var userCookie *string = flag.String("cookie", "1092990552=903f5Z2FA411DxfPYORhRDmNqohZzkNsnLuvj76u; PHPSESSID=vq5be8aobr23gdt68ig0mmud31; 1092990552=0; _gat=1; jdna=596e6fb28c1bb47f949e65e1ae03f7f5#1465139210979; Hm_lvt_fd93b7fb546adcfbcf80c4fc2b54da2c=1465130953,1465137904; Hm_lpvt_fd93b7fb546adcfbcf80c4fc2b54da2c=1465139211; _ga=GA1.2.1645811469.1465130953", "")
var client *http.Client = &http.Client{}

func init() {
	flag.Parse()
	if !strings.HasSuffix(*dir, "/") {
		*dir = *dir + "/"
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	meizar.New(*dir, *startPage, rule.RuleProvider(), *userCookie, client).Start()
}
