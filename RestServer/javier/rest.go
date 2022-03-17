package javier

import (
	ep "TRIAND/RestServer/helpers"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Rutas declaracion de rutas
func Rutas(ruta ep.TipoRuta) {

	ruta("/javier", funcion, "GET", "Ruta de javier")
	ruta("/javierEnvio", MyEndPoint, "POST", "Ruta de javier para prueba de envio de datos")
	ruta("/calcular", calcular, "POST", "Ruta para calculo")
	ruta("/calcular/{num1}/{num2}/{op}", calcularsin, "POST", "Ruta para calculo")

}

func calcularsin(w http.ResponseWriter, r *http.Request) {
	res := ep.NewResponse("Myendpoint", w)

	datos := mux.Vars(r)

	num1, _ := strconv.ParseFloat(datos["num1"], 64)
	num2, _ := strconv.ParseFloat(datos["num2"], 64)
	op := datos["op"]

	var resultado float64

	switch op {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2

	case "*":
		resultado = num1 * num2
	case "/":
		resultado = num1 / num2
	case "c":
		resultado = num1 - num1

	default:
		res.ErrSend("El operador no existe")
		return
	}

	res.DatoSend(resultado)

}

func calcular(w http.ResponseWriter, r *http.Request) {

	res := ep.NewResponse("Myendpoint", w)
	var datos Datos
	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	var resultado float64

	switch datos.Operador {
	case "+":
		resultado = datos.Numero1 + datos.Numero2
	case "-":
		resultado = datos.Numero1 - datos.Numero2
	case "*":
		resultado = datos.Numero1 * datos.Numero2
	case "/":
		resultado = datos.Numero1 / datos.Numero2
	case "c":
		resultado = datos.Numero1 - datos.Numero1

	default:
		res.ErrSend("El operador no existe")
		return
	}

	res.DatoSend(resultado)

}

//Datos Modelo de datos para calculadors
type Datos struct {
	Numero1  float64 `json:"num1"`
	Numero2  float64 `json:"num2"`
	Operador string  `json:"op"`
}

//Endpoint de ejemplo para devolver un jason simple
func funcion(w http.ResponseWriter, r *http.Request) {
	res := ep.NewResponse("Prueba de javier", w)

	res.Dato("Este es el dato").Msg("El nombre debe tener 10 letras").ErrSend("Nombre invalido")

}

//MyEndPoint funciond de end de prueba
func MyEndPoint(w http.ResponseWriter, r *http.Request) {

	res := ep.NewResponse("Myendpoint", w)
	var usuario Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	res.DatoSend(usuario)

}

//Usuario Modelo de usuario
type Usuario struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	NumeroDoc int    `json:"numero_doc"`
}
