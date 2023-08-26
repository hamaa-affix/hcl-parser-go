package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
	// "github.com/hashicorp/hcl2/hclwrite"
)

const (
	labelName = "datadog_monitor"
)

type Monitor struct {
	Type              string   `hcl:"type"`
	Name              string   `hcl:"name"`
	Query             string   `hcl:"query"`
	Message           string   `hcl:"message"`
	NotifyAudit       bool     `hcl:"notify_audit"`
	NewHostDelay      int      `hcl:"new_host_delay"`
	TimeoutH          int      `hcl:"timeout_h"`
	NotifyNoData      bool     `hcl:"notify_no_data"`
	NoDataTimeFrame   int      `hcl:"no_data_timeframe"`
	RenotifyInterval  int      `hcl:"renotify_interval"`
	Tags              []string `hcl:"tags"`
	IncludeTags       bool     `hcl:"include_tags"`
	MonitorThresholds `hcl:"monitor_thresholds,block"`
}

type MonitorThresholds struct {
	Ok       int `hcl:"ok"`
	Warning  int `hcl:"warning"`
	Critical int `hcl:"critical"`
}

type MapResource struct {
	Name  string
	Query string
}

func (m *MapResource) addName(monitor Monitor) {
	m.Name = monitor.Name
}

func (m *MapResource) addQuery(monitor Monitor) {
	m.Query = monitor.Query
}

func main() {
	data, err := ioutil.ReadFile("./monitor.tf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	// Run()
}

func Run() {
	parser := hclparse.NewParser()
	file, diag := parser.ParseHCLFile("./monitor.tf")
	if diag != nil {
		log.Fatalf("parse error: %s", diag.Error())
	}

	// 読み込んだblock.bodyの不要なattributeを削りたかった。
	// 調べないとわからないが、下記はFile{}のblock一覧をスライスに格納、
	// それをloopで回してblock.Bodyのattributeを精査しようとした・
	// なぜやろうとしたか、スキーマ定義(構造体)をミニマムにしたかったのでやろうとした
	// blocks := file.BlocksAtPos(hcl.Pos{})
	// for _, block := range blocks {
	// 	atr, diag := block.Body.JustAttributes()
	// 	if diag != nil {
	// 		log.Fatal(diag)
	// 	}

	// 	for key, value := range atr {
	// 		// if value.Name == "monitor_thresholds" {
	// 		// 	continue
	// 		// }
	// 		fmt.Println(key)
	// 		fmt.Println(value.Expr)
	// 	}
	// }

	var monitor Monitor
	//BlocksAtPos は、指定された位置を含むすべてのブロックを、
	//最も外側のブロックが最初で最も内側のブロックが最後になるように順序付けして検索しようとします。
	//これはベストエフォート型の方法であるため、すべての位置またはすべての HCL 構文に対して完全な結果を生成できるわけではありません。
	// hcl.Pos{}.Line(int) これは0となるから、file型の先頭からblockを抽出してスライスに格納している
	main := file.BlocksAtPos(hcl.Pos{})
	for _, b := range main {
		if b.Type != "resource" {
			continue
		}

		label := b.Labels[0]
		if label != labelName {
			continue
		}

		diag = gohcl.DecodeBody(b.Body, nil, &monitor)
		if diag != nil {
			log.Fatal(diag)
		}
	}

	var mapResource MapResource
	mapResource.addName(monitor)
	mapResource.addQuery(monitor)

	records := [][]string{
		{"name", "query"},
		{mapResource.Name, mapResource.Query},
	}

	csvFile, err := os.Create("demo.csv")
	if err != nil {
		log.Fatalf("error: create csv file => %v", err)
	}
	defer csvFile.Close()

	cw := csv.NewWriter(csvFile)
	cw.WriteAll(records)

	// hclFile := hclwrite.NewFile()
	// gohcl.EncodeIntoBody(&monitor, hclFile.Body())

}
