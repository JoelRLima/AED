package sorting

// out of place e Ele tá estável
func CountingSort(v []int) []int {
	// Encontra o menor e o maior valor
	menor := v[0]
	maior := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < menor {
			menor = v[i]
		}
		if v[i] > maior {
			maior = v[i]
		}
	}

	// Cria o vetor de contagem
	tamC := maior - menor + 1
	c := make([]int, tamC)

	// Conta a frequência de cada elemento
	for i := 0; i < len(v); i++ {
		c[v[i]-menor]++
	}

	// Acumula as contagens
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	// Gera o vetor ordenado
	o := make([]int, len(v))
	for i := len(v) - 1; i >= 0; i-- {
		c[v[i]-menor]--
		o[c[v[i]-menor]] = v[i]
	}
	return o
}

//Counting Sort é um algoritmo não comparativo que funciona apenas para inteiros dentro de um intervalo
//conhecido. Ele conta quantas vezes cada valor ocorre e usa essas contagens para determinar a posição
//correta de cada elemento no vetor ordenado. Seu tempo é linear (O(n + k)),
//onde k é a amplitude dos valores, sendo muito eficiente quando k não é muito maior que n.
//Não faz comparações e é instável se não for adaptado.
