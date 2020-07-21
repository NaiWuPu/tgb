package main

import (
	"encoding/json"
	"fmt"
)

type Xgs_Msg_Type int

type Xgs_Msg_ST struct {
	Msg_type     Xgs_Msg_Type
	Msg_content  []byte
	Inner_policy bool
	File_desc    map[string]interface{}
}

type commandType struct {
	Type  string            `json:"type" description:""`
	Cmd   string            `json:"cmd" description:""`
	Param commandParamsType `json:"param" description:""`
}

type commandParamsType struct {
	Module             string `json:"module" description:""`
	PlugPolicyFilename string `json:"plug_policy_filename" description:""`
	Xgs_Msg_ST_File_desc
}

type Xgs_Msg_ST_File_desc struct {
	PlugId            interface{} `json:"plug_id" description:""`
	Cpu               interface{} `json:"cpu" description:""`
	Disk              interface{} `json:"disk" description:""`
	Mem               interface{} `json:"mem" description:""`
	PlugConfigVersion interface{} `json:"plug_config_version" description:""`
}

func main() {
	f1()
	f2()
}

func f1() {
	var work Xgs_Msg_ST
	command := make([]map[string]interface{}, 1)
	command[0] = make(map[string]interface{})
	command[0]["type"] = interface{}("plug")
	command[0]["cmd"] = interface{}("updates")
	param := make(map[string]interface{})
	param["module"] = interface{}("plug_policy")
	param["plug_policy_filename"] = interface{}("plug_policy_filename")
	param["plug_id"] = work.File_desc["plug_id"]
	param["plug_config_version"] = work.File_desc["plug_config_version"]
	param["cpu"] = work.File_desc["cpu"]
	param["mem"] = work.File_desc["mem"]
	param["disk"] = work.File_desc["disk"]
	command[0]["param"] = interface{}(param)
	ret_json, _ := json.Marshal(command)
	fmt.Printf("%s \n", ret_json)

}

func f2() {
	var work Xgs_Msg_ST
	var command = make([]commandType, 1)
	command[0].Type = "plug"
	command[0].Cmd = "updates"
	command[0].Param.Module = "plug_policy"
	command[0].Param.PlugPolicyFilename = "plug_policy_filename"
	command[0].Param.PlugConfigVersion = work.File_desc["plug_config_version"]
	command[0].Param.Cpu = work.File_desc["cpu"]
	command[0].Param.Mem = work.File_desc["mem"]
	command[0].Param.Disk = work.File_desc["disk"]
	retJson, _ := json.Marshal(command)
	fmt.Printf("%s \n", retJson)

}
