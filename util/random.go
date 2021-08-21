package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt 生成随机整数
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0->max-min
}

//RandomString 生成长度为n的随机字符串
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner 随机生成owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney 随机生成金额
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency 生成随机货币
func RandomCurrency() string {
	currencies := []string{USD, CAD, EUR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
