package restserver

import (
	helpers "TRIAND/RestServer/helpers"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//StartAPI Activa el servidor HTTP APIS Rest
func StartAPI() {

	Port := "3000"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	//colorReset := "\033[0m"

	rout := mux.NewRouter().StrictSlash(false)
	rutas := helpers.NewRouter(rout)
	Rest(rutas.Ruta)
	routes := cors.AllowAll().Handler(rout)

	http.DefaultClient.Timeout = time.Minute * time.Duration(1)
	fmt.Println(string(colorGreen), "Ejecutando API TRIAND-Pruebas: URL = http://localhost:"+Port+"/")

	err := http.ListenAndServe(":"+Port, routes)
	if err != nil {
		fmt.Println(string(colorYellow), "Error en la ejecucion de APIS puerto:"+Port)
		fmt.Println(string(colorRed), err.Error())
	}

}
