package api

import "testing"

func TestNewRouter(T *testing.T) {

	if len(routes) != 3 {
		T.Error("Las rutas deberían ser 3")
	}
}
