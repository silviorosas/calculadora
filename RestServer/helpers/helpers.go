package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Res hola
type Res struct {
	tag string
	w   http.ResponseWriter
	e   interface{}
	d   interface{}
	m   string
}

/*NewResponse Crea un nuevo responce
Ej: res := NewResponse(tag string, w http.ResponseWriter)
Parametros:
	@Tag: Coresponde a la accion a realizar o solicitada por el request
	@w: La variable del response donde se ralizara la respuesta
*/
func NewResponse(tag string, w http.ResponseWriter) *Res {

	return &Res{tag: tag, w: w}

}

//Msg configura un mensage en el response
func (r *Res) Msg(mensaje string) *Res {
	r.m = mensaje
	return r
}

//Err configura un error en el responce
func (r *Res) Err(err interface{}) *Res {
	r.e = err
	return r
}

//Dato configura el dato del responce
func (r *Res) Dato(dato interface{}) *Res {
	r.d = dato
	return r
}

//ErrSend configura y envia el error
func (r *Res) ErrSend(err interface{}) {
	r.e = err
	r.send()
}

//DatoSend configura y envia el dato
func (r *Res) DatoSend(dato interface{}) {
	r.d = dato
	r.send()
}

//MsgSend configura un mensage en el response y lo envia
func (r *Res) MsgSend(mensaje string) {
	r.m = mensaje
	r.send()
}

//send envia toda la configuracion al request
func (r *Res) send() {
	if r.d != nil || r.e != nil || r.m != "" {
		responseMSG(r.tag, r.e, r.d, r.w, r.m)
	} else {
		colorRed := "\033[31m"
		colorReset := "\033[0m"
		fmt.Print(string(colorRed), "ATENCION ERROR EN RESPONCE: ")
		fmt.Println(string(colorReset), "Ningun parametro seteado con informacion a enviar")
	}

}

/*
ResponseMSG funcion que genera la respuesta de la API incorporando un mensaje especial
solo funcona con la variab de entrono NOTIFICCIONES en true
	@tag: Es la tarea o funcion que se solicito ej: Nuevo Usuario
	@datos: Son los datos que se utilizaron para realizar el tag solisitado
	@w: Corresponde al ResponseWriter "w http.ResponseWriter"
	@msg: Corresponde al mensaje el mismo aprecesra con 1 - !ATENCION¡ al comienzo
*/
func responseMSG(tag string, errores interface{}, datos interface{}, w http.ResponseWriter, msg string) {
	defer errorControl("shared/helpers/helpersResponse")
	if errores != nil {
		createRespnse(tag, errores, datos, w, http.StatusBadRequest, false, msg)
	} else {
		createRespnse(tag, errores, datos, w, http.StatusOK, true, msg)
	}
}

func createRespnse(tag string, errores interface{}, datos interface{}, w http.ResponseWriter, stausHead int, status bool, msg string) {

	defer errorControl("shared/helpers/helpersCreateResponse")
	response := Response{}
	response.Datos = datos
	response.Tag = tag
	response.Error = errores
	response.Status = status

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(stausHead)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("Error al decodificar json para crear response " + err.Error())
	}

}

/*
ErrorControl Controla los errores en las funciones
@pila: Corresponde a donde se esta llamando el controlador de errores. Generalmente en el nombre de la funcion contenedora
@w: El ResponseWriter ( w http.ResponseWriter)
@r: El Request (r *http.Request)
*/
func errorControl(pila string) {

	panic := false // True para activar el panic

	if panic {
		reco := recover()
		if reco != nil {
			errorSend(reco, pila, nil, nil)
		}
	}
}

func errorSend(reco interface{}, pila string, w http.ResponseWriter, r *http.Request) {

	if w != nil {
		w.Header().Set("context-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode("Erro interno")
		if err != nil {
			panic("Error al decodificar jason para enviar a Discord " + err.Error())
		}
	}

}

//Response estructura de respuestas de la API
type Response struct {
	//El Msg solo se enviara si hay mensajes de los desarrolladores y
	//si estan activados de las variables de entorno
	Msg map[string]string `json:"MSG,omitempty"`
	//Er parametro Error se enviara junto con los errores si no hay
	//errores este no se envia
	Error interface{} `json:"error,omitempty"`
	//Datos solicitados o enviados
	Datos interface{} `json:"datos,omitempty"`
	//Si es True la transaccion se ralizo si es False fallo
	Status bool `json:"status" validate:"required"`
	//Accion o transaccion que se realizo
	Tag string `json:"tag" validate:"required"`
}

//TipoRuta tipo ruta
type TipoRuta func(string, endpoint, string, string)

type endpoint func(http.ResponseWriter, *http.Request)

//RoutersStruct estructura de rutas
type RoutersStruct struct {
	Router *mux.Router
}

//NewRouter Contructor de rutas
func NewRouter(rout *mux.Router) RoutersStruct {
	return RoutersStruct{Router: rout}
}

/*Ruta Crea una ruta: Las RUTAS requieresn bd y poseen seguridad
Las mismas comprenden el menu del usuario
*/
func (r *RoutersStruct) Ruta(ruta string, endpoint endpoint, method string, description string) {
	r.Router.HandleFunc(ruta, endpoint).Methods(method)
}
