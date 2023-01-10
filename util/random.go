package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghigklmnopqrstuvwxyz"

func init() {
	//make sure each random element is different
	//if don't use seed and don't stop the process, will get different random number
	//but when restart the process, will get the same random element like last time
	//used seed --> different every time (Seed(seed int64), seed(this int64 number) will deceide the deterministic)
	rand.Seed(time.Now().UnixNano())
}

// random number from [min,max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //[0,n) --> so need +1
}

func RadomStr(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	owner := RadomStr(9)
	return owner
}

func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currency_types := []string{"RMB", "EUR", "USD", "AED", "ETC"}
	return currency_types[rand.Intn(len(currency_types))]
}
