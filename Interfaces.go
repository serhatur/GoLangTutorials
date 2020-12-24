package main

func main() {

}

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

//base struct's

type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

//Ferrari
type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}
