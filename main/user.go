package main

/*
#include <stdio.h>
*/
import "C"
import "open_im_sdk/open_im_sdk"

//export GetUsersInfo
func GetUsersInfo(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetUsersInfo",
	}
	open_im_sdk.GetUsersInfo(callBack, id, C.GoString(userIDList))
}

//export GetUsersInfoFromSrv
func GetUsersInfoFromSrv(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetUsersInfoFromSrv",
	}
	open_im_sdk.GetUsersInfo(callBack, id, C.GoString(userIDList))
}

//export SetSelfInfo
func SetSelfInfo(operationID *C.char, userInfo *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetSelfInfo",
	}
	open_im_sdk.SetSelfInfo(callBack, id, C.GoString(userInfo))
}

//export GetSelfUserInfo
func GetSelfUserInfo(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetSelfUserInfo",
	}
	open_im_sdk.GetSelfUserInfo(callBack, id)
}

//export UpdateMsgSenderInfo
func UpdateMsgSenderInfo(operationID *C.char, nickname *C.char, faceURL *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "UpdateMsgSenderInfo",
	}
	open_im_sdk.UpdateMsgSenderInfo(callBack, id, C.GoString(nickname), C.GoString(faceURL))
}
