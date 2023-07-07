package cola_prioridad

type heap[T comparable] struct {
	arreglo     []T
	funcionCmp  func(T, T) int
	tamanioHeap int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.funcionCmp = funcion_cmp
	return heap
}
func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.arreglo = make([]T, len(arreglo))
	copiarArreglo(&heap.arreglo, arreglo)
	heap.tamanioHeap = len(arreglo)
	heap.funcionCmp = funcion_cmp
	heapify(heap.arreglo, funcion_cmp)

	return heap
}
func copiarArreglo[T comparable](copiar *[]T, original []T) {
	for i := 0; i < len(original); i++ {
		(*copiar)[i] = original[i]
	}
}
func heapify[T comparable](arreglo []T, funcionCmp func(T, T) int) {

	for i := ((len(arreglo) - 1) / 2); i >= 0; i-- {
		downHeap(arreglo, len(arreglo), i, funcionCmp)
	}
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)

	i := len(elementos) - 1
	for i >= 0 {

		swap(&elementos[0], &elementos[i])
		downHeap(elementos, i, 0, funcion_cmp)
		i -= 1
	}
}

func (h *heap[T]) EstaVacia() bool {
	return len(h.arreglo) == 0
}
func swap[T comparable](viejo, nuevo *T) {
	aux := *viejo
	*viejo = *nuevo
	*nuevo = aux
}

func upHeap[T comparable](arr *[]T, hijo int, funcionCmp func(T, T) int) {
	if hijo == 0 {
		return
	}
	posPadre := uint((hijo - 1) / 2)
	if funcionCmp((*arr)[int(posPadre)], (*arr)[hijo]) < 0 {
		swap(&(*arr)[int(posPadre)], &(*arr)[hijo])
		upHeap(arr, int(posPadre), funcionCmp)
	}
}
func downHeap[T comparable](arr []T, tam, padre int, funcionCmp func(T, T) int) {
	if padre >= tam-1 {
		return
	}
	hIzq := 2*padre + 1
	hDer := 2*padre + 2
	max := padre

	if hDer < (tam) {

		if funcionCmp(arr[padre], arr[hDer]) > 0 && funcionCmp(arr[padre], arr[hIzq]) > 0 {
			max = padre
		} else if funcionCmp(arr[hDer], arr[hIzq]) > 0 {
			max = hDer
		} else {
			max = hIzq
		}
	} else if hIzq < (tam) {
		if funcionCmp(arr[padre], arr[hIzq]) < 0 {
			max = hIzq
		} else {
			//caso base
			return
		}
	}

	if max != padre {
		swap(&(arr)[max], &(arr)[padre])
		downHeap(arr, tam, max, funcionCmp)
	}

}

func (h *heap[T]) Encolar(nuevo T) {
	h.arreglo = append(h.arreglo, nuevo)
	upHeap(&h.arreglo, len(h.arreglo)-1, h.funcionCmp)
	h.tamanioHeap += 1
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.arreglo[0]
}
func (h *heap[T]) redimensionar() {
	aux := make([]T, h.Cantidad())
	copiarArreglo(&aux, h.arreglo)
	h.arreglo = aux
	h.tamanioHeap = len(h.arreglo)

}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := h.arreglo[0]
	swap(&h.arreglo[len(h.arreglo)-1], &h.arreglo[0])

	h.arreglo = h.arreglo[:(len(h.arreglo) - 1)]

	if !h.EstaVacia() {
		downHeap(h.arreglo, len(h.arreglo), 0, h.funcionCmp)
	}
	if h.Cantidad() < (h.tamanioHeap / 2) {
		h.redimensionar()
	}

	return dato
}

func (h *heap[T]) Cantidad() int {
	return len(h.arreglo)
}
