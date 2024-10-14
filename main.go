package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* Latihan :
	harga jual := 150000.0 harga jual per unit dalam Rp
	harga beli := 100000.0 harga beli per unit dalam Rp
	biayaOperasional := 1000.0 biaya operasional per unit dalam rp
	diskon := 15.0 diskon dalam persen
	jumlahTerjual := 100 jumlah produk yg terjual


	tampilkan :
	- harga jual setelah diskon
	- total pendapatan
	- total biaya (modal yang dikeluarkan)
	- total keuntungan

	*/

	var hargaJual, hargaBeli, biayaOperasional float64 = 150000.0, 100000.0, 1000.0
	var diskon float64 = 0.15
	var jumlahTerjual int = 100

	var hargaJualDiskon float64 = hargaJual - (hargaJual * float64(diskon))

	var totalPendapatan float64 = hargaJualDiskon * float64(jumlahTerjual)

	var biayaModal float64 = hargaBeli + biayaOperasional

	var totalKeuntungan float64 = (hargaJualDiskon + biayaOperasional - hargaBeli) * float64(jumlahTerjual)

	var strhargaJualDiskon string = strconv.FormatFloat(hargaJualDiskon, 'f', 2, 64)
	fmt.Println("Harga jual diskon = Rp", strhargaJualDiskon)

	var strtotalPendapatan string = strconv.FormatFloat(totalPendapatan, 'f', 2, 64)
	fmt.Println("Total Pendapatan = Rp", strtotalPendapatan)

	var strbiayaModal string = strconv.FormatFloat(biayaModal, 'f', 2, 64)
	fmt.Println("Biaya Modal = Rp", strbiayaModal)

	var strtotalKeuntungan string = strconv.FormatFloat(totalKeuntungan, 'f', 2, 64)
	fmt.Println("Total Keuntungan = Rp", strtotalKeuntungan)
}
