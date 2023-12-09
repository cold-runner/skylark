package util

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
)

func RandSmsCode(number int) (string, error) {
	if number < 1 {
		return "", errors.New("随机位数至少为1")
	}
	if number > 10 {
		return "", errors.New("随机位数至多为10")
	}

	factor := int(math.Pow10(number - 1))
	return strconv.Itoa(rand.Intn(number*factor) + factor), nil
}
