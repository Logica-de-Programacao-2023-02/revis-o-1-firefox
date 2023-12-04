package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

package main

import (
	"strings"
)

func ProcessString(s string) string {
	var result strings.Builder

	for _, char := range s {
		if strings.Contains("BCDFGHJKLMNPQRSTVWXYZ", strings.ToUpper(string(char))) {
			result.WriteString(strings.ToLower(string(char)))
		}

		if strings.Contains("BCDFGHJKLMNPQRSTVWXYZ", strings.ToUpper(string(char))) {
			result.WriteString(".")
		}

		if !strings.Contains("AEIOUaeiou", string(char)) {
			result.WriteString(string(char))
		}
	}

	return result.String()
}
