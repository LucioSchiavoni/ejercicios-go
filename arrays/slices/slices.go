package slices

import (
	"fmt"
)

var tablas []int = []int{10, 20, 5}
var arry [10]int = [10]int{6, 24, 77, 45, 78, 23, 55}

func MuestroSlice() {
	fmt.Println(tablas)

	slice3 := arry[:4]

	fmt.Println(slice3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, capadidad %d", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	// fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))

	paises := map[string]string{
		"luxemburgo":   "tecnologia",
		"lienchestein": "paz",
		"suiza":        "oro",
		"alemania":     "futbol",
		"ibiza":        "fiesta",
		"argentina":    "crisis",
		"Uruguay":      "Precios altos",
	}

	//%s para devolver string \ %d para decimal / %t para bool
	for lugares, descripcion := range paises {
		fmt.Printf("%s tiene %s\n", lugares, descripcion)
	}
}
