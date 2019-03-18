package main

import (
	"fmt"
	"github.com/weilaihui/fdfs_client"
	"io/ioutil"
)

func main() {

	ff, _ := ioutil.ReadFile("shop-service/test/1.jpg")
	fmt.Println("image len:", len(ff))
	/*
	hosts := []string{"10.0.1.32"}
	port := 22122
	minConns := 10
	maxConns := 150
	connPool,_ := fdfs_client.NewConnectionPool(hosts, port, minConns, maxConns)
	*/
	path := "shop-service/test/client.conf"
	fds, error := fdfs_client.NewFdfsClient(path)

	if fds == nil {
		fmt.Println("conn error: %s", error)
		var test string
		fmt.Scanln(&test)
		return
	}
	uploadResponse, err := fds.UploadByBuffer(ff, "jpg")

	if uploadResponse == nil {
		fmt.Println("upload error: %s", err)
		var test string
		fmt.Scanln(&test)
		return
	}
	fmt.Println("group name:", uploadResponse.GroupName)
	fmt.Println("remote file id:", uploadResponse.RemoteFileId)

	}
