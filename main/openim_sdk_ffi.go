package main

/*
#include <stdio.h>
#include "include/dart_api_dl.h"

typedef struct {
    void (*onMethodChannel)(Dart_Port_DL port, char*, char*, char*, double*, char*);
} CGO_OpenIM_Listener;

static void callOnMethodChannel(CGO_OpenIM_Listener *listener, Dart_Port_DL port, char* methodName, char* operationID, char* callMethodName, double* errCode, char* message) {
    listener->onMethodChannel(port, methodName, operationID, callMethodName, errCode, message);
}
*/
import "C"
import (
	"open_im_sdk/open_im_sdk"
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

//export WakeUp
func WakeUp(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "WakeUp",
	}
	open_im_sdk.WakeUp(callBack, id)
}

//export NetworkChanged
func NetworkChanged(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "NetworkChanged",
	}
	open_im_sdk.NetworkChanged(callBack, id)
}

//export UploadImage
func UploadImage(operationID *C.char, filePath *C.char, token *C.char, obj *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "UploadImage",
	}
	open_im_sdk.UploadImage(callBack, id, C.GoString(filePath), C.GoString(token), C.GoString(obj))
}

//export UploadFile
func UploadFile(operationID *C.char, filePath *C.char) {
	id := C.GoString(operationID)
	callBack := &SendMsgCallBackListener{
		operationID: id,
		methodName:  "UploadFile",
	}
	open_im_sdk.UploadFile(callBack, id, C.GoString(filePath))
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

//export GetLoginStatus
func GetLoginStatus() C.int32_t {
	return C.int32_t(open_im_sdk.GetLoginStatus())
}

//export GetLoginUser
func GetLoginUser() *C.char {
	return C.CString(open_im_sdk.GetLoginUser())
}

//export GetUsersInfo
func GetUsersInfo(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetUsersInfo",
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

//export CreateGroup
func CreateGroup(operationID *C.char, groupBaseInfo *C.char, memberList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "CreateGroup",
	}
	open_im_sdk.CreateGroup(callBack, id, C.GoString(groupBaseInfo), C.GoString(memberList))
}

//export JoinGroup
func JoinGroup(operationID *C.char, groupID *C.char, reqMsg *C.char, joinSource C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "JoinGroup",
	}
	open_im_sdk.JoinGroup(callBack, id, C.GoString(groupID), C.GoString(reqMsg), int32(joinSource))
}

//export QuitGroup
func QuitGroup(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "QuitGroup",
	}
	open_im_sdk.QuitGroup(callBack, id, C.GoString(groupID))
}

//export DismissGroup
func DismissGroup(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DismissGroup",
	}
	open_im_sdk.DismissGroup(callBack, id, C.GoString(groupID))
}

//export ChangeGroupMute
func ChangeGroupMute(operationID *C.char, groupID *C.char, isMute C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ChangeGroupMute",
	}
	open_im_sdk.ChangeGroupMute(callBack, id, C.GoString(groupID), bool(isMute))
}

//export ChangeGroupMemberMute
func ChangeGroupMemberMute(operationID *C.char, groupID *C.char, userID *C.char, mutedSeconds C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ChangeGroupMute",
	}
	open_im_sdk.ChangeGroupMemberMute(callBack, id, C.GoString(groupID), C.GoString(userID), int(mutedSeconds))
}

//export SetGroupMemberRoleLevel
func SetGroupMemberRoleLevel(operationID *C.char, groupID *C.char, userID *C.char, roleLevel C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ChangeGroupMute",
	}
	open_im_sdk.SetGroupMemberRoleLevel(callBack, id, C.GoString(groupID), C.GoString(userID), int(roleLevel))
}

//export SetGroupMemberInfo
func SetGroupMemberInfo(operationID *C.char, groupMemberInfo *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGroupMemberInfo",
	}
	open_im_sdk.SetGroupMemberInfo(callBack, id, C.GoString(groupMemberInfo))
}

//export GetJoinedGroupList
func GetJoinedGroupList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetJoinedGroupList",
	}
	open_im_sdk.GetJoinedGroupList(callBack, id)
}

//export GetGroupsInfo
func GetGroupsInfo(operationID *C.char, groupIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetGroupsInfo",
	}
	open_im_sdk.GetGroupsInfo(callBack, id, C.GoString(groupIDList))
}

