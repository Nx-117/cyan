package randMath

/**
随机数相关
*/
import (
	"math/rand"
	"time"
)

/**
生成4位随机数
*/
func RandMath4() interface{} {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
}

/**
生成6位随机数
*/
func RandMath6() interface{} {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
}

/**
生成 0 至 age指定范围 的随机数
*/
func RandMath(age int32) interface{} {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(age)
}
