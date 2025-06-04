package sorting

//in place
func BubbleSort(v []int){
	for i := 0; i < len(v)-1; i++{
		trocou := false
		for j := 0; j < len(v)-i-1; j++{
			if v[j] > v[j+1]{
				v[j+1], v[j] = v[j], v[j+1]
				trocou = true
			}
		}
		if !trocou{return}
	}
}

