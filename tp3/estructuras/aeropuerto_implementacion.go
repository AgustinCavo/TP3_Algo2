package estructuras

import "strconv"

type aeropuerto struct {
	ciudad           string
	codigoAeropuerto string
	latitud          int
	longitud         int
}

func CrearAeropuerto(datos []string) Aeropuerto {

	aeropuerto := new(aeropuerto)
	aeropuerto.ciudad = datos[0]
	aeropuerto.codigoAeropuerto = datos[1]
	aeropuerto.latitud, _ = strconv.Atoi(datos[2])
	aeropuerto.longitud, _ = strconv.Atoi(datos[3])

	return aeropuerto

}

func (a *aeropuerto) InfoAeropuerto() []string {
	results := []string{a.ciudad, a.codigoAeropuerto, strconv.Itoa(a.latitud), strconv.Itoa(a.longitud)}
	return results
}
