package main

import "C"
import (
	"encoding/json"
	"open_im_sdk/open_im_sdk"
)

type UploadFileCallback struct {
	operationID string
	methodName  string
}

func (u *UploadFileCallback) Open(size int64) {
	data := map[string]int64{"size": size}
	jsonData, _ := json.Marshal(data)
	callBack("Open", u.operationID, u.methodName, nil, jsonData)
}

func (u *UploadFileCallback) PartSize(partSize int64, num int) {
	data := map[string]interface{}{
		"partSize": partSize,
		"num":      num,
	}
	jsonData, _ := json.Marshal(data)
	callBack("PartSize", u.operationID, u.methodName, nil, jsonData)
}

func (u *UploadFileCallback) HashPartProgress(index int, size int64, partHash string) {
	data := map[string]interface{}{
		"index":    index,
		"size":     size,
		"partHash": partHash,
	}
	jsonData, _ := json.Marshal(data)
	callBack("HashPartProgress", u.operationID, u.methodName, nil, jsonData)
}

func (u *UploadFileCallback) HashPartComplete(partsHash string, fileHash string) {
	data := map[string]interface{}{
		"fileHash":  fileHash,
		"partsHash": partsHash,
	}
	jsonData, _ := json.Marshal(data)
	callBack("HashPartComplete", u.operationID, u.methodName, nil, jsonData)
}

func (u *UploadFileCallback) UploadID(uploadID string) {
	callBack("UploadID", u.operationID, u.methodName, nil, uploadID)
}

func (u *UploadFileCallback) UploadPartComplete(index int, partSize int64, partHash string) {
	data := map[string]interface{}{
		"index":    index,
		"partSize": partSize,
		"partHash": partHash,
	}
	jsonData, _ := json.Marshal(data)
	callBack("UploadPartComplete", u.operationID, u.methodName, nil, jsonData)
}

func (u *UploadFileCallback) UploadComplete(fileSize int64, streamSize int64, storageSize int64) {
	data := map[string]interface{}{
		"fileSize":    fileSize,
		"streamSize":  streamSize,
		"storageSize": storageSize,
	}
	jsonData, _ := json.Marshal(data)
	callBack("UploadComplete", u.operationID, u.methodName, nil, jsonData)
}

func (u *UploadFileCallback) Complete(size int64, url string, typ int) {
	data := map[string]interface{}{
		"size": size,
		"url":  url,
		"typ":  typ,
	}
	jsonData, _ := json.Marshal(data)
	callBack("Complete", u.operationID, u.methodName, nil, jsonData)

}

//export UploadFile
func UploadFile(operationID *C.char, req *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "UploadFile",
	}
	uploadFileCallback := &UploadFileCallback{
		operationID: id,
		methodName:  "UploadFile",
	}
	open_im_sdk.UploadFile(callBack, id, C.GoString(req), uploadFileCallback)
}
