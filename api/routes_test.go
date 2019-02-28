package api

import (
	"testing"
)

func TestRoutes(t *testing.T) {
	router := NewRouter()
	if router.Get("mutant") == nil {
		t.Errorf("Debería existir la ruta mutant")
	}
	if router.Get("Stats") == nil {
		t.Errorf("Debería existir la ruta stats")
	}
	if router.Get("health") == nil {
		t.Errorf("Debería existir la ruta mutant")
	}
}
