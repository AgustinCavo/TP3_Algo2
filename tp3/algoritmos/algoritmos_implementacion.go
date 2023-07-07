package algoritmos

import (
	TDACola "tdas/cola"
	TDAColaPrioridades "tdas/cola_prioridad"
	TDADicionario "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

const (
	INFF float64 = 999999999999999999999999999999999 * 2.9
	INFI int     = 999999999999999999
)

type heapStructDikjstra struct {
	vertice   string
	distancia float64
}
type heapStructMST[K, T comparable] struct {
	verticeIni K
	verticeDes K
	distancia  T
}
type parClaveValor struct {
	vertice string
	valor   float64
}

func fCmpStructDihjstra(A, B heapStructDikjstra) int {

	if A.distancia > B.distancia {
		return -1
	} else if A.distancia < B.distancia {
		return 1
	} else if A.distancia == B.distancia {
		return 0
	}
	return 0
}

func fcmpBFS(A, B int) int {
	if A > B {
		return 1
	} else if A < B {
		return -1
	} else if A == B {
		return 0
	}
	return 0
}
func fCmpStructMST(A, B heapStructMST[string, float64]) int {
	if A.distancia > B.distancia {
		return -1
	} else if A.distancia < B.distancia {
		return 1
	} else if A.distancia == B.distancia {
		return 0
	}
	return 0
}
func Dikjstra(grafoPesadoDirigido TDAGrafo.Grafo[string, float64], inicio, fin string) (TDADicionario.Diccionario[string, string], TDADicionario.Diccionario[string, float64]) {

	padre := TDADicionario.CrearHash[string, string]()
	valores := TDADicionario.CrearHash[string, float64]()
	heapMinimos := TDAColaPrioridades.CrearHeap(fCmpStructDihjstra)
	vertices := grafoPesadoDirigido.ObtenerVertices()

	for _, v := range vertices {
		valores.Guardar(v, INFF)
	}
	valores.Guardar(inicio, 0)
	padre.Guardar(inicio, "None")

	origen := heapStructDikjstra{inicio, valores.Obtener(inicio)}
	heapMinimos.Encolar(origen)
	for !heapMinimos.EstaVacia() {
		verticeMinimo := heapMinimos.Desencolar()
		if fin != "" && verticeMinimo.vertice == fin {
			return padre, valores
		}
		for _, a := range grafoPesadoDirigido.ObtenerAdyacentes(verticeMinimo.vertice) {

			distanciaActual := valores.Obtener(verticeMinimo.vertice) + float64(grafoPesadoDirigido.ObtenerPeso(verticeMinimo.vertice, a))
			if distanciaActual < valores.Obtener(a) {
				valores.Guardar(a, distanciaActual)
				padre.Guardar(a, verticeMinimo.vertice)
				actualizar := heapStructDikjstra{a, valores.Obtener(a)}
				heapMinimos.Encolar(actualizar)
			}
		}

	}

	return padre, valores

}
func BFSCM(grafoPesado TDAGrafo.Grafo[string, float64], inicio, fin string) (TDADicionario.Diccionario[string, string], TDADicionario.Diccionario[string, float64]) {

	padre := TDADicionario.CrearHash[string, string]()
	distancia := TDADicionario.CrearHash[string, float64]()
	visitados := TDADicionario.CrearHash[string, string]()
	cola := TDACola.CrearColaEnlazada[string]()

	cola.Encolar(inicio)

	padre.Guardar(inicio, "None")
	visitados.Guardar(inicio, inicio)

	for _, v := range grafoPesado.ObtenerVertices() {
		distancia.Guardar(v, INFF)
	}
	distancia.Guardar(inicio, 0)
	for !cola.EstaVacia() {
		verticeMinimo := cola.Desencolar()
		if fin != "" && fin == verticeMinimo {
			return padre, distancia
		}
		for _, a := range grafoPesado.ObtenerAdyacentes(verticeMinimo) {
			if !visitados.Pertenece(a) {
				distanciaActual := distancia.Obtener(verticeMinimo) + 1
				if distanciaActual < distancia.Obtener(a) {
					distancia.Guardar(a, distanciaActual)
					padre.Guardar(a, verticeMinimo)
					visitados.Guardar(a, a)
					cola.Encolar(a)
				}
			}

		}

	}
	return padre, distancia
}

func MSTPrim(grafoNodirigidoPesado TDAGrafo.Grafo[string, float64]) TDAGrafo.Grafo[string, float64] {
	vertice := grafoNodirigidoPesado.VerticeAleatorio()
	visitados := TDADicionario.CrearHash[string, string]()
	visitados.Guardar(vertice, vertice)
	arbol := TDAGrafo.CrearGrafo[string, float64](false)

	heapMinimos := TDAColaPrioridades.CrearHeap(fCmpStructMST)

	for _, w := range grafoNodirigidoPesado.ObtenerAdyacentes(vertice) {
		actual := heapStructMST[string, float64]{vertice, w, grafoNodirigidoPesado.ObtenerPeso(vertice, w)}
		heapMinimos.Encolar(actual)
	}

	for _, v := range grafoNodirigidoPesado.ObtenerVertices() {
		arbol.AgregarVertice(v)
	}

	for !heapMinimos.EstaVacia() {
		minimo := heapMinimos.Desencolar()
		if visitados.Pertenece(minimo.verticeDes) {
			continue
		}
		arbol.AgregarArista(minimo.verticeIni, minimo.verticeDes, minimo.distancia)

		visitados.Guardar(minimo.verticeDes, minimo.verticeDes)

		for _, x := range grafoNodirigidoPesado.ObtenerAdyacentes(minimo.verticeDes) {
			if !visitados.Pertenece(x) {
				arista := heapStructMST[string, float64]{minimo.verticeDes, x, grafoNodirigidoPesado.ObtenerPeso(minimo.verticeDes, x)}
				heapMinimos.Encolar(arista)

			}
		}
	}

	return arbol
}

func ContarAristas[K comparable](grafo TDAGrafo.Grafo[K, float64], dirigido bool) float64 {
	var resultado float64

	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			resultado += grafo.ObtenerPeso(v, w)
		}

	}

	if dirigido {
		return resultado
	} else {
		return resultado / 2
	}
}

