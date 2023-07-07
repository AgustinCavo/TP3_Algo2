package diccionario

import (
	TDAPila "tdas/pila"
)

type abb[K comparable, V any] struct {
	raiz        *nodo[K, V]
	cantidad    int
	funcion_cmp func(K, K) int
}

type nodo[K comparable, V any] struct {
	der *nodo[K, V]
	izq *nodo[K, V]
	par *parClaveValor[K, V]
}

type iteradorArbol[K comparable, V any] struct {
	arbol  *abb[K, V]
	actual *nodo[K, V]
	pila   TDAPila.Pila[*nodo[K, V]]
	desde  *K
	hasta  *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.funcion_cmp = funcion_cmp
	return abb
}

func crearNodo[K comparable, V any](clave K, valor V) *nodo[K, V] {
	nodo := new(nodo[K, V])
	nodo.par = crearClaveValor(clave, valor)
	return nodo
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	nodoR := crearNodo(clave, dato)

	if ab.raiz == nil {
		ab.raiz = nodoR
		ab.cantidad += 1
	} else {
		ab.guardarNodo(ab.raiz, nodoR)
	}
}

func (ab *abb[K, V]) guardarNodo(padre *nodo[K, V], aguardar *nodo[K, V]) {

	if ab.funcion_cmp(padre.par.clave, aguardar.par.clave) > 0 {
		if padre.izq == nil {
			padre.izq = aguardar
			ab.cantidad += 1
			return
		} else {
			ab.guardarNodo(padre.izq, aguardar)
		}
	} else if ab.funcion_cmp(padre.par.clave, aguardar.par.clave) < 0 {
		if padre.der == nil {
			padre.der = aguardar
			ab.cantidad += 1
			return
		} else {
			ab.guardarNodo(padre.der, aguardar)
		}
	} else {
		padre.par.dato = aguardar.par.dato
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	pertenece, _ := ab.busquedaNodo(ab.raiz, clave)
	return pertenece
}

func (ab *abb[K, V]) busquedaNodo(nodoR *nodo[K, V], clave K) (bool, *nodo[K, V]) {
	if nodoR == nil {
		return false, nil
	} else if ab.funcion_cmp(nodoR.par.clave, clave) < 0 {
		return (ab.busquedaNodo(nodoR.der, clave))
	} else if ab.funcion_cmp(nodoR.par.clave, clave) > 0 {
		return (ab.busquedaNodo(nodoR.izq, clave))
	} else {
		return true, nodoR
	}
}
func (ab *abb[K, V]) Obtener(clave K) V {
	_, nodo := ab.busquedaNodo(ab.raiz, clave)
	if nodo != nil {
		return nodo.par.dato
	} else {
		panic("La clave no pertenece al diccionario")
	}
}
func (ab *abb[K, V]) Borrar(clave K) V {
	_, nodo := ab.busquedaNodo(ab.raiz, clave)
	if nodo != nil {
		dato := nodo.par.dato
		ab.raiz = ab.borrarNodo(&ab.raiz, clave)
		ab.cantidad -= 1
		return dato
	} else {
		panic("La clave no pertenece al diccionario")
	}
}
func (ab *abb[K, V]) borrarNodo(nodo **nodo[K, V], clave K) *nodo[K, V] {
	if *nodo == nil {
		return *nodo
	} else if ab.funcion_cmp((*nodo).par.clave, clave) < 0 {
		(*nodo).der = (ab.borrarNodo(&(*nodo).der, clave))
	} else if ab.funcion_cmp((*nodo).par.clave, clave) > 0 {
		(*nodo).izq = (ab.borrarNodo(&(*nodo).izq, clave))
	} else {
		if (*nodo).izq != nil && (*nodo).der != nil {
			nodoMinimo := ab.mayorNodoArbol((*nodo).izq)
			(*nodo).par = nodoMinimo.par
			(*nodo).izq = ab.borrarNodo(&(*nodo).izq, (*nodo).par.clave)
		} else if (*nodo).der != nil {
			return (*nodo).der
		} else if (*nodo).izq != nil {
			return (*nodo).izq
		} else {
			return nil
		}
	}
	return *nodo
}

//Iterador interno

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if ab.raiz == nil {
		return
	}
	desde, hasta := ab.menorMayorArbol(ab.raiz)
	ab.IterarRango(&desde.par.clave, &hasta.par.clave, visitar)
}

