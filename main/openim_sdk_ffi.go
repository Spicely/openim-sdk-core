package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "include/dart_api_dl.h"

typedef struct {
    void (*onMethodChannel)(Dart_Port_DL port, const char*, const char*, const char*, int32_t*, const char*);
} CGO_OpenIM_Listener;

static void callOnMethodChannel(CGO_OpenIM_Listener *listener, Dart_Port_DL port, const char* methodName, const char* operationID,const char* callMethodName, int32_t* errCode, const char* message) {
    listener->onMethodChannel(port, methodName, operationID, callMethodName, errCode, message);
}
*/
import "C"
import (
	"open_im_sdk/open_im_sdk"
	"unsafe"
)

var openIMListener *C.CGO_OpenIM_Listener

var main_isolate_send_port C.Dart_Port_DL

//export RegisterCallback
func RegisterCallback(callback *C.CGO_OpenIM_Listener, port C.Dart_Port_DL) {
	openIMListener = callback
	main_isolate_send_port = port

	conversationListener := &ConversationListener{}
	open_im_sdk.SetConversationListener(conversationListener)

	advancedMsgListener := &AdvancedMsgListener{}
	open_im_sdk.SetAdvancedMsgListener(advancedMsgListener)

	messageKvInfoListener := &MessageKvInfoListener{}
	open_im_sdk.SetMessageKvInfoListener(messageKvInfoListener)

	batchMsgListener := &BatchMsgListener{}
	open_im_sdk.SetBatchMsgListener(batchMsgListener)

	friendshipListener := &FriendshipListener{}
	open_im_sdk.SetFriendListener(friendshipListener)

	groupListener := &GroupListener{}
	open_im_sdk.SetGroupListener(groupListener)

	organizationListener := &OrganizationListener{}
	open_im_sdk.SetOrganizationListener(organizationListener)

	userListener := &UserListener{}
	open_im_sdk.SetUserListener(userListener)

	signalingListener := &SignalingListener{}
	open_im_sdk.SetSignalingListener(signalingListener)

	signalingListenerForService := &SignalingListener{}
	open_im_sdk.SetSignalingListenerForService(signalingListenerForService)

	listenerForService := &ListenerForService{}
	open_im_sdk.SetListenerForService(listenerForService)

	workMomentsListener := &WorkMomentsListener{}
	open_im_sdk.SetWorkMomentsListener(workMomentsListener)
}

func callBack(methodName string, operationID interface{}, callMethodName interface{}, errCode interface{}, message interface{}) {
	cMethodName := C.CString(methodName)
	defer C.free(unsafe.Pointer(cMethodName))

	var cOperationID *C.char
	if operationID != nil {
		cOperationID = C.CString(operationID.(string))
		defer C.free(unsafe.Pointer(cOperationID))
	}
	var cCallMethodName *C.char
	if callMethodName != nil {
		cCallMethodName = C.CString(callMethodName.(string))
		defer C.free(unsafe.Pointer(cCallMethodName))
	}

	var cErrCode *C.int32_t
	if errCode != nil {
		cErrCode = (*C.int32_t)(unsafe.Pointer(&errCode))
	}

	var cMessage *C.char
	if message != nil {
		cMessage = C.CString(message.(string))
		defer C.free(unsafe.Pointer(cMessage))
	}

	C.callOnMethodChannel(openIMListener, main_isolate_send_port, cMethodName, cOperationID, cCallMethodName, cErrCode, cMessage)
}

type ListenerForService struct{}

func (ls ListenerForService) OnGroupApplicationAdded(groupApplication string) {
	callBack("OnGroupApplicationAdded", nil, nil, nil, groupApplication)
}
func (ls ListenerForService) OnGroupApplicationAccepted(groupApplication string) {
	callBack("OnGroupApplicationAccepted", nil, nil, nil, groupApplication)
}
func (ls ListenerForService) OnFriendApplicationAdded(friendApplication string) {
	callBack("OnFriendApplicationAdded", nil, nil, nil, friendApplication)
}
func (ls ListenerForService) OnFriendApplicationAccepted(groupApplication string) {
	callBack("OnFriendApplicationAccepted", nil, nil, nil, groupApplication)
}
func (ls ListenerForService) OnRecvNewMessage(message string) {
	callBack("OnRecvNewMessage", nil, nil, nil, message)
}

type OnConnListener struct{}

func (c *OnConnListener) OnConnecting() {
	callBack("OnConnecting", nil, nil, nil, nil)
}

func (c *OnConnListener) OnConnectSuccess() {
	callBack("OnConnectSuccess", nil, nil, nil, nil)
}

