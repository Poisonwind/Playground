package main

import (
	"fmt"
	"hw/electronic"
)

func main() {

	iPhone := electronic.NewApplePhone("iPhone 11", "iOS 11")
	s10 := electronic.NewAndroidPhone("samsung", "s10", "android")
	KX10 := electronic.NewRadioPhone("LG", "KX-10", 12)

	printCharacteristics(iPhone)
	printCharacteristics(s10)
	printCharacteristics(KX10)


}

func printCharacteristics(phone electronic.Phone)  {

	fmt.Println(phone.Type())
	fmt.Println(phone.Brand())
	fmt.Println(phone.Model())

	switch t := phone.(type) {
	case *electronic.ApplePhone:
		fmt.Println(t.OS())
	case *electronic.AndroidPhone:
		fmt.Println(t.OS())
	case *electronic.RadioPhone:
		fmt.Println(t.ButtonCount())		
	}

}