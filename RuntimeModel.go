package andflow_client

import (
	"time"
)


type ActionContentModel struct {
	ActionId string 		`bson:"action_id" json:"action_id""`
	ContentType string      `bson:"content_type" json:"content_type"`
	Content string          `bson:"content" json:"content"`
}

type ActionDataModel struct{
	Name 	string 		 `bson:"name" json:"name"`
	Value  	interface{}  `bson:"value" json:"value"`   //执行结果
}

type ActionStateModel struct {
	ActionId string 			`bson:"action_id" json:"action_id""`
	ActionName string 			`bson:"action_name" json:"action_name"`
	ActionTitle string		 	`bson:"action_title" json:"action_title"`
	ActionDes string 			`bson:"action_des" json:"action_des"`
	ActionIcon string 			`bson:"action_icon" json:"action_icon"`

	IsError int 				`bson:"is_error" json:"is_error"`
	State int 					`bson:"state" json:"state"`
	Data    []*ActionDataModel 	`bson:"data" json:"data"`			//执行结果
	Content *ActionContentModel `bson:"content" json:"content"`     //界面显示的内容
	BeginTime time.Time 		`bson:"begin_time" json:"begin_time"`
	EndTime time.Time 			`bson:"end_time" json:"end_time"`
}

type LinkStateModel struct {
	SourceActionId string 	`bson:"source_action_id" json:"source_action_id"`
	TargetActionId string	`bson:"target_action_id" json:"target_action_id"`
	IsError int				`bson:"is_error" json:"is_error"`
	State int				`bson:"state" json:"state"`
	BeginTime  time.Time	`bson:"begin_time" json:"begin_time"`
	EndTime    time.Time 	`bson:"end_time" json:"end_time"`

}

type LogModel struct {
	Tp   string 		`bson:"tp" json:"tp"`
	Id   string 		`bson:"id" json:"id"`
	Name string 		`bson:"name" json:"name"`
	Title string	 	`bson:"title" json:"title"`
	Time time.Time	 	`bson:"time" json:"time"`
	Tag    string 	 	`bson:"tag" json:"tag"`
	Content string 	 	`bson:"content" json:"content"`
}

type ParamInfo struct {
	TypeName string `bson:"type_name" json:"type_name"`
	Size     int    `bson:"size" json:"size"`
	ExpireMillisecond int64   `bson:"expire_millisecond" json:"expire_millisecond"`
}


type RuntimeDataModel struct{
	Name 	string 		 `bson:"name" json:"name"`
	Value  	interface{}  `bson:"value" json:"value"`   //执行结果
}

type RuntimeModel struct {
	Id               string                `bson:"_id" json:"id"`
	ContextId        string                `bson:"context_id" json:"context_id"`
	Des              string                `bson:"des" json:"des"`
	BeginTime        time.Time               `bson:"start_time" json:"start_time"`
	EndTime          time.Time               `bson:"end_time" json:"end_time"`
	CurrentStartTime time.Time             `bson:"current_start_time" json:"current_start_time"`
	CurrentStopTime  time.Time             `bson:"current_stop_time" json:"current_stop_time"`
	CurrentTimeUsed  int64                 `bson:"current_time_used" json:"current_time_used"`
	IsRunning        int                   `bson:"is_running" json:"is_running"`
	IsError          int                   `bson:"is_error" json:"is_error"`
	Flow             *FlowModel            `bson:"flow" json:"flow"`
 	NextLinks        map[string]*LinkModel `bson:"next_links" json:"next_links"`
	ParamInfos       map[string]*ParamInfo `bson:"param_infos" json:"param_infos"`
	FlowState        int                   `bson:"flow_state" json:"flow_state"`
	ActionStates     []*ActionStateModel   `bson:"action_states" json:"action_states"`
	LinkStates       []*LinkStateModel     `bson:"link_states" json:"link_states"`
	Logs             []*LogModel           `bson:"logs" json:"logs"`
	ClientId         string                `bson:"client_id" json:"client_id"`
	NodeId           string                `bson:"node_id" json:"node_id"`
	LogPath        string                  `bson:"log_path" json:"log_path"`
	Data  		   []*RuntimeDataModel		 `bson:"data" json:"data"`
	CreateUser     int                       `bson:"create_user" json:"create_user"`
	CreateTime     time.Time              	 `bson:"create_time" json:"create_time"`
	UpdateUser     int						 `bson:"update_user" json:"update_user"`
	UpdateTime     time.Time              	 `bson:"update_time" json:"update_time"`
}

