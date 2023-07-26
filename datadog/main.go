package main

import (
	"log"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/hashicorp/hcl2/hclwrite"
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

func main() {
	Run()
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
		label := b.Labels[0]
		if label != labelName {
			continue
		}

		diag = gohcl.DecodeBody(b.Body, nil, &monitor)
		if diag != nil {
			log.Fatal(diag)
		}

	}

	hclFile := hclwrite.NewFile()
	gohcl.EncodeIntoBody(&monitor, hclFile.Body())

}
