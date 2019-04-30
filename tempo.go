package tempo

import (
	"bytes"
	"fmt"
	"html/template"
	"path"
	"strings"
)

// GetTemplate has file path, data for the template file, path for the template directory
func GetTemplate(filePath string, data interface{}, viewPath string) string {
	var tplBuffer bytes.Buffer

	filePath = strings.Trim(filePath, "/")

	var fileFullPath string = fmt.Sprintf("%s/%s", viewPath, filePath)
	parseName := strings.Replace(strings.Trim(fileFullPath, "/"), "/", "_", -1)
	t, err := template.New(parseName).Funcs(template.FuncMap{
		"addOne": func(val int) int {
			return val + 1
		},
	}).ParseFiles("./" + fileFullPath)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(&tplBuffer, path.Base("./"+fileFullPath), data)
	if err != nil {
		panic(err)
	}

	return tplBuffer.String()
}
