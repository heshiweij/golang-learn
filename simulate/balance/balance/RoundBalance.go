package balance

import (
	"errors"
	"fmt"
)

func init() {
	//mgr.allBalance["round"] = &RoundBalance{}
	RegisterBalance("round", &RoundBalance{})
	fmt.Println("round")
}

func init()  {
	fmt.Println("round2")
}

type RoundBalance struct {
	current int // current index
}

func (r *RoundBalance) DoBalance(instances []*Instance) (instance *Instance, err error) {
	length := len(instances)
	if length <= 0 {
		err = errors.New("can't find any instances")
	}
	r.current ++
	index := r.current % length
	instance = instances[index]
	return
}
