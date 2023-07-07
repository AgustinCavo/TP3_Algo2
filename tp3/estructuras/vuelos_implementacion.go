package estructuras

import "strconv"

type vuelo struct {
	aeropuertoInicio    string
	aeropuertoDestino   string
	tiempoPromedio      int
	precio              int
	cantVuelosEntreAero int
}

func CrearVuelo(datos []string) Vuelos {

	vuelo := new(vuelo)
	vuelo.aeropuertoInicio = datos[0]
	vuelo.aeropuertoDestino = datos[1]
	vuelo.tiempoPromedio, _ = strconv.Atoi(datos[2])
	vuelo.precio, _ = strconv.Atoi(datos[3])
	vuelo.cantVuelosEntreAero, _ = strconv.Atoi(datos[4])

	return vuelo

}
func (v *vuelo) InfoVuelos() []string {
	results := []string{v.aeropuertoInicio, v.aeropuertoDestino, strconv.Itoa(v.tiempoPromedio), strconv.Itoa(v.precio), strconv.Itoa(v.cantVuelosEntreAero)}
	return results
}
func (v *vuelo) Recorrido() []string {
	results := []string{v.aeropuertoInicio, v.aeropuertoDestino}
	return results

}
func (v *vuelo) Costo() int {
	return v.precio
}
func (v *vuelo) Tiempo() int {
	return v.tiempoPromedio
}
func (v *vuelo) Frecuencia() int {
	return v.cantVuelosEntreAero
}
func (v *vuelo) InfoCompletaVuelo() string {
	return (v.aeropuertoInicio + "," + v.aeropuertoDestino + "," + strconv.Itoa(v.tiempoPromedio) + "," + strconv.Itoa(v.precio) + "," + strconv.Itoa(v.cantVuelosEntreAero))
}
