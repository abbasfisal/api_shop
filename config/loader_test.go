package config

import "testing"

func TestSingleton(t *testing.T) {
	Load("config")

	first := C
	second := C

	if first != second {
		t.Errorf("config is not singleton")
	}
}

func TestCheckAppNameFromDefault(t *testing.T) {
	Load("config")

	if C.App.Name != C.App.Name {
		t.Error("App name is wrong", C.App.Name)
	}
}
