package pkg_math

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GetRandomCode(n int) string {
	nStr := strconv.Itoa(n)
	format := "%0" + nStr + "v"
	num := int32(math.Pow10(n))
	code := fmt.Sprintf(format, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(num))
	return code
}

// layout 20060102150405
// RandomNowCode(6, "20060102_150405_", "_", false, "userid")
func RandomNowCode(n int, layout, suffix string, leftOrRight bool, prefix ...string) string {
	pre := ""
	timeStr := time.Now().Format(layout)
	randomCode := GetRandomCode(n)
	code := timeStr + randomCode
	if leftOrRight {

		code = pre + code
	} else {
		if len(prefix) > 0 {
			for _, s := range prefix {
				pre += suffix
				pre += s
			}
		}
		code = code + pre
	}
	return code
}
