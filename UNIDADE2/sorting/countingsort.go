package sorting

//out of place e Ele tá estável
func CountingSort(v []int) []int {
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
	tamC := maior - menor + 1
	c := make([]int, tamC)
	for i := 0; i < len(v); i++ {
		c[v[i]-menor]++
	}
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}
	o := make([]int, len(v))
	for i := len(v) - 1; i >= 0; i-- {
		c[v[i]-menor]--
		o[c[v[i]-menor]] = v[i]
	}
	return o
}
