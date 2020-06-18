num := [5]int{1, 2, 3, 4, 5}
slice1 := num[:]
slice2 := num[1:4]
slice3 := num[2:]
slice4 := num[:2]
fmt.Println(num[0])
fmt.Println(slice1)
fmt.Println(slice2)
fmt.Println(slice3)
fmt.Println(slice4)
