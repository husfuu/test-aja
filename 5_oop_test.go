package main

import "testing"

func TestOopSepeda(t *testing.T) {
	sepeda := Sepeda{Kendaraan{"Swoosh"}, "Normal"}
	sepeda.Akselerasi()

	if sepeda.Rantai != "Perlu perbaikan" {
		t.Errorf("Expected rantai to be 'Perlu perbaikan', got %s", sepeda.Rantai)
	}
}

func TestOopMobil(t *testing.T) {
	mobil := Mobil{Kendaraan{"Vroom"}, "Penuh"}
	mobil.Akselerasi()

	if mobil.Bensin != "Kosong" {
		t.Errorf("Expected bensin to be 'Kosong', got %s", mobil.Bensin)
	}
}
