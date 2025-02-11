package controller

import (
	"github.com/holdno/gopherCron/app"
	"github.com/holdno/gopherCron/cmd/service/response"
	"github.com/holdno/gopherCron/common"
	"github.com/holdno/gopherCron/utils"

	"github.com/gin-gonic/gin"
)

type CreateTemporaryTaskReq struct {
	ProjectID    int64  `json:"project_id" form:"project_id" binding:"required"`
	TaskID       string `json:"task_id" form:"task_id" binding:"required"`
	ScheduleTime int64  `json:"schedule_time" form:"schedule_time" binding:"required"`
	Command      string `json:"command" form:"command" binding:"required"`
	Timeout      int    `json:"timeout" form:"timeout" binding:"required"`
	Remark       string `json:"remark" form:"remark" binding:"required"`
	Noseize      int    `json:"noseize" form:"noseize"`
}

func CreateTemporaryTask(c *gin.Context) {
	var (
		err error
		req CreateTemporaryTaskReq
	)

	if err = utils.BindArgsWithGin(c, &req); err != nil {
		response.APIError(c, err)
		return
	}

	srv := app.GetApp(c)
	uid := utils.GetUserID(c)
	if err = srv.CreateTemporaryTask(common.TemporaryTask{
		ProjectID:      req.ProjectID,
		TaskID:         req.TaskID,
		UserID:         uid,
		ScheduleTime:   req.ScheduleTime,
		Command:        req.Command,
		ScheduleStatus: common.TEMPORARY_TASK_SCHEDULE_STATUS_WAITING,
		Noseize:        req.Noseize,
		Remark:         req.Remark,
		Timeout:        req.Timeout,
	}); err != nil {
		response.APIError(c, err)
		return
	}
	response.APISuccess(c, nil)
}

type GetTemporaryTaskListReq struct {
	ProjectID int64 `json:"project_id" form:"project_id" binding:"required"`
}

func GetTemporaryTaskList(c *gin.Context) {
	var (
		err error
		req GetTemporaryTaskListReq
	)
	if err = utils.BindArgsWithGin(c, &req); err != nil {
		response.APIError(c, err)
		return
	}

	srv := app.GetApp(c)
	list, err := srv.GetTemporaryTaskListWithUser(req.ProjectID)
	if err != nil {
		response.APIError(c, err)
		return
	}

	response.APISuccess(c, list)
}

type DeleteTemporaryTaskReq struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

func DeleteTemporaryTask(c *gin.Context) {
	var (
		err error
		req DeleteTemporaryTaskReq
	)
	if err = utils.BindArgsWithGin(c, &req); err != nil {
		response.APIError(c, err)
		return
	}

	srv := app.GetApp(c)
	if err := srv.DeleteTemporaryTask(req.ID); err != nil {
		response.APIError(c, err)
		return
	}

	response.APISuccess(c, nil)
}
