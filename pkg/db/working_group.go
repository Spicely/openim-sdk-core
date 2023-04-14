//go:build !js
// +build !js

package db

import (
	"context"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/utils"
)

func (d *DataBase) GetJoinedWorkingGroupIDList(ctx context.Context) ([]string, error) {
	groupList, err := d.GetJoinedGroupListDB(ctx)
	if err != nil {
		return nil, utils.Wrap(err, "")
	}
	var groupIDList []string
	for _, v := range groupList {
		if v.GroupType == constant.WorkingGroup {
			groupIDList = append(groupIDList, v.GroupID)
		}
	}
	return groupIDList, nil
}

func (d *DataBase) GetJoinedWorkingGroupList(ctx context.Context) ([]*model_struct.LocalGroup, error) {
	groupList, err := d.GetJoinedGroupListDB(ctx)
	var transfer []*model_struct.LocalGroup
	for _, v := range groupList {
		if v.GroupType == constant.WorkingGroup {
			transfer = append(transfer, v)
		}
	}
	return transfer, utils.Wrap(err, "GetJoinedSuperGroupList failed ")
}
