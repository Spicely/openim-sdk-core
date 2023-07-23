package main

/*
#include <stdio.h>
#include "include/dart_api_dl.h"

typedef struct {
    void (*onMethodChannel)(Dart_Port_DL port, char*, char*, char*, double*, char*);
	void (*onNativeMethodChannel)(char*, char*, char*, double*, char*);
} CGO_OpenIM_Listener;

static void callOnMethodChannel(CGO_OpenIM_Listener *listener, Dart_Port_DL port, char* methodName, char* operationID, char* callMethodName, double* errCode, char* message) {
	if (listener->onMethodChannel != NULL) {
		listener->onMethodChannel(port, methodName, operationID, callMethodName, errCode, message);
	}
}
static void callOnNativeMethodChannel(CGO_OpenIM_Listener *listener, char* methodName, char* operationID, char* callMethodName, double* errCode, char* message) {
	if (listener->onNativeMethodChannel != NULL) {
		listener->onNativeMethodChannel(methodName, operationID, callMethodName, errCode, message);
	}
}
*/
import "C"
import (
	"open_im_sdk/open_im_sdk"
)

var openIMListener *C.CGO_OpenIM_Listener

var main_isolate_send_port C.Dart_Port_DL

var isListenerInit = false
var initSDK = false

//export RegisterCallback
func RegisterCallback(callback *C.CGO_OpenIM_Listener, port C.Dart_Port_DL) {
	openIMListener = callback
	main_isolate_send_port = port
	initListener()
}

func initListener() {
	if isListenerInit {
		return
	}
	isListenerInit = true

	groupListener := &GroupListener{}
	open_im_sdk.SetGroupListener(groupListener)

	conversationListener := &ConversationListener{}
	open_im_sdk.SetConversationListener(conversationListener)

	advancedMsgListener := &AdvancedMsgListener{}
	open_im_sdk.SetAdvancedMsgListener(advancedMsgListener)

	batchMsgListener := &BatchMsgListener{}
	open_im_sdk.SetBatchMsgListener(batchMsgListener)

	userListener := &UserListener{}
	open_im_sdk.SetUserListener(userListener)

	friendshipListener := &FriendshipListener{}
	open_im_sdk.SetFriendListener(friendshipListener)

	messageKvInfoListener := &MessageKvInfoListener{}
	open_im_sdk.SetMessageKvInfoListener(messageKvInfoListener)

}

func callBack(methodName string, operationID interface{}, callMethodName interface{}, errCode interface{}, message interface{}) {
	cMethodName := C.CString(methodName)

	var cOperationID *C.char
	if operationID != nil {
		cOperationID = C.CString(operationID.(string))
	}
	var cCallMethodName *C.char
	if callMethodName != nil {
		cCallMethodName = C.CString(callMethodName.(string))
	}

	var cErrCode *C.double
	if errCode != nil {
		if methodName == "OnProgress" {
			cErrCode = (*C.double)(C.malloc(C.sizeof_double))
			*cErrCode = C.double(errCode.(int))
		} else {
			cErrCode = (*C.double)(C.malloc(C.sizeof_double))
			*cErrCode = C.double(errCode.(int32))
		}
	}

	var cMessage *C.char
	if message != nil {
		cMessage = C.CString(message.(string))
	}
	C.callOnMethodChannel(openIMListener, main_isolate_send_port, cMethodName, cOperationID, cCallMethodName, cErrCode, cMessage)
	C.callOnNativeMethodChannel(openIMListener, cMethodName, cOperationID, cCallMethodName, cErrCode, cMessage)

}

type OnConnListener struct{}

func (c *OnConnListener) OnConnecting() {
	callBack("OnConnecting", nil, nil, nil, nil)
}

func (c *OnConnListener) OnConnectSuccess() {
	callBack("OnConnectSuccess", nil, nil, nil, nil)
}

func (c *OnConnListener) OnConnectFailed(errCode int32, errMsg string) {
	callBack("OnConnectFailed", nil, nil, errCode, errMsg)
}

func (c *OnConnListener) OnKickedOffline() {
	callBack("OnKickedOffline", nil, nil, nil, nil)
}

func (c *OnConnListener) OnUserTokenExpired() {
	callBack("OnUserTokenExpired", nil, nil, nil, nil)
}

type BaseListener struct {
	operationID string
	methodName  string
}

func (b BaseListener) OnError(errCode int32, errMsg string) {
	callBack("OnError", b.operationID, b.methodName, errCode, errMsg)
}

