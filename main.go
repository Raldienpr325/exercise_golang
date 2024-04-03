package main

import (
	"fmt"
)

type Roda interface {
	Jenis() string
}

type BanKaret struct{}

func (b BanKaret) Jenis() string {
	return "ban karet"
}

type BanKayu struct{}

func (b BanKayu) Jenis() string {
	return "ban kayu"
}

type BanBesi struct{}

func (b BanBesi) Jenis() string {
	return "ban besi"
}

type Pintu struct {
	SuaraKetuk string
	SuaraBuka  string
}

type Mobil struct {
	RodaDepan    Roda
	RodaBelakang Roda
	PintuKanan   Pintu
	PintuKiri    Pintu
}

func (m Mobil) GantiRodaDepan(roda Roda) {
	m.RodaDepan = roda
}

func (m Mobil) GantiRodaBelakang(roda Roda) {
	m.RodaBelakang = roda
}

func main() {
	mobil := Mobil{
		RodaDepan:    BanKaret{},
		RodaBelakang: BanBesi{},
		PintuKanan: Pintu{
			SuaraKetuk: "Knock",
			SuaraBuka:  "beep",
		},
		PintuKiri: Pintu{
			SuaraKetuk: "beep",
			SuaraBuka:  "Knock",
		},
	}

	// Ganti roda
	mobil.GantiRodaDepan(BanKayu{})
	mobil.GantiRodaBelakang(BanBesi{})

	// Pintu kanan
	fmt.Println("Ketuk pintu kanan:", mobil.PintuKanan.SuaraKetuk)
	fmt.Println("Buka pintu kanan:", mobil.PintuKanan.SuaraBuka)

	// Pintu kiri
	fmt.Println("Ketuk pintu kiri:", mobil.PintuKiri.SuaraKetuk)
	fmt.Println("Buka pintu kiri:", mobil.PintuKiri.SuaraBuka)
}
