package render

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Render struct {
}
type ProjectParams struct {
	ProjectName string
	SpiderName  string
	OutputDir   string
	MoudleName  string
	LogDir      string
	LogLevel    string
	IsNew       bool
}

func NewRender() *Render {
	return &Render{}
}
func cleanDir(dirs []string) {
	for _, dir := range dirs {
		os.RemoveAll(dir)

	}
}

// CreateNewProject 创建新的项目
// name 项目名称
// outputDir 输出目录
// spiderName 爬虫名称
func (r *Render) CreateNewProject(project *ProjectParams) (string, error) {
	// 创建项目目录
	outputDir := project.OutputDir
	name := project.ProjectName
	moudleName := project.MoudleName
	spiderName := project.SpiderName
	dirs := make([]string, 0)
	if project.IsNew {
		outputDir = filepath.Join(outputDir, name)

	}
	spidersDir := filepath.Join(outputDir, "spiders")
	pipelinesDir := filepath.Join(outputDir, "pipelines")
	middlewaresDir := filepath.Join(outputDir, "middlewares")
	dirs = append(dirs, outputDir, spidersDir, pipelinesDir, middlewaresDir)
	caser := cases.Title(language.English)
	spiderName = caser.String(spiderName)
	_, err := r.CreateNewSpider(spiderName, spidersDir, "spider")
	defer func() {
		if err != nil {
			cleanDir(dirs)
		}
		if p := recover(); p != nil {
			err = p.(error)
			cleanDir(dirs)
		}
	}()

	if err != nil {
		return "", err
	}
	_, err = r.CreateNewPipeline(NamedToCamelCase(name,true), pipelinesDir, "pipeline")
	if err != nil {
		return "", err
	}
	_, err = r.CreateNewMiddleware(NamedToCamelCase(name,true), middlewaresDir, "middleware")
	if err != nil {
		return outputDir, err
	}
	_, err = r.createMainFile(spiderName, outputDir, moudleName,name)
	if err != nil {
		return outputDir, err

	}
	_, err = r.createSettings(project.LogDir, project.LogLevel, outputDir)
	err = r.GOModInit(outputDir, moudleName)
	return "", nil
}
func (r *Render) GOModInit(dir, moduleName string) error {
	// 更改当前工作目录
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("changing directory failed: %w", err)
	}

	// 执行 `go mod init`
	cmdInit := exec.Command("go", "mod", "init", moduleName)
	cmdInit.Stdout = os.Stdout
	cmdInit.Stderr = os.Stderr
	err = cmdInit.Run()
	if err != nil {
		return fmt.Errorf("`go mod init` failed: %w", err)
	}

	// 执行 `go mod tidy`
	cmdTidy := exec.Command("go", "mod", "tidy")
	cmdTidy.Stdout = os.Stdout
	cmdTidy.Stderr = os.Stderr
	err = cmdTidy.Run()
	if err != nil {
		return fmt.Errorf("`go mod tidy` failed: %w", err)
	}

	return nil

}
func (r *Render) extractFileType(fileName string) string {
	// 去掉最后的 ".tpl" 后缀
	if strings.HasSuffix(fileName, ".tpl") {
		return filepath.Ext(fileName[:len(fileName)-len(".tpl")])
	}
	return filepath.Ext(fileName)
}

// RenderTemplate 渲染模板
func (r *Render) RenderTemplate(fileName string, outputDir string, tpl string, params interface{}) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	tmpl, err := template.ParseFiles(filepath.Join(currentDir, "../templates", tpl))
	if err != nil {
		panic(err)
	}
	// 确保输出目录存在
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	lastFolder := filepath.Base(outputDir)
	paramsMap, err := StructToMap(params)
	if err != nil {
		panic("params is not map[string]interface{}")

	}
	paramsMap["PackageName"] = lastFolder
	// 创建输出文件
	suffix := r.extractFileType(tpl)
	outputFilePath := filepath.Join(outputDir, fileName+suffix)
	file, err := os.Create(outputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	err = tmpl.Execute(file, paramsMap)
	if err != nil {
		return "", err
	}
	return outputFilePath, nil
}

// CreateNewSpider 创建新的爬虫
func (r *Render) CreateNewSpider(spiderName string, outputDir string, fileName string) (string, error) {
	params := struct {
		SpiderName string
	}{
		SpiderName: NamedToCamelCase(spiderName,true),
	}
	return r.RenderTemplate(fileName, outputDir, "spider.go.tpl", params)
}

// CreateNewPipeline 创建新的管道
func (r *Render) CreateNewPipeline(pipelineName string, outputDir string, fileName string) (string, error) {
	params := struct {
		Pipeline string
	}{
		Pipeline: NamedToCamelCase(pipelineName,true),
	}
	return r.RenderTemplate(fileName, outputDir, "pipeline.go.tpl", params)
}

// CreateNewMiddleware 创建新的中间件
func (r *Render) CreateNewMiddleware(middlerwareName string, outputDir string, fileName string) (string, error) {
	params := struct {
		Middlerware string
	}{
		Middlerware: NamedToCamelCase(middlerwareName,true),
	}
	return r.RenderTemplate(fileName, outputDir, "middlerware.go.tpl", params)
}

func (r *Render) createMainFile(spiderName string, outputDir string, moduleName string,projectName string) (string, error) {
	params := struct {
		SpiderName string
		ModuleName string
		ProjectName string
	}{
		SpiderName: spiderName,
		ModuleName: moduleName,
		ProjectName: NamedToCamelCase(projectName,true),
	}
	return r.RenderTemplate("main", outputDir, "main.go.tpl", params)
}
func (r *Render) createSettings(logDir string, logLevel string, outputDir string) (string, error) {
	params := struct {
		LogDir   string
		LogLevel string
	}{
		LogDir:   logDir,
		LogLevel: logLevel,
	}
	return r.RenderTemplate("settings", outputDir, "settings.yaml.tpl", params)
}