//export SearchGroups
func SearchGroups(operationID *C.char, searchParam *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SearchGroups",
	}
	open_im_sdk.SearchGroups(callBack, id, C.GoString(searchParam))
}

//export SetGroupInfo
func SetGroupInfo(operationID *C.char, groupID *C.char, groupInfo *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGroupInfo",
	}
	open_im_sdk.SetGroupInfo(callBack, id, C.GoString(groupID), C.GoString(groupInfo))
}

//export SetGroupVerification
func SetGroupVerification(operationID *C.char, groupID *C.char, verification C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGroupVerification",
	}
	open_im_sdk.SetGroupVerification(callBack, id, C.GoString(groupID), int32(verification))
}

//export SetGroupLookMemberInfo
func SetGroupLookMemberInfo(operationID *C.char, groupID *C.char, rule C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGroupLookMemberInfo",
	}
	open_im_sdk.SetGroupLookMemberInfo(callBack, id, C.GoString(groupID), int32(rule))
}

//export SetGroupApplyMemberFriend
func SetGroupApplyMemberFriend(operationID *C.char, groupID *C.char, rule C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGroupApplyMemberFriend",
	}
	open_im_sdk.SetGroupApplyMemberFriend(callBack, id, C.GoString(groupID), int32(rule))
}

//export GetGroupMemberList
func GetGroupMemberList(operationID *C.char, groupID *C.char, filter C.int32_t, offset C.int32_t, count C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetGroupMemberList",
	}
	open_im_sdk.GetGroupMemberList(callBack, id, C.GoString(groupID), int32(filter), int32(offset), int32(count))
}

//export GetGroupMemberOwnerAndAdmin
func GetGroupMemberOwnerAndAdmin(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetGroupMemberOwnerAndAdmin",
	}
	open_im_sdk.GetGroupMemberOwnerAndAdmin(callBack, id, C.GoString(groupID))
}

//export GetGroupMemberListByJoinTimeFilter
func GetGroupMemberListByJoinTimeFilter(operationID *C.char, groupID *C.char, offset C.int32_t, count C.int32_t, joinTimeBegin C.int64_t, joinTimeEnd C.int64_t, filterUserIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetGroupMemberListByJoinTimeFilter",
	}
	open_im_sdk.GetGroupMemberListByJoinTimeFilter(callBack, id, C.GoString(groupID), int32(offset), int32(count), int64(joinTimeBegin), int64(joinTimeEnd), C.GoString(filterUserIDList))
}

//export GetGroupMembersInfo
func GetGroupMembersInfo(operationID *C.char, groupID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetGroupMembersInfo",
	}
	open_im_sdk.GetGroupMembersInfo(callBack, id, C.GoString(groupID), C.GoString(userIDList))
}

//export KickGroupMember
func KickGroupMember(operationID *C.char, groupID *C.char, reason *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "KickGroupMember",
	}
	open_im_sdk.KickGroupMember(callBack, id, C.GoString(groupID), C.GoString(reason), C.GoString(userIDList))
}

//export TransferGroupOwner
func TransferGroupOwner(operationID *C.char, groupID *C.char, newOwnerUserID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "TransferGroupOwner",
	}
	open_im_sdk.TransferGroupOwner(callBack, id, C.GoString(groupID), C.GoString(newOwnerUserID))
}

//export InviteUserToGroup
func InviteUserToGroup(operationID *C.char, groupID *C.char, reason *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "InviteUserToGroup",
	}
	open_im_sdk.InviteUserToGroup(callBack, id, C.GoString(groupID), C.GoString(reason), C.GoString(userIDList))
}

//export GetRecvGroupApplicationList
func GetRecvGroupApplicationList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetRecvGroupApplicationList",
	}
	open_im_sdk.GetRecvGroupApplicationList(callBack, id)
}

//export GetSendGroupApplicationList
func GetSendGroupApplicationList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetSendGroupApplicationList",
	}
	open_im_sdk.GetSendGroupApplicationList(callBack, id)
}

//export AcceptGroupApplication
func AcceptGroupApplication(operationID *C.char, groupID *C.char, fromUserID *C.char, handleMsg *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "AcceptGroupApplication",
	}
	open_im_sdk.AcceptGroupApplication(callBack, id, C.GoString(groupID), C.GoString(fromUserID), C.GoString(handleMsg))
}

