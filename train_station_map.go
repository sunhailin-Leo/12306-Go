package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var stationToCode map[string]string
var codeToStation map[string]string

type TrainStation struct {
	Station map[string]string `json:"station"`
}

type JsonStruct struct{}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (j *JsonStruct) Load(fileName string, v interface{}) {
	// ReadFile 函数会读取文件的全部内容，并将结果以 []byte 类型返回
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func loadJsonStation() {
	JsonParse := NewJsonStruct()
	stationMap := TrainStation{}
	JsonParse.Load(stationJson, &stationMap)
	// 车站 - Code
	stationToCode = stationMap.Station
	// Code - 车站
	tempMap := make(map[string]string)
	for k, v := range stationToCode {
		tempMap[v] = k
	}
	codeToStation = tempMap
}
