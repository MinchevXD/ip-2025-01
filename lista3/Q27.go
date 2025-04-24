import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	k := 1.00
	m := 0.00
	s := 1.00
	for i := 0.00; i <= 20; i++ {

		s *= x * x
		k *= ((float64(i) + float64(i+1)) * (float64(i) + float64(i+2)))
		if int(i)%2 == 0 {
			m += s / k
		} else {
			m -= s / k
		}

	}
	y := math.Cos(x)
	d := (1 - m) - y
	fmt.Printf("%.4f %.4f %.4f\n", 1-m, y, d)
}
