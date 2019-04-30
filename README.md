# TEMPO
TEMPO is a lightweight template parser for Go templates

### Inputs
* template folder name (with path which relative to the project folder)
* template file name (with path which relative to the template folder)
* data for the template file

### Example
```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, tempo.GetTemplate("template", "example.html", nil))
}
```
