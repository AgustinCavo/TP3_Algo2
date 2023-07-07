package sistema

import (
	"bufio"
	ALGORITMOS "flybondi/algoritmos"
	TDAEstructuras "flybondi/estructuras"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAColaPrioridades "tdas/cola_prioridad"
	TDADicionario "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
)

const (
	RAPIDO     string  = "rapido"
	COSTO      string  = "barato"
	ESCALA     string  = "escala"
	FRECUENCIA string  = "frecuencia"
	INF        float64 = 999999999999999999
)

type parClaveValor struct {
	vertice string
	valor   float64
}

type sistemaImplementacion struct {
	grafo           TDAGrafo.Grafo[string, TDAEstructuras.Vuelos]
	infoAeropuertos TDADicionario.Diccionario[string, TDAEstructuras.Aeropuerto]
	infoCiudades    TDADicionario.Diccionario[string, TDADicionario.Diccionario[string, TDAEstructuras.Aeropuerto]]
}
type heapStruct struct {
	valor            float64
	claveAeropuerto  string
	claveAeropuertoD string
	padres           TDADicionario.Diccionario[string, string]
}
type heapStructMST[K, T comparable] struct {
	verticeIni K
	verticeDes K
	distancia  T
}

func fCmpStructMST[K, T comparable](A, B heapStructMST[string, float64]) int {
	if A.distancia < B.distancia {
		return 1
	} else if A.distancia > B.distancia {
		return -1
	}
	return 0
}
func CrearSistemaImplementacion() Sistema {
	sistema := new(sistemaImplementacion)
	sistema.grafo = TDAGrafo.CrearGrafo[string, TDAEstructuras.Vuelos](true)
	sistema.infoAeropuertos = TDADicionario.CrearHash[string, TDAEstructuras.Aeropuerto]()
	sistema.infoCiudades = TDADicionario.CrearHash[string, TDADicionario.Diccionario[string, TDAEstructuras.Aeropuerto]]()
	return sistema
}
func abrirArchivo(path string) (*os.File, error) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error al agregar archivo")
		os.Exit(1)
		defer file.Close()
	}
	return file, err
}

func parsearArchivosAeropuerto(path string, hashAeropuerto TDADicionario.Diccionario[string, TDAEstructuras.Aeropuerto], hashCiudades TDADicionario.Diccionario[string, TDADicionario.Diccionario[string, TDAEstructuras.Aeropuerto]]) {
	file, _ := abrirArchivo(path)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		datos := strings.Split(linea, ",")
		aeropuerto := TDAEstructuras.CrearAeropuerto(datos)
		hashAeropuerto.Guardar(datos[1], aeropuerto)
		if hashCiudades.Pertenece(datos[0]) {
			hashAux := hashCiudades.Obtener(datos[0])
			hashAux.Guardar(datos[1], aeropuerto)
			hashCiudades.Guardar(datos[0], hashAux)
		} else {
			hashAux := TDADicionario.CrearHash[string, TDAEstructuras.Aeropuerto]()
			hashAux.Guardar(datos[1], aeropuerto)
			hashCiudades.Guardar(datos[0], hashAux)
		}

	}
	defer file.Close()
}
func parsearArchivoVuelos(path string, hashVuelo TDADicionario.Diccionario[string, TDAEstructuras.Vuelos]) {
	file, _ := abrirArchivo(path)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		datos := strings.Split(linea, ",")
		vuelos := TDAEstructuras.CrearVuelo(datos)
		clave := datos[0] + datos[1] + datos[3]
		hashVuelo.Guardar(clave, vuelos)
	}
	defer file.Close()
}
func (s *sistemaImplementacion) cargarGrafo(infoVuelos TDADicionario.Diccionario[string, TDAEstructuras.Vuelos]) {
	iterAerpuertos := s.infoAeropuertos.Iterador()
	iterVuelos := infoVuelos.Iterador()

	for iterAerpuertos.HaySiguiente() {
		aeropuerto, _ := iterAerpuertos.VerActual()
		s.grafo.AgregarVertice(aeropuerto)
		iterAerpuertos.Siguiente()
	}
	for iterVuelos.HaySiguiente() {
		_, vuelo := iterVuelos.VerActual()
		datos := vuelo.Recorrido()
		s.grafo.AgregarArista(datos[0], datos[1], vuelo)
		iterVuelos.Siguiente()
	}
}
func (s *sistemaImplementacion) InicializarSistema(aeropuertos, vuelos string) string {

	infoVuelos := TDADicionario.CrearHash[string, TDAEstructuras.Vuelos]()

	parsearArchivosAeropuerto(aeropuertos, s.infoAeropuertos, s.infoCiudades)
	parsearArchivoVuelos(vuelos, infoVuelos)
	s.cargarGrafo(infoVuelos)

	return "OK"
}

