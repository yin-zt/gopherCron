package common

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gorhill/cronexpr"
)

const (
	REMOTE_EVENT_PUT                  = "remote_event_put"
	REMOTE_EVENT_DELETE               = "remote_event_delete"
	REMOTE_EVENT_UPDATE               = "remote_event_update"
	REMOTE_EVENT_TMP_SCHEDULE         = "remote_event_tmp_schedule"
	REMOTE_EVENT_WORKFLOW_SCHEDULE    = "remote_event_workflow_schedule"
	REMOTE_EVENT_TASK_STOP            = "remote_event_task_stop"
	REMOTE_EVENT_CHECK_TASK_ISRUNNING = "remote_event_check_task_isrunning"

	CENTER_EVENT_WORKFLOW_REFRESH = "workflow_refresh"
)

// TaskInfo 任务详情
type TaskInfo struct {
	TaskID    string `json:"task_id"`
	Name      string `json:"name"`
	ProjectID int64  `json:"project_id"`

	Command    string        `json:"command"`
	Cron       string        `json:"cron"`
	Remark     string        `json:"remark"`
	Timeout    int           `json:"timeout"` // 任务超时时间 单位 秒(s)
	CreateTime int64         `json:"create_time"`
	Status     int           `json:"status"`
	IsRunning  int           `json:"is_running"`
	Noseize    int           `json:"noseize"`
	Exclusion  int           `json:"exclusion"` // 互斥规则
	ClientIP   string        `json:"client_ip"`
	TmpID      string        `json:"tmp_id"` // 每次任务执行的唯一标识
	FlowInfo   *WorkflowInfo `json:"flow_info,omitempty"`
}

type TaskListItemWithWorkflows struct {
	TaskID    string `json:"task_id"`
	Name      string `json:"name"`
	ProjectID int64  `json:"project_id"`

	Command    string  `json:"command"`
	Cron       string  `json:"cron"`
	Remark     string  `json:"remark"`
	Timeout    int     `json:"timeout"` // 任务超时时间 单位 秒(s)
	CreateTime int64   `json:"create_time"`
	Status     int     `json:"status"`
	IsRunning  int     `json:"is_running"`
	Noseize    int     `json:"noseize"`
	Exclusion  int     `json:"exclusion"` // 互斥规则
	Workflows  []int64 `json:"workflows,omitempty"`
}

type WorkflowInfo struct {
	WorkflowID int64  `json:"workflow_id"`
	TmpID      string `json:"tmp_id"`
}

type TaskRunningInfo struct {
	Status    string `json:"status"`
	TmpID     string `json:"tmp_id"`
	Timestamp int64  `json:"timestamp,omitempty"`
	AgentIP   string `json:"agent_ip"`
}

// TaskSchedulePlan 任务调度计划
type TaskSchedulePlan struct {
	Task     *TaskInfo
	Expr     *cronexpr.Expression // 解析后的cron表达式
	TmpID    string
	NextTime time.Time
	Type     PlanType
}

type PlanType string

const (
	NormalPlan   PlanType = "normal"
	WorkflowPlan PlanType = "workflow"
)

// TaskExecutingInfo 任务执行状态
type TaskExecutingInfo struct {
	Task     *TaskInfo `json:"task"`
	PlanTime time.Time `json:"plan_time"` // 理论上的调度时间
	PlanType PlanType  `json:"plan_type"`
	RealTime time.Time `json:"real_time"` // 实际调度时间
	TmpID    string    `json:"tmp_id"`

	CancelCtx  context.Context    `json:"-"`
	CancelFunc context.CancelFunc `json:"-"` // 用来取消Command执行的cancel函数
}

// TaskExecuteResult 任务执行结果
type TaskExecuteResult struct {
	ExecuteInfo *TaskExecutingInfo `json:"execute_info"`
	Output      string             `json:"output"`     // 程序输出
	Err         string             `json:"error"`      // 是否发生错误
	StartTime   time.Time          `json:"start_time"` // 开始时间
	EndTime     time.Time          `json:"end_time"`   // 结束时间
}

// TaskResultLog 任务执行结果日志
type TaskResultLog struct {
	Result      string `json:"result"`
	SystemError string `json:"system_error"`
	Error       string `json:"error"`
}

// ETCD_PREFIX topic prefix  default: /cron
var (
	ETCD_PREFIX             = "/cron"
	TEMPORARY               = "t_scheduler"
	WORKFLOW                = "t_flow"
	WORKFLOW_ACK            = "t_flow_ack"
	WORKFLOW_MASTER         = "t_flow_master"
	WEBHOOK_MASTER          = "t_webhook_master"
	CLEANUP_MASTER          = "t_cleanup_master"
	CALC_CONSISTENCY_MASTER = "t_calc_consistency_master"
	TEMPORARY_MASTER        = "t_temporary_master"
	STATUS                  = "t_status"
)

