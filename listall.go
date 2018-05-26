package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"gitlab.qiniu.io/pingan/libqiniu"
)

var chinaEastZone = libqiniu.Zone{
	RsHost:  "http://rs.qiniu.com",
	RsfHost: "http://rsf.qiniu.com",
	UpHosts: []string{"http://up.qiniup.com", "http://up-nb.qiniup.com", "http://up-xs.qiniup.com"},
}

func main() {
	var (
		file       libqiniu.File
		httpClient http.Client
		aksk       libqiniu.AccessKeySecretKey
	)

	aksk.AccessKey = "dGx9Fiz8DZOIP_ActZdmJs61_qSadWaF6biynyDB"
	aksk.SecretKey = "XhOFnbWCQeCmw6l3TMCBOc7Ej9jkAbRpK0m8NJMg"

	bucket := libqiniu.NewClient(&httpClient).Authorize(aksk).Zone(&chinaEastZone).Bucket(os.Args[1])
	iter := bucket.List(context.Background(), "")
	for iter.Next(&file) {
		fmt.Printf("%s\n", file.Key)
	}
	if err := iter.Err(); err != nil {
		log.Fatalln(err)
	}
}
