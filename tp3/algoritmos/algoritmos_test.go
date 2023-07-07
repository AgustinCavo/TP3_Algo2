package algoritmos

import (
	TDADiccionario "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBFS(t *testing.T) {
	t.Log("Comprueba que el recorrido bfs que vimos en clase tenga los mismos resultados de distancia")
	grafo := TDAGrafo.CrearGrafo[string, float64](false)
	padres := TDADiccionario.CrearHash[string, string]()
	distancia := TDADiccionario.CrearHash[string, float64]()

	grafo.AgregarVertice("A")
	grafo.AgregarVertice("B")
	grafo.AgregarVertice("C")
	grafo.AgregarVertice("D")
	grafo.AgregarVertice("E")
	grafo.AgregarVertice("F")
	grafo.AgregarVertice("G")
	grafo.AgregarVertice("H")
	grafo.AgregarVertice("J")

	require.EqualValues(t, 9, grafo.CantidadVertices())

	grafo.AgregarArista("A", "B", 1)
	grafo.AgregarArista("A", "C", 1)

	grafo.AgregarArista("B", "D", 1)
	grafo.AgregarArista("B", "C", 1)
	grafo.AgregarArista("B", "E", 1)

	grafo.AgregarArista("C", "J", 1)
	grafo.AgregarArista("C", "E", 1)

	grafo.AgregarArista("D", "G", 1)

	grafo.AgregarArista("E", "F", 1)

	grafo.AgregarArista("F", "G", 1)
	grafo.AgregarArista("F", "H", 1)

	grafo.AgregarArista("G", "H", 1)

	require.EqualValues(t, true, grafo.EstanUnidosEstosVertices("A", "B"))
	require.EqualValues(t, true, grafo.EstanUnidosEstosVertices("B", "A"))

	padres, distancia = BFSCM(grafo, "A", "")
	require.EqualValues(t, 0.00, distancia.Obtener("A"))
	require.EqualValues(t, 1.00, distancia.Obtener("B"))
	require.EqualValues(t, 1.00, distancia.Obtener("C"))
	require.EqualValues(t, 2.00, distancia.Obtener("D"))
	require.EqualValues(t, 2.00, distancia.Obtener("E"))
	require.EqualValues(t, 3.00, distancia.Obtener("F"))
	require.EqualValues(t, 3.00, distancia.Obtener("G"))
	require.EqualValues(t, 4.00, distancia.Obtener("H"))
	require.EqualValues(t, 2.00, distancia.Obtener("J"))
	require.EqualValues(t, "None", padres.Obtener("A"))
	require.EqualValues(t, "A", padres.Obtener("B"))
	require.EqualValues(t, "A", padres.Obtener("C"))
	require.EqualValues(t, "B", padres.Obtener("D"))
}
func TestDijsktra(t *testing.T) {
	t.Log("Comprueba que el recorrido minimo echo con dijsktra que vimos en clase tenga los mismos resultados de distancia")
	grafo := TDAGrafo.CrearGrafo[string, float64](false)
	padres := TDADiccionario.CrearHash[string, string]()
	distancia := TDADiccionario.CrearHash[string, float64]()

	grafo.AgregarVertice("A")
	grafo.AgregarVertice("B")
	grafo.AgregarVertice("C")
	grafo.AgregarVertice("D")
	grafo.AgregarVertice("E")
	grafo.AgregarVertice("F")

	grafo.AgregarArista("A", "B", 3)
	grafo.AgregarArista("A", "C", 5)

	grafo.AgregarArista("B", "D", 1)
	grafo.AgregarArista("B", "E", 4)

	grafo.AgregarArista("C", "E", 1)
	grafo.AgregarArista("C", "F", 4)

	grafo.AgregarArista("E", "D", 6)
	grafo.AgregarArista("E", "F", 2)

	grafo.AgregarArista("F", "D", 7)

	padres, distancia = Dikjstra(grafo, "A", "")

	require.EqualValues(t, 0.00, distancia.Obtener("A"))
	require.EqualValues(t, 3.00, distancia.Obtener("B"))
	require.EqualValues(t, 5.00, distancia.Obtener("C"))
	require.EqualValues(t, 4.00, distancia.Obtener("D"))
	require.EqualValues(t, 6.00, distancia.Obtener("E"))
	require.EqualValues(t, 8.00, distancia.Obtener("F"))
	require.EqualValues(t, "None", padres.Obtener("A"))
	require.EqualValues(t, "A", padres.Obtener("B"))
	require.EqualValues(t, "A", padres.Obtener("C"))
	require.EqualValues(t, "B", padres.Obtener("D"))
	require.EqualValues(t, "C", padres.Obtener("E"))
	require.EqualValues(t, "E", padres.Obtener("F"))
}

func TestMST(t *testing.T) {
	grafo := TDAGrafo.CrearGrafo[string, float64](false)
	vertices := TDADiccionario.CrearHash[string, string]()

	grafo.AgregarVertice("A")
	grafo.AgregarVertice("B")
	grafo.AgregarVertice("C")
	grafo.AgregarVertice("D")
	grafo.AgregarVertice("E")
	grafo.AgregarVertice("F")
	grafo.AgregarVertice("G")

	for _, v := range grafo.ObtenerVertices() {
		vertices.Guardar(v, v)
	}

	grafo.AgregarArista("A", "B", 4)
	grafo.AgregarArista("A", "E", 3)
	grafo.AgregarArista("A", "F", 5)

	grafo.AgregarArista("B", "C", 7)
	grafo.AgregarArista("B", "G", 4)
	grafo.AgregarArista("B", "D", 3)

	grafo.AgregarArista("C", "D", 2)

	grafo.AgregarArista("D", "E", 4)
	grafo.AgregarArista("D", "G", 8)
	require.EqualValues(t, true, grafo.EstanUnidosEstosVertices("F", "A"))

	arbol := MSTPrim(grafo)

	for _, v := range arbol.ObtenerVertices() {
		require.EqualValues(t, true, vertices.Pertenece(v))
	}

	resultado := ContarAristas(arbol, false)

	require.EqualValues(t, 21, resultado)

}
func TestMSTAristas(t *testing.T) {
	grafo := TDAGrafo.CrearGrafo[string, float64](false)
	vertices := TDADiccionario.CrearHash[string, string]()

	grafo.AgregarVertice("A")
	grafo.AgregarVertice("B")
	grafo.AgregarVertice("C")
	grafo.AgregarVertice("D")
	grafo.AgregarVertice("E")
	grafo.AgregarVertice("F")
	grafo.AgregarVertice("G")

	for _, v := range grafo.ObtenerVertices() {
		vertices.Guardar(v, v)
	}

	grafo.AgregarArista("A", "B", 4)
	grafo.AgregarArista("A", "E", 3)
	grafo.AgregarArista("A", "F", 5)

	grafo.AgregarArista("B", "C", 7)
	grafo.AgregarArista("B", "G", 4)
	grafo.AgregarArista("B", "D", 3)

	grafo.AgregarArista("C", "D", 2)

	grafo.AgregarArista("D", "E", 4)
	grafo.AgregarArista("D", "G", 8)
	require.EqualValues(t, true, grafo.EstanUnidosEstosVertices("F", "A"))

	arbol := MSTPrim(grafo)

	for _, v := range arbol.ObtenerVertices() {
		for _, w := range arbol.ObtenerAdyacentes(v) {

			require.EqualValues(t, true, grafo.EstanUnidosEstosVertices(v, w))

		}
	}
}
