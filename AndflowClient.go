package andflow_client

import (
 	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code int `json:"code"`
	Msg  string                `json:"msg"`
	Obj  interface{}           `json:"obj"`  //返回值对象

	Cipher string `json:"cipher"`
	Sign string	`json:"sign"`
	Key string  `json:"key"`
}
// 请求参数的模型基类
type JsonParam  struct {
	Chipher string `json:"chipher"`
	Sign    string `json:"sign"`
	Key     string `json:"key"`
}

type FlowRunParam struct {
	Async bool   `json:"async"`
	ClientId string  `json:"client_id"`
	RuntimeId string `json:"runtime_id"`
	FlowCode string  `json:"flow_code"`
	Params map[string]interface{} `json:"params"`
}

type AndflowClient struct {
	Url string
	AppId string
	Secret string
	TokenId string
 	Expires int64
}

func (c *AndflowClient)post(p string, content_type string,content string) (string,error){
	address :=  c.Url+"/"+p+"?TokenId="+c.TokenId
	if len(content_type)==0{
		content_type = "application/json"
	}




	resp, err := http.Post(address, content_type, strings.NewReader(content) )
	if err != nil {
		return "",err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}

	return string(body),nil
}

func (c *AndflowClient)login()error{
	values := url.Values{}
	values.Add("AppId",c.AppId)
	values.Add("Secret",c.Secret)

	res,err := c.post("token/login","application/x-www-form-urlencoded", values.Encode())
	if err!=nil{

		return err
	}

	result := JsonResult{}
	err = json.Unmarshal([]byte(res),&result)
	if err!=nil{
		return err
	}
	if result.Code!=0 {
		return errors.New(result.Msg)
	}
	if result.Obj==nil{
		return errors.New("登陆失败")
	}

	obj,ok := result.Obj.(map[string]interface{})
	if !ok{
		return errors.New("登陆失败")
	}

	c.TokenId = obj["Id"].(string)
	c.Expires = int64(obj["Expires"].(float64))

	return nil
}

func NewClient(url string, appid string , secret string )(*AndflowClient,error){
	client:=AndflowClient{Url:url,AppId:appid,Secret:secret}
	err:=client.login()
	return &client,err
}

//启动流程
func (c *AndflowClient) StartFlow(flowCode string,params map[string]interface{},async bool)(*RuntimeModel, error){

	if len(c.TokenId)==0 || c.Expires==0 || time.Now().Unix()>=c.Expires{
		err := c.login()
		if err!=nil{
			return nil,err
		}
	}

	p := FlowRunParam{Async:async, FlowCode:flowCode, Params:params}

	data, err := json.Marshal(p)
	if err!=nil{
		return nil, err
	}

	res,err := c.post("flow/run","application/json",string(data) )
	if err!=nil{
		return nil, err
	}
	result := JsonResult{}
	err = json.Unmarshal([]byte(res),&result)
	if err!=nil{
		return nil, err
	}


	if result.Code!=0 {
		return nil,errors.New(result.Msg)
	}
	if result.Obj==nil{
		return nil,errors.New("启动流程失败")
	}

	d,err:=json.Marshal(result.Obj)

	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	runtimeModel:=RuntimeModel{}

	err = json.Unmarshal(d, &runtimeModel)
	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	return &runtimeModel,nil


}

//运行流程
func (c *AndflowClient) RunFlow(runtimeId string, params map[string]interface{},async bool)(*RuntimeModel, error){
	if len(c.TokenId)==0 || c.Expires==0 || time.Now().Unix()>=c.Expires{
		err := c.login()
		if err!=nil{
			return nil,err
		}
	}

	p := FlowRunParam{Async:async,RuntimeId:runtimeId, Params:params}

	data, err := json.Marshal(p)
	if err!=nil{
		return nil, err
	}

	res,err := c.post("flow/run","application/json",string(data) )
	if err!=nil{
		return nil, err
	}
	result := JsonResult{}
	err = json.Unmarshal([]byte(res),&result)
	if err!=nil{
		return nil, err
	}


	if result.Code!=0 {
		return nil,errors.New(result.Msg)
	}
	if result.Obj==nil{
		return nil,errors.New("执行流程失败")
	}

	d,err:=json.Marshal(result.Obj)

	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	runtimeModel:=RuntimeModel{}

	err = json.Unmarshal(d, &runtimeModel)
	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	return &runtimeModel,nil

}