//export RefuseGroupApplication
func RefuseGroupApplication(operationID *C.char, groupID *C.char, fromUserID *C.char, handleMsg *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "RefuseGroupApplication",
	}
	open_im_sdk.RefuseGroupApplication(callBack, id, C.GoString(groupID), C.GoString(fromUserID), C.GoString(handleMsg))
}

//export SetGroupMemberNickname
func SetGroupMemberNickname(operationID *C.char, groupID *C.char, userID *C.char, groupMemberNickname *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGroupMemberNickname",
	}
	open_im_sdk.SetGroupMemberNickname(callBack, id, C.GoString(groupID), C.GoString(userID), C.GoString(groupMemberNickname))
}

//export SearchGroupMembers
func SearchGroupMembers(operationID *C.char, searchParam *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SearchGroupMembers",
	}
	open_im_sdk.SearchGroupMembers(callBack, id, C.GoString(searchParam))
}

// //////////////////////////friend/////////////////////////////////////
//
//export GetDesignatedFriendsInfo
func GetDesignatedFriendsInfo(operationID *C.char, userIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetDesignatedFriendsInfo",
	}
	open_im_sdk.GetDesignatedFriendsInfo(callBack, id, C.GoString(userIDList))
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

//export GetRecvFriendApplicationList
func GetRecvFriendApplicationList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetRecvFriendApplicationList",
	}
	open_im_sdk.GetRecvFriendApplicationList(callBack, id)
}

//export GetSendFriendApplicationList
func GetSendFriendApplicationList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetSendFriendApplicationList",
	}
	open_im_sdk.GetSendFriendApplicationList(callBack, id)
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

///////////////////////conversation////////////////////////////////////

//export GetAllConversationList
func GetAllConversationList(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetAllConversationList",
	}
	open_im_sdk.GetAllConversationList(callBack, id)
}

//export GetConversationListSplit
func GetConversationListSplit(operationID *C.char, offset C.int32_t, count C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetConversationListSplit",
	}
	open_im_sdk.GetConversationListSplit(callBack, id, int(offset), int(count))
}

//export GetOneConversation
func GetOneConversation(operationID *C.char, sessionType C.int, sourceID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetOneConversation",
	}
	open_im_sdk.GetOneConversation(callBack, id, int(sessionType), C.GoString(sourceID))
}

//export GetMultipleConversation
func GetMultipleConversation(operationID *C.char, conversationIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetMultipleConversation",
	}
	open_im_sdk.GetMultipleConversation(callBack, id, C.GoString(conversationIDList))
}

//export SetOneConversationPrivateChat
func SetOneConversationPrivateChat(operationID *C.char, conversationID *C.char, isPrivate C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetOneConversationPrivateChat",
	}
	open_im_sdk.SetOneConversationPrivateChat(callBack, id, C.GoString(conversationID), bool(isPrivate))
}

//export SetOneConversationBurnDuration
func SetOneConversationBurnDuration(operationID *C.char, conversationID *C.char, burnDuration C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetOneConversationBurnDuration",
	}
	open_im_sdk.SetOneConversationBurnDuration(callBack, id, C.GoString(conversationID), int32(burnDuration))
}

//export SetOneConversationRecvMessageOpt
func SetOneConversationRecvMessageOpt(operationID *C.char, conversationID *C.char, opt C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetOneConversationRecvMessageOpt",
	}
	open_im_sdk.SetOneConversationRecvMessageOpt(callBack, id, C.GoString(conversationID), int(opt))
}

//export SetConversationRecvMessageOpt
func SetConversationRecvMessageOpt(operationID *C.char, conversationIDList *C.char, opt C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationRecvMessageOpt",
	}
	open_im_sdk.SetConversationRecvMessageOpt(callBack, id, C.GoString(conversationIDList), int(opt))
}

//export SetGlobalRecvMessageOpt
func SetGlobalRecvMessageOpt(operationID *C.char, opt C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGlobalRecvMessageOpt",
	}
	open_im_sdk.SetGlobalRecvMessageOpt(callBack, id, int(opt))
}

//export HideConversation
func HideConversation(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "HideConversation",
	}
	open_im_sdk.HideConversation(callBack, id, C.GoString(conversationID))
}

//export GetConversationRecvMessageOpt
func GetConversationRecvMessageOpt(operationID *C.char, conversationIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetConversationRecvMessageOpt",
	}
	open_im_sdk.GetConversationRecvMessageOpt(callBack, id, C.GoString(conversationIDList))
}

