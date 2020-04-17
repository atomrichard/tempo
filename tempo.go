package tempo

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
	"io/ioutil"
	"path/filepath"
	"os"
	"log"
)

// GetTemplate has 3 inputs:
// * template folder name (with path which relative to the project folder)
// * template file name (with path which relative to the template folder)
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

func visit(files *[]string, ftype string) filepath.WalkFunc {
		ftype = "." + ftype
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
				if filepath.Ext(path) == ftype {

	        *files = append(*files, path)
				}
				return nil
    }
}

// MergeFiles is for merging files in a folder (eg. js files) has 4 inputs:
// * folderPath (with path which relative to the project folder)
// * fileExtension the extension of the files which should be included (rn. it doesn't suport multiple extensions)
// * contentType for the proper response header (eg.: "text/javascript; charset=utf-8")
// * cacheAge cache max age as string
func MergeFiles(folderPath string, fileExtension string, contentType string, cacheAge string) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=" + cacheAge)

    var files []string
    root := fmt.Sprintf("./%s/", folderPath)
		err := filepath.Walk(root, visit(&files, fileExtension))
    if err != nil {
			fmt.Println(err)
    }
    for _, file := range files {
        //fmt.Println(file)

				b, err := ioutil.ReadFile(file) // just pass the file name
				if err != nil {
					fmt.Print(err)
				}
				str := string(b) // convert content to a 'string'

				fmt.Fprintf(w, str)
    }
	})
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

// ServeFiles is a dead-simple file serving helper
func ServeFiles(w http.ResponseWriter, r *http.Request) {
  if fileExists(r.URL.Path[1:]) {
    http.ServeFile(w, r, r.URL.Path[1:])
  } else {
    http.Error(w, "Not Found", 404)
  }
}
