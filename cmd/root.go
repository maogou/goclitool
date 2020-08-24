package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "",
	Short: "",
	Long: "",
}

//初始化注册子命令
func init() {
	//单词处理工具
	rootCmd.AddCommand(wordCmd)
	//时间处理工具
	rootCmd.AddCommand(timeCmd)
	//sql转结构体
	rootCmd.AddCommand(sql2Cmd)
	//curl请求
	rootCmd.AddCommand(httpCmd)
}


//执行子命令
func Execute() error  {
	return rootCmd.Execute()
}


