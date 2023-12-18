package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wetrycode/teg-cli/render"
	"github.com/wetrycode/tegenaria"
)

var logger = tegenaria.GetLogger("command")
var outputDir string         // 用于存储输出目录的变量
var spiderName string        // 用于存储爬虫名称的变量
var projectMoudleName string // 用于存储项目模块名称的变量
var name, filename, output string
var logDir, logLevel string
var IsNew bool = true

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tg-cli [commad] [options]",
	Short: "tegenaria is a crawler framework based on golang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func AddCmdFlags(cmd *cobra.Command) {

	// 添加 --spider 标志
	cmd.Flags().StringVarP(&spiderName, "spider", "s", "", "Specify the spider name")

	// 添加 --moudle 标志
	cmd.Flags().StringVarP(&projectMoudleName, "moudle", "m", "", "Specify the project moudle name")

	// 添加 --logdir 标志
	cmd.Flags().StringVarP(&logDir, "logdir", "l", filepath.Join("/var/log", "tegenaria"), "Specify the log dir")

	// 添加 --loglevel 标志
	cmd.Flags().StringVarP(&logLevel, "loglevel", "v", "INFO", "Specify the log level")

	cmd.MarkFlagRequired("spider")
	cmd.MarkFlagRequired("moudle")
}

// ExecuteCmd manage engine by command
func ExecuteCmd() {

	var newCmd = &cobra.Command{
		Use:   "new projectName",
		Short: "Create a new tegenaria project",
		// Uncomment the following line if your bare application
		// has an action associated with it:

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
				return
			}

			projectName := args[0]
			// 假设你有一个函数来处理创建项目的逻辑
			if len(strings.TrimSpace(spiderName)) == 0 {
				logger.Error("spider name can not be empty")
				return
			}
			// 创建项目
			//
			if len(strings.TrimSpace(projectMoudleName)) == 0 {
				logger.Error("project moudle name can not be empty")
				return
			}
			// 创建项目
			project := &render.ProjectParams{
				ProjectName: projectName,
				SpiderName:  spiderName,
				OutputDir:   outputDir,
				MoudleName:  projectMoudleName,
				LogDir:      logDir,
				LogLevel:    logLevel,
				IsNew:       IsNew,
			}
			projectDir, err := render.NewRender().CreateNewProject(project)
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Infof("The project %s was created successfully.", projectDir)
		},
	}
	AddCmdFlags(newCmd)
	// 添加 --output 标志
	newCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Set the output directory")
	newCmd.MarkFlagRequired("output")

	var addCmd = &cobra.Command{
		Use:   "add spider|pipeline|mid",
		Short: "Add a new spider|pipeline|mid to project",
		Run: func(cmd *cobra.Command, args []string) {
			filePath := ""
			var err error = nil
			defer func() {
				if err != nil {
					os.RemoveAll(filePath)
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
					logger.Error(err)
					return
				}
			case "pipeline":
				filePath, err = render.NewRender().CreateNewPipeline(name, output, filename)
				if err != nil {
					logger.Error(err)
				}
			case "mid":
				filePath, err = render.NewRender().CreateNewMiddleware(name, output, filename)
				if err != nil {
					logger.Error(err)
					return
				}
			default:
				cmd.Help()
			}

		},
	}
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init a new tegenaria project",
		Run: func(cmd *cobra.Command, args []string) {
			workDir, _ := os.Getwd()
			lastFolder := filepath.Base(workDir)

			projectName := lastFolder
			args = make([]string, 0)
			args = append(args, projectName)
			outputDir = workDir
			IsNew = false
			newCmd.Run(cmd, args)

		},
	}
	AddCmdFlags(initCmd)
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
	rootCmd.AddCommand(initCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}

}

func main() {
	ExecuteCmd()
}
