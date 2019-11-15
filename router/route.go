package mvc_route

import (
	"srencat/mvc_controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func InitRouter()(*httprouter.Router){
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
		w.WriteHeader(200)
	})

	router.ServeFiles("/static/*filepath", http.Dir("./public"))
	router.GET("/", mvc_controller.Index)
	router.POST("/login", mvc_controller.Login)
	return router
}