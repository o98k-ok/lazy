package doc

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkdrive "github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
)

type LarkDocer interface {
	Upload(filename string, dat []byte) (string, error)
	FormatEdit(file string, newname string) error
}

type Larkdoc struct {
	AppId  string
	Secret string
	Parent string

	Client *lark.Client
}

func Addr[T any](val T) *T {
	return &val
}

func NewLarkDocer(appid, sec string, folder string) LarkDocer {
	cli := lark.NewClient(appid, sec)
	return &Larkdoc{AppId: appid, Secret: sec, Parent: folder, Client: cli}
}

func (ld *Larkdoc) Upload(filename string, dat []byte) (string, error) {
	req := larkdrive.NewUploadAllFileReqBuilder().Body(&larkdrive.UploadAllFileReqBody{
		FileName:   Addr(filename),
		ParentType: Addr("explorer"),
		ParentNode: Addr(ld.Parent),
		Size:       Addr(len(dat)),
		File:       bytes.NewReader(dat),
	}).Build()
	resp, err := ld.Client.Drive.File.UploadAll(context.Background(), req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK || resp.Code != 0 {
		return "", errors.New("status code is invalid")
	}
	return *resp.Data.FileToken, nil
}

func (ld *Larkdoc) FormatEdit(file string, newname string) error {
	req := larkdrive.NewCreateImportTaskReqBuilder().ImportTask(&larkdrive.ImportTask{
		FileExtension: Addr("md"),
		FileToken:     &file,
		Type:          Addr("docx"),
		FileName:      Addr(newname),
		Point: &larkdrive.ImportTaskMountPoint{
			MountType: Addr(1),
			MountKey:  &ld.Parent,
		},
	}).Build()
	resp, err := ld.Client.Drive.ImportTask.Create(context.Background(), req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK || resp.Code != 0 {
		return errors.New("status code is invalid")
	}

	// DEL file
	delRsp, err := ld.Client.Drive.File.Delete(context.Background(), larkdrive.NewDeleteFileReqBuilder().FileToken(file).Type("file").Build())
	if err != nil {
		return err
	}
	if delRsp.StatusCode != http.StatusOK || delRsp.Code != 0 {
		return errors.New("status code is invalid")
	}
	return nil
}
