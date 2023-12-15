//Задача 1
//Получение максимального значения для типа

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Если значение не нулевое, вернуть максимальное значение для типа
func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	var mInt8 int8
	var mInt16 int16
	var mInt32 int32
	var mInt64 int64
	if in8 != 0 {
		mInt8 = int8(1<<(8-1) - 1)
	}
	if in16 != 0 {
		mInt16 = int16(1<<(16-1) - 1)
	}
	if in32 != 0 {
		mInt32 = int32(1<<(32-1) - 1)
	}
	if in64 != 0 {
		mInt64 = int64(1<<(64-1) - 1)
	}
	return mInt8, mInt16, mInt32, mInt64
}

// Если значение не нулевое, вернуть максимальное значение для типа
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	var muin8 uint8
	var mInt16 uint16
	var mInt32 uint32
	var mInt64 uint64
	if uin8 != 0 {
		muin8 = uint8(1<<(8) - 1)
	}
	if uin16 != 0 {
		mInt16 = uint16(1<<(16) - 1)
	}
	if uin32 != 0 {
		mInt32 = uint32(1<<(32) - 1)
	}
	if uin64 != 0 {
		mInt64 = uint64(1<<(64) - 1)
	}
	return muin8, mInt16, mInt32, mInt64
}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)
	fmt.Println(bits)
	return bits
}
func main() {
	fmt.Println(getIntMaxValue(0, 5, 3, 0))
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