func (b BaseListener) OnSuccess(data string) {
	callBack("OnSuccess", b.operationID, b.methodName, nil, data)
}

type SendMsgCallBackListener struct {
	operationID string
	methodName  string
	clientMsgID string
}

func (c SendMsgCallBackListener) OnProgress(progress int) {
	callBack("OnProgress", c.operationID, c.methodName, progress, c.clientMsgID)
}

func (c SendMsgCallBackListener) OnError(errCode int32, errMsg string) {
	callBack("OnError", c.operationID, c.methodName, errCode, errMsg)
}

func (c SendMsgCallBackListener) OnSuccess(data string) {
	callBack("OnSuccess", c.operationID, c.methodName, nil, data)
}

type UserListener struct{}

func (o UserListener) OnSelfInfoUpdated(userInfo string) {
	callBack("OnSelfInfoUpdated", nil, nil, nil, userInfo)
}

type AdvancedMsgListener struct{}

func (a AdvancedMsgListener) OnRecvNewMessage(message string) {
	callBack("OnRecvNewMessage", nil, nil, nil, message)
}

func (a AdvancedMsgListener) OnRecvC2CReadReceipt(msgReceiptList string) {
	callBack("OnRecvC2CReadReceipt", nil, nil, nil, msgReceiptList)
}

func (a AdvancedMsgListener) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
	callBack("OnRecvGroupReadReceipt", nil, nil, nil, groupMsgReceiptList)
}

func (a AdvancedMsgListener) OnNewRecvMessageRevoked(messageRevoked string) {
	callBack("OnNewRecvMessageRevoked", nil, nil, nil, messageRevoked)
}

func (a AdvancedMsgListener) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {
	callBack("OnRecvMessageExtensionsChanged", msgID, nil, nil, reactionExtensionList)
}

func (a AdvancedMsgListener) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
	callBack("OnRecvMessageExtensionsDeleted", msgID, nil, nil, reactionExtensionKeyList)
}

func (a AdvancedMsgListener) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {
	callBack("OnRecvMessageExtensionsAdded", msgID, nil, nil, reactionExtensionList)
}

func (a AdvancedMsgListener) OnRecvOfflineNewMessage(messageList string) {
	callBack("OnRecvOfflineNewMessage", nil, nil, nil, messageList)
}

func (a AdvancedMsgListener) OnMsgDeleted(message string) {
	callBack("OnMsgDeleted", nil, nil, nil, message)
}

func (a AdvancedMsgListener) OnRecvMessageRevoked(message string) {
	callBack("OnRecvMessageRevoked", nil, nil, nil, message)
}

type FriendshipListener struct{}

func (f FriendshipListener) OnFriendApplicationAdded(friendApplication string) {
	callBack("OnFriendApplicationAdded", nil, nil, nil, friendApplication)
}

func (f FriendshipListener) OnFriendApplicationDeleted(friendApplication string) {
	callBack("OnFriendApplicationDeleted", nil, nil, nil, friendApplication)
}

func (f FriendshipListener) OnFriendApplicationAccepted(friendApplication string) {
	callBack("OnFriendApplicationAccepted", nil, nil, nil, friendApplication)
}

func (f FriendshipListener) OnFriendApplicationRejected(friendApplication string) {
	callBack("OnFriendApplicationRejected", nil, nil, nil, friendApplication)
}

func (f FriendshipListener) OnFriendAdded(friendInfo string) {
	callBack("OnFriendAdded", nil, nil, nil, friendInfo)
}

func (f FriendshipListener) OnFriendDeleted(friendInfo string) {
	callBack("OnFriendDeleted", nil, nil, nil, friendInfo)
}

func (f FriendshipListener) OnFriendInfoChanged(friendInfo string) {
	callBack("OnFriendInfoChanged", nil, nil, nil, friendInfo)
}

func (f FriendshipListener) OnBlackAdded(blackInfo string) {
	callBack("OnBlackAdded", nil, nil, nil, blackInfo)
}

func (f FriendshipListener) OnBlackDeleted(blackInfo string) {
	callBack("OnBlackDeleted", nil, nil, nil, blackInfo)
}

type GroupListener struct{}

func (gl GroupListener) OnJoinedGroupAdded(groupInfo string) {
	callBack("OnJoinedGroupAdded", nil, nil, nil, groupInfo)
}
func (gl GroupListener) OnJoinedGroupDeleted(groupInfo string) {
	callBack("OnJoinedGroupDeleted", nil, nil, nil, groupInfo)
}
func (gl GroupListener) OnGroupMemberAdded(groupMemberInfo string) {
	callBack("OnGroupMemberAdded", nil, nil, nil, groupMemberInfo)
}

