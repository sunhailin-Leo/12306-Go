package main

import (
	"fmt"
	"github.com/go-resty/resty"
	"github.com/liyu4/tablewriter"
	"github.com/tidwall/gjson"
	"os"
	"strings"
	"time"
)

const (
	baseSearchTrainInfoURL string = "https://search.12306.cn/search/v1/train/search?keyword={code}&date={date}"
	baseGetTrainInfoURL    string = "https://kyfw.12306.cn/otn/queryTrainInfo/query?leftTicketDTO.train_no={trainCode}&leftTicketDTO.train_date={trainDate}&rand_code="
)

type TrainInfoTable struct {
	TrainNo string
	Date    string

	RestClient *resty.Client

	SearchURL    string
	RealTrainNo  string
	TrainInfoURL string
}

func (t *TrainInfoTable) searchByTrainNo() *TrainInfoTable {
	dataResp, err := t.RestClient.R().
		SetHeader("Origin", "https://kyfw.12306.cn").
		SetHeader("Referer", "https://kyfw.12306.cn/otn/queryTrainInfo/init").
		SetHeader("Sec-Fetch-Mode", "cors").
		SetHeader("User-Agent", userAgent).
		Get(t.SearchURL)
	if err != nil {
		logger.Fatalf("[12306查询助手]车次搜索异常, 错误原因: %v", err)
	}
	if dataResp != nil {
		jsonData := gjson.Parse(dataResp.String())
		for _, data := range jsonData.Get("data").Array() {
			if t.TrainNo == data.Get("station_train_code").String() {
				t.RealTrainNo = data.Get("train_no").String()
			}
		}
		if t.RealTrainNo == "" {
			logger.Infof("[12306查询助手]无法查询到当前车次信息!")
			os.Exit(1)
		}
	}
	return t
}

func (t *TrainInfoTable) searchTrainInfoByRealTrainNo() *TrainInfoTable {
	t.TrainInfoURL = stringFormat(baseGetTrainInfoURL, "{trainCode}", t.RealTrainNo, "{trainDate}", t.Date)
	dataResp, err := t.RestClient.R().
		SetHeader("Referer", "https://kyfw.12306.cn/otn/queryTrainInfo/init").
		SetHeader("User-Agent", userAgent).
		Get(t.TrainInfoURL)
	if err != nil {
		logger.Fatalf("[12306查询助手]车次信息查询异常, 错误原因: %v", err)
	}
	if dataResp != nil {
		jsonData := gjson.Parse(dataResp.String())
		table := tablewriter.NewColorWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader([]string{"站序", "车站", "出发时间", "到达时间", "停站时间", "运行时间"})
		for index, station := range jsonData.Get("data").Get("data").Array() {
			gap := fmt.Sprintf("0 分钟")
			if index > 0 {
				arrive, _ := time.Parse("15:04", station.Get("arrive_time").String())
				start, _ := time.Parse("15:04", station.Get("start_time").String())
				gap = fmt.Sprintf("%.0f 分钟", start.Sub(arrive).Minutes())
			}
			row := []string{
				station.Get("station_no").String(),
				station.Get("station_name").String(),
				station.Get("arrive_time").String(),
				station.Get("start_time").String(),
				gap,
				station.Get("running_time").String(),
			}
			table.Append(row)
		}
		table.Render()
	}
	return t
}

func (t *TrainInfoTable) runParser() *TrainInfoTable {
	t.searchByTrainNo()
	time.Sleep(time.Second)
	t.searchTrainInfoByRealTrainNo()
	return t
}

func (t *TrainInfoTable) initTrainInfoTable() *TrainInfoTable {
	t.RestClient = resty.New()
	t.SearchURL = stringFormat(baseSearchTrainInfoURL, "{code}", t.TrainNo, "{date}", strings.Replace(t.Date, "-", "", -1))
	return t
}

func NewTrainInfoTable(trainNo, date string) *TrainInfoTable {
	trainInfo := &TrainInfoTable{
		TrainNo: strings.ToUpper(trainNo),
		Date:    date,
	}
	trainInfo.initTrainInfoTable()
	return trainInfo
}
