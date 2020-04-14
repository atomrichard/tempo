# TEMPO
TEMPO is a lightweight toolset for handling templates and files

### GetTemplate Inputs
* template folder name (with path which relative to the project folder)
* template file name (with path which relative to the template folder)
* data for the template file

### GetTemplate Example
```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, tempo.GetTemplate("template", "example.html", nil))
}
```


### ServeFile Inputs
* No inputs required

### ServeFile Example
```golang
http.HandleFunc("/favicon.ico", tempo.ServeFiles)
```


### MergeFiles Inputs
* folderPath (with path which relative to the project folder)
* fileExtension the extension of the files which should be included (rn. it doesn't suport multiple extensions)
* contentType for the proper response header (eg.: "text/javascript; charset=utf-8")
* cacheAge cache max age as string

### MergeFiles Example
```golang
http.HandleFunc("/js.js", tempo.MergeFiles("src/js", "js", "text/javascript; charset=utf-8", "604800"))
```
