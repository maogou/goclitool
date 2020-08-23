package cmd

import (
	"github.com/spf13/cobra"
	"gocli/internal/sql2struct"
	"log"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sql2Cmd = &cobra.Command{
	Use: "sql",
	Short: "sql转换结构体",
	Long: "sql转换为Go结构体",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use: "struct",
	Short: "sql转换结构体",
	Long: "sql转换为Go结构体",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType: dbType,
			Host: host,
			UserName: username,
			Password: password,
			Charset: charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()

		if err != nil {
			log.Fatalf("dbModel.Connect err: %v",err)
		}

		columns,err := dbModel.GetColumns(dbName,tableName)

		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v",err)
		}

		template := sql2struct.NewStructTemplate()

		templateColumns := template.AssemblyColumns(columns)

		err = template.Generate(tableName,templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v",err)
		}
	},
}

func init()  {
	sql2Cmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}