//export DeleteConversation
func DeleteConversation(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteConversation",
	}
	open_im_sdk.DeleteConversation(callBack, id, C.GoString(conversationID))
}

//export DeleteAllConversationFromLocal
func DeleteAllConversationFromLocal(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteAllConversationFromLocal",
	}
	open_im_sdk.DeleteAllConversationFromLocal(callBack, id)
}

//export SetConversationDraft
func SetConversationDraft(operationID *C.char, conversationID *C.char, draftText *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationDraft",
	}
	open_im_sdk.SetConversationDraft(callBack, id, C.GoString(conversationID), C.GoString(draftText))
}

//export ResetConversationGroupAtType
func ResetConversationGroupAtType(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ResetConversationGroupAtType",
	}
	open_im_sdk.ResetConversationGroupAtType(callBack, id, C.GoString(conversationID))
}

//export PinConversation
func PinConversation(operationID *C.char, conversationID *C.char, isPinned C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "PinConversation",
	}
	open_im_sdk.PinConversation(callBack, id, C.GoString(conversationID), bool(isPinned))
}

//export GetTotalUnreadMsgCount
func GetTotalUnreadMsgCount(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetTotalUnreadMsgCount",
	}
	open_im_sdk.GetTotalUnreadMsgCount(callBack, id)
}

//export CreateAdvancedTextMessage
func CreateAdvancedTextMessage(operationID *C.char, text *C.char, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(messageEntityList)))
}

//export CreateTextMessage
func CreateTextMessage(operationID *C.char, text *C.char) *C.char {
	return C.CString(open_im_sdk.CreateTextMessage(C.GoString(operationID), C.GoString(text)))
}

//export CreateTextAtMessage
func CreateTextAtMessage(operationID *C.char, text *C.char, atUserList *C.char, atUsersInfo *C.char, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateTextAtMessage(C.GoString(operationID), C.GoString(text), C.GoString(atUserList), C.GoString(atUsersInfo), C.GoString(message)))
}

//export CreateLocationMessage
func CreateLocationMessage(operationID *C.char, description *C.char, longitude C.double, latitude C.double) *C.char {
	return C.CString(open_im_sdk.CreateLocationMessage(C.GoString(operationID), C.GoString(description), float64(longitude), float64(latitude)))
}

//export CreateCustomMessage
func CreateCustomMessage(operationID *C.char, data *C.char, extension *C.char, description *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCustomMessage(C.GoString(operationID), C.GoString(data), C.GoString(extension), C.GoString(description)))
}

//export CreateQuoteMessage
func CreateQuoteMessage(operationID *C.char, text *C.char, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message)))
}

//export CreateAdvancedQuoteMessage
func CreateAdvancedQuoteMessage(operationID *C.char, text *C.char, message *C.char, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message), C.GoString(messageEntityList)))
}

//export CreateCardMessage
func CreateCardMessage(operationID *C.char, cardInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCardMessage(C.GoString(operationID), C.GoString(cardInfo)))
}

//export CreateVideoMessageFromFullPath
func CreateVideoMessageFromFullPath(operationID *C.char, videoFullPath *C.char, videoType *C.char, duration C.int64_t, snapshotFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageFromFullPath(C.GoString(operationID), C.GoString(videoFullPath), C.GoString(videoType), int64(duration), C.GoString(snapshotFullPath)))
}

//export CreateImageMessageFromFullPath
func CreateImageMessageFromFullPath(operationID *C.char, imageFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageFromFullPath(C.GoString(operationID), C.GoString(imageFullPath)))
}

//export CreateSoundMessageFromFullPath
func CreateSoundMessageFromFullPath(operationID *C.char, soundPath *C.char, duration C.int64_t) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageFromFullPath(C.GoString(operationID), C.GoString(soundPath), int64(duration)))
}

//export CreateFileMessageFromFullPath
func CreateFileMessageFromFullPath(operationID *C.char, fileFullPath *C.char, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageFromFullPath(C.GoString(operationID), C.GoString(fileFullPath), C.GoString(fileName)))
}

//export CreateImageMessage
func CreateImageMessage(operationID *C.char, imagePath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessage(C.GoString(operationID), C.GoString(imagePath)))
}

//export CreateImageMessageByURL
func CreateImageMessageByURL(operationID *C.char, sourcePicture *C.char, bigPicture *C.char, snapshotPicture *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageByURL(C.GoString(operationID), C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture)))
}

