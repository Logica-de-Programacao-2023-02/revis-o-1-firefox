package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	if len(barLengths) == 0 {
		return 0, 0
	}

	countMap := make(map[int]int)

	for _, length := range barLengths {
		countMap[length]++
	}

	maxLength := 0
	maxCount := 0
	for length, count := range countMap {
		if count > maxCount || (count == maxCount && length > maxLength) {
			maxLength = length
			maxCount = count
		}
	}

	totalTowers := len(countMap)

	return maxLength, totalTowers
}
