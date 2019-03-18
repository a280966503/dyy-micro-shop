package upload

import (
	"dyy-micro-shop/manager-service/proto/upload"
	"dyy-micro-shop/utils/path"
	"fmt"
	"github.com/weilaihui/fdfs_client"
	"golang.org/x/net/context"
)

type Upload struct {
}


func (h *Upload) UploadFile(ctx context.Context, in *upload.ReqUpload, out *upload.Resp) error {

	fmt.Println(len(in.Img))
	tracker := &fdfs_client.Tracker{
		[]string{"192.168.25.133"},
		22122,
	}
	client, err := fdfs_client.NewFdfsClientByTracker(tracker)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err != nil {
		fmt.Println("1",err)
		return err
	}

	response, err := client.UploadByBuffer(in.Img, "jpg")

	if err != nil {
		fmt.Println("2",err)
		return err
	}

	*out=upload.Resp{Flag:true,Message:path.IMG_PATH+response.RemoteFileId}
	return nil
}

//
//func (h *Upload) UploadFile(ctx context.Context, in *upload.ReqUpload, out *upload.Resp) error {
//
//	tracker := &fdfs_client.Tracker{
//		[]string{"192.168.25.133"},
//		22122,
//	}
//	client, err := fdfs_client.NewFdfsClientByTracker(tracker)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	if err != nil {
//		fmt.Println("1",err)
//		return err
//	}
//
//	response, err := client.UploadByBuffer(in.Img, "jpg")
//
//	if err != nil {
//		fmt.Println("2",err)
//		return err
//	}
//
//	*out=upload.Resp{Flag:true,Message:path.IMG_PATH+response.RemoteFileId}
//	return nil
//}