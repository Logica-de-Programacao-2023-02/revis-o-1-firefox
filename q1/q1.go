package q1

//Em um dia quente de verão, Pete e seu amigo Billy decidiram comprar uma melancia. Eles escolheram a maior e mais
//saborosa, na opinião deles, e, em seguida, pesaram a fruta nas balanças, obtendo seu peso em quilos. Morrendo de sede,
//correram para casa com a melancia e decidiram dividi-la. No entanto, enfrentaram um problema difícil.
//
//Como grandes fãs de números pares, Pete e Billy querem dividir a melancia de tal forma que cada uma das duas partes pese
//um número par de quilos, sem que as partes necessariamente tenham pesos iguais. Mas os meninos estão extremamente
//cansados e querem começar a refeição o mais rápido possível. Portanto, precisam de ajuda para descobrir se é possível
//dividir a melancia da maneira que desejam. É importante destacar que cada um deles deve receber uma parte de peso
//positivo.
//
//A função deve retornar um valor booleano, indicando se é possível ou não dividir a melancia da forma desejada. Se o peso
//da melancia for menor ou igual a 0, a função deve retornar um erro.

import (
	"errors"
	"fmt"
)

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, errors.New("O peso da melancia deve ser maior que 0")
	}

	if weight%2 == 0 && weight > 2 {
		return true, nil
	}

	return false, nil
}

func main() {
	weight := 10

	result, err := DivideWatermelon(weight)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	if result {
		fmt.Println("É possível dividir a melancia da forma desejada.")
	} else {
		fmt.Println("Não é possível dividir a melancia da forma desejada.")
	}
}
