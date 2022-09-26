package utils

import (
	"math/rand"
	"time"
)

func RandWindValue() int8 {
	rand.Seed(time.Now().UnixMicro())
	res := rand.Intn((100 - 1) + 1)
	return int8(res)
}

func RandWaterValue() int8 {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn((100 - 1) + 1)
	return int8(res)
}
