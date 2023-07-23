package main

/*
#include <stdio.h>
#include <stdint.h>
*/
import "C"
import "open_im_sdk/open_im_sdk"

//export GetSpecifiedFriendsInfo
func GetSpecifiedFriendsInfo(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetSpecifiedFriendsInfo",
	}
	open_im_sdk.GetSpecifiedFriendsInfo(callBack, id, C.GoString(userIDList))
}

//export GetFriendList
func GetFriendList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetFriendList",
	}
	open_im_sdk.GetFriendList(callBack, id)
}

//export GetFriendListPage
func GetFriendListPage(operationID *C.char, offset C.int32_t, count C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetFriendListPage",
	}
	open_im_sdk.GetFriendListPage(callBack, id, int32(offset), int32(count))
}

//export SearchFriends
func SearchFriends(operationID *C.char, searchParam *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SearchFriends",
	}
	open_im_sdk.SearchFriends(callBack, id, C.GoString(searchParam))
}

//export CheckFriend
func CheckFriend(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "CheckFriend",
	}
	open_im_sdk.CheckFriend(callBack, id, C.GoString(userIDList))
}

//export AddFriend
func AddFriend(operationID *C.char, userIDReqMsg *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "AddFriend",
	}
	open_im_sdk.AddFriend(callBack, id, C.GoString(userIDReqMsg))
}

//export SetFriendRemark
func SetFriendRemark(operationID *C.char, userIDRemark *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetFriendRemark",
	}
	open_im_sdk.SetFriendRemark(callBack, id, C.GoString(userIDRemark))
}

//export DeleteFriend
func DeleteFriend(operationID *C.char, friendUserID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteFriend",
	}
	open_im_sdk.DeleteFriend(callBack, id, C.GoString(friendUserID))
}

//export GetFriendApplicationListAsRecipient
func GetFriendApplicationListAsRecipient(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetFriendApplicationListAsRecipient",
	}
	open_im_sdk.GetFriendApplicationListAsRecipient(callBack, id)
}

//export GetFriendApplicationListAsApplicant
func GetFriendApplicationListAsApplicant(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetFriendApplicationListAsApplicant",
	}
	open_im_sdk.GetFriendApplicationListAsApplicant(callBack, id)
}

//export AcceptFriendApplication
func AcceptFriendApplication(operationID *C.char, userIDHandleMsg *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "AcceptFriendApplication",
	}
	open_im_sdk.AcceptFriendApplication(callBack, id, C.GoString(userIDHandleMsg))
}

//export RefuseFriendApplication
func RefuseFriendApplication(operationID *C.char, userIDHandleMsg *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "RefuseFriendApplication",
	}
	open_im_sdk.RefuseFriendApplication(callBack, id, C.GoString(userIDHandleMsg))
}

//export AddBlack
func AddBlack(operationID *C.char, blackUserID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "AddBlack",
	}
	open_im_sdk.AddBlack(callBack, id, C.GoString(blackUserID))
}

//export GetBlackList
func GetBlackList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetBlackList",
	}
	open_im_sdk.GetBlackList(callBack, id)
}

//export RemoveBlack
func RemoveBlack(operationID *C.char, removeUserID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "RemoveBlack",
	}
	open_im_sdk.RemoveBlack(callBack, id, C.GoString(removeUserID))
}
