package service

type RelationSvc struct {
	MyId 			int64
}
// GetUserFocus 根据userId得到用户的关注列表
func (rs *RelationSvc) GetUserFocus(userId int64) []dao.Follow {
	follow, _ := dao.GetUserFocus(userId)
	return follow
}
// GetUserFans 根据userId得到用户的粉丝列表
func (rs *RelationSvc) GetUserFans(userId int64) []dao.Follow {
	follower, _ := dao.GetUserFans(userId)
	return follower
}

// LenUserFocus 根据userId得到用户的关注数量
func (rs *RelationSvc) LenUserFocus(userId int64) int64 {
	follow, _ := dao.GetUserFocus(userId)
	return int64(len(follow))
}

// LenUserFans 根据userId得到用户的粉丝数量
func (rs *RelationSvc) LenUserFans(userId int64) int64 {
	follower, _ := dao.GetUserFans(userId)
	return int64(len(follower))
}

// RelationAction 关注或取消关注操作，传入当前用户，目标用户，操作类型
func (rs *RelationSvc) RelationAction(userId, toUserId, actionType int32) {
	if actionType == 1 {
		_ = dao.InsertFocus(&dao.Follow{UserId: userId, FocusId: toUserId})
	} else {
		dao.DeleteFocus(&dao.Follow{UserId: userId, FocusId: toUserId})
	}
}

// IsFollow 判断myId有没有关注toId
func (rs *RelationSvc) IsFollow(toId int64) bool {
	return dao.IsFollow(rs.MyId, toId)
}