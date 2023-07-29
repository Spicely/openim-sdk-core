package main

/*
#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
*/
import "C"
import "open_im_sdk/open_im_sdk"

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
func GetConversationListSplit(operationID *C.char, offset C.int, count C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetConversationListSplit",
	}
	open_im_sdk.GetConversationListSplit(callBack, id, int(offset), int(count))
}

//export GetOneConversation
func GetOneConversation(operationID *C.char, sessionType C.int32_t, sourceID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "GetOneConversation",
	}
	open_im_sdk.GetOneConversation(callBack, id, int32(sessionType), C.GoString(sourceID))
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

//export SetGlobalRecvMessageOpt
func SetGlobalRecvMessageOpt(operationID *C.char, opt C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetGlobalRecvMessageOpt",
	}
	open_im_sdk.SetGlobalRecvMessageOpt(callBack, id, int(opt))
}

//export SetConversationMsgDestructTime
func SetConversationMsgDestructTime(operationID *C.char, conversationID *C.char, msgDestructTime C.int64_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationMsgDestructTime",
	}
	open_im_sdk.SetConversationMsgDestructTime(callBack, id, C.GoString(conversationID), int64(msgDestructTime))
}

//export SetConversationIsMsgDestruct
func SetConversationIsMsgDestruct(operationID *C.char, conversationID *C.char, isMsgDestruct C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationIsMsgDestruct",
	}
	open_im_sdk.SetConversationIsMsgDestruct(callBack, id, C.GoString(conversationID), bool(isMsgDestruct))
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

//export SetConversationPrivateChat
func SetConversationPrivateChat(operationID *C.char, conversationID *C.char, isPrivate C.bool) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationPrivateChat",
	}
	open_im_sdk.SetConversationPrivateChat(callBack, id, C.GoString(conversationID), bool(isPrivate))
}

//export SetConversationBurnDuration
func SetConversationBurnDuration(operationID *C.char, conversationID *C.char, duration C.int32_t) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationBurnDuration",
	}
	open_im_sdk.SetConversationBurnDuration(callBack, id, C.GoString(conversationID), int32(duration))
}

//export SetConversationRecvMessageOpt
func SetConversationRecvMessageOpt(operationID *C.char, conversationID *C.char, opt C.int) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetConversationRecvMessageOpt",
	}
	open_im_sdk.SetConversationRecvMessageOpt(callBack, id, C.GoString(conversationID), int(opt))
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

//export GetAtAllTag
func GetAtAllTag(operationID *C.char) *C.char {
	id := C.GoString(operationID)
	return C.CString(open_im_sdk.GetAtAllTag(id))
}

//export CreateAdvancedTextMessage
func CreateAdvancedTextMessage(operationID *C.char, text *C.char, messageEntityList *C.char) *C.char {
	id := C.GoString(operationID)

	result := open_im_sdk.CreateAdvancedTextMessage(id, C.GoString(text), C.GoString(messageEntityList))
	return C.CString(result)
}

//export CreateTextAtMessage
func CreateTextAtMessage(operationID *C.char, text *C.char, atUserList *C.char, atUsersInfo *C.char, message *C.char) *C.char {
	id := C.GoString(operationID)

	result := open_im_sdk.CreateTextAtMessage(id, C.GoString(text), C.GoString(atUserList), C.GoString(atUsersInfo), C.GoString(message))
	return C.CString(result)
}

//export CreateTextMessage
func CreateTextMessage(operationID *C.char, text *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateTextMessage(id, C.GoString(text))
	return C.CString(result)
}

//export CreateLocationMessage
func CreateLocationMessage(operationID *C.char, description *C.char, longitude C.double, latitude C.double) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateLocationMessage(id, C.GoString(description), float64(longitude), float64(latitude))
	return C.CString(result)
}

//export CreateCustomMessage
func CreateCustomMessage(operationID *C.char, data *C.char, extension *C.char, description *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateCustomMessage(id, C.GoString(data), C.GoString(extension), C.GoString(description))
	return C.CString(result)
}

//export CreateQuoteMessage
func CreateQuoteMessage(operationID *C.char, text *C.char, message *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateQuoteMessage(id, C.GoString(text), C.GoString(message))
	return C.CString(result)
}

