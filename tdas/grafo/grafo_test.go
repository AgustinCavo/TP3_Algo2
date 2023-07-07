package grafo_test

import (
	TDADiccionario "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrafoVacio(t *testing.T) {
	t.Log("Comprueba que se agregue un vertice al grafo y cantidad de vertices")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	require.EqualValues(t, 0, grafo.CantidadVertices())
	grafo.AgregarVertice(10)
	require.EqualValues(t, true, grafo.ExisteEsteVertice(10))
	require.EqualValues(t, 1, grafo.CantidadVertices())
	grafo.AgregarVertice(12)
	grafo.AgregarVertice(14)
	grafo.AgregarVertice(16)
	require.EqualValues(t, false, grafo.ExisteEsteVertice(13))
	require.EqualValues(t, 4, grafo.CantidadVertices())
}
func TestBorrarVerticeSolitario(t *testing.T) {
	t.Log("Comprueba que se agregue un vertice al grafo y se pueda eliminar")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	grafo.AgregarVertice(10)
	require.EqualValues(t, 1, grafo.CantidadVertices())
	require.EqualValues(t, true, grafo.ExisteEsteVertice(10))
	grafo.RemoverVertice(10)
	require.EqualValues(t, false, grafo.ExisteEsteVertice(10))
	require.EqualValues(t, 0, grafo.CantidadVertices())
}

func TestObtenerVertices(t *testing.T) {
	t.Log("Comprueba que se agreguen vertices y sean todos retornados")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	hashResu := TDADiccionario.CrearHash[int, int]()

	for i := 10; i < 17; i++ {
		hashResu.Guardar(i, i)
		grafo.AgregarVertice(i)
	}

	require.EqualValues(t, 7, grafo.CantidadVertices())
	require.EqualValues(t, true, grafo.ExisteEsteVertice(10))

	vertices := grafo.ObtenerVertices()

	for i := 0; i < hashResu.Cantidad(); i++ {
		require.EqualValues(t, true, hashResu.Pertenece(vertices[i]))
	}
	grafo.RemoverVertice(16)
	hashResu.Borrar(16)

	vertices = grafo.ObtenerVertices()

	for i := 0; i < hashResu.Cantidad(); i++ {
		require.EqualValues(t, true, hashResu.Pertenece(vertices[i]))
	}

}
func TestVerticesSeperados(t *testing.T) {
	t.Log("Comprueba que dos vertices sin arista no la tengan los une y comprueba de nuevo")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	grafo.AgregarVertice(10)
	require.EqualValues(t, 1, grafo.CantidadVertices())
	require.EqualValues(t, true, grafo.ExisteEsteVertice(10))
	grafo.AgregarVertice(11)
	require.EqualValues(t, true, grafo.ExisteEsteVertice(11))
	require.EqualValues(t, 2, grafo.CantidadVertices())
	require.EqualValues(t, false, grafo.EstanUnidosEstosVertices(10, 11))
	grafo.AgregarArista(10, 11, 1)
	require.EqualValues(t, true, grafo.EstanUnidosEstosVertices(10, 11))
}

type estado struct {
	costo  int
	tiempo int
}

func TestVerticesPesosAristas(t *testing.T) {
	t.Log("Comprueba que dos vertices sin arista no la tengan los une y comprueba de nuevo junto a su peso")
	grafo := TDAGrafo.CrearGrafo[int, estado](true)
	grafo.AgregarVertice(10)
	require.EqualValues(t, 1, grafo.CantidadVertices())
	require.EqualValues(t, true, grafo.ExisteEsteVertice(10))
	grafo.AgregarVertice(11)
	require.EqualValues(t, true, grafo.ExisteEsteVertice(11))
	require.EqualValues(t, 2, grafo.CantidadVertices())
	require.EqualValues(t, false, grafo.EstanUnidosEstosVertices(10, 11))
	peso1011 := estado{15, 26}
	grafo.AgregarArista(10, 11, peso1011)
	require.EqualValues(t, true, grafo.EstanUnidosEstosVertices(10, 11))
	resultado := grafo.ObtenerPeso(10, 11)
	require.EqualValues(t, 15, resultado.costo)
	require.EqualValues(t, 26, resultado.tiempo)
}
func TestAdyacentesAUnvertice(t *testing.T) {
	t.Log("Comprueba que todos los adyacentes a un vertices sean correctos")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	hashResu := TDADiccionario.CrearHash[int, int]()
	for i := 10; i < 17; i++ {
		grafo.AgregarVertice(i)
	}
	hashResu.Guardar(12, 12)
	hashResu.Guardar(14, 14)
	hashResu.Guardar(16, 16)

	grafo.AgregarArista(10, 12, 1)
	grafo.AgregarArista(10, 14, 1)
	grafo.AgregarArista(10, 16, 1)

	results := grafo.ObtenerAdyacentes(10)
	for i := 0; i < hashResu.Cantidad(); i++ {
		require.EqualValues(t, true, hashResu.Pertenece(results[i]))
	}
}
func TestRemoverAristas(t *testing.T) {
	t.Log("Comprueba que todos los adyacentes a un vertices sean correctos")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	hashResu := TDADiccionario.CrearHash[int, int]()
	for i := 10; i < 17; i++ {
		grafo.AgregarVertice(i)
	}
	hashResu.Guardar(12, 12)
	hashResu.Guardar(14, 14)
	hashResu.Guardar(16, 16)

	grafo.AgregarArista(10, 12, 1)
	grafo.AgregarArista(10, 14, 1)
	grafo.AgregarArista(10, 16, 1)

	results := grafo.ObtenerAdyacentes(10)
	for i := 0; i < hashResu.Cantidad(); i++ {
		require.EqualValues(t, true, hashResu.Pertenece(results[i]))
	}

	hashResu.Borrar(12)
	grafo.RemoverArista(10, 12)
	require.EqualValues(t, false, grafo.EstanUnidosEstosVertices(10, 12))

}
func TestVerticeAleatorio(t *testing.T) {
	t.Log("Comprueba que el vertice aleatorio pertenezca al grafo y no se repita ")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	hashResu := TDADiccionario.CrearHash[int, int]()
	for i := 10; i < 17; i++ {
		grafo.AgregarVertice(i)
		hashResu.Guardar(i, i)
	}

	results1 := grafo.VerticeAleatorio()
	require.EqualValues(t, true, hashResu.Pertenece(results1))
	results2 := grafo.VerticeAleatorio()
	require.EqualValues(t, true, hashResu.Pertenece(results2))

}
func TestVerticeAleatorioVolumen(t *testing.T) {
	t.Log("Comprueba que el vertice aleatorio pertenezca al grafo y no se repita ")
	grafo := TDAGrafo.CrearGrafo[int, int](true)
	hashResu := TDADiccionario.CrearHash[int, int]()
	for i := 10; i < 17; i++ {
		grafo.AgregarVertice(i)
		hashResu.Guardar(i, i)
	}

	for v := 0; v <= 1000; v++ {
		results := grafo.VerticeAleatorio()
		require.EqualValues(t, true, hashResu.Pertenece(results))
	}

}
