package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) { //Mostrando como retorna dois parametros
	return "Retorno 1", "Retorno 2" //Em GO não precisa criar um objeto para retornar dois parametros
}

func main() {
	f1()

	f2("Param1", "Param2")

	r3 := f3()
	fmt.Println(f3(), r3) //pode printar das duas formas

	r4 := f4("Param1", "Param2")
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
	fmt.Println(f5()) //também pode printar assim, as variaveis vem concatenadas.
}
