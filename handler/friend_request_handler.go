package handler

import (
	db "Butterfly/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendFriendReq struct {
	ToUserId    int32  `json:"to_user_id" binding:"required"`
	RequestDesc string `json:"request_desc"`
}

// 发送好友申请
func (handler *Handler) sendFriendRequest(ctx *gin.Context) {
	var req sendFriendReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// handler.ParamsError(ctx)
		return
	}

	userId := ctx.GetInt("")

	// 不能添加自己
	if req.ToUserId == int32(userId) {
		// handler.ParamsError(ctx, "不能添加自己")
		return
	}

	// 检查目标用户是否存在
	targetUser, _ := handler.store.GetUserById(ctx, req.ToUserId)
	if targetUser.ID < 1 {
		// handler.ParamsError(ctx, "用户不存在")
		return
	}

	// 检查是否已是好友
	if exists, _ := handler.store.ExistsFriendship(ctx, &db.ExistsFriendshipParams{
		UserID:   int32(userId),
		FriendID: req.ToUserId,
	}); exists {
		ctx.JSON(http.StatusOK, "已经是好友")
		return
	}

	// 检查申请是否存在(当前用户发送给目标用户的)
	exists, _ := handler.store.ExistsFriendRequest(ctx, &db.ExistsFriendRequestParams{
		FromUserID: int32(userId),
		ToUserID:   req.ToUserId,
	})
	if exists {
		// handler.Success(ctx, "已发送申请, 等待对方处理")
		return
	}

	// 检查是否有来自目标用户的待处理申请(反向申请)
	// incomingReqExists, _ := handler.Store.GetFriendRequest(ctx, &db.GetFriendRequestParams{
	// 	FromUserID: req.ToUserId,
	// 	ToUserID:   currentUserId,
	// })

	// 如果存在反向申请, 则自动接受
	// if incomingReqExists {
	// 	if err := handler.store.FriendRequestTx(ctx, &db.FriendRequestTxParams{
	// 		Status:     config.Accepted, // 同意申请
	// 		FromUserId: req.ToUserId,    // 申请发起者是目标用户
	// 		ToUserId:   currentUserId,   // 当前用户是接收者
	// 		FromNote:   handler.CurrentUserInfo.Nickname,
	// 		ToNote:     targetUser.Nickname,
	// 	}); err != nil {
	// 		// logs.Errorf("自动接受好友申请失败: %v", err)
	// 		// handler.ServerError(ctx)
	// 		return
	// 	}
	// 	handler.Success(ctx, "success")
	// 	return
	// }

	// 添加申请记录
	if err := handler.store.CreateFriendRequest(ctx, &db.CreateFriendRequestParams{
		FromUserID:  int32(userId),
		ToUserID:    req.ToUserId,
		RequestDesc: req.RequestDesc,
	}); err != nil {
		// logs.Errorf("添加申请记录失败: %v", err)
		// handler.ServerError(ctx)
		return
	}

	// handler.Success(ctx, "申请成功")
}

// 获取当前用户收到的待处理好友申请
func (handler *Handler) listFriendRequest(ctx *gin.Context) {

	userId := ctx.GetInt("")

	handler.store.ListFriendRequestByPending(ctx, int32(userId))
}

type processFriendRequest struct {
	RequestId int32  `json:"request_id" binding:"required"`
	ToUserId  int32  `json:"to_user_id" binding:"required"`
	Action    string `json:"action" binding:"required,oneof=accept reject"`
	Remark    string `json:"remark"`
}

// 处理好友申请
func (handler *Handler) processFriendRequest(ctx *gin.Context) {
	var req processFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// handler.ParamsError(ctx)
		return
	}

	userId := ctx.GetInt("")

	exists, _ := handler.store.ExistsFriendRequest(ctx, &db.ExistsFriendRequestParams{
		FromUserID: req.ToUserId,
		ToUserID:   int32(userId),
	})
	if !exists {
		// logs.Errorf("get friend request error: %v", err.Error())
		// handler.Error(ctx, http.StatusNotFound, "不存在的申请记录")
		return
	}

	// 校验当前用户是否是接收者
	// if fr.ToUserID != handler.CurrentUserInfo.ID {
	// 	logs.Errorf("无权处理此请求, senderId: %d, receiveId: %d", fr.ToUserID, handler.CurrentUserInfo.ID)
	// 	handler.ParamsError(ctx)
	// 	return
	// }

	switch req.Action {
	case "accept": // 接受好友申请
		// err := handler.acceptedUserProcess(ctx, req)
		// if err != nil {
		// 	logs.Errorf("accepted user error: %v\n", err)
		// 	handler.ServerError(ctx)
		// 	return
		// }
	case "reject": // 拒绝好友申请
		// err = handler.Store.UpdateFriendRequest(ctx, &db.UpdateFriendRequestParams{
		// 	FromUserID: req.FromUserId,
		// 	ToUserID:   handler.CurrentUserInfo.ID,
		// 	Status:     config.Rejected,
		// })
		// if err != nil {
		// 	logs.Errorf("rejected user error: %v\n", err)
		// 	handler.ServerError(ctx)
		// 	return
		// }
	default:
		// handler.ParamsError(ctx)
		return
	}

}