func normalCompare(A, B heapStruct) int {
	if A.valor > B.valor {
		return -1
	} else if A.valor < B.valor {
		return 1
	} else {
		return 0
	}
}

func reconstrucion(padres TDADicionario.Diccionario[string, string], inicio, fin string) []string {
	resultados := make([]string, 0)

	pila := TDAPila.CrearPilaDinamica[string]()
	actual := fin
	pila.Apilar(fin)

	for actual != inicio {
		pila.Apilar(padres.Obtener(actual))
		actual = padres.Obtener(actual)
	}
	for !pila.EstaVacia() {
		resultados = append(resultados, pila.Desapilar())
	}

	return resultados
}
func (s *sistemaImplementacion) aeropuertoEnCiudad(ciudad string) []string {
	iterCiudad := s.infoCiudades.Obtener(ciudad).Iterador()
	var resultado []string

	for iterCiudad.HaySiguiente() {
		aeropuerto, _ := iterCiudad.VerActual()
		resultado = append(resultado, aeropuerto)
		iterCiudad.Siguiente()
	}
	return resultado
}
func crearGrafoSubArista(grafoPrincipal TDAGrafo.Grafo[string, TDAEstructuras.Vuelos], tipoPeso string) TDAGrafo.Grafo[string, float64] {
	nuevoGrafo := TDAGrafo.CrearGrafo[string, float64](true)
	for _, v := range grafoPrincipal.ObtenerVertices() {
		nuevoGrafo.AgregarVertice(v)
	}

	for _, v := range grafoPrincipal.ObtenerVertices() {
		for _, w := range grafoPrincipal.ObtenerAdyacentes(v) {
			vuelo := grafoPrincipal.ObtenerPeso(v, w)
			if tipoPeso == RAPIDO || tipoPeso == ESCALA {
				nuevoGrafo.AgregarArista(v, w, float64(vuelo.Tiempo()))
			} else if tipoPeso == COSTO {
				nuevoGrafo.AgregarArista(v, w, float64(vuelo.Costo()))
			} else if tipoPeso == FRECUENCIA {
				nuevoGrafo.AgregarArista(v, w, 1/float64(vuelo.Frecuencia()))
			}
		}
	}
	return nuevoGrafo
}
func (s *sistemaImplementacion) todosContraTodos(datos []string) []string {

	heapMinimo := TDAColaPrioridades.CrearHeap(normalCompare)
	inicio := s.aeropuertoEnCiudad(datos[1])
	fin := s.aeropuertoEnCiudad(datos[2])

	subGrafo := crearGrafoSubArista(s.grafo, datos[0])

	for _, aeroIni := range inicio {
		for _, aeroFin := range fin {

			padres := TDADicionario.CrearHash[string, string]()
			valores := TDADicionario.CrearHash[string, float64]()

			if datos[0] == RAPIDO || datos[0] == COSTO || datos[0] == FRECUENCIA {
				padres, valores = ALGORITMOS.Dikjstra(subGrafo, aeroIni, aeroFin)
			} else if datos[0] == ESCALA {
				padres, valores = ALGORITMOS.BFSCM(subGrafo, aeroIni, aeroFin)
			}

			paraHeap := heapStruct{valores.Obtener(aeroFin), aeroIni, aeroFin, padres}
			heapMinimo.Encolar(paraHeap)

		}
	}

	mejorI, mejorD, camino := heapMinimo.VerMax().claveAeropuerto, heapMinimo.VerMax().claveAeropuertoD, heapMinimo.VerMax().padres
	resultadoCM := reconstrucion(camino, mejorI, mejorD)
	return resultadoCM

}
func (s *sistemaImplementacion) CaminoMas(parametros string) string {

	datos := strings.Split(parametros, ",")
	resultadoCM := s.todosContraTodos(datos)
	salidafuncion := resultadoCM[0]
	for i := 1; i < len(resultadoCM); i++ {
		salidafuncion = salidafuncion + " -> " + resultadoCM[i]
	}

	return salidafuncion
}
func (s *sistemaImplementacion) CaminoEscalas(parametros string) string {
	datos := strings.Split(parametros, ",")
	resultadoCM := s.todosContraTodos(datos)
	salidafuncion := resultadoCM[0]
	for i := 1; i < len(resultadoCM); i++ {
		salidafuncion = salidafuncion + " -> " + resultadoCM[i]
	}

	return salidafuncion
}

