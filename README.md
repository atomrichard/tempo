# TEMPO
TEMPO is a lightweight template parser for Go templates

* template folder name (with path which relative to the project folder)
* template file anem (with path which relative to the template folder)
* data for the template file

### Exapmle
```golang
func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, tempo.GetTemplate("template", "example.html", nil))
}
```
