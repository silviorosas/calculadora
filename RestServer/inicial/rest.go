package inicial

import (
	ep "TRIAND/RestServer/helpers"
	"net/http"
)

//Rutas declaracion de rutas
func Rutas(ruta ep.TipoRuta) {

	ruta("/ejemplo", ejempoEPSimple, "GET", "Ruta de ejemplo json simple")
	ruta("/ejemploArray", ejempoEPArray, "GET", "Ruta de ejemplo array")

}

//Endpoint de ejemplo para devolver un jason simple
func ejempoEPSimple(w http.ResponseWriter, r *http.Request) {
	res := ep.NewResponse("Ejemplo", w)

	type Usuario struct {
		Nombre   string
		Usuario  string
		Password string
		Email    string
	}

	usuario := &Usuario{
		Nombre:   "Francisco",
		Usuario:  "franz",
		Password: "123",
		Email:    "correo@prueba.com"}

	res.DatoSend(usuario)
}

//ejempoEPArray de ejemplo para devolver un jason simple
func ejempoEPArray(w http.ResponseWriter, r *http.Request) {
	res := ep.NewResponse("Ejemplo de Array", w)

	type UsuarioArray struct {
		Nombre   string
		Usuario  string
		Password string
		Email    string
	}

	var usuarioArray []UsuarioArray

	for i := 0; i < 5; i++ {

		usuario := UsuarioArray{
			Nombre:   "Francisco",
			Usuario:  "franz",
			Password: "123",
			Email:    "correo@prueba.com"}

		usuarioArray = append(usuarioArray, usuario)

	}

	res.DatoSend(usuarioArray)
}
