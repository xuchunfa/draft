package main

import (
	"fmt"
)

func main() {

	/*
		输出:
		false
		true
		true
		true
		false
		false
		false
		true
	*/
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = a1
	a2[0] = 99
	//fmt.Println(a1, a2)
	fmt.Println("a1[0]==a2[0]? ", a1[0] == a2[0])
	var sa = a1[:]
	sa[1] = 88
	//fmt.Println(a1, sa)
	fmt.Println("a1[1]==sa[1]? ", a1[1] == sa[1])
	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = s1
	s2[0] = 99
	//fmt.Println(s1, s2)
	fmt.Println("s1[0]==s2[0]? ", s1[0] == s2[0])
	var s3 = s2[1:4]
	s3[0] = 88
	//fmt.Println(s1, s3)
	fmt.Println("s1[1]==s3[0]? ", s1[1] == s3[0])
	var s4 = append(s2, 6) // s1=s2=[99,88,3,4,5] len=5 cap=5
	s4[2] = 77             //  s4=[99,88,77,4,5,6] len=6 cap=10
	//fmt.Println(s1, s4)
	fmt.Println("s1[2]==s4[2]? ", s1[2] == s4[2])
	var oldLen int
	oldLen = len(s1)
	//fmt.Println(s1)
	appendInt(s1, 6)
	//fmt.Println(s1)
	fmt.Println("len(s1)==oldLen+1?", len(s1) == oldLen+1) // len(s1)=5 oldLen=5
	oldLen = len(s1)
	//fmt.Println(s1, s2)
	appendIntPtr(&s2, 6) // s2=[99,88,3,4,5,6] len=6 cap=10
	//fmt.Println(s1, s2)
	fmt.Println("len(s1)==oldLen+1?", len(s1) == oldLen+1) // len(s1)=5 oldLen=5
	oldLen = len(s1)
	//fmt.Println(s1)
	appendIntPtr(&s1, 6) // s1=[99,88,3,4,5,6] len=6 cap=10 oldLen=5
	//fmt.Println(s1)
	fmt.Println("len(s1)==oldLen+1?", len(s1) == oldLen+1)
}

func appendInt(s []int, elems ...int) {
	s = append(s, elems...)
}

func appendIntPtr(ps *[]int, elems ...int) {
	*ps = append(*ps, elems...)
}
