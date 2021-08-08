package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type CommonReturnType struct {
	Status string `json:"status"`
	// 若status == success 则data内返回前端需要的数据
	// 若status == fail 则data内使用通用的错误格式
	Data interface{} `json:"data"`
}

func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	// 指定视图所在路径
	pagePath := filepath.Join("./web/tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("创建模板实例错误: %v", err)
		return
	}

	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("在模板中融合数据时发生错误: %v", err)
		return
	}

}
