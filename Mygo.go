package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello, world")
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "SpaceX 1231", 94)

	// const lightSpeed = 12312
	// var distance = 560000

	// fmt.Println(distance/lightSpeed, "seconds")

	var distance, speed = 5600.0, 1000.0

	distance++

	fmt.Println(distance/speed, "seconds")

	var num = rand.Intn(10) + 1 // rand.Intn(10) [0,10)
	fmt.Println("随机数是：", num)

	var command = "walk outside"
	var exit = strings.Contains(command, "outside") // 判断包含

	fmt.Println(exit)

	var command1 = "go east"
	if command1 == "go east" {
		fmt.Println("123")
	} else if command == "go inside" {
		fmt.Println("321")
	} else {
		fmt.Println("111")
	}

	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("111")
	case room == "lake":
		fmt.Println("222")
		fallthrough
	case room == "underwater":
		fmt.Println("333")
	}

	// var count = 10
	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	fmt.Println("Liftoff!")

	ssp := 10

	fmt.Println(ssp, "start")

	for i := 0; i < 10; i++ {
		fmt.Println(ssp)
	}

	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64)
	fmt.Println(pi32)

	fmt.Printf("type %T for %[1]v\n", pi32, pi32)

	lightSpeed1 := big.NewInt(299999)
	secondsPerDay := big.NewInt(123)
	fmt.Println(lightSpeed1, secondsPerDay)

	distance1 := new(big.Int)
	distance1.SetString("30000000000000000000000000000000000000000000000", 10)
	fmt.Println(distance1)

	fmt.Println("123\n123")
	fmt.Println(`123\n123`) // 原始字面量

	var p rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v \n", p, alpha, omega, bang)
	fmt.Printf("%c %c %c %c \n", p, alpha, omega, bang)

	//grade := 'A'

	str := "沈书鹏"
	fmt.Println(len(str), "bytes")                    //实际存的字节
	fmt.Println(utf8.RuneCountInString(str), "runes") //返回字符数

	str = str + "10"

	for i, v := range str {
		fmt.Printf("%2v %03c %d\n", i, v, v)
	}

	k := 100.0
	fmt.Println(myTransform(k), k)

	kelvin := celsius(300)
	c := kelvin.celsius()
	fmt.Println(c)

	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	F()

	func() { // 匿名该函数
		fmt.Println("你好")
	}()

	c1 := counter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	var planets [8]string
	planets[0] = "M"
	planets[1] = "F"
	planets[2] = "C"

	earth := planets[2]
	fmt.Println(earth)

}

func myTransform(k float64) float64 {
	k -= 273.15
	return k
}

type celsius float64

func (k celsius) celsius() celsius {
	return celsius(k - 273.15)
}

// 一等函数

type kelvin_ float64

func fakeSensor() kelvin_ {
	return kelvin_(rand.Intn(151) + 150)
}

func realSensor() kelvin_ {
	return 0
}

var F = func() { // 匿名该函数
	fmt.Println("你好")
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
