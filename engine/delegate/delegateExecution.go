package delegate

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/variable"
)

type DelegateExecution interface {
	SetBusinessKey(businessKey string)

	GetCurrentFlowElement() bpmn.FlowElement

	SetCurrentFlowElement(flow bpmn.FlowElement)

	GetDeploymentId() int

	SetDeploymentId(deploymentId int)

	GetProcessInstanceId() int64

	SetProcessInstanceId(processInstanceId int64)

	GetProcessDefineId() int64

	SetProcessDefineId(processDefineId int64)

	GetCurrentActivityId() string

	SetCurrentActivityId(currentActivityId string)

	//SetVariable(execution ExecutionEntity,variables map[string]interface{}) error

	GetSpecificVariable(variableName string) (variable.Variable, error)

	SetScope(variable *variable.Variable)

	GetVariable() map[string]interface{}

	GetProcessVariable() map[string]interface{}

	GetTaskId() int64

	SetTaskId(taskId int64)
}