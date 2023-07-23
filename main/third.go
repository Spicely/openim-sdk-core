package main

/*
#include <stdio.h>
#include <stdint.h>
*/
import "C"
import "open_im_sdk/open_im_sdk"

//export UpdateFcmToken
func UpdateFcmToken(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "UpdateFcmToken",
	}
	open_im_sdk.UpdateFcmToken(callBack, id, C.GoString(userIDList))
}

//export SetAppBadge
func SetAppBadge(operationID *C.char, appUnreadCount C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetAppBadge",
	}
	open_im_sdk.SetAppBadge(callBack, id, int32(appUnreadCount))
}
