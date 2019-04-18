// print out all XYYX combos
// looking for things like abba, zappa, and maybe kappa
// dotto, datta, gabba, gamma, hanna, jabba, kazza, kijji, kommo?, lotto, notto?, razza?,
// tadda?, talla?, tappa?, vocco?, vonno?, votto?,
package main

import "fmt"

func main() {

	//var l = "bcdfghjklmnpqrstvwxz"
	var l = "ckst"
	var v = "aeiouy"

	for _, valI := range l {

		for _, valJ := range v {

			for _, valK := range l {

				for _, valL := range v {

					fmt.Println(fmt.Sprintf("%s%s%s%s%s", string(valI), string(valJ), string(valK), string(valK), string(valL)))
					//fmt.Println(string(valJ))
					//fmt.Println(string(valK))

				}

			}

		}
	}

}
