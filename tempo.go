package tempo

import (
	"bytes"
	"fmt"
	"html/template"
	"path"
	"strings"
)

// GetTemplate has 3 inputs:
// * template folder name (with path which relative to the project folder)
// * template file anem (with path which relative to the template folder)
// * data for the template file
func GetTemplate(viewPath string, filePath string, data interface{}) string {
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
