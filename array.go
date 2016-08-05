/*************************************************************************
 * File Name: array.go
 * Description:
 * Author: raojia
 * Mail: raojia@le.com
 * Created Time: 2016-08-05 18时40分24秒
 * Last Modified: 2016-08-05 18时42分42秒
************************************************************************/
package main

import "fmt"

func PrintArr() {

	var a [3]int = [3]int{0, 1, 2}
	var b [3]int = [3]int{}
	var c [3]int
	c = [3]int{0, 1, 2}
	d := [3]int{}
	fmt.Printf("test", a, b, c, d)
}