// BuildTaskUpdateKey 任务更新锁的key
func BuildTaskUpdateKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%s/update/%d/%s", ETCD_PREFIX, projectID, taskID)
}

// BuildKey etcd 保存任务的key
func BuildKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%s/%d/%s", ETCD_PREFIX, projectID, taskID)
}

func BuildTaskPrefixKey(projectID int64) string {
	return fmt.Sprintf("%s/%d/", ETCD_PREFIX, projectID)
}

func BuildWorkflowAddUserLockKey(workflowID, userID int64) string {
	return fmt.Sprintf("%s/lock/workflow_user_%d_%d", ETCD_PREFIX, workflowID, userID)
}

func BuildCleanupMasterKey() string {
	return fmt.Sprintf("%s/%s", ETCD_PREFIX, CLEANUP_MASTER)
}

func BuildWebhookMasterKey() string {
	return fmt.Sprintf("%s/%s", ETCD_PREFIX, WEBHOOK_MASTER)
}

func BuildWorkflowMasterKey() string {
	return fmt.Sprintf("%s/%s", ETCD_PREFIX, WORKFLOW_MASTER)
}

func BuildCalaConsistencyMasterKey() string {
	return fmt.Sprintf("%s/%s", ETCD_PREFIX, CALC_CONSISTENCY_MASTER)
}

func BuildTemporaryMasterKey() string {
	return fmt.Sprintf("%s/%s", ETCD_PREFIX, TEMPORARY_MASTER)
}

func GetTaskStatusPrefixKey() string {
	return fmt.Sprintf("%s/%s/", ETCD_PREFIX, STATUS)
}

// func ParseTaskStatusKey(key string) (int64, string) {
// 	projectTask := strings.Replace(key, GetTaskStatusPrefixKey(), "", -1)
// 	result := strings.Split(projectTask, "/")
// 	if len(result) != 2 {
// 		return 0, ""
// 	}
// 	projectID, _ := strconv.ParseInt(result[0], 10, 64)
// 	return projectID, result[1]
// }

func BuildTaskStatusKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%s/%d/%s/%s", ETCD_PREFIX, projectID, taskID, STATUS)
}

// func BuildTaskRunningKeyPrefix(projectID int64, taskID string) string {
// 	return fmt.Sprintf("%s/status/running/%d/%s", ETCD_PREFIX, projectID, taskID)
// }

// func BuildTaskRunningKey(agentIP string, projectID int64, taskID string) string {
// 	return fmt.Sprintf("%s/%s", BuildTaskRunningKeyPrefix(projectID, taskID), agentIP)
// }

func BuildWorkflowTaskStatusKey(workflowID, projectID int64, taskID string) string {
	return fmt.Sprintf("%s%d/%s", BuildWorkflowTaskStatusKeyPrefix(workflowID), projectID, taskID)
}

func BuildWorkflowTaskStatusKeyPrefix(workflowID int64) string {
	return fmt.Sprintf("%s/%s/%d/", ETCD_PREFIX, WORKFLOW, workflowID)
}

// BuildSchedulerKey 临时调度的key
func BuildSchedulerKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%s/%d/%s/%s", ETCD_PREFIX, projectID, TEMPORARY, taskID)
}

// BuildWorkflowSchedulerKey workflow调度的key
func BuildWorkflowSchedulerKey(workflowID, projectID int64, taskID string) string {
	return fmt.Sprintf("%s/%d/%s/%d/%s", ETCD_PREFIX, projectID, WORKFLOW, workflowID, taskID)
}

// BuildWorkflowSchedulerKey agent返回接收调度确认的key
func BuildWorkflowAckKey(workflowID, projectID int64, taskID, tmpID string) string {
	return fmt.Sprintf("%s/%d/%s/%d/%s/%s", ETCD_PREFIX, projectID, WORKFLOW_ACK, workflowID, taskID, tmpID)
}

func BuildTaskResultQueuePrefixKey() string {
	return fmt.Sprintf("%s/%s/%s/", ETCD_PREFIX, "queue", "task_result")
}

func BuildTaskResultQueueProjectKey(projectID int64) string {
	return fmt.Sprintf("%s%d/", BuildTaskResultQueuePrefixKey(), projectID)
}

// IsTemporaryKey 检测是否为临时调度key
func IsTemporaryKey(key string) bool {
	return strings.Contains(key, "/"+TEMPORARY+"/")
}

func IsWorkflowKey(key string) bool {
	return strings.Contains(key, "/"+WORKFLOW+"/")
}

func IsStatusKey(key string) bool {
	return strings.Contains(key, "/"+STATUS)
}

func IsAckKey(key string) bool {
	return strings.Contains(key, "/"+WORKFLOW_ACK)
}

