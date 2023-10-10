package util

import (
	"math/rand"
	"strconv"
	"strings"
)

func RandCode(number int) string {
	var randCode strings.Builder
	for i := 0; i < number; i++ {
		randCode.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return randCode.String()
}
