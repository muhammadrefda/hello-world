package main

import (
	"fmt"
	_ "strings"
)

func main() {
	//var firstName string = "John"
	//var middleName = "Paul"
	//lastName := "Doe"
	//lastName = "Smith"

	//fmt.Printf("Hello %s %s %s!\n", firstName, middleName, lastName)

	/*
		var first, second, third string
		first, second, third = "John", "Paul", "Doe"
		fmt.Println(first, second, third)
	*/

	//var fourth, fifth, sixth string = "mick", "stanley", "harrison"
	//fmt.Println(fourth, fifth, sixth)

	//seventh, eighth, ninth := "jagger", "jhonny", "james"
	//fmt.Println(seventh, eighth, ninth)
	//
	//number, day, year := 1, "Monday", 2019
	//fmt.Println(number, day, year)

	//_ = "doni"
	//_ = "dodi"
	//name, _ := "anwar", "sudarsono"
	//fmt.Println(name, _)

	//name := new(string)
	//fmt.Println(name)
	//fmt.Println(*name)

	/*
		var positiveNumber uint8 = 89
		var negativeNumber int8 = -89
		fmt.Println("bilangan positif:", positiveNumber)
		fmt.Println("bilangan negatif:", negativeNumber)
	*/

	//var decimalNumber = 2.62
	//fmt.Printf("bilangan decimal: %f\n", decimalNumber)
	//fmt.Printf("bilangan decimal: %.3f\n", decimalNumber)
	//fmt.Println("bilangan decimal:", decimalNumber)

	//	var exist bool = true
	//	fmt.Printf("exist? %t \n", exist)
	//
	//	var message = `My Name is
	//jhonny`
	//
	//	fmt.Println(message)

	//fmt.Println("jhonny deep")
	//fmt.Println("john", "wick")

	//fmt.Print("deddy\n")
	//fmt.Print("dodo ", "wick")
	//fmt.Print("john", " ", "wick\n")

	//var value = (((2+6)%3)*4 - 2) / 3
	//var isEqual = (value == 2)

	//fmt.Printf("nilai %d (%t) \n", value, isEqual)

	//var left = false
	//var right = true

	//var leftAndRight = left && right
	//fmt.Printf("left && right \t(%t) \n", leftAndRight)
	//
	//var leftOrRight = left || right
	//fmt.Printf("left || right \t(%t) \n", leftOrRight)
	//
	//var leftReverse = !left
	//fmt.Printf("!left \t\t(%t) \n", leftReverse)

	//var point = 3
	//
	//if point == 10 {
	//	fmt.Println("lulus dgn nilai sempurna")
	//} else if point > 5 {
	//	fmt.Println("lulus")
	//} else if point == 4 {
	//	fmt.Println("hampir lulus")
	//} else {
	//	fmt.Println("tidak lulus")
	//}

	//var point = 8840.0
	//if percent := point / 100; percent >= 100 {
	//	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	//} else if percent >= 70 {
	//	fmt.Printf("%.1f%s good\n", percent, "%")
	//} else {
	//	fmt.Printf("%.1f%s not bad\n", percent, "%")
	//}

	//var point = 6

	//switch point {
	//case 8:
	//	fmt.Println("lulus dgn nilai sempurna")
	//case 7:
	//	fmt.Println("lulus")
	//default:
	//	fmt.Println("tidak lulus")
	//}

	//switch {
	//case point == 3:
	//	fmt.Println("lulus dgn nilai sempurna")
	//case (point > 5) && (point < 8):
	//	fmt.Println("lulus")
	//default:
	//	{
	//		fmt.Println("tidak lulus")
	//		fmt.Println("you failed")
	//	}
	//}

	//switch {
	//case point == 8:
	//	fmt.Println("lulus dgn nilai sempurna")
	//case (point < 8) && (point > 3):
	//	fmt.Println("lulus")
	//	fallthrough
	//case point < 5:
	//	fmt.Println("hampir lulus")
	//default:
	//	{
	//		fmt.Println("tidak lulus")
	//		fmt.Println("you failed")
	//	}
	//}

	//var point = 1

	//if point > 7 {
	//	switch point {
	//	case 10:
	//		fmt.Println("perfect")
	//	default:
	//		fmt.Println("good")
	//	}
	//} else {
	//	if point == 5 {
	//		fmt.Println("not bad")
	//	} else if point == 3 {
	//		fmt.Println("keep trying")
	//	} else {
	//		fmt.Println("you failed")
	//		if point == 0 {
	//			fmt.Println("you got zero")
	//		}
	//	}
	//}

	//for i := 0; i <= 5; i++ {
	//	fmt.Println(i)
	//}

	//var i = 0

	//for i < 5 {
	//	fmt.Println("Angka", i)
	//	i++
	//	if i == 5 {
	//	}
	//}

	//var i = 0

	//for {
	//	fmt.Println("Angka", i)
	//
	//	i++
	//
	//	if i == 55 {
	//		break
	//	}
	//}

	//for i := 0; i < 10; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//
	//	if i > 8 {
	//		break
	//	}
	//
	//	fmt.Println("Angka", i)
	//}

	//for i := 0; i < 5; i++ {
	//	for j := i; j < 5; j++ {
	//		fmt.Print(j, " ")
	//	}
	//	fmt.Println()
	//}

	//outerLoop:
	//	for i := 0; i < 5; i++ {
	//		for j := i; j < 5; j++ {
	//			if i == 3 {
	//				break outerLoop
	//			}
	//			fmt.Print("matriks [", i, "][", j, "]", "\n")
	//		}
	//	}

	//var names [4]string
	//names[0] = "jhonny"
	//names[1] = "wick"
	//names[2] = "deddy"
	//names[3] = "dodo"
	//
	//fmt.Println(names[0], names[1], names[2], names[3])

	//var fruits [4]string
	//
	//fruits = [4]string{
	//	"apple",
	//	"banana",
	//	"orange",
	//	"grape"}
	//
	//fmt.Println("Jumlah Element \t\t", len(fruits))
	//fmt.Println("Isi Element \t\t", fruits)

	//var numbers = [...]int{1, 2, 3, 4, 5}

	//fmt.Println("data array \t:", numbers)
	//fmt.Println("jumlah array \t:", len(numbers))

	//var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	//var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	//fmt.Println("numbers1:", numbers1)
	//fmt.Println("numbers2:", numbers2)

	//var fruits = [4]string{"apple", "banana", "orange", "grape"}

	//for i := 0; i < len(fruits); i++ {
	//	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	//}
	//
	//var fruits = [4]string{"apple", "grape", "banana", "melon"}
	//
	//for i, fruit := range fruits {
	//	fmt.Printf("elemen %d : %s\n", i, fruit)
	//}

	// penggunaan underscore
	//var fruits = [4]string{"apple", "grape", "banana", "melon"}

	//for _, fruit := range fruits {
	//	fmt.Println(fruit)
	//}

	//var fruits = make([]string, 2)
	//fruits[0] = "apple"
	//fruits[1] = "grape"

	//fmt.Println(fruits)

	//var fruits = []string{"apple", "grape", "banana", "melon"}
	//var cars = [...]string{"mazda", "honda", "toyota", "nissan"}

	//fmt.Println(fruits[0])
	//fmt.Println(cars)

	//var fruits = []string{"apple", "grape", "banana", "melon"}
	//var newFruits = fruits[0:2]
	//fmt.Println(newFruits)

	//var fruits = []string{"apple", "grape", "banana", "melon"}

	//var aFruits = fruits[0:3]
	//var bFruits = fruits[1:4]

	//var aaFruits = aFruits[1:2]
	//var bbFruits = bFruits[0:2]
	//
	//fmt.Println(fruits)
	//fmt.Println(aFruits)
	//fmt.Println(bFruits)
	//fmt.Println(aaFruits)
	//fmt.Println(bbFruits)
	//
	//bbFruits[0] = "pineapple"
	//
	//fmt.Println(fruits)

	//var fruits = []string{"apple", "grape", "banana", "melon"}
	//
	//fmt.Println(len(fruits))
	//
	//fmt.Println("fruits value using cap", cap(fruits))
	//
	//var aFruits = fruits[0:2]
	//fmt.Println(len(aFruits))
	//fmt.Println(cap(aFruits))

	//var fruits = []string{"apple", "grape", "banana"}
	//var bFruits = fruits[0:2]
	//
	//fmt.Println("ini cap", cap(bFruits))
	//fmt.Println("ini len", len(bFruits))
	//
	//fmt.Println(fruits)
	//fmt.Println(bFruits)
	//
	//var cFruits = append(bFruits, "papaya")
	//
	//fmt.Println(fruits)
	//fmt.Println(bFruits)
	//fmt.Println(cFruits)

	//dst := make([]string, 3)
	//src := []string{"watermelon", "pineapple", "apple", "orange"}
	//n := copy(dst, src)

	//fmt.Println(dst)
	//fmt.Println(src)
	//fmt.Println(n)

	//dst := []string{"potato", "tomato", "onion"}
	//src := []string{"watermelon", "pineapple"}
	//n := copy(dst, src)
	//
	//fmt.Println(dst)
	//fmt.Println(src)
	//fmt.Println(n)

	//var fruits = []string{"apple", "grape", "banana"}
	//var aFruits = fruits[0:2]
	//var bFruits = fruits[0:2:2]
	//
	//fmt.Println(fruits)
	//fmt.Println(len(fruits))
	//fmt.Println(cap(fruits))
	//
	//fmt.Println(aFruits)
	//fmt.Println(len(aFruits))
	//fmt.Println(cap(aFruits))
	//
	//fmt.Println(bFruits)
	//fmt.Println(len(bFruits))
	//fmt.Println(cap(bFruits))

	//var chicken map[string]int
	//chicken = map[string]int{}

	//chicken["january"] = 50
	//chicken["february"] = 40

	//fmt.Println("january", chicken["january"])
	//fmt.Println("february", chicken["february"])
	//fmt.Println("mei", chicken["mei"])

	//var data map[string]int

	//data = map[string]int{}
	//data["one"] = 1

	//fmt.Println(data["one"])

	//var chicken1 = map[string]int{"january": 50, "february": 40}

	//var chickenVertical = map[string]int{
	//	"january":  50,
	//	"february": 40,
	//}

	//fmt.Println(chicken1)
	//fmt.Println(chickenVertical)

	//var chicken3 = map[string]int{}
	//var chicken4 = make(map[string]int)
	//var chicken5 = *new(map[string]int)

	//var chicken = map[string]int{
	//	"januari":  50,
	//	"februari": 40,
	//	"maret":    34,
	//	"april":    67,
	//}
	//for key, val := range chicken {
	//	fmt.Println(key, " \t:", val)
	//}

	//var chicken = map[string]int{"januari": 50, "februari": 40}

	//fmt.Println(len(chicken))
	//fmt.Println(chicken)

	//delete(chicken, "januari")
	//fmt.Println(len(chicken))
	//fmt.Println(chicken)

	//var chicken = map[string]int{"januari": 50, "februari": 40}
	//var value, isExist = chicken["januari"]
	//
	//if isExist {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("item is not exists")
	//}

	//var chickens = []map[string]string{
	//	{
	//		"name":   "chicken blue",
	//		"gender": "male",
	//	},
	//	{
	//		"name":   "chicken red",
	//		"gender": "male"},
	//
	//	{
	//		"name":   "chicken yellow",
	//		"gender": "female",
	//	},
	//}

	//for _, chicken := range chickens {
	//	fmt.Println(chicken["gender"], chicken["name"])
	//}

	//var data = []map[string]string{
	//	{"name": "chicken blue", "gender": "male", "color": "brown"},
	//	{"address": "mango street", "id": "k001"},
	//	{"community": "chicken lovers"},
	//}

	//fmt.Println(data)

	//var names = []string{"john", "wick"}
	//printMessage("halo", names)

	//rand.Seed(time.Now().Unix())
	//var randomValue int

	//randomValue = randomWithRange(2, 10)
	//fmt.Println("random number:", randomValue)

	//randomValue = randomWithRange(2, 10)
	//fmt.Println("random number:", randomValue)

	//randomValue = randomWithRange(2, 10)
	//fmt.Println("random number:", randomValue)

	//divideNumber(10, 2)
	//divideNumber(4, 0)
	//divideNumber(8, -4)

	//var diameter float64 = 7
	//var area, circumference = calculate(diameter)
	//
	//fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	//fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("average is %.2f", avg)
	fmt.Println(msg)
}

//func printMessage(message string, arr []string) {
//	var nameString = strings.Join(arr, " ")
//	fmt.Println(message, nameString)
//}

//func randomWithRange(min, max int) int {
//	var value = rand.Int()%(max-min+1) + min
//	return value
//}

//func divideNumber(m, n int) {
//	if n == 0 {
//		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
//		return
//	}
//
//	var res = m / n
//	fmt.Printf("%d / %d = %d\n", m, n, res)
//}

//func calculate(d float64) (float64, float64) {
//	var area = math.Pi * math.Pow(d/2, 2)
//	var circumference = math.Pi * d
//
//	return area, circumference
//}

//Fungsi Dengan Predefined Return Value
//func calculate(d float64) (area float64, circumference float64) {
//	area = math.Pi * math.Pow(d/2, 2)
//	circumference = math.Pi * d
//
//	return
//}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
