//Задача 1
//Получение максимального значения для типа

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Если значение не нулевое, вернуть максимальное значение для типа
func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	var mInt8 int8 = math.MaxInt8
	var mInt16 int16 = math.MaxInt16
	var mInt32 int32 = math.MaxInt32
	var mInt64 int64 = math.MaxInt64
	if in8 != 0 {
		return mInt8, 0, 0, 0
	}
	if in16 != 0 {
		return 0, mInt16, 0, 0
	}
	if in32 != 0 {
		return 0, 0, mInt32, 0
	}
	if in64 != 0 {
		return 0, 0, 0, mInt64
	}
	return 0, 0, 0, 0
}

// Если значение не нулевое, вернуть максимальное значение для типа
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	var muin8 uint8 = math.MaxUint8
	var mInt16 uint16 = math.MaxUint16
	var mInt32 uint32 = math.MaxUint32
	var mInt64 uint64 = math.MaxUint64
	if uin8 != 0 {
		return muin8, 0, 0, 0
	}
	if uin16 != 0 {
		return 0, mInt16, 0, 0
	}
	if uin32 != 0 {
		return 0, 0, mInt32, 0
	}
	if uin64 != 0 {
		return 0, 0, 0, mInt64
	}
	return 0, 0, 0, 0
}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)
	fmt.Println(bits)
	return bits
}
func main() {
	fmt.Println(getIntMaxValue(0, 5, 0, 0))
	fmt.Println(getUintMaxValue(0, 0, 3, 0))
	//g1 := getBits(42241)
	//g2 := getBits(12)
	fmt.Println(equalbits(22, 22))
	var x int = 0
	increment(&x)
	fmt.Println(x)
}

// Задача 2
// Функция при передаче одинаковых переменных возвращает true используя bitwise операции
func equalbits(bits1, bits2 int) bool {
	x := (bits1 ^ bits2)
	var res bool
	if x == 0 {
		res = true
	} else {
		res = false
	}
	return res
}

// Задача 3
// Написать функцию которая принимает указатель на переменную и увеличивает ее значение на 1, с помощью инкремента
func increment(x *int) {
	*x++
}