//export CreateSoundMessageByURL
func CreateSoundMessageByURL(operationID *C.char, soundBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageByURL(C.GoString(operationID), C.GoString(soundBaseInfo)))
}

//export CreateSoundMessage
func CreateSoundMessage(operationID *C.char, soundPath *C.char, duration C.int64_t) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessage(C.GoString(operationID), C.GoString(soundPath), int64(duration)))
}

//export CreateVideoMessageByURL
func CreateVideoMessageByURL(operationID *C.char, videoBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageByURL(C.GoString(operationID), C.GoString(videoBaseInfo)))
}

//export CreateVideoMessage
func CreateVideoMessage(operationID *C.char, videoPath *C.char, videoType *C.char, duration C.int64_t, snapshotPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessage(C.GoString(operationID), C.GoString(videoPath), C.GoString(videoType), int64(duration), C.GoString(snapshotPath)))
}

//export CreateFileMessageByURL
func CreateFileMessageByURL(operationID *C.char, fileBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageByURL(C.GoString(operationID), C.GoString(fileBaseInfo)))
}

//export CreateFileMessage
func CreateFileMessage(operationID *C.char, filePath *C.char, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessage(C.GoString(operationID), C.GoString(filePath), C.GoString(fileName)))
}

//export CreateMergerMessage
func CreateMergerMessage(operationID *C.char, mergerElemList *C.char, title *C.char, summaryList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateMergerMessage(C.GoString(operationID), C.GoString(mergerElemList), C.GoString(title), C.GoString(summaryList)))
}

//export CreateFaceMessage
func CreateFaceMessage(operationID *C.char, index C.int, data *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFaceMessage(C.GoString(operationID), int(index), C.GoString(data)))
}

//export CreateForwardMessage
func CreateForwardMessage(operationID *C.char, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateForwardMessage(C.GoString(operationID), C.GoString(message)))
}

//export SendMessage
func SendMessage(operationID *C.char, message *C.char, recvID, groupID *C.char, offlinePushInfo *C.char, clientMsgID *C.char) {
	id := C.GoString(operationID)
	callBack := &SendMsgCallBackListener{
		operationID: id,
		methodName:  "SendMessage",
		clientMsgID: C.GoString(clientMsgID),
	}
	open_im_sdk.SendMessage(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}

//export SendMessageNotOss
func SendMessageNotOss(operationID *C.char, message *C.char, recvID, groupID *C.char, offlinePushInfo *C.char, clientMsgID *C.char) {
	id := C.GoString(operationID)
	callBack := &SendMsgCallBackListener{
		operationID: id,
		methodName:  "SendMessageNotOss",
		clientMsgID: C.GoString(clientMsgID),
	}
	open_im_sdk.SendMessageNotOss(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}

// //export SendMessageByBuffer
// func SendMessageByBuffer(operationID *C.char, message *C.char, recvID *C.char, groupID *C.char, offlinePushInfo *C.char, buffer1 *C.char, buffer1Size C.int, buffer2 *C.char, buffer2Size C.int) {
// 	id := C.GoString(operationID)
// 	callBack := &SendMsgCallBackListener{
// 		operationID: id,
// 		methodName:  "SendMessageByBuffer",
// 	}

// 	// Convert buffer1 from C string to Go byte slice
// 	buffer1Go := C.GoBytes(unsafe.Pointer(buffer1), buffer1Size)
// 	// Convert buffer2 from C string to Go byte slice
// 	buffer2Go := C.GoBytes(unsafe.Pointer(buffer2), buffer2Size)

// 	open_im_sdk.SendMessageByBuffer(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo), buffer1Go, buffer2Go)
// }

//export FindMessageList
func FindMessageList(operationID *C.char, findMessageOptions *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "FindMessageList",
	}
	open_im_sdk.FindMessageList(callBack, id, C.GoString(findMessageOptions))
}

//export GetHistoryMessageList
func GetHistoryMessageList(operationID *C.char, getMessageOptions *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetHistoryMessageList",
	}
	open_im_sdk.GetHistoryMessageList(callBack, id, C.GoString(getMessageOptions))
}

//export GetAdvancedHistoryMessageList
func GetAdvancedHistoryMessageList(operationID *C.char, getMessageOptions *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetAdvancedHistoryMessageList",
	}
	open_im_sdk.GetAdvancedHistoryMessageList(callBack, id, C.GoString(getMessageOptions))
}

