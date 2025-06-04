package sorting

//in place
func InsertionSort(v []int){
	for i := 1; i < len(v); i++{
		temp := v[i]
		j := i
		for ;j >= 1 && v[j-1] > temp;j--{
			v[j] = v[j-1]
		}
		v[j] = temp
	}
}