func (ab *abb[K, V]) menorMayorArbol(nodo *nodo[K, V]) (*nodo[K, V], *nodo[K, V]) {
	desde := ab.menorNodoArbol(ab.raiz)
	hasta := ab.mayorNodoArbol(ab.raiz)
	return desde, hasta
}
func (ab *abb[K, V]) mayorNodoArbol(nodo *nodo[K, V]) *nodo[K, V] {
	if nodo.der == nil {
		return nodo
	} else {
		return (ab.mayorNodoArbol(nodo.der))
	}
}
func (ab *abb[K, V]) menorNodoArbol(nodo *nodo[K, V]) *nodo[K, V] {
	if nodo.izq == nil {
		return nodo
	} else {
		return (ab.menorNodoArbol(nodo.izq))
	}
}
func iterarRango[K comparable, V any](nodoR *nodo[K, V], visitar func(clave K, dato V) bool, desde *K, hasta *K, funcion_cmp func(K, K) int, corte *bool) {

	if nodoR == nil {
		return
	}
	if !*corte && funcion_cmp(nodoR.par.clave, *desde) >= 0 {
		iterarRango(nodoR.izq, visitar, desde, hasta, funcion_cmp, corte)
	}
	if !*corte && funcion_cmp(nodoR.par.clave, *desde) >= 0 && funcion_cmp(nodoR.par.clave, *hasta) <= 0 {
		if !visitar(nodoR.par.clave, nodoR.par.dato) {
			*corte = true
			return
		}
	}
	if !*corte && funcion_cmp(nodoR.par.clave, *hasta) <= 0 {
		iterarRango(nodoR.der, visitar, desde, hasta, funcion_cmp, corte)
	}
}
func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if ab.raiz == nil || desde != nil && hasta != nil && ab.funcion_cmp(*desde, *hasta) > 0 {
		return
	}

	corte := false
	yaRecorri := false

	if hasta == nil || desde == nil {
		menorNodo, mayorNodo := ab.menorMayorArbol(ab.raiz)
		menorclave, mayorclave := menorNodo.par.clave, mayorNodo.par.clave
		if desde == nil && hasta == nil {
			iterarRango(ab.raiz, visitar, &menorclave, &mayorclave, ab.funcion_cmp, &corte)
			yaRecorri = true
			return
		} else if hasta == nil {
			hasta = &mayorclave
		} else {
			desde = &menorclave
		}
	}
	if !yaRecorri {
		iterarRango(ab.raiz, visitar, desde, hasta, ab.funcion_cmp, &corte)
		return
	}
}

// Iterador externo

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorArbol[K, V])

	if ab.raiz != nil {
		menorNodo, mayorNodo := ab.menorMayorArbol(ab.raiz)
		desde := &menorNodo.par.clave
		hasta := &mayorNodo.par.clave
		return ab.IteradorRango(desde, hasta)
	} else {
		iter.pila = TDAPila.CrearPilaDinamica[*nodo[K, V]]()
		iter.actual = ab.raiz
	}
	return iter
}

func (i *iteradorArbol[K, V]) HaySiguiente() bool {
	if (i.pila).EstaVacia() {
		return false
	}
	nodo_act := i.pila.VerTope()
	if i.arbol.funcion_cmp(nodo_act.par.clave, *i.hasta) <= 0 && i.arbol.funcion_cmp(nodo_act.par.clave, *i.desde) >= 0 {
		return true
	}
	if i.arbol.funcion_cmp(nodo_act.par.clave, *i.hasta) > 0 {
		return false
	}
	return true
}

func (i *iteradorArbol[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.pila.VerTope().par.clave, i.pila.VerTope().par.dato
}

func (i *iteradorArbol[K, V]) Siguiente() {

	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.actual = i.pila.Desapilar()
	if i.actual.der != nil {
		i.actual = i.actual.der
		if i.desde == i.hasta && i.hasta == nil {
			apilacionIzq(i.actual, &i.pila)
		} else {
			i.buscarPrimeroIterRango(i.actual, &i.pila, i.desde, i.arbol.funcion_cmp)
		}
	}
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iter := new(iteradorArbol[K, V])
	iter.arbol = ab
	iter.pila = TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	iter.actual = ab.raiz
	iter.desde = desde
	iter.hasta = hasta

	if iter.actual == nil || (iter.desde == iter.hasta && iter.hasta == nil) { //Caso arbol vacio o no hay iteracion con rango
		return ab.Iterador()
	}

	if iter.desde == nil {
		menorNodo := ab.menorNodoArbol(iter.actual)
		iter.desde = &menorNodo.par.clave
	} else if iter.hasta == nil {
		mayorNodo := ab.mayorNodoArbol(iter.actual)
		iter.hasta = &mayorNodo.par.clave
	}

	iter.buscarPrimeroIterRango(iter.actual, &iter.pila, iter.desde, ab.funcion_cmp)

	return iter
}
func apilacionIzq[K comparable, V any](nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]]) {
	if nodoR == nil {
		return
	}
	(*pila).Apilar(nodoR)
	apilacionIzq(nodoR.izq, pila)
}

func (i *iteradorArbol[K, V]) buscarPrimeroIterRango(nodoR *nodo[K, V], pila *TDAPila.Pila[*nodo[K, V]], desde *K, compar func(K, K) int) {
	if i.actual == nil {
		return
	}
	if compar(nodoR.par.clave, *desde) < 0 {
		i.actual = i.actual.der
		i.buscarPrimeroIterRango(i.actual, pila, desde, compar)
	} else if compar(nodoR.par.clave, *desde) >= 0 {
		(*pila).Apilar(nodoR)
		i.actual = i.actual.izq
	}
	i.buscarPrimeroIterRango(i.actual, pila, desde, compar)
}