func (c *OnConnListener) OnConnectFailed(errCode int32, errMsg string) {
	callBack("OnConnectFailed", nil, errCode, nil, errMsg)
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
	callBack("OnError", b.operationID, nil, errCode, errMsg)
}

func (b BaseListener) OnSuccess(data string) {
	callBack("OnSuccess", b.operationID, b.methodName, nil, data)
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
	callBack("OnRecvMessageExtensionsChanged", nil, nil, nil, reactionExtensionList)
}

func (a AdvancedMsgListener) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
	callBack("OnRecvMessageExtensionsDeleted", nil, nil, nil, reactionExtensionKeyList)
}

func (a AdvancedMsgListener) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {
	callBack("OnRecvMessageExtensionsAdded", nil, nil, nil, reactionExtensionList)
}

func (a AdvancedMsgListener) OnRecvOfflineNewMessages(messageList string) {
	callBack("OnRecvOfflineNewMessages", nil, nil, nil, messageList)
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

type OrganizationListener struct{}

func (ol OrganizationListener) OnOrganizationUpdated() {
	callBack("OnOrganizationUpdated", nil, nil, nil, nil)
}

type WorkMomentsListener struct{}

func (wml WorkMomentsListener) OnRecvNewNotification() {
	callBack("OnRecvNewNotification", nil, nil, nil, nil)
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

type SignalingListener struct {
}

func (s SignalingListener) OnReceiveNewInvitation(receiveNewInvitationCallback string) {
	callBack("OnReceiveNewInvitation", nil, nil, nil, receiveNewInvitationCallback)
}

func (s SignalingListener) OnInviteeAccepted(inviteeAcceptedCallback string) {
	callBack("OnInviteeAccepted", nil, nil, nil, inviteeAcceptedCallback)
}

func (s SignalingListener) OnInviteeAcceptedByOtherDevice(inviteeAcceptedCallback string) {
	callBack("OnInviteeAcceptedByOtherDevice", nil, nil, nil, inviteeAcceptedCallback)
}

func (s SignalingListener) OnInviteeRejected(inviteeRejectedCallback string) {
	callBack("OnInviteeRejected", nil, nil, nil, inviteeRejectedCallback)
}

func (s SignalingListener) OnInviteeRejectedByOtherDevice(inviteeRejectedCallback string) {
	callBack("OnInviteeRejectedByOtherDevice", nil, nil, nil, inviteeRejectedCallback)
}

func (s SignalingListener) OnInvitationCancelled(invitationCancelledCallback string) {
	callBack("OnInvitationCancelled", nil, nil, nil, invitationCancelledCallback)
}

func (s SignalingListener) OnInvitationTimeout(invitationTimeoutCallback string) {
	callBack("OnInvitationTimeout", nil, nil, nil, invitationTimeoutCallback)
}

func (s SignalingListener) OnHangUp(hangUpCallback string) {
	callBack("OnHangUp", nil, nil, nil, hangUpCallback)
}

func (s SignalingListener) OnRoomParticipantConnected(onRoomParticipantConnectedCallback string) {
	callBack("OnRoomParticipantConnected", nil, nil, nil, onRoomParticipantConnectedCallback)
}

func (s SignalingListener) OnRoomParticipantDisconnected(onRoomParticipantDisconnectedCallback string) {
	callBack("OnRoomParticipantDisconnected", nil, nil, nil, onRoomParticipantDisconnectedCallback)
}

//export GetSdkVersion
func GetSdkVersion() *C.char {
	return C.CString(open_im_sdk.SdkVersion())
}

//export InitSDK
func InitSDK(operationID *C.char, config *C.char) C.bool {
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

//export GetUsersInfo
func GetUsersInfo(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetUsersInfo",
	}
	open_im_sdk.GetUsersInfo(callBack, C.GoString(operationID), C.GoString(userIDList))
}

//export GetSelfUserInfo
func GetSelfUserInfo(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetSelfUserInfo",
	}
	open_im_sdk.GetSelfUserInfo(callBack, C.GoString(operationID))
}

//export GetAllConversationList
func GetAllConversationList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetAllConversationList",
	}
	open_im_sdk.GetAllConversationList(callBack, C.GoString(operationID))
}

//export GetConversationListSplit
func GetConversationListSplit(operationID *C.char, offset *C.int32_t, count *C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetConversationListSplit",
	}
	open_im_sdk.GetConversationListSplit(callBack, C.GoString(operationID), int(*offset), int(*count))
}

func main() {}
