package cola_prioridad_test

import (
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	cmpFunc := func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0
	}
	t.Log("Hacemos pruebas con lista vacia")
	heap := TDAHeap.CrearHeap(cmpFunc)
	require.True(t, heap.EstaVacia())

	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestInsertarPrimero(t *testing.T) {
	t.Log("Hacemos pruebas insertando algunos elementos")
	cmpFunc := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}
	t.Log("Hacemos pruebas con lista vacia")
	heap := TDAHeap.CrearHeap(cmpFunc)
	heap.Encolar(2)
	require.Equal(t, 2, heap.VerMax())
	heap.Encolar(1)
	require.Equal(t, 2, heap.VerMax())
	heap.Encolar(15)
	require.Equal(t, 15, heap.VerMax())
	heap.Encolar(5)
	heap.Encolar(14)
	heap.Encolar(23)
	require.Equal(t, 23, heap.VerMax())
	require.Equal(t, 6, heap.Cantidad())
}

func TestBorrar(t *testing.T) {
	t.Log("Hacemos pruebas Borrando algunos elementos")
	cmpFunc := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}
	heap := TDAHeap.CrearHeap(cmpFunc)
	heap.Encolar(9)
	heap.Encolar(3)
	heap.Encolar(8)
	heap.Encolar(5)
	require.Equal(t, 9, heap.VerMax())
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 9, heap.Desencolar())
	require.Equal(t, 3, heap.Cantidad())
	heap.Desencolar()
	heap.Desencolar()
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.Equal(t, true, heap.EstaVacia())
}

func TestHeapify(t *testing.T) {
	t.Log("Hacemos pruebas Heapify")
	arr := []int{5, 6, 9, 10}
	heap := TDAHeap.CrearHeapArr(arr[:], func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})

	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 9, heap.VerMax())
	require.Equal(t, 9, heap.Desencolar())
	require.Equal(t, 6, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())

	arr2 := []int{5, 3, 1, 2}
	heap2 := TDAHeap.CrearHeapArr(arr2[:], func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	require.Equal(t, 5, heap2.VerMax())
}

func TestHeapsort(t *testing.T) {
	t.Log("Hacemos pruebas HeapSort")
	arr := []int{5, 3, 1, 2, 4, 8, 9, 6, 7}

	TDAHeap.HeapSort(arr[:], func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	for i := 1; i < len(arr)-2; i++ {
		require.Equal(t, i, arr[i-1])
	}

}
func TestRedim(t *testing.T) {
	cmpFunc := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}
	t.Log("Hacemos pruebas con lista vacia")
	heap := TDAHeap.CrearHeap(cmpFunc)

	for i := 1; i <= 1000000; i++ {
		heap.Encolar(i)
	}

	for i := 1000000; i > 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	require.Equal(t, 0, heap.Cantidad())
}
