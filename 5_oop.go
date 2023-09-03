package main

import "fmt"

type Kendaraan struct {
	Suara string
}

func (k *Kendaraan) Akselerasi() string {
	res := fmt.Sprintf("Menghasilkan suara: %s", k.Suara)
	return res
}

type Sepeda struct {
	Kendaraan
	Rantai string
}

func (s *Sepeda) Akselerasi() string {
	res := fmt.Sprintf("Menghasilkan suara: %s", s.Suara)
	s.Rantai = "Perlu perbaikan"
	return res
}

type Mobil struct {
	Kendaraan
	Bensin string
}

func (m *Mobil) Akselerasi() string {
	res := fmt.Sprintf("Menghasilkan suara: %s", m.Suara)
	m.Bensin = "Kosong"
	return res
}
