package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"leek-api/app/http/requests"
	"leek-api/app/http/resources"
	"os"
	"strings"
	"time"
)

type Builder struct {
}

type generateStruct struct {
	path     string
	position string
	tmplName string
}

type routeStruct struct {
	Uri        string
	FuncName   string
	StructName string
}

type controllerStruct struct {
	Name       string
	Short      string
	StructName string
}

type modelStruct struct {
	Name       string
	Short      string
	StructName string
}

func (b Builder) Scaffold(c *gin.Context) {

	// 获取参数
	params := requests.BuilderForm{}

	// 参数校验
	if err := c.ShouldBind(&params); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	// 路由写入
	routes(params)

	// 创建控制器
	controller(params.Name)

	// 创建模型
	model(params.Name)
}

// routes 路由写入
func routes(params requests.BuilderForm) (string, bool) {
	path := "routes/api.go"

	// 打开 api.go，可读写 - 可追加内容
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		console("ERROR", "ROUTE", err.Error())
		return path, false
	}

	// 创建模板内容
	data := routeStruct{
		Uri:        "/" + strings.TrimLeft(params.Uri, "/"),
		FuncName:   lowerCamelCase(params.Uri, "/"),
		StructName: strings.Title(params.Name),
	}

	generateParams := generateStruct{
		path:     path,
		position: "ROUTE",
		tmplName: "route",
	}
	return generate(file, generateParams, data)
}

// controller 创建控制器
func controller(s string) (string, bool) {
	// 创建文件
	fileName := strings.ToLower(s)
	path := "app/http/controllers/" + fileName + ".go"
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		console("ERROR", "CONTROLLER", err.Error())
		return path, false
	}

	// 创建模板内容
	data := controllerStruct{
		Name:       s,
		Short:      strings.ToLower(string(s[0])),
		StructName: strings.Title(s),
	}

	generateParams := generateStruct{
		path:     path,
		position: "CONTROLLER",
		tmplName: "controller",
	}
	return generate(file, generateParams, data)
}

// model 创建模型
func model(s string) (string, bool) {
	folder, ok := modelFolder(s)
	if !ok {
		return folder, ok
	}

	// 创建文件
	lowerName := strings.ToLower(s)
	path := folder + "/" + lowerName + ".go"
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		console("ERROR", "MODEL", err.Error())
		return path, false
	}

	// 创建模板内容
	data := modelStruct{
		Name:       lowerName,
		Short:      string(lowerName[0]),
		StructName: strings.Title(s),
	}

	generateParams := generateStruct{
		path:     path,
		position: "MODEL",
		tmplName: "model",
	}
	if _, ok = generate(file, generateParams, data); !ok {
		return path, ok
	}

	return modelCurd(data)
}

// modelFolder 创建模型文件夹
func modelFolder(s string) (string, bool) {
	folder := "app/models/" + strings.ToLower(s)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		// 创建文件夹
		if err = os.Mkdir(folder, 0777); err != nil {
			console("ERROR", "MODEL_FOLDER", err.Error())
			return folder, false
		}
	}

	console("SUCCESS", "MODEL_FOLDER", folder)
	return folder, true
}

// modelCurd 创建模型 CURD 文件
func modelCurd(params modelStruct) (string, bool) {
	// 创建文件
	path := "app/models/" + params.Name + "/curd.go"
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		console("ERROR", "MODEL_CURD", err.Error())
		return path, false
	}

	generateParams := generateStruct{
		path:     path,
		position: "MODEL_CURD",
		tmplName: "curd",
	}
	return generate(file, generateParams, params)
}

// generate 生成文件
func generate(file *os.File, params generateStruct, data interface{}) (string, bool) {
	// 模板文件
	params.tmplName = "templates/" + params.tmplName + ".stub"
	// 替换模板内容
	tmpl, err := template.ParseFiles(params.tmplName)
	if err != nil {
		console("ERROR", params.position, params.path)
		return params.path, false
	}

	// 输入内容到文件中
	err = tmpl.Execute(file, data)
	if err != nil {
		console("ERROR", params.position, params.path)
		return params.path, false
	}

	console("SUCCESS", params.position, params.path)
	return params.path, true
}

// lowerCamelCase 小驼峰 eg: camelCase
func lowerCamelCase(s, sep string) string {
	// 每个单词首字母都转换成大写
	s = strings.Title(s)

	// 字符串中的第一个字符转换成小写
	s = strings.ToLower(string(s[0])) + s[1:]

	return strings.Replace(s, sep, "", -1)
}

// console 打印内容到控制台
func console(status, position, content string) {

	now := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("[%v] [%v] [%v] %v \n", now, status, position, content)
}
