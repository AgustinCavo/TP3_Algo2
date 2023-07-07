package estructuras

type Vuelos interface {
	InfoVuelos() []string
	Recorrido() []string
	Costo() int
	Tiempo() int
	Frecuencia() int
	InfoCompletaVuelo() string
}
