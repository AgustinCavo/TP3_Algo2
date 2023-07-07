package grafo

import (
	"fmt"
	"math/rand"
	TDADiccionario "tdas/diccionario"
	"time"
)

type vertice[K comparable, T any] struct {
	dato    K
	aristas TDADiccionario.Diccionario[K, T]
}

type grafo[K comparable, T any] struct {
	vertices TDADiccionario.Diccionario[K, vertice[K, T]]
	dirigido bool
}

func CrearGrafo[K comparable, T any](dirigido bool) Grafo[K, T] {
	g := new(grafo[K, T])
	g.vertices = TDADiccionario.CrearHash[K, vertice[K, T]]()
	g.dirigido = dirigido

	return g
}
func (g *grafo[K, T]) AgregarVertice(nombreVertice K) {

	if !g.vertices.Pertenece(nombreVertice) {
		nuevoVertice := new(vertice[K, T])
		nuevoVertice.dato = nombreVertice
		nuevoVertice.aristas = TDADiccionario.CrearHash[K, T]()

		g.vertices.Guardar(nombreVertice, *nuevoVertice)
	}
}
func (g *grafo[K, T]) AgregarArista(inicio, destino K, peso T) {
	if g.dirigido {
		g.vertices.Obtener(inicio).aristas.Guardar(destino, peso)
	} else {
		g.vertices.Obtener(inicio).aristas.Guardar(destino, peso)
		g.vertices.Obtener(destino).aristas.Guardar(inicio, peso)
	}

}
func (g *grafo[K, T]) CantidadVertices() int {
	return g.vertices.Cantidad()
}
func (g *grafo[K, T]) EstanUnidosEstosVertices(inicio, destino K) bool {
	if g.vertices.Pertenece(inicio) && g.vertices.Pertenece(destino) {
		if g.vertices.Obtener(inicio).aristas.Pertenece(destino) {
			return true
		}
	}
	return false
}
func (g *grafo[K, T]) ExisteEsteVertice(vertice K) bool {
	return g.vertices.Pertenece(vertice)
}
func (g *grafo[K, T]) ObtenerPeso(inicio, destino K) T {
	return g.vertices.Obtener(inicio).aristas.Obtener(destino)
}

func (g *grafo[K, T]) ObtenerAdyacentes(vertice K) []K {
	infoVertice := g.vertices.Obtener(vertice)
	aristas := make([]K, infoVertice.aristas.Cantidad())
	iterAristas := infoVertice.aristas.Iterador()
	i := 0
	for iterAristas.HaySiguiente() {
		verticeAdyacente, _ := iterAristas.VerActual()
		aristas[i] = verticeAdyacente
		i++
		iterAristas.Siguiente()
	}
	return aristas
}
func (g *grafo[K, T]) ObtenerVertices() []K {
	iter := g.vertices.Iterador()
	vertices := make([]K, g.vertices.Cantidad())
	i := 0
	for iter.HaySiguiente() {
		vertice, _ := iter.VerActual()
		vertices[i] = vertice
		i++
		iter.Siguiente()
	}
	return vertices
}
func (g *grafo[K, T]) RemoverArista(inicio, destino K) {
	g.vertices.Obtener(inicio).aristas.Borrar(destino)
}

func (g *grafo[K, T]) RemoverVertice(eliminado K) {
	iter := g.vertices.Iterador()

	if g.vertices.Pertenece(eliminado) {
		for iter.HaySiguiente() {

			_, datos := iter.VerActual()

			if datos.aristas.Pertenece(eliminado) {
				datos.aristas.Borrar(eliminado)
			}
			iter.Siguiente()
		}
		g.vertices.Borrar(eliminado)
	} else {
		fmt.Print("No pertenece al grafo")
	}
}

func (g *grafo[K, T]) VerticeAleatorio() K {

	rand.Seed(time.Now().UnixNano())

	min := 0
	max := len(g.ObtenerVertices()) - 1

	// Generar un número aleatorio dentro del rango
	randomNumber := rand.Intn(max-min+1) + min

	return g.ObtenerVertices()[randomNumber]

}