//export CreateAdvancedQuoteMessage
func CreateAdvancedQuoteMessage(operationID *C.char, text *C.char, message *C.char, messageEntityList *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateAdvancedQuoteMessage(id, C.GoString(text), C.GoString(message), C.GoString(messageEntityList))
	return C.CString(result)
}

//export CreateCardMessage
func CreateCardMessage(operationID *C.char, cardInfo *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateCardMessage(id, C.GoString(cardInfo))
	return C.CString(result)
}

//export CreateVideoMessageFromFullPath
func CreateVideoMessageFromFullPath(operationID *C.char, videoFullPath *C.char, videoType *C.char, duration C.int64_t, snapshotFullPath *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateVideoMessageFromFullPath(id, C.GoString(videoFullPath), C.GoString(videoType), int64(duration), C.GoString(snapshotFullPath))
	return C.CString(result)
}

//export CreateImageMessageFromFullPath
func CreateImageMessageFromFullPath(operationID *C.char, imageFullPath *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateImageMessageFromFullPath(id, C.GoString(imageFullPath))
	return C.CString(result)
}

//export CreateSoundMessageFromFullPath
func CreateSoundMessageFromFullPath(operationID *C.char, soundFullPath *C.char, duration C.int64_t) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateSoundMessageFromFullPath(id, C.GoString(soundFullPath), int64(duration))
	return C.CString(result)
}

//export CreateFileMessageFromFullPath
func CreateFileMessageFromFullPath(operationID *C.char, fileFullPath *C.char, fileName *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateFileMessageFromFullPath(id, C.GoString(fileFullPath), C.GoString(fileName))
	return C.CString(result)
}

//export CreateImageMessage
func CreateImageMessage(operationID *C.char, imagePath *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateImageMessage(id, C.GoString(imagePath))
	return C.CString(result)
}

//export CreateImageMessageByURL
func CreateImageMessageByURL(operationID *C.char, sourcePicture *C.char, bigPicture *C.char, snapshotPicture *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateImageMessageByURL(id, C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture))
	return C.CString(result)
}

//export CreateSoundMessageByURL
func CreateSoundMessageByURL(operationID *C.char, soundBaseInfo *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateSoundMessageByURL(id, C.GoString(soundBaseInfo))
	return C.CString(result)
}

//export CreateSoundMessage
func CreateSoundMessage(operationID *C.char, soundPath *C.char, duration C.int64_t) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateSoundMessage(id, C.GoString(soundPath), int64(duration))
	return C.CString(result)
}

//export CreateVideoMessageByURL
func CreateVideoMessageByURL(operationID *C.char, videoBaseInfo *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateVideoMessageByURL(id, C.GoString(videoBaseInfo))
	return C.CString(result)
}

//export CreateVideoMessage
func CreateVideoMessage(operationID *C.char, videoPath *C.char, videoType *C.char, duration C.int64_t, snapshotPath *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateVideoMessage(id, C.GoString(videoPath), C.GoString(videoType), int64(duration), C.GoString(snapshotPath))
	return C.CString(result)
}

//export CreateFileMessageByURL
func CreateFileMessageByURL(operationID *C.char, fileBaseInfo *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateFileMessageByURL(id, C.GoString(fileBaseInfo))
	return C.CString(result)
}

//export CreateFileMessage
func CreateFileMessage(operationID *C.char, filePath *C.char, fileName *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateFileMessage(id, C.GoString(filePath), C.GoString(fileName))
	return C.CString(result)
}

//export CreateMergerMessage
func CreateMergerMessage(operationID *C.char, messageList *C.char, title *C.char, summaryList *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateMergerMessage(id, C.GoString(messageList), C.GoString(title), C.GoString(summaryList))
	return C.CString(result)
}

//export CreateFaceMessage
func CreateFaceMessage(operationID *C.char, index C.int, data *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateFaceMessage(id, int(index), C.GoString(data))
	return C.CString(result)
}

//export CreateForwardMessage
func CreateForwardMessage(operationID *C.char, m *C.char) *C.char {
	id := C.GoString(operationID)
	result := open_im_sdk.CreateForwardMessage(id, C.GoString(m))
	return C.CString(result)
}

//export GetConversationIDBySessionType
func GetConversationIDBySessionType(operationID *C.char, sourceID *C.char, sessionType C.int) string {
	return open_im_sdk.GetConversationIDBySessionType(C.GoString(operationID), C.GoString(sourceID), int(sessionType))
}

