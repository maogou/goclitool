package cmd

import (
	"github.com/spf13/cobra"
	"gocli/internal/curl"
	"log"
	"strings"
)

var url string
var method string


var httpCmd = &cobra.Command{
	Use: "http",
	Short: "http请求",
	Long: "http方式的请求",
	Run: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(method) {
		case "get":
			curl.Get(url)
		default:
			log.Fatalf("暂不支持该方法: %s",method)
		}
	},
}

func init()  {
	httpCmd.Flags().StringVarP(&url,"url","u","localhost","请输入url")
	httpCmd.Flags().StringVarP(&method,"method","m","get","请求方法")
}