func (c *AndflowClient) GetRuntimesByFlowCode(flowCode string)([]*RuntimeModel,error){


	if len(c.TokenId)==0 || c.Expires==0 || time.Now().Unix()>=c.Expires{
		err := c.login()
		if err!=nil{
			return nil,err
		}
	}

	values := url.Values{}
	values.Add("flow_code",flowCode)

	res,err := c.post("flow/runtime_noend_flowcode","application/x-www-form-urlencoded", values.Encode())

	if err!=nil{
		return nil, err
	}
	result := JsonResult{}
	err = json.Unmarshal([]byte(res),&result)
	if err!=nil{
		return nil, err
	}

	if result.Code!=0 {
		return nil,errors.New(result.Msg)
	}
	if result.Obj==nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	d,err:=json.Marshal(result.Obj)

	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	runtimeModels:=make([]*RuntimeModel,0)

	err = json.Unmarshal(d, &runtimeModels)
	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	return runtimeModels,nil
}

func (c *AndflowClient) GetRuntimesByFlowCodeAndNextActionParam(flowCode string , key,value string)([]*RuntimeModel,error){



	if len(c.TokenId)==0 || c.Expires==0 || time.Now().Unix()>=c.Expires{
		err := c.login()
		if err!=nil{
			return nil,err
		}
	}

	values := url.Values{}
	values.Add("flow_code",flowCode)
	values.Add("key",key)
	values.Add("value",value)

	res,err := c.post("flow/runtime_noend_flowcode_next_action_param","application/x-www-form-urlencoded", values.Encode())

	if err!=nil{
		return nil, err
	}
	result := JsonResult{}
	err = json.Unmarshal([]byte(res),&result)
	if err!=nil{
		return nil, err
	}

	if result.Code!=0 {
		return nil,errors.New(result.Msg)
	}
	if result.Obj==nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	d,err:=json.Marshal(result.Obj)

	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	runtimeModels:=make([]*RuntimeModel,0)

	err = json.Unmarshal(d, &runtimeModels)
	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	return runtimeModels,nil
}

func (c *AndflowClient) GetRuntime(runtimeId string)(*RuntimeModel, error){

	if len(c.TokenId)==0 || c.Expires==0 || time.Now().Unix()>=c.Expires{
		err := c.login()
		if err!=nil{
			return nil,err
		}
	}

	values := url.Values{}
	values.Add("id",runtimeId)

	res,err := c.post("flow/getruntime","application/x-www-form-urlencoded", values.Encode())

	if err!=nil{
		return nil, err
	}
	result := JsonResult{}
	err = json.Unmarshal([]byte(res),&result)
	if err!=nil{
		return nil, err
	}

	if result.Code!=0 {
		return nil,errors.New(result.Msg)
	}
	if result.Obj==nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	d,err:=json.Marshal(result.Obj)

	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	runtimeModel:=RuntimeModel{}

	err = json.Unmarshal(d, &runtimeModel)
	if err!=nil{
		return nil,errors.New("获取流程运行信息失败")
	}

	return &runtimeModel,nil
}


func (r *RuntimeModel) GetData()map[string]interface{}{

	res := make(map[string]interface{})
	for _,d:=range r.Data{
		res[d.Name]=d.Value
	}
	return res
}

func (r *RuntimeModel) GetNextActions()[]*ActionModel{
	actions:=make([]*ActionModel,0)
	for _,link:=range r.NextLinks{
		id := link.TargetId
		action := r.Flow.GetActionModel(id)
		actions=append(actions, action)

	}
	return actions
}

