package balance

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//mgr.allBalance["rand"] = &RandomBalance{}
	RegisterBalance("rand", &RandomBalance{})

	fmt.Println("rand")
}

type RandomBalance struct {
}

// [min,max)
func makeRandomNum(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if min > max {
		return max
	}

	return r.Intn(max-min) + min
}

func (r *RandomBalance) DoBalance(instances []*Instance) (instance *Instance, err error) {
	if len(instances) <= 0 { // include instance = nil
		//panic("Can't find any instance")
		err = errors.New("can't find any instance")
	}

	instance = instances[makeRandomNum(0, len(instances))]
	return
}
