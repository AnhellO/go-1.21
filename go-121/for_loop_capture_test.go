package go121

import "testing"

// Esta prueba tiene como objetivo comprobar que todos los casos de prueba son pares (¡y no lo son!),
// pero pasa sin GOEXPERIMENT=loopvar. El problema es que t.Parallel detiene el cierre y deja que
// el bucle continúe, y luego ejecuta todos los cierres en paralelo cuando TestAllEvenBuggy retorna.
// Para cuando se ejecuta la instrucción if en el cierre, el ciclo finaliza y v tiene su valor de
// iteración final = 6. Las cuatro subpruebas ahora continúan ejecutándose en paralelo y todas
// verifican que 6 sea par, en lugar de verificar cada uno de las casos de prueba.
func TestAllEvenBuggy(t *testing.T) {
	testCases := []int{1, 2, 4, 6}
	for _, v := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			if v%2 != 0 {
				t.Fatal("odd v", v)
			}
		})
	}
}
