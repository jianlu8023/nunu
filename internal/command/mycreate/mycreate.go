package mycreate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/duke-git/lancet/v2/strutil"
	"github.com/jianlu8023/nunu/internal/pkg/helper"
	"github.com/jianlu8023/nunu/tpl"
	"github.com/spf13/cobra"
)

type MyCreate struct {
	ProjectName          string // projectName demo-app
	CreateType           string // request service
	FilePath             string // filepath ./src/middleware/ctrl/demo or demo
	PackageName          string // packageName demo
	StructName           string
	StructNameLowerFirst string
	StructNameFirstChar  string
	StructNameSnakeCase  string
	IsFull               bool
}

func (c *MyCreate) String() string {
	if jsonBytes, err := json.Marshal(c); err != nil {
		return ""
	} else {
		return string(jsonBytes)
	}

}

func NewMyCreate() *MyCreate {
	return &MyCreate{}
}

var CmdMyCreate = &cobra.Command{
	Use:     "mycreate [type] [handler-name]",
	Short:   "MyCreate a new handler/service/request/response/model/repository",
	Example: "nunu mycreate handler user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var (
	tplPath string
)

func init() {
	CmdMyCreateRequest.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdMyCreateService.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdMyCreateHandler.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdMyCreateModel.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdMyCreateRepository.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdMyCreateResponse.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdMyCreateAll.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
}

var CmdMyCreateService = &cobra.Command{
	Use:     "service",
	Short:   "MyCreate a new service",
	Example: "nunu mycreate service user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdMyCreateHandler = &cobra.Command{
	Use:     "handler",
	Short:   "MyCreate a new handler",
	Example: "nunu mycreate handler user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdMyCreateResponse = &cobra.Command{
	Use:     "response",
	Short:   "MyCreate a new response",
	Example: "nunu mycreate response user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdMyCreateRequest = &cobra.Command{
	Use:     "request",
	Short:   "MyCreate a new request",
	Example: "nunu mycreate request user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdMyCreateModel = &cobra.Command{
	Use:     "model",
	Short:   "MyCreate a new model",
	Example: "nunu mycreate model user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

var CmdMyCreateRepository = &cobra.Command{
	Use:     "repository",
	Short:   "MyCreate a new repository",
	Example: "nunu mycreate repository user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

var CmdMyCreateAll = &cobra.Command{
	Use:     "all",
	Short:   "MyCreate a new service & request & model & response & repository",
	Example: "nunu mycreate all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	c := NewMyCreate()
	// 此时c 什么都咩有
	c.ProjectName = helper.GetProjectName(".")
	// 此时c 有了ProjectName go.mod文件中module 的后面
	c.CreateType = cmd.Use
	// 此时c 有了CreateType request/model/service/all等等
	c.FilePath, c.PackageName = filepath.Split(args[0])
	// 此时c 有了FilePath PackageName
	// 若 传入时 指定了路经 则filePath有值
	// packageName 为路经的最后一个值
	// example nunu mycreate all user
	// filePath ""
	// packageName user
	// example nunu mycreate all ./ctrl/user/login
	// FilePath ./ctrl/user/
	// PackageName login

	c.StructName = strutil.UpperFirst(strutil.CamelCase(c.PackageName))
	// 此时c 有了StructName 首字母大写的PackageName
	c.StructNameLowerFirst = strutil.LowerFirst(c.PackageName)
	// 此时c 有了StructNameLowerFirst 首字母小写的PackageName
	c.StructNameFirstChar = string(c.StructNameLowerFirst[0])
	// 此时c 有了StructNameFirstChar 首字母小写的PackageName的第一个字母
	c.StructNameSnakeCase = strutil.SnakeCase(c.StructName)
	// 此时c 有了StructNameSnakeCase 首字母小写的PackageName的蛇形命名
	switch c.CreateType {
	case "request", "service", "handler", "response", "model", "repository":
		c.genFile()
	case "all":

		c.CreateType = "handler"
		c.genFile()

		c.CreateType = "model"
		c.genFile()

		c.CreateType = "request"
		c.genFile()

		c.CreateType = "service"
		c.genFile()

		c.CreateType = "repository"
		c.genFile()

		c.CreateType = "response"
		c.genFile()
	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}

}
func (c *MyCreate) genFile() {
	filePath := c.FilePath
	// 如果路经为空 则使用ctrl/PackageName
	// 如果路经不为空 则使用路经/PackageName
	// example : ctrl/user
	if filePath == "" {
		filePath = fmt.Sprintf("ctrl/%s", c.PackageName)
	} else {
		if strings.HasSuffix(filePath, "/") {
			filePath = strings.TrimSuffix(filePath, "/")
		}
		filePath = fmt.Sprintf("%s/%s", filePath, c.PackageName)
	}

	if !strings.HasSuffix(filePath, "/") {
		filePath = fmt.Sprintf("%s/", filePath)
	}
	f := createFile(filePath, strings.ToLower(c.CreateType)+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(c.CreateType)+".go", "already exists.")
		return
	}
	defer f.Close()
	var t *template.Template
	var err error
	if tplPath == "" {
		t, err = template.ParseFS(tpl.MyCreateTemplateFS, fmt.Sprintf("mycreate/%s.tpl", c.CreateType))
	} else {
		t, err = template.ParseFiles(path.Join(tplPath, fmt.Sprintf("%s.tpl", c.CreateType)))
	}
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, filePath+strings.ToLower(c.CreateType)+".go")

}
func createFile(dirPath string, filename string) *os.File {
	filePath := filepath.Join(dirPath, filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}
