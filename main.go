package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wetrycode/teg-cli/render"
	"github.com/wetrycode/tegenaria"
)

var logger = tegenaria.GetLogger("command")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tg-cli [commad] [options]",
	Short: "tegenaria is a crawler framework based on golang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// ExecuteCmd manage engine by command
func ExecuteCmd() {
	var outputDir string         // 用于存储输出目录的变量
	var spiderName string        // 用于存储爬虫名称的变量
	var projectMoudleName string // 用于存储项目模块名称的变量
	var name, filename, output string

	var newCmd = &cobra.Command{
		Use:   "new projectName",
		Short: "Create a new tegenaria project",
		// Uncomment the following line if your bare application
		// has an action associated with it:

		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("准备启动%s爬虫", args[0])
			if len(args) < 1 {
				cmd.Help()
				return
			}

			projectName := args[0]
			logger.Infof("准备启动%s爬虫", projectName)
			// 假设你有一个函数来处理创建项目的逻辑
			if len(strings.TrimSpace(spiderName)) == 0 {
				panic("spider name can not be empty")
			}
			// 创建项目
			//
			if len(strings.TrimSpace(projectMoudleName)) == 0 {
				panic("project moudle name can not be empty")
			}
			projectDir, err := render.NewRender().CreateNewProject(projectName, outputDir, spiderName, projectMoudleName)
			if err != nil {
				panic(err)
			}
			logger.Infof("项目%s创建成功", projectDir)
		},
	}
	var addCmd = &cobra.Command{
		Use:   "add spider|pipeline|mid",
		Short: "Add a new spider|pipeline|mid to project",
		Run: func(cmd *cobra.Command, args []string) {
			filePath := ""
			var err error = nil
			defer func() {
				if err != nil {
					os.RemoveAll(filePath)
					panic(err)
				}
				logger.Infof("文件%s创建成功", filePath)
			}()
			if len(args) < 1 {
				cmd.Help()
				return
			}
			addType := args[0]
			switch addType {
			case "spider":
				filePath, err = render.NewRender().CreateNewSpider(name, output, filename)
				if err != nil {
					panic(err)
				}
			case "pipeline":
				filePath, err = render.NewRender().CreateNewPipeline(name, output, filename)
				if err != nil {
					panic(err)
				}
			case "mid":
				filePath, err = render.NewRender().CreateNewMiddleware(name, output, filename)
			default:
				cmd.Help()
			}

		},
	}
	// 添加 --output 标志
	newCmd.Flags().StringVarP(&outputDir, "output", "o", "./", "Set the output directory")

	// 添加 --spider 标志
	newCmd.Flags().StringVarP(&spiderName, "spider", "s", "", "Specify the spider name")

	// 添加 --moudle 标志
	newCmd.Flags().StringVarP(&projectMoudleName, "moudle", "m", "", "Specify the project moudle name")

	newCmd.MarkFlagRequired("spider")
	newCmd.MarkFlagRequired("moudle")
	newCmd.MarkFlagRequired("output")
	// 添加 --name 标志
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the component")

	// 添加 --filename 标志
	addCmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename of the component")
	// 添加 --output 标志
	addCmd.Flags().StringVarP(&output, "output", "o", "", "Output directory")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("filename")
	addCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(addCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	ExecuteCmd()
}