//export GetAdvancedHistoryMessageListReverse
func GetAdvancedHistoryMessageListReverse(operationID *C.char, getMessageOptions *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetAdvancedHistoryMessageListReverse",
	}
	open_im_sdk.GetAdvancedHistoryMessageListReverse(callBack, id, C.GoString(getMessageOptions))
}

//export GetHistoryMessageListReverse
func GetHistoryMessageListReverse(operationID *C.char, getMessageOptions *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetHistoryMessageListReverse",
	}
	open_im_sdk.GetHistoryMessageListReverse(callBack, id, C.GoString(getMessageOptions))
}

//export RevokeMessage
func RevokeMessage(operationID *C.char, message *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "RevokeMessage",
	}
	open_im_sdk.RevokeMessage(callBack, id, C.GoString(message))
}

//export NewRevokeMessage
func NewRevokeMessage(operationID *C.char, message *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "NewRevokeMessage",
	}
	open_im_sdk.NewRevokeMessage(callBack, id, C.GoString(message))
}

//export TypingStatusUpdate
func TypingStatusUpdate(operationID *C.char, recvID *C.char, msgTip *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "TypingStatusUpdate",
	}
	open_im_sdk.TypingStatusUpdate(callBack, id, C.GoString(recvID), C.GoString(msgTip))
}

//export MarkC2CMessageAsRead
func MarkC2CMessageAsRead(operationID *C.char, userID *C.char, msgIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "MarkC2CMessageAsRead",
	}
	open_im_sdk.MarkC2CMessageAsRead(callBack, id, C.GoString(userID), C.GoString(msgIDList))
}

//export MarkMessageAsReadByConID
func MarkMessageAsReadByConID(operationID *C.char, conversationID *C.char, msgIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "MarkMessageAsReadByConID",
	}
	open_im_sdk.MarkMessageAsReadByConID(callBack, id, C.GoString(conversationID), C.GoString(msgIDList))
}

//export MarkGroupMessageHasRead
func MarkGroupMessageHasRead(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "MarkGroupMessageHasRead",
	}
	open_im_sdk.MarkGroupMessageHasRead(callBack, id, C.GoString(groupID))
}

//export MarkGroupMessageAsRead
func MarkGroupMessageAsRead(operationID *C.char, groupID *C.char, msgIDList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "MarkGroupMessageAsRead",
	}
	open_im_sdk.MarkGroupMessageAsRead(callBack, id, C.GoString(groupID), C.GoString(msgIDList))
}

//export DeleteMessageFromLocalStorage
func DeleteMessageFromLocalStorage(operationID *C.char, message *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteMessageFromLocalStorage",
	}
	open_im_sdk.DeleteMessageFromLocalStorage(callBack, id, C.GoString(message))
}

//export DeleteMessageFromLocalAndSvr
func DeleteMessageFromLocalAndSvr(operationID *C.char, message *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteMessageFromLocalAndSvr",
	}
	open_im_sdk.DeleteMessageFromLocalAndSvr(callBack, id, C.GoString(message))
}

//export DeleteConversationFromLocalAndSvr
func DeleteConversationFromLocalAndSvr(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteConversationFromLocalAndSvr",
	}
	open_im_sdk.DeleteConversationFromLocalAndSvr(callBack, id, C.GoString(conversationID))
}

//export DeleteAllMsgFromLocalAndSvr
func DeleteAllMsgFromLocalAndSvr(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteAllMsgFromLocalAndSvr",
	}
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(callBack, id)
}

//export DeleteAllMsgFromLocal
func DeleteAllMsgFromLocal(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteAllMsgFromLocal",
	}
	open_im_sdk.DeleteAllMsgFromLocal(callBack, id)
}

//export ClearC2CHistoryMessage
func ClearC2CHistoryMessage(operationID *C.char, userID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ClearC2CHistoryMessage",
	}
	open_im_sdk.ClearC2CHistoryMessage(callBack, id, C.GoString(userID))
}

//export ClearC2CHistoryMessageFromLocalAndSvr
func ClearC2CHistoryMessageFromLocalAndSvr(operationID *C.char, userID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ClearC2CHistoryMessageFromLocalAndSvr",
	}
	open_im_sdk.ClearC2CHistoryMessageFromLocalAndSvr(callBack, id, C.GoString(userID))
}

//export ClearGroupHistoryMessage
func ClearGroupHistoryMessage(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ClearGroupHistoryMessage",
	}
	open_im_sdk.ClearGroupHistoryMessage(callBack, id, C.GoString(groupID))
}