func PatchProjectIDTaskIDFromStatusKey(key string) (string, string) {
	sp := strings.Split(key, "/")
	if len(sp) != 5 {
		return "", ""
	}
	return sp[2], sp[3]
}

// BuildLockKey etcd 分布式锁key
func BuildLockKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%s/lock/%d/%s", ETCD_PREFIX, projectID, taskID)
}

// BuildLockKey etcd 分布式锁key
func BuildKillKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%s/kill/%d/%s", ETCD_PREFIX, projectID, taskID)
}

// BuildRegisterKey etcd 服务发现key
func BuildRegisterKey(projectID int64, ip string) string {
	return fmt.Sprintf("%s/register/%d/%s", ETCD_PREFIX, projectID, ip)
}

func BuildAgentCommandKey(host, command string) string {
	return BuildAgentRegisteKey(host) + command
}

// BuildAgentRegisteKey agent 注册
func BuildAgentRegisteKey(ip string) string {
	return fmt.Sprintf("%s/agent/%s/", ETCD_PREFIX, ip)
}

// BuildMonitorKey 构建监控信息存储的key
func BuildMonitorKey(ip string) string {
	return ETCD_PREFIX + "/monitor/" + ip
}

// BuildWorkflowPlanKey 构建
func BuildWorkflowPlanKey(workflowID int64) string {
	return fmt.Sprintf("%s/workflow_plan/%d", ETCD_PREFIX, workflowID)
}

// BuildTableKey 构建scheduler 关系表中的key
func (t *TaskInfo) SchedulerKey() string {
	return GenTaskSchedulerKey(t.ProjectID, t.TaskID)
}

func GenTaskSchedulerKey(projectID int64, taskID string) string {
	return fmt.Sprintf("%d_%s", projectID, taskID)
}

func Unmarshal(value []byte) (*TaskInfo, error) {
	task := new(TaskInfo)
	err := json.Unmarshal(value, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// 从etcd的key中提取任务名称
func ExtractTaskID(project int64, key string) string {
	return strings.TrimPrefix(key, BuildKey(project, ""))
}

// 从etcd的key中提取节点ip
func ExtractWorkerIP(project int64, key string) string {
	return strings.TrimPrefix(key, BuildRegisterKey(project, ""))
}

// 从etcd的key中提取任务名称
func ExtractKillID(project int64, key string) string {
	return strings.TrimPrefix(key, BuildKillKey(project, ""))
}

func ExtractAgentCommand(key string) string {
	keys := strings.Split(key, "/")
	return keys[len(keys)-1]
}

type TaskEvent struct {
	EventType int // save delete
	Task      *TaskInfo
}

func BuildTaskEvent(eventType int, task *TaskInfo) *TaskEvent {
	return &TaskEvent{
		EventType: eventType,
		Task:      task,
	}
}

// 构造执行计划
func BuildTaskSchedulerPlan(task *TaskInfo) (*TaskSchedulePlan, error) {
	var (
		expr *cronexpr.Expression
		err  error
	)

	if expr, err = cronexpr.Parse(task.Cron); err != nil {
		return nil, err
	}

	return &TaskSchedulePlan{
		Task:     task,
		Expr:     expr,
		NextTime: expr.Next(time.Now()),
		Type:     NormalPlan,
		TmpID:    task.TmpID,
	}, nil
}

// 构造执行计划
func BuildWorkflowTaskSchedulerPlan(task *TaskInfo) (*TaskSchedulePlan, error) {
	return &TaskSchedulePlan{
		Task:  task,
		Type:  WorkflowPlan,
		TmpID: task.TmpID,
	}, nil
}

// BuildTaskExecuteInfo 构建 executer
func BuildTaskExecuteInfo(plan *TaskSchedulePlan) *TaskExecutingInfo {
	if plan.TmpID == "" {
		plan.TmpID = plan.Task.TmpID
	}
	info := &TaskExecutingInfo{
		Task:     plan.Task,
		PlanTime: plan.NextTime, // 计划调度时间
		PlanType: plan.Type,
		RealTime: time.Now(), // 真实执行时间
		TmpID:    plan.TmpID,
	}

	if plan.Task.Timeout != 0 {
		info.CancelCtx, info.CancelFunc = context.WithTimeout(context.Background(), time.Duration(plan.Task.Timeout)*time.Second)
	} else {
		info.CancelCtx, info.CancelFunc = context.WithCancel(context.Background())
	}

	return info
}

type AckResponse struct {
	Version string          `json:"version"`
	Data    json.RawMessage `json:"data"`
}

type AckResponseV1 struct {
	Type     string `json:"ack"`
	ClientIP string `json:"client_ip"`
	TmpID    string `json:"tmp_id"`
}
