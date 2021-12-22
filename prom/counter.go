package prom

var counter = 0

func Add(step int) {
	counter += step
}

func Get() int {
	return counter
}
