package grafo

type Grafo[K comparable, T any] interface {

	// Agrega una Arista entre dos vertices con un peso al grafo
	AgregarArista(K, K, T)

	//Agrega un vertice al grafo
	AgregarVertice(K)

	//Devuelve un vector con todos los vertices del grafo
	ObtenerVertices() []K

	//Devuelve un vector con todos los vertices adyacentes
	ObtenerAdyacentes(K) []K

	// Devuelve la cantidad de vertices presentes en el grafo
	CantidadVertices() int

	//Remueve el vertice y sus conexiones del grafo
	RemoverVertice(K)

	//Remueva la arista entre dos de los vertices
	RemoverArista(K, K)

	//Devuelve el peso entre dos vertices
	ObtenerPeso(K, K) T

	//Devuelve si en el grafo dos vertices estan unidos por una arista en sentido inicio-fin
	EstanUnidosEstosVertices(K, K) bool

	//Devuelve si en el grafo existe el vertice
	ExisteEsteVertice(K) bool

	//Devuelve un vertice aleatorio
	VerticeAleatorio() K
}
