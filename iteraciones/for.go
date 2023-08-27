package iteraciones

import (
	"fmt"
)

func Iterar() {
	// += 5 o -=5 va a pegar saltos de a 5
	for i := 100; i > 10; i -= 5 {
		if i == 40 {
			continue
		} //se va a saltar el 40
		fmt.Println(i)
	}
}