func (s *sistemaImplementacion) Centralidad(parametro string) string {
	cent := TDADicionario.CrearHash[string, float64]()
	subGrafo := crearGrafoSubArista(s.grafo, FRECUENCIA)

	for _, v := range s.grafo.ObtenerVertices() {
		cent.Guardar(v, 0)
	}
	for _, v := range s.grafo.ObtenerVertices() {

		padres, distancia := ALGORITMOS.Dikjstra(subGrafo, v, "")
		centAux := TDADicionario.CrearHash[string, float64]()

		for _, w := range subGrafo.ObtenerVertices() {
			centAux.Guardar(w, 0)
		}

		verticesOrdenados := ALGORITMOS.OrdenarVertices(subGrafo, distancia)

		for _, k := range verticesOrdenados {

			padre := padres.Obtener(k)
			if padre == "None" {
				continue
			}

			valorActual := centAux.Obtener(k) + 1
			valor := centAux.Obtener(padre) + valorActual

			centAux.Guardar(padres.Obtener(k), valor)
		}

		for _, l := range subGrafo.ObtenerVertices() {
			if l == v {
				continue
			}
			cent.Guardar(l, centAux.Obtener(l)+cent.Obtener(l))
		}
	}

	return reconstrucionCentral(cent, parametro)
}
func fCmpParClaveValor(a, b parClaveValor) int {
	if a.valor > b.valor {
		return 1
	} else if a.valor > b.valor {
		return -1
	} else if a.valor > b.valor {
		return 0
	}
	return 0
}
func reconstrucionCentral(centrales TDADicionario.Diccionario[string, float64], parametro string) string {
	cantidad, _ := strconv.Atoi(parametro)
	datos := make([]parClaveValor, 0)
	iter := centrales.Iterador()

	for iter.HaySiguiente() {
		actual := new(parClaveValor)
		actual.vertice, actual.valor = iter.VerActual()
		datos = append(datos, *actual)
		iter.Siguiente()
	}

	heapMaximos := TDAColaPrioridades.CrearHeapArr(datos, fCmpParClaveValor)

	agregar := heapMaximos.Desencolar()

	resultado := agregar.vertice

	for i := cantidad - 1; i > 0; i-- {
		agregar := heapMaximos.Desencolar()

		resultado = resultado + "," + agregar.vertice
	}
	return resultado
}
func (s *sistemaImplementacion) NuevaAerolinea(parametro string) string {

	file, err := os.OpenFile(parametro, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Problema con el archivo de salida")
	}
	defer file.Close()

	subGrafo := crearGrafoSubArista(s.grafo, COSTO)
	minimoRecorridos := ALGORITMOS.MSTPrim(subGrafo)

	for _, v := range minimoRecorridos.ObtenerVertices() {
		for _, w := range minimoRecorridos.ObtenerAdyacentes(v) {

			vuelo := new(TDAEstructuras.Vuelos)
			if s.grafo.EstanUnidosEstosVertices(v, w) {

				*vuelo = s.grafo.ObtenerPeso(v, w)
			} else if s.grafo.EstanUnidosEstosVertices(w, v) {

				*vuelo = s.grafo.ObtenerPeso(w, v)
			}

			_, err = file.WriteString((*vuelo).InfoCompletaVuelo() + "\n")
			if err != nil {
				panic("Error al escribir en el archivo")
			}

		}
	}

	return "OK"
}

func (s *sistemaImplementacion) Itenerario(parametro string) string {

	return "OK"
}

func (s *sistemaImplementacion) ExportarKML(parametro string) string {
	return "OK"
}
