package main

import "C"
import "open_im_sdk/open_im_sdk"

type UploadFileCallback struct {
	operationID string
	methodName  string
}

func (u *UploadFileCallback) Open(size int64) {
	callBack("Open", u.operationID, u.methodName, size, nil)
}

func (u *UploadFileCallback) PartSize(partSize int64, num int) {
	callBack("PartSize", u.operationID, u.methodName, nil, nil)
}

func (u *UploadFileCallback) HashPartProgress(index int, size int64, partHash string) {
	callBack("HashPartProgress", u.operationID, u.methodName, nil, nil)
}

func (u *UploadFileCallback) HashPartComplete(partsHash string, fileHash string) {
	callBack("HashPartComplete", u.operationID, u.methodName, nil, nil)
}

func (u *UploadFileCallback) UploadID(uploadID string) {
	callBack("UploadID", u.operationID, u.methodName, nil, nil)
}

func (u *UploadFileCallback) UploadPartComplete(index int, partSize int64, partHash string) {
	callBack("UploadPartComplete", u.operationID, u.methodName, nil, nil)
}

func (u *UploadFileCallback) UploadComplete(fileSize int64, streamSize int64, storageSize int64) {
	callBack("UploadComplete", u.operationID, u.methodName, nil, nil)
}

func (u *UploadFileCallback) Complete(size int64, url string, typ int) {
	callBack("Complete", u.operationID, u.methodName, nil, nil)

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