func (gl GroupListener) OnGroupMemberDeleted(groupMemberInfo string) {
	callBack("OnGroupMemberDeleted", nil, nil, nil, groupMemberInfo)
}

func (gl GroupListener) OnGroupApplicationAdded(groupApplication string) {
	callBack("OnGroupApplicationAdded", nil, nil, nil, groupApplication)
}
func (gl GroupListener) OnGroupApplicationDeleted(groupApplication string) {
	callBack("OnGroupApplicationDeleted", nil, nil, nil, groupApplication)
}
func (gl GroupListener) OnGroupInfoChanged(groupInfo string) {
	callBack("OnGroupInfoChanged", nil, nil, nil, groupInfo)
}
func (gl GroupListener) OnGroupDismissed(groupInfo string) {
	callBack("OnGroupDismissed", nil, nil, nil, groupInfo)
}
func (gl GroupListener) OnGroupMemberInfoChanged(groupMemberInfo string) {
	callBack("OnGroupMemberInfoChanged", nil, nil, nil, groupMemberInfo)
}
func (gl GroupListener) OnGroupApplicationAccepted(groupApplication string) {
	callBack("OnGroupApplicationAccepted", nil, nil, nil, groupApplication)
}
func (gl GroupListener) OnGroupApplicationRejected(groupApplication string) {
	callBack("OnGroupApplicationRejected", nil, nil, nil, groupApplication)
}

type BatchMsgListener struct{}

func (bml BatchMsgListener) OnRecvNewMessages(messageList string) {
	callBack("OnRecvNewMessages", nil, nil, nil, messageList)
}

func (bml BatchMsgListener) OnRecvOfflineNewMessages(messageList string) {
	callBack("OnRecvOfflineNewMessages", nil, nil, nil, messageList)
}

type MessageKvInfoListener struct{}

func (mkl MessageKvInfoListener) OnMessageKvInfoChanged(messageChangedList string) {
	callBack("OnMessageKvInfoChanged", nil, nil, nil, messageChangedList)
}

type ConversationListener struct{}

func (c ConversationListener) OnSyncServerStart() {
	callBack("OnSyncServerStart", nil, nil, nil, nil)
}

func (c ConversationListener) OnSyncServerFinish() {
	callBack("OnSyncServerFinish", nil, nil, nil, nil)
}

func (c ConversationListener) OnSyncServerFailed() {
	callBack("OnSyncServerFailed", nil, nil, nil, nil)
}

func (c ConversationListener) OnNewConversation(conversationList string) {
	callBack("OnNewConversation", nil, nil, nil, conversationList)
}

func (c ConversationListener) OnConversationChanged(conversationList string) {
	callBack("OnConversationChanged", nil, nil, nil, conversationList)
}

func (c ConversationListener) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	callBack("OnTotalUnreadMessageCountChanged", nil, nil, totalUnreadCount, nil)
}

//export GetSdkVersion
func GetSdkVersion() *C.char {
	return C.CString(open_im_sdk.GetSdkVersion())
}

//export InitSDK
func InitSDK(operationID *C.char, config *C.char) C.bool {
	if initSDK {
		return true
	}
	initSDK = true
	listener := &OnConnListener{}
	return C.bool(open_im_sdk.InitSDK(listener, C.GoString(operationID), C.GoString(config)))
}

//export Login
func Login(operationID *C.char, userID *C.char, token *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "Login",
	}
	open_im_sdk.Login(callBack, id, C.GoString(userID), C.GoString(token))
}

//export Logout
func Logout(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "Logout",
	}
	open_im_sdk.Logout(callBack, id)
}

//export SetAppBackgroundStatus
func SetAppBackgroundStatus(operationID *C.char, isBackground C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetAppBackgroundStatus",
	}
	open_im_sdk.SetAppBackgroundStatus(callBack, id, bool(isBackground))
}

//export NetworkStatusChanged
func NetworkStatusChanged(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "NetworkStatusChanged",
	}
	open_im_sdk.NetworkStatusChanged(callBack, id)
}

//export GetLoginStatus
func GetLoginStatus(operationID *C.char) C.int {
	id := C.GoString(operationID)
	return C.int(open_im_sdk.GetLoginStatus(id))

}

//export GetLoginUserID
func GetLoginUserID() *C.char {
	return C.CString(open_im_sdk.GetLoginUserID())
}

func main() {}
