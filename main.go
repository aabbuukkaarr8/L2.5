package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test() // тут мы присвоили err - type = *customError и value = nil - а это уже не равно nil, по этому мы проходим
	// if err != nil
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