//export SendMessage
func SendMessage(operationID *C.char, message *C.char, recvID *C.char, groupID *C.char, offlinePushInfo *C.char) {
	id := C.GoString(operationID)
	callBack := &SendMsgCallBackListener{
		operationID: id,
		methodName:  "SendMessage",
	}
	open_im_sdk.SendMessage(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}

//export SendMessageNotOss
func SendMessageNotOss(operationID *C.char, message *C.char, recvID *C.char, groupID *C.char, offlinePushInfo *C.char) {
	id := C.GoString(operationID)
	callBack := &SendMsgCallBackListener{
		operationID: id,
		methodName:  "SendMessageNotOss",
	}
	open_im_sdk.SendMessageNotOss(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}

// //export SendMessageByBuffer
// func SendMessageByBuffer(operationID *C.char, message *C.char, recvID *C.char, groupID *C.char, offlinePushInfo *C.char, buffer1 *C.char, buffer2 *C.char) {
// 	id := C.GoString(operationID)
// 	callBack := &SendMsgCallBackListener{
// 		operationID: id,
// 		methodName:  "SendMessageByBuffer",
// 	}
// 	open_im_sdk.SendMessageByBuffer(callBack, id, C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo), C.GoString(buffer1), C.GoString(buffer2))
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

//export RevokeMessage
func RevokeMessage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "RevokeMessage",
	}
	open_im_sdk.RevokeMessage(callBack, id, C.GoString(conversationID), C.GoString(clientMsgID))
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

//export MarkConversationMessageAsRead
func MarkConversationMessageAsRead(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "MarkConversationMessageAsRead",
	}
	open_im_sdk.MarkConversationMessageAsRead(callBack, id, C.GoString(conversationID))
}

//export MarkMessagesAsReadByMsgID
func MarkMessagesAsReadByMsgID(operationID *C.char, conversationID *C.char, clientMsgIDs *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "MarkMessagesAsReadByMsgID",
	}
	open_im_sdk.MarkMessagesAsReadByMsgID(callBack, id, C.GoString(conversationID), C.GoString(clientMsgIDs))
}

//export DeleteMessageFromLocalStorage
func DeleteMessageFromLocalStorage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteMessageFromLocalStorage",
	}
	open_im_sdk.DeleteMessageFromLocalStorage(callBack, id, C.GoString(conversationID), C.GoString(clientMsgID))
}

//export DeleteMessage
func DeleteMessage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteMessage",
	}
	open_im_sdk.DeleteMessage(callBack, id, C.GoString(conversationID), C.GoString(clientMsgID))
}

//export DeleteConversationFromLocal
func DeleteConversationFromLocal(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteConversationFromLocal",
	}
	open_im_sdk.DeleteConversationFromLocal(callBack, id, C.GoString(conversationID))
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

//export ClearConversationAndDeleteAllMsg
func ClearConversationAndDeleteAllMsg(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "ClearConversationAndDeleteAllMsg",
	}
	open_im_sdk.ClearConversationAndDeleteAllMsg(callBack, id, C.GoString(conversationID))
}

//export DeleteConversationAndDeleteAllMsg
func DeleteConversationAndDeleteAllMsg(operationID *C.char, conversationID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "DeleteConversationAndDeleteAllMsg",
	}
	open_im_sdk.DeleteConversationAndDeleteAllMsg(callBack, id, C.GoString(conversationID))
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
func InsertGroupMessageToLocalStorage(operationID *C.char, message *C.char, groupID *C.char, sendID *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "InsertGroupMessageToLocalStorage",
	}
	open_im_sdk.InsertGroupMessageToLocalStorage(callBack, id, C.GoString(message), C.GoString(groupID), C.GoString(sendID))
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

//export SetMessageLocalEx
func SetMessageLocalEx(operationID *C.char, conversationID *C.char, clientMsgID *C.char, localEx *C.char) {
	id := C.GoString(operationID)
	callBack := &BaseListener{
		operationID: id,
		methodName:  "SetMessageLocalEx",
	}
	open_im_sdk.SetMessageLocalEx(callBack, id, C.GoString(conversationID), C.GoString(clientMsgID), C.GoString(localEx))
}
