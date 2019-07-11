package main

import (
	"fmt"
	"strings"
	"unicode"
)

const xiaomiId = "9hf0P6MBHM7OY6UOWm"
const expected = "9hf0_p6_m_b_h_m7-_y6_u_o_"

func main()  {


	fmt.Println(converter(xiaomiId))

}

func converter(id string) string  {

	new_rune := []rune{}

	runes := []rune(id)
	for i:= 0; i < len(runes); i++ {


		if(unicode.IsUpper(runes[i])){

			temp_rune := string(runes[i])

			temp_rune = strings.ToLower(temp_rune)

			temp_rune = "_" + temp_rune

			temp_runes := []rune(temp_rune)

			for k := 0; k < len(temp_runes); k++{

				new_rune = append(new_rune, temp_runes[k])

			}


		} else{

			new_rune = append(new_rune, runes[i])

		}
	}

	new_string :=  string(new_rune)



	return new_string

}



