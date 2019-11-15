package mvc_controller

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"srencat/mvc_model"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	usr := struct {
		Name string
		Age  int
	}{"Hello", 0}

	tpls := []string{"tpl/inc/layout.tmpl", "tpl/index.tmpl", "tpl/inc/partial.tmpl"}
	rnd.FuncMap(template.FuncMap{
		"toUpper": toUpper,
	})
	err := rnd.Template(w, http.StatusOK, tpls, usr)
	if err != nil {
		log.Fatal(err)
	}
	mvc_model.Model()
}