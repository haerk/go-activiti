package model

//
//import (
//	"encoding/xml"
//	"github.com/lios/go-activiti/engine"
//	"github.com/lios/go-activiti/errs"
//)
//
//var (
//	//将元素存入map
//	processMap = make(map[int64]Process, 0)
//)
//
////流程定义对象
//type Definitions struct {
//	DefinitionsName    xml.Name  `xml:"definitions"`
//	Xmlns              string    `xml:"xmlns,attr"`
//	Xsi                string    `xml:"xsi,attr"`
//	Xsd                string    `xml:"xsd,attr"`
//	Activiti           string    `xml:"activiti,attr"`
//	Bpmndi             string    `xml:"bpmndi,attr"`
//	Omgdc              string    `xml:"omgdc,attr"`
//	Omgdi              string    `xml:"omgdi,attr"`
//	TypeLanguage       string    `xml:"typeLanguage,attr"`
//	RgetNamespace      string    `xml:"rgetNamespace,attr"`
//	ExpressionLanguage string    `xml:"expressionLanguage,attr"`
//	TargetNamespace    string    `xml:"targetNamespace,attr"`
//	Process            []Process `xml:"process"`
//	Message            []Message `xml:"message"`
//}
//type Process struct {
//	BaseElement
//	ProcessName            xml.Name                 `xml:"process"`
//	Id                     string                   `xml:"id,attr"`
//	Name                   string                   `xml:"name,attr"`
//	Documentation          string                   `xml:"documentation"`
//	IsExecutable           string                   `xml:"isExecutable,attr"`
//	StartEvent             []StartEvent             `xml:"startEvent"`
//	EndEvent               []EndEvent               `xml:"endEvent"`
//	UserTask               []UserTask               `xml:"userTask"`
//	SequenceFlow           []SequenceFlow           `xml:"sequenceFlow"`
//	ExclusiveGateway       []ExclusiveGateway       `xml:"exclusiveGateway"`
//	InclusiveGateway       []InclusiveGateway       `xml:"inclusiveGateway"`
//	ParallelGateway        []ParallelGateway        `xml:"parallelGateway"`
//	BoundaryEvent          []BoundaryEvent          `xml:"boundaryEvent"`
//	IntermediateCatchEvent []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
//	SubProcess             []SubProcess             `xml:"subProcess"`
//	Flow                   []FlowElement
//	InitialFlowElement     FlowElement
//	FlowMap                map[string]FlowElement
//}
//
////子流程
//type SubProcess struct {
//	*Process
//	SubProcessName xml.Name `xml:"subProcess"`
//}
//
////消息订阅
//type Message struct {
//	BaseElement
//	MessageName xml.Name `xml:"message"`
//}
//
////父类实现体
//type Flow struct {
//	BaseElement
//	Id                string `xml:"id,attr"`
//	Name              string `xml:"name,attr"`
//	IncomingFlow      []FlowElement
//	OutgoingFlow      []FlowElement
//	SourceFlowElement FlowElement
//	TargetFlowElement FlowElement
//	Behavior          engine.ActivityBehavior
//}
//
////开始节点
//type StartEvent struct {
//	*Flow
//	StartEventName xml.Name `xml:"startEvent"`
//	Initiator      string   `xml:"initiator,attr"`
//	FormKey        string   `xml:"formKey,attr"`
//}
//
////结束节点
//type EndEvent struct {
//	*Flow
//	EndEventName xml.Name `xml:"endEvent"`
//}
//
////用户任务
//type UserTask struct {
//	*Flow
//	UserTaskName      xml.Name          `xml:"userTask"`
//	Assignee          string            `xml:"assignee,attr"`
//	CandidateUsers    []string          `xml:"candidateUsers,attr"`
//	CandidateGroups   []string          `xml:"candidateGroups,attr"`
//	ExtensionElements ExtensionElements `xml:"extensionElements"`
//	MultiInstance     MultiInstance      `xml:"multiInstanceLoopCharacteristics"`
//}
//
//type MultiInstance struct {
//	IsSequential bool `xml:"isSequential,attr"`
//	Collection string `xml:"collection,attr"`
//	CompletionCondition string `xml:"completionCondition"`
//}
//type ExtensionElements struct {
//	ExtensionElementName xml.Name       `xml:"extensionElements"`
//	TaskListener         []TaskListener `xml:"taskListener"`
//}
//type TaskListener struct {
//	TaskListenerName xml.Name `xml:"taskListener"`
//	EventType        string   `xml:"event,attr"`
//}
//
////连线
//type SequenceFlow struct {
//	*Flow
//	SequenceFlowName    xml.Name `xml:"sequenceFlow"`
//	Id                  string   `xml:"id,attr"`
//	SourceRef           string   `xml:"sourceRef,attr"`
//	TargetRef           string   `xml:"targetRef,attr"`
//	ConditionExpression string   `xml:"conditionExpression"`
//}
//type Gateway struct {
//	*Flow
//	DefaultFlow string
//}
//
////排他网关
//type ExclusiveGateway struct {
//	*Gateway
//}
//
////包容网关
//type InclusiveGateway struct {
//	*Gateway
//}
//
////并行网关
//type ParallelGateway struct {
//	*Gateway
//}
//
////边界事件
//type BoundaryEvent struct {
//	*Flow
//	BoundaryEventName    xml.Name             `xml:"boundaryEvent"`
//	AttachedToRef        string               `xml:"attachedToRef,attr"`
//	CancelActivity       string               `xml:"cancelActivity,attr"`
//	TimerEventDefinition TimerEventDefinition `xml:"timerEventDefinition"`
//}
//
////定时任务
//type TimerEventDefinition struct {
//	TimerEventDefinitionName xml.Name `xml:"timerEventDefinition"`
//	TimeDuration             string   `xml:"timeDuration"`
//}
//
////中间抛出事件
//type IntermediateCatchEvent struct {
//	*Flow
//	IntermediateCatchEventName xml.Name               `xml:"intermediateCatchEvent"`
//	MessageEventDefinition     MessageEventDefinition `xml:"messageEventDefinition"`
//}
//
////消息事件
//type MessageEventDefinition struct {
//	MessageEventDefinitionName xml.Name `xml:"messageEventDefinition"`
//	MessageRef                 string   `xml:"messageRef,attr"`
//}
//
////接口
//type FlowElement interface {
//	SetIncoming(f []FlowElement)
//	SetOutgoing(f []FlowElement)
//	GetIncoming() []FlowElement
//	GetOutgoing() []FlowElement
//
//	SetSourceFlowElement(f FlowElement)
//	SetTargetFlowElement(f FlowElement)
//	GetSourceFlowElement() FlowElement
//	GetTargetFlowElement() FlowElement
//
//	GetBehavior() engine.ActivityBehavior
//	SetBehavior(behavior engine.ActivityBehavior)
//
//}
//
//func SetProcess(defineId int64, process Process) {
//	//_,err := processMap[process.Id]
//	processMap[defineId] = process
//}
//
//func GetProcess(id int64) (Process, error) {
//	process, ok := processMap[id]
//	if !ok {
//		return Process{}, errs.ProcessError{}
//	}
//	return process, nil
//}
//
//func (pocess Process) GetFlowElement(flowElementId string) FlowElement {
//	return pocess.FlowMap[flowElementId]
//}
