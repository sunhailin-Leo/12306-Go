package main

import (
	"fmt"
	"github.com/go-resty/resty"
	"github.com/liyu4/tablewriter"
	"github.com/tidwall/gjson"
	"os"
	"strings"
)

const (
	baseTicketParseURL string = "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date={date}&leftTicketDTO.from_station={from}&leftTicketDTO.to_station={to}&purpose_codes={ticketType}"
)

const (
	HXHBed       int = 33 // 动卧
	SeatBusiness int = 32 // 商务座
	SeatFirst    int = 31 // 一等座
	SeatSecond   int = 30 // 二等座
	HardSeat     int = 29 // 硬座
	HardBed      int = 28 // 硬卧
	NoSeat       int = 26 // 无座
	SeatSpecial  int = 25 // 特等座
	SoftSeat     int = 23 // 软座
)

type TrainTable struct {
	FromStation  string
	ToStation    string
	Date         string
	PurposeCodes string

	RestClient *resty.Client
	TableURL   string
}

func (t *TrainTable) runParser() *TrainTable {
	dataResp, err := t.RestClient.R().
		SetHeader("Host", "kyfw.12306.cn").
		SetHeader("Referer", "https://kyfw.12306.cn/otn/leftTicket/init").
		SetHeader("Cookie", "JSESSIONID=962FAEAB08FA8C00A10A471326B20B4E;").
		SetHeader("User-Agent", userAgent).
		SetHeader("X-Requested-With", "XMLHttpRequest").
		Get(t.TableURL)
	if err != nil {
		logger.Fatalf("[12306查询助手]查询异常, 错误原因: %v", err)
	}
	if dataResp != nil {
		jsonData := gjson.Parse(dataResp.String())
		table := tablewriter.NewColorWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader([]string{"车次", "始发站", "终到站", "出发站", "到达站", "出发时间", "到达时间", "历时", "商务/特等座", "一等座", "二等座", "动卧", "硬卧", "软座", "硬座", "无座"})
		for _, data := range jsonData.Get("data").Get("result").Array() {
			rowData := strings.Split(data.String(), "|")
			// 解析表格
			trainCode := fmt.Sprintf("%s", rowData[3])
			BeginStation := fmt.Sprintf("\033[31m(始)\033[0m:%s", codeToStation[rowData[4]])
			EndStation := fmt.Sprintf("\033[32m(终)\033[0m:%s", codeToStation[rowData[5]])
			FromStation := fmt.Sprintf("%s", codeToStation[rowData[6]])
			ToStation := fmt.Sprintf("%s", codeToStation[rowData[7]])
			FromStationStartTime := fmt.Sprintf("出发时间: %s", rowData[8])
			ToStationArriveTime := fmt.Sprintf("到达时间: %s", rowData[9])
			DurationTime := fmt.Sprintf("历时: %s", rowData[10])
			TrainHeadSeat := fmt.Sprintf("商务/特等座: %s", If(len(rowData[SeatBusiness]) > 0, rowData[SeatBusiness], "--"))
			TrainFirstSeat := fmt.Sprintf("一等座: %s", If(len(rowData[SeatFirst]) > 0, rowData[SeatFirst], "--"))
			TrainSecondSeat := fmt.Sprintf("二等座: %s", If(len(rowData[SeatSecond]) > 0, rowData[SeatSecond], "--"))
			TrainSpecialHXHBed := fmt.Sprintf("动卧: %s", If(len(rowData[HXHBed]) > 0, rowData[HXHBed], "--"))
			TrainHardBed := fmt.Sprintf("硬卧: %s", If(len(rowData[HardBed]) > 0, rowData[HardBed], "--"))
			TrainSoftSeat := fmt.Sprintf("软座: %s", If(len(rowData[SoftSeat]) > 0, rowData[SoftSeat], "--"))
			TrainHardSeat := fmt.Sprintf("硬座: %s", If(len(rowData[HardSeat]) > 0, rowData[HardSeat], "--"))
			TrainNoSeat := fmt.Sprintf("无座: %s", If(len(rowData[NoSeat]) > 0, rowData[NoSeat], "--"))
			row := []string{
				trainCode, BeginStation, EndStation, FromStation,
				ToStation, FromStationStartTime, ToStationArriveTime, DurationTime,
				TrainHeadSeat, TrainFirstSeat, TrainSecondSeat, TrainSpecialHXHBed,
				TrainHardBed, TrainSoftSeat, TrainHardSeat, TrainNoSeat,
			}
			table.Append(row)
		}
		table.Render()
	}
	return t
}

func (t *TrainTable) initTrainTable() *TrainTable {
	t.FromStation = stationToCode[t.FromStation]
	if t.FromStation == "" {
		logger.Fatalf("[12306查询助手]起始站输入有误!")
	}
	t.ToStation = stationToCode[t.ToStation]
	if t.ToStation == "" {
		logger.Fatalf("[12306查询助手]到达站输入有误!")
	}
	t.RestClient = resty.New()
	t.TableURL = stringFormat(baseTicketParseURL, "{date}", t.Date, "{from}", t.FromStation, "{to}", t.ToStation, "{ticketType}", t.PurposeCodes)
	return t
}

func NewTrainTable(from, to, date, purposeCodes string) *TrainTable {
	if purposeCodes == "" {
		purposeCodes = "ADULT"
	}
	table := &TrainTable{
		FromStation:  from,
		ToStation:    to,
		Date:         date,
		PurposeCodes: purposeCodes,
	}
	table.initTrainTable()
	return table
}
