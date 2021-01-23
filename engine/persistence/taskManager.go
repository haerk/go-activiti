package persistence

import (
	"fmt"
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/engine"
	. "github.com/lios/go-activiti/entity"
	"github.com/lios/go-activiti/errs"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
	"time"
)

type TaskManager struct {
	Task *Task
}

func (taskManager TaskManager) Insert(execution engine.ExecutionEntity) (err error) {
	err = db.DB().Create(taskManager.Task).Error
	if err == nil {
		err = taskManager.recordTaskCreated(taskManager.Task, execution)
	}
	//dispatcher := event.GetEventDispatcher()
	//dispatcher.DispatchEvent(CreateEntityEvent())
	return err
}

func (taskManager TaskManager) recordTaskCreated(task *Task, entity engine.ExecutionEntity) (err error) {
	historicTaskManager := HistoricTaskManager{}
	historicTask := taskManager.createHistoricTask(task)
	historicTaskManager.HistoricTask = historicTask
	err = historicTaskManager.Insert()
	if err != nil {
		historicActinstManager := HistoricActinstManager{}
		actinst, err := historicActinstManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), task.TaskDefineKey)
		if err == nil {
			actinst.Assignee = task.Assignee
			actinst.TaskId = task.Id
			historicActinstManager.HistoricActinst = actinst
			err = historicActinstManager.Update()
		}
	}
	return err
}

func (taskManager TaskManager) createHistoricTask(task *Task) HistoricTask {
	historicTask := HistoricTask{}
	//historicTask.TaskEntity = task.TaskEntity
	historicTask.TaskId = task.Id
	historicTask.ProcessInstanceId = task.ProcessInstanceId
	historicTask.StartTime = task.StartTime
	historicTask.TenantId = task.TenantId
	historicTask.Assignee = task.Assignee
	historicTask.TaskDefineKey = task.TaskDefineKey
	historicTask.DeploymentId = task.DeploymentId
	historicTask.TaskDefineName = task.TaskDefineName
	return historicTask
}

func (taskManager TaskManager) FindById(taskId int) (Task, error) {
	task := Task{}
	err := db.DB().Where("id= ?", taskId).First(&task).Error
	if err != nil {
		log.Infoln("Select FindById Err ", err)
		return task, err
	}
	return task, nil
}

func (taskManager TaskManager) FindByProcessInstanceId(processInstanceId int64) (task []Task, err error) {
	task = make([]Task, 0)
	err = db.DB().Where("proc_inst_id=?", processInstanceId).Find(&task).Error
	if err != nil {
		log.Infoln("Select FindByProcessInstanceId err ", err)
	}
	if task == nil || len(task) <= 0 {
		return task, errs.ProcessError{Code: "1001", Msg: "Not find"}
	}
	return task, err
}

func (taskManager TaskManager) DeleteTask(task Task) (err error) {
	err = db.DB().Where("id = ?", task.Id).Delete(&task).Error
	if err != nil {
		return err
	}
	identityLinkManager := IdentityLinkManager{}
	identityLinks, errSelect := identityLinkManager.SelectByTaskId(task.Id)
	if errSelect == nil {
		for _, identityLink := range identityLinks {
			err = identityLinkManager.Delete(identityLink.Id)
			if err != nil {
				return err
			}
		}
	}
	variableManager := VariableManager{}
	variables, errSelect := variableManager.SelectByTaskId(task.Id)
	if errSelect == nil {
		for _, variable := range variables {
			err = variableManager.Delete(variable.Id)
			if err != nil {
				return err
			}
		}
	}
	err = recordTaskEnd(task)
	return err
}

func recordTaskEnd(task Task) (err error) {
	historicTaskManager := HistoricTaskManager{}
	historicTask := HistoricTask{}
	historicTask.TaskId = task.Id
	historicTask.EndTime = time.Now()
	historicTaskManager.HistoricTask = historicTask
	err = historicTaskManager.MarkEnded()
	if err != nil {
		return err
	}

	historicActinst := HistoricActinst{}
	historicActinst.EndTime = historicTask.EndTime
	historicActinst.TaskId = historicTask.TaskId
	historicActinstManager := HistoricActinstManager{}
	historicActinstManager.HistoricActinst = historicActinst
	return historicActinstManager.UpdateTaskId()
}

func (taskManager TaskManager) QueryUndoTask(userId, groupId string) (taskResult []TaskEntity, err error) {
	taskResult = make([]TaskEntity, 0)
	var sql = "SELECT  t.*,i.user_id,i.group_id FROM task t " +
		"LEFT JOIN identity_link i on t.id = i.task_id " +
		"WHERE 1=1 "
	if userId != "" {
		sql += fmt.Sprintf("AND i.user_id = '%s' ", userId)
	}
	if groupId != "" {
		sql += fmt.Sprintf("AND i.group_id = '%s' ", groupId)
	}
	err = db.DB().Raw(sql).Find(&taskResult).Error
	if err != nil {
		return taskResult, err
	}
	return taskResult, nil
}
