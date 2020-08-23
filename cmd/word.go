package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gocli/internal/word"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1  //全部转为大写
	ModeLower  //全部转为小写
	ModeUnderscoreToUpperCamelcase //下划线单词转大写单词
	ModeUnderscoreToLowerCamelcase //下划线单词转小写单词
	ModeCamelcaseToUnderscore //驼峰单词转为下划线单词
)

//描述信息
var desc = strings.Join([]string{
	"该命令支持各种单词格式转换,模式如下:",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
	"3:下划线单词转为大写驼峰单词",
	"4:下划线单词转为小写驼峰单词",
	"5:驼峰单词转为下划线单词",
},"\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换格式,请执行 help word 查看帮助")
		}

		fmt.Println("转换结果为:",content)
	},
}

//初始化参数绑定
func init()  {
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入要转换的单词")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",1,"请输入单词转换的模式")
}


