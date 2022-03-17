package restserver

import (
	helpers "TRIAND/RestServer/helpers"
	"TRIAND/RestServer/inicial"
	"TRIAND/RestServer/javier"
)

//Rest importa los endpoints de las rutas
func Rest(ruotes helpers.TipoRuta) {
	inicial.Rutas(ruotes)
	javier.Rutas(ruotes)

}
