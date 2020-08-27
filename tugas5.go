package main

import "fmt"
import "math/rand"


func main() {
  var n1, n2, n3 pelajar;
  n1.nama = "Ahmad"
  n2.nama = "Abi"
  n3.nama = "Zaka"

  n1.tampil_mahasiswa()
  n2.tampil_mahasiswa()
  n3.tampil_mahasiswa()
}

func (s pelajar) tampil_mahasiswa() {
    s.umur = rand.Intn(100)
    fmt.Printf("%s %d", s.nama, s.umur)
    fmt.Println()
}

type pelajar struct {
  nama string
  umur int
}