//export ClearGroupHistoryMessageFromLocalAndSvr
func ClearGroupHistoryMessageFromLocalAndSvr(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ClearGroupHistoryMessageFromLocalAndSvr",
	}
	open_im_sdk.ClearGroupHistoryMessageFromLocalAndSvr(callBack, id, C.GoString(groupID))
}

//export InsertSingleMessageToLocalStorage
func InsertSingleMessageToLocalStorage(operationID *C.char, message *C.char, recvID *C.char, sendID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "InsertSingleMessageToLocalStorage",
	}
	open_im_sdk.InsertSingleMessageToLocalStorage(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(sendID))
}

//export InsertGroupMessageToLocalStorage
func InsertGroupMessageToLocalStorage(operationID *C.char, message *C.char, groupID *C.char, senderID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "InsertGroupMessageToLocalStorage",
	}
	open_im_sdk.InsertGroupMessageToLocalStorage(callBack, id, C.GoString(message), C.GoString(groupID), C.GoString(senderID))
}

//export SearchLocalMessages
func SearchLocalMessages(operationID *C.char, searchParam *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SearchLocalMessages",
	}
	open_im_sdk.SearchLocalMessages(callBack, id, C.GoString(searchParam))
}

//export GetConversationIDBySessionType
func GetConversationIDBySessionType(sourceID *C.char, sessionType C.int) *C.char {
	return C.CString(open_im_sdk.GetConversationIDBySessionType(C.GoString(sourceID), int(sessionType)))
}

//export GetAtAllTag
func GetAtAllTag() *C.char {
	return C.CString(open_im_sdk.GetAtAllTag())
}

//export SetMessageReactionExtensions
func SetMessageReactionExtensions(operationID *C.char, message *C.char, reactionExtensionList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetMessageReactionExtensions",
	}
	open_im_sdk.SetMessageReactionExtensions(callBack, id, C.GoString(message), C.GoString(reactionExtensionList))
}

//export AddMessageReactionExtensions
func AddMessageReactionExtensions(operationID *C.char, message *C.char, reactionExtensionList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "AddMessageReactionExtensions",
	}
	open_im_sdk.AddMessageReactionExtensions(callBack, id, C.GoString(message), C.GoString(reactionExtensionList))
}

//export DeleteMessageReactionExtensions
func DeleteMessageReactionExtensions(operationID *C.char, message *C.char, reactionExtensionList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteMessageReactionExtensions",
	}
	open_im_sdk.DeleteMessageReactionExtensions(callBack, id, C.GoString(message), C.GoString(reactionExtensionList))
}

//export GetMessageListReactionExtensions
func GetMessageListReactionExtensions(operationID *C.char, messageList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetMessageListReactionExtensions",
	}
	open_im_sdk.GetMessageListReactionExtensions(callBack, id, C.GoString(messageList))
}

//export GetMessageListSomeReactionExtensions
func GetMessageListSomeReactionExtensions(operationID *C.char, messageList *C.char, reactionExtensionList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetMessageListSomeReactionExtensions",
	}
	open_im_sdk.GetMessageListSomeReactionExtensions(callBack, id, C.GoString(messageList), C.GoString(reactionExtensionList))
}

//export SetTypeKeyInfo
func SetTypeKeyInfo(operationID *C.char, message *C.char, typeKey *C.char, ex *C.char, isCanRepeat C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetTypeKeyInfo",
	}
	open_im_sdk.SetTypeKeyInfo(callBack, id, C.GoString(message), C.GoString(typeKey), C.GoString(ex), bool(isCanRepeat))
}

//export GetTypeKeyListInfo
func GetTypeKeyListInfo(operationID *C.char, messageList *C.char, typeKeyList *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetTypeKeyListInfo",
	}
	open_im_sdk.GetTypeKeyListInfo(callBack, id, C.GoString(messageList), C.GoString(typeKeyList))
}

//export GetAllTypeKeyInfo
func GetAllTypeKeyInfo(operationID *C.char, message *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetAllTypeKeyInfo",
	}
	open_im_sdk.GetAllTypeKeyInfo(callBack, C.GoString(message), id)
}

//export SignalingInviteInGroup
func SignalingInviteInGroup(operationID *C.char, signalInviteInGroupReq *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingInviteInGroup",
	}
	open_im_sdk.SignalingInviteInGroup(callBack, id, C.GoString(signalInviteInGroupReq))
}

