package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocli/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

//当前时间子命令
var nowTimeCmd = &cobra.Command{
	Use: "now",
	Short: "获取当前时间",
	Long: "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		s :=fmt.Sprintf("输出结果: %s, %d",nowTime.Format("2006-01-02 15:04:05"),nowTime.Unix())

		fmt.Println(s)
	},
}

var calculateTimeCmd = &cobra.Command{
	Use: "calc",
	Short: "计算所需时间",
	Long: "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime," "){
				layout = "2006-01-02"
			}

			currentTimer,err = time.Parse(layout,calculateTime)
			if err != nil {
				t,_ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t),0)
			}
		}

		calculateTime,err:= timer.GetCalculateTime(currentTimer,duration)

		if err != nil {
			log.Fatalf("timer.GetCalculateTime err:%v",err)
		}

		log.Printf("输出结果:%s, %d",calculateTime.Format(layout),calculateTime.Unix())
	},
}

//初始化方法
func init()  {
	//注册子命令的子命令
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime,"calculate","c","","需要计算的实际,有效单为时间戳或已经格式化的时间")
	calculateTimeCmd.Flags().StringVarP(&duration,"duration","d","","持续时间,有效时间单位 ns,us ,ms,s,m,h")
}
