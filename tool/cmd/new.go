package cmd

import (
	"errors"
	module "github.com/LCY2013/paas/tool/module"
	"github.com/spf13/cobra"
)

// 自动生成 service 目录
var new = &cobra.Command{
	Use:   "new",
	Short: "Customize the service and automatically generate the service directory",
	Long:  `This command can automatically generate the project directory, which is convenient to quickly create the basic project code`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please enter the project name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := module.NewServiceProject(cmd, args)
		if err != nil {
			return
		}
	},
}

// 自动生成 service 目录
var newService = &cobra.Command{
	Use:   "newService",
	Short: "Customize the service and automatically generate the service directory",
	Long:  `This command can automatically generate the project directory, which is convenient to quickly create the basic project code`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please enter the project name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := module.NewServiceProject(cmd, args)
		if err != nil {
			return
		}
	},
}

//自动生成基础接口程序

var createApi = &cobra.Command{
	Use:   "createApi",
	Short: "Customize the service and automatically generate the API directory",
	Long:  `This command can automatically generate the project directory, which is convenient to quickly create the basic project code`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please enter the project name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := module.NewApiProject(cmd, args)
		if err != nil {
			return
		}
	},
}

func init() {

	//添加命令
	rootCmd.AddCommand(new)
	rootCmd.AddCommand(newService)
	rootCmd.AddCommand(createApi)
	//设置flags
	//newCmd.PersistentFlags().String("service", "s", "自动生成 service 项目代码")

}
