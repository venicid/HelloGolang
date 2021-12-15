package async

import (
	"errors"
	"time"
)

func Sum(args []int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	time.Sleep(20 * time.Second)
	return sum, errors.New("sum errors")

}