func OrdenarVertices(grafoPesado TDAGrafo.Grafo[string, float64], distancias TDADicionario.Diccionario[string, float64]) []string {
	verticeDistancia := make([]parClaveValor, 0)
	resultado := make([]string, 0)
	for _, v := range grafoPesado.ObtenerVertices() {
		dato := distancias.Obtener(v)
		if dato == INFF {
			continue
		}
		par := new(parClaveValor)
		par.vertice = v
		par.valor = dato
		verticeDistancia = append(verticeDistancia, *par)
	}

	verticeDistancia = mergeSort(verticeDistancia)

	for _, v := range verticeDistancia {
		resultado = append(resultado, v.vertice)
	}

	return resultado
}
func mergeSort(verticeDistancia []parClaveValor) []parClaveValor {
	if len(verticeDistancia) <= 1 {
		return verticeDistancia
	}

	medio := len(verticeDistancia) / 2
	izq := mergeSort(verticeDistancia[:medio])
	der := mergeSort(verticeDistancia[medio:])

	return merge(izq, der)
}
func merge(izq, der []parClaveValor) []parClaveValor {
	i, j := 0, 0
	final := []parClaveValor{}
	for i < len(izq) && j < len(der) {
		if izq[i].valor < der[j].valor {
			final = append(final, izq[i])
			i++
		} else {
			final = append(final, der[j])
			j++
		}
	}
	for i < len(izq) {
		final = append(final, izq[i])
		i++
	}
	for j < len(der) {
		final = append(final, der[j])
		j++
	}
	return final
}
