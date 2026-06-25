package main

import "fmt"

func main() {
	var str1 string = "хелло ворлд"
	var str2 string = "hello world"
	var str3 string = "你好世界"
	fmt.Println(str1, str2, str3)
	fmt.Println(len(str3))
	for _, v := range str3 {
		fmt.Printf("%c\n ", v)
	}
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%d", str1[i])
	}
	binary_str := []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 32, 208, 188, 208, 184, 209, 128}
	str := string(binary_str)
	fmt.Println(str)
	var value rune = 0xd09f
	fmt.Printf("%c", value)
	fmt.Printf("%c", value)
	var value2 rune = 'П'
	fmt.Printf("%c \n", value2) // П
	fmt.Printf("%x \n", value2) // 41f
	fmt.Printf("%d \n", value2) // 1055
	fmt.Println(value2)
	rune_slice := []rune(str2) // преобразуем str2 в срез рун
	for _, v := range rune_slice {
		fmt.Printf("%c ", v)
	}
	myrune := []rune{0x41f, 0x440, 0x438, 0x432, 0x435, 0x442, 0x020, 0x43c, 0x438, 0x440}
	str4 := string(myrune)
	fmt.Println(str4)
	str5 := "YT||YT"
	rune1 := []rune(str5)
	rune1[2] = '/'
	str5 = string(rune1)
	fmt.Println(str5)

}