//export SignalingInvite
func SignalingInvite(operationID *C.char, signalInviteReq *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingInvite",
	}
	open_im_sdk.SignalingInvite(callBack, id, C.GoString(signalInviteReq))
}

//export SignalingAccept
func SignalingAccept(operationID *C.char, signalAcceptReq *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingAccept",
	}
	open_im_sdk.SignalingAccept(callBack, id, C.GoString(signalAcceptReq))
}

//export SignalingReject
func SignalingReject(operationID *C.char, signalRejectReq *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingReject",
	}
	open_im_sdk.SignalingReject(callBack, id, C.GoString(signalRejectReq))
}

//export SignalingCancel
func SignalingCancel(operationID *C.char, signalCancelReq *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingCancel",
	}
	open_im_sdk.SignalingCancel(callBack, id, C.GoString(signalCancelReq))
}

//export SignalingHungUp
func SignalingHungUp(operationID *C.char, signalHungUpReq *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingHungUp",
	}
	open_im_sdk.SignalingHungUp(callBack, id, C.GoString(signalHungUpReq))
}

//export SignalingGetRoomByGroupID
func SignalingGetRoomByGroupID(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingGetRoomByGroupID",
	}
	open_im_sdk.SignalingGetRoomByGroupID(callBack, id, C.GoString(groupID))
}

//export SignalingGetTokenByRoomID
func SignalingGetTokenByRoomID(operationID *C.char, groupID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SignalingGetTokenByRoomID",
	}
	open_im_sdk.SignalingGetTokenByRoomID(callBack, id, C.GoString(groupID))
}

//export GetSubDepartment
func GetSubDepartment(operationID *C.char, departmentID *C.char, offset C.int, size C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetSubDepartment",
	}
	open_im_sdk.GetSubDepartment(callBack, id, C.GoString(departmentID), int(offset), int(size))
}

//export GetDepartmentMember
func GetDepartmentMember(operationID *C.char, departmentID *C.char, offset C.int, size C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetDepartmentMember",
	}
	open_im_sdk.GetDepartmentMember(callBack, id, C.GoString(departmentID), int(offset), int(size))
}

//export GetUserInDepartment
func GetUserInDepartment(operationID *C.char, userID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetUserInDepartment",
	}
	open_im_sdk.GetUserInDepartment(callBack, id, C.GoString(userID))
}

//export GetDepartmentMemberAndSubDepartment
func GetDepartmentMemberAndSubDepartment(operationID *C.char, departmentID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetDepartmentMemberAndSubDepartment",
	}
	open_im_sdk.GetDepartmentMemberAndSubDepartment(callBack, id, C.GoString(departmentID))
}

//export GetParentDepartmentList
func GetParentDepartmentList(operationID *C.char, departmentID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetParentDepartmentList",
	}
	open_im_sdk.GetParentDepartmentList(callBack, id, C.GoString(departmentID))
}

//export GetDepartmentInfo
func GetDepartmentInfo(operationID *C.char, departmentID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetDepartmentInfo",
	}
	open_im_sdk.GetDepartmentInfo(callBack, id, C.GoString(departmentID))
}

//export SearchOrganization
func SearchOrganization(operationID *C.char, input *C.char, offset C.int, count C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SearchOrganization",
	}
	open_im_sdk.SearchOrganization(callBack, id, C.GoString(input), int(offset), int(count))
}

//export GetWorkMomentsUnReadCount
func GetWorkMomentsUnReadCount(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetWorkMomentsUnReadCount",
	}
	open_im_sdk.GetWorkMomentsUnReadCount(callBack, id)
}

//export GetWorkMomentsNotification
func GetWorkMomentsNotification(operationID *C.char, offset C.int, count C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetWorkMomentsNotification",
	}
	open_im_sdk.GetWorkMomentsNotification(callBack, id, int(offset), int(count))
}

//export ClearWorkMomentsNotification
func ClearWorkMomentsNotification(operationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ClearWorkMomentsNotification",
	}
	open_im_sdk.ClearWorkMomentsNotification(callBack, id)
}

//export UpdateFcmToken
func UpdateFcmToken(operationID *C.char, fmcToken *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "UpdateFcmToken",
	}
	open_im_sdk.UpdateFcmToken(callBack, id, C.GoString(fmcToken))
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

func main() {}
