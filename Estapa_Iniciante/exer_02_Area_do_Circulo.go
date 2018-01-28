/*
A fórmula para calcular a área de uma circunferência é: area = π . raio. Considerando para este problema que π = 3.14159:

- Efetue o cálculo da área, elevando o valor de Raio ao quadrado e multiplicando por π.

Entrada
A entrada contém um valor de ponto flutuante (dupla precisão), no caso, a variável raio.

Saída
Apresentar a mensagem "A=" seguido pelo valor da variável area, conforme exemplo abaixo, com 4 casas após o ponto decimal.
Utilize variáveis de dupla precisão (double). Como todos os problemas, não esqueça de imprimir o fim de linha após o resultado,
caso contrário, você receberá "Presentation Error".
*/

package main

import (
	"fmt"
	"strconv"
)

const r = 3.14159

func Area(raio float64) float64 {

	var Area float64

	Area = r * (raio * raio)

	return Area
}

func main() {
	var (
		Raio float64
	)

	fmt.Scan(&Raio)

	fmt.Printf("A=%s\n", strconv.FormatFloat(Area(Raio), 'f', 4, 64))

}
