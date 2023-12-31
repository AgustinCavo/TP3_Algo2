package diccionario_test

import (
	"fmt"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN1 = []int{100, 200, 400, 800, 1600, 3200}

func TestUnElementABB(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("C", 10)
	dic.Guardar("D", 12)
	dic.Guardar("A", 12)
	require.EqualValues(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
}

func TestDiccionarioGuardarABB(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoABB(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	dic.Guardar(clave, "miau")
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDatoHopscotchABB(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	for i := 0; i < 500; i++ {
		dic.Guardar(strconv.Itoa(i), i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(strconv.Itoa(i), 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(strconv.Itoa(i)) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrarABB(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorradosABB(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestClaveVaciaABB(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloABB(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func buscarABB(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClavesABB(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burro"
	clave5 := "Castor"
	claves := []string{clave1, clave2, clave3, clave4, clave5}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)
	dic.Guardar(claves[3], nil)
	dic.Guardar(claves[4], nil)

	cs := []string{"", "", "", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, _ *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 5, cantidad)
	require.NotEqualValues(t, -1, buscarABB(cs[0], claves))
	require.NotEqualValues(t, -1, buscarABB(cs[1], claves))
	require.NotEqualValues(t, -1, buscarABB(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])

}

func TestIteradorInternoValoresABB(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})
	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoValoresConBorradosABB(t *testing.T) { ////NO PASA EL TEST, CHEQUEAR EL BORRAR
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	require.EqualValues(t, 7, dic.Borrar(clave0))
	require.EqualValues(t, 5, dic.Cantidad())

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {

		*ptrFactorial *= dato

		return true
	})

	require.EqualValues(t, 720, factorial)
}
func TestIteradorInternoValoresConBorrad(t *testing.T) { ////NO PASA EL TEST, CHEQUEAR EL BORRAR
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave0, clave0)
	dic.Guardar(clave1, clave1)
	dic.Guardar(clave2, clave2)
	dic.Guardar(clave3, clave3)
	dic.Guardar(clave4, clave4)
	dic.Guardar(clave5, clave5)

	require.EqualValues(t, clave0, dic.Borrar(clave0))

	dic.Iterar(func(_ string, dato string) bool {

		return true
	})
	require.EqualValues(t, 5, dic.Cantidad())
}
func TestIteradorInternoValoresConBorradosclavesOrdenadas(t *testing.T) { ////NO PASA EL TEST, CHEQUEAR EL BORRAR
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "A"
	clave1 := "B"
	clave2 := "C"
	clave3 := "D"
	clave4 := "E"
	clave5 := "F"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	require.EqualValues(t, 7, dic.Borrar(clave0))
	require.EqualValues(t, 5, dic.Cantidad())

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {

		*ptrFactorial *= dato

		return true
	})

	require.EqualValues(t, 720, factorial)
}
func TestIteradorInternoSumaDeTodosLosElementosABB(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno, luego se hace una suma de todos los elementos")

	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	suma := 0
	ptrsuma := &suma
	dic.Iterar(func(_ string, dato int) bool {
		*ptrsuma += dato
		return true
	})

	require.EqualValues(t, 27, suma)
}

func TestIteradorInternoSumaDeTodosLosElementosConBorradoABB(t *testing.T) { //NO PASA EL TEST, CHEQUEAR EL BORRAR
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno, luego se hace una suma de todos los elementos")

	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)
	dic.Borrar(clave0) //Chequear el borrar

	suma := 0
	ptrsuma := &suma
	dic.Iterar(func(_ string, dato int) bool {
		*ptrsuma += dato
		return true
	})

	require.EqualValues(t, 20, suma)
}

func comparacion(c1, c2 int) int {
	if c1 < c2 {
		return c1 - c2
	} else if c1 > c2 {
		return c1 + c2
	} else {
		return 0
	}
}
func TestIteradorInternoABBVacio(t *testing.T) {
	t.Log("Prueba de iterador interno, con un ABB Vacio")

	dic := TDADiccionario.CrearABB[int, int](comparacion)
	suma := 0
	ptrsuma := &suma
	dic.Iterar(func(_ int, dato int) bool {
		*ptrsuma += dato
		return true
	})

	require.EqualValues(t, 0, suma)

}
func TestIteradorInternoCorteABB(t *testing.T) {
	t.Log("Prueba de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](comparacion)

	// Inserta 'n' parejas en el hash
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, _ int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%2 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el hash
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que devuelva los valores correctos
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que borre y devuelva los valores correctos
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkABB(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN1 {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

//////////////////////////////////////PRUEBAS ITERADOR EXTERNO////////////////////

func TestIterardorDiccionarioVacioABB(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIteradorABB(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")

	claves := []int{6, 15, 1, 4, 10, 16, 8, 13, 11, 14}
	valores := []int{6, 15, 1, 4, 10, 16, 8, 13, 11, 14}
	inorder := []int{1, 4, 6, 8, 10, 11, 13, 14, 15, 16}
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	for i := 0; i < len(claves); i++ {
		claveActual, _ := iter.VerActual()
		require.EqualValues(t, inorder[i], claveActual)
		iter.Siguiente()
	}
}

func TestIteradorNoLlegaAlFinalABB(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarABB(primero, claves))
	require.NotEqualValues(t, -1, buscarABB(segundo, claves))
	require.NotEqualValues(t, -1, buscarABB(tercero, claves))
}

func TestPruebaIterarTrasBorradosABB(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el hash
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorABB(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN1 {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorABB(b, n)
			}
		})
	}
}
func Par(num int, otro int) int {
	if num%2 == 0 {
		return 1
	}
	return 0
}

// Iterador Con RANGO
func TestIterarDiccionarioVacioABBConRango(t *testing.T) {
	t.Log("Iterar con rango sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	A := "A"
	B := "B"
	desde := &A
	hasta := &B
	iter := dic.IteradorRango(desde, hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarDiccionarioABBAconRango(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	// Inserta 'n' parejas en el hash
	dic.Guardar(6, 6)
	dic.Guardar(15, 15)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(10, 10)
	dic.Guardar(16, 16)
	dic.Guardar(8, 8)
	dic.Guardar(13, 13)
	dic.Guardar(11, 11)
	dic.Guardar(14, 14)
	D := 10
	H := 14
	desde := &D
	hasta := &H
	iter := dic.IteradorRango(desde, hasta)

	require.True(t, iter.HaySiguiente())

	primero, _ := iter.VerActual()
	require.EqualValues(t, 10, primero)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	segundo, _ := iter.VerActual()

	require.NotEqualValues(t, primero, segundo)
	require.EqualValues(t, 11, segundo)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	tercero, _ := iter.VerActual()
	require.EqualValues(t, 13, tercero)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 14, cuarto)

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIterarDiccionarioABBAconRangosinDesde(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	dic.Guardar(6, 6)
	dic.Guardar(15, 15)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(10, 10)
	dic.Guardar(16, 16)
	dic.Guardar(8, 8)
	dic.Guardar(13, 13)
	dic.Guardar(11, 11)
	dic.Guardar(14, 14)

	D := 6
	H := 14
	desde := &D
	hasta := &H
	iter := dic.IteradorRango(desde, hasta)

	require.True(t, iter.HaySiguiente())

	primero, _ := iter.VerActual()
	require.EqualValues(t, 6, primero)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	segundo, _ := iter.VerActual()

	require.NotEqualValues(t, primero, segundo)
	require.EqualValues(t, 8, segundo)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	tercero, _ := iter.VerActual()
	require.EqualValues(t, 10, tercero)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 11, cuarto)
	iter.Siguiente()

	quinto, _ := iter.VerActual()
	require.EqualValues(t, 13, quinto)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	sexto, _ := iter.VerActual()
	require.EqualValues(t, 14, sexto)

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIterarDiccionarioABBAconRangosinDesde_2(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	dic.Guardar(6, 6)
	H := 14
	hasta := &H
	iter := dic.IteradorRango(nil, hasta)

	require.True(t, iter.HaySiguiente())

	primero, _ := iter.VerActual()
	require.EqualValues(t, 6, primero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIterarDiccionarioABBAconRangosinHasta_2(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	dic.Guardar(6, 6)
	H := 14
	hasta := &H
	iter := dic.IteradorRango(nil, hasta)

	require.True(t, iter.HaySiguiente())

	primero, _ := iter.VerActual()
	require.EqualValues(t, 6, primero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIterarDiccionarioABBAconRangosinDesde_3(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)
	dic.Guardar(6, 6)
	dic.Guardar(5, 5)
	D := 6
	desde := &D
	iter := dic.IteradorRango(desde, nil)

	require.True(t, iter.HaySiguiente())

	primero, _ := iter.VerActual()
	require.EqualValues(t, 6, primero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIterarDiccionarioABBAconRangosArbol7nodos(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)
	dic.Guardar(3, 3)
	dic.Guardar(1, 1)
	dic.Guardar(2, 2)
	dic.Guardar(4, 4)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)

	D := 2
	H := 5
	desde := &D
	hasta := &H
	iter := dic.IteradorRango(desde, hasta)

	require.True(t, iter.HaySiguiente())

	primero, _ := iter.VerActual()
	require.EqualValues(t, 2, primero)
	iter.Siguiente()

	segundo, _ := iter.VerActual()
	require.EqualValues(t, 3, segundo)
	iter.Siguiente()

	tercero, _ := iter.VerActual()
	require.EqualValues(t, 4, tercero)
	iter.Siguiente()

	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 5, cuarto)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIterarDiccionarioABBAconRangosSinDesdeNiHasta(t *testing.T) {
	t.Log("Se comporta como un iter normal")
	dic := TDADiccionario.CrearABB[int, int](comparacion)
	dic.Guardar(3, 3)
	dic.Guardar(1, 1)
	dic.Guardar(2, 2)
	dic.Guardar(4, 4)

	suma := 0
	ptrsuma := &suma
	dic.IterarRango(nil, nil, func(_, dato int) bool {
		*ptrsuma += dato
		return true
	})

	require.EqualValues(t, 10, suma)
}

func TestIteradorInternoDiccionarioABBAconRangosinDesde(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	dic.Guardar(3, 3)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)

	H := 10
	hasta := &H

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.IterarRango(nil, hasta, func(c int, _ int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%2 != 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}
func TestIteradorInternoDiccionarioABBAconRangosinHasta(t *testing.T) {
	t.Log("Guardamos valores en un Diccionario, e iteramos en un rango válido, se debe iterar correctamente")
	dic := TDADiccionario.CrearABB[int, int](comparacion)

	dic.Guardar(3, 3)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)
	D := 5

	desde := &D

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.IterarRango(desde, nil, func(c int, _ int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%2 != 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

//*/
