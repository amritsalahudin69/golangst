// turunan func dari suatu struck!
package main

import "fmt"

type KaryawanSalt struct {
	name         string
	gaji         int
	potonganBPJS int
	tunjangan    int
}

func (ks KaryawanSalt) Gajibersih() int {
	return ks.gaji - ks.potonganBPJS
}

func (ks KaryawanSalt) GajiBruto() int {
	return ks.gaji + ks.tunjangan
}

func main() {
	KaryawanSalt1 := KaryawanSalt{
		name:         "Boris",
		gaji:         10000000,
		potonganBPJS: 300211,
		tunjangan:    200000,
	}
	fmt.Println(KaryawanSalt1.Gajibersih())
	fmt.Println(KaryawanSalt1.GajiBruto())
}
