/* map遍历只能用for range,因为for需要下标 */
package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"

	for k, v := range cities { //k，v分别对应key和value
		fmt.Println(k, v)

	}

	//再遍历一个较复杂的双层map，双层map中的v对应一个map，则需要再range一次
	students := make(map[string]map[string]string)
	students["01"] = make(map[string]string)
	students["01"]["name"] = "allen"
	students["01"]["sex"] = "1"
	students["01"]["adress"] = "abc"

	students["02"] = make(map[string]string)
	students["02"]["name"] = "bob"
	students["02"]["sex"] = "2"
	students["02"]["adress"] = "cdrf"

	for k1, v1 := range students {
		fmt.Printf("k1=%v", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()

	}

}
