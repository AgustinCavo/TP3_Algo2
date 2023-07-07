package main

import (
	"bufio"
	TDASistema "flybondi/sistema"
	"fmt"
	"os"
	"strings"
)

const (
	CAMINOMAS       string = "camino_mas"
	CAMINOESCALAS   string = "camino_escalas"
	CENTRALIDAD     string = "centralidad"
	NUVEVAAEROLINEA string = "nueva_aerolinea"
	ITINERARIO      string = "itinerario"
	EXPORTARKML     string = "exportar_kml"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("cantidad incorrecta parametros")
		os.Exit(1)
	}

	archivoListaPath := os.Args[1]
	archivoPadronPath := os.Args[2]

	sistema := TDASistema.CrearSistemaImplementacion()
	sistema.InicializarSistema(archivoListaPath, archivoPadronPath)

	loopEjecucion(sistema)
}

func loopEjecucion(sistema TDASistema.Sistema) {

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {

		comando, parametros := parsearComando(s.Text())

		switch comando {
		case CAMINOMAS:
			if len(parametros) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando camino_mas")
			} else {
				fmt.Println(sistema.CaminoMas(parametros[0]))
			}
		case CAMINOESCALAS:
			if len(parametros) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando camino_escalas")
			} else {
				parametros[0] = "escala," + parametros[0]
				fmt.Println(sistema.CaminoEscalas(parametros[0]))
			}
		case CENTRALIDAD:
			if len(parametros) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando centralidad")
			} else {
				fmt.Println(sistema.Centralidad(parametros[0]))
			}
		case NUVEVAAEROLINEA:
			if len(parametros) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando nueva_aerolinea")
			} else {
				fmt.Println(sistema.NuevaAerolinea(parametros[0]))
			}
		case ITINERARIO:
			//fmt.Println(sistema.SiguienteVuelo(parametros))
		case EXPORTARKML:
			if len(parametros) != 2 {
				fmt.Fprintln(os.Stderr, "Error en comando exportar_kml")
			} else {
				//	fmt.Println(sistema.Borrar(parametros))
			}
		default:
			fmt.Println("Error en comando")
		}
	}
}

func parsearComando(linea string) (string, []string) {
	inputs := strings.SplitN(linea, " ", 2)
	return inputs[0], inputs[1:]
}
