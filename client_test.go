package andflow_client

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {

	NewClient("http://localhost:8090/api","ed96458c4ad4455dbd41498648e19aca","1413b415cc274458bbc43b3c1ed6c821")


}

//启动流程
func TestStartFlow(t *testing.T) {

	client,err:=NewClient("http://localhost:8090/api","ed96458c4ad4455dbd41498648e19aca","1413b415cc274458bbc43b3c1ed6c821")
	if err!=nil{
		panic(err)

	}
	params:=make(map[string]interface{})
	params["user_id"] = "user1"
	params["state"] = "yes"
	params["description"] = "user1 test"

	runtimeId,err:=client.StartFlow("0f64609e32f447738c82605c9e688b40",params,false)

	fmt.Println(runtimeId)
}
//执行流程
func TestRunFlow(t *testing.T) {

	client,err:=NewClient("http://localhost:8090/api","ed96458c4ad4455dbd41498648e19aca","1413b415cc274458bbc43b3c1ed6c821")
	if err!=nil{
		panic(err)

	}
	params:=make(map[string]interface{})
	params["user_id"] = "user1"
	params["state"] = "yes"
	params["description"] = "user1 test"

	runtimeId,err:=client.RunFlow("53aeb697ee0e42b7b4069c686a56c0a9",params,false)

	fmt.Println(runtimeId)
}

//加载还未完成的流程
func TestLoadRuntimeFlowCode(t *testing.T) {

	client,err:=NewClient("http://localhost:8090/api","ed96458c4ad4455dbd41498648e19aca","1413b415cc274458bbc43b3c1ed6c821")
	if err!=nil{
		panic(err)

	}

	runtimes,err:=client.GetRuntimesByFlowCode("0f64609e32f447738c82605c9e688b40")

	fmt.Println(runtimes)
}

//加载还未完成的用户代办流程
func TestLoadRuntimeFlowCodeAndNextActionParam(t *testing.T) {

	client,err:=NewClient("http://localhost:8090/api","ed96458c4ad4455dbd41498648e19aca","1413b415cc274458bbc43b3c1ed6c821")
	if err!=nil{
		panic(err)

	}

	runtimes,err:=client.GetRuntimesByFlowCodeAndNextActionParam("0f64609e32f447738c82605c9e688b40","user_id","user2")

	fmt.Println(runtimes)
}