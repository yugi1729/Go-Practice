package main 

import "fmt" 

func main(){
	var t int
	fmt.Scanf("%d\n",&t)
	for t > 0{
		var length int 
		fmt.Println("length, arr :")
		fmt.Scanf("%d\n",&length)
		arr := make([]int,length)
		for i, _ := range arr{
			fmt.Scanf("%d",&arr[i])
		}
		var target int
		fmt.Scanf("%d",&target)
		m := make(map[int]int)
		for _, v := range arr{
			m[target - v] = v
			
		}
		for _, v := range arr{
			if val, ok := m[v]; ok{
				fmt.Println(v, val)
				return
			}	
		}	
		t--
	}

}