package main

import (
	"fmt"
	"go_learner/simulate/balance/balance"
	"strconv"
	"time"
)

//var logger *logrus.Logger
//var logFile = "./log"

func init() {
	if false {
		//logger = logrus.New()
		//logger.SetLevel(logrus.DebugLevel)
		//
		//file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		//if err != nil {
		//	panic(err)
		//}
		//logger.SetOutput(file)
	}
}

const (
	RAND = iota
	ROUND
)

/*
func makeBalance(bType int) (b balance.Balance) {
	switch bType {
	case RAND:
		b = &balance.RandomBalance{}
	case ROUND:
		b = &balance.RoundBalance{}
	}

	return
}
*/

func main() {

	// create instances
	instances := make([]*balance.Instance, 0)

	for i := 0; i < 5; i++ {
		one := balance.NewInstance("192.168.1."+strconv.Itoa(i), 1100+i, 1)
		instances = append(instances, one)
	}

	for {

		inst, err := balance.DoBalance("hash", instances)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Found: %s:%d\n", inst.GetHost(), inst.GetPort())

		<-time.After(2 * time.Second)
	}

	/*
	var balanceInst = makeBalance(RAND)

	for {
		inst, err := balanceInst.DoBalance(instances)

		if err != nil {
			fmt.Println("select host error: ", err)
			continue
		}

		//fmt.Println("found: ", inst.GetHost(), ":", inst.GetPort())

		fmt.Println("found: ", inst.String())

		<-time.After(time.Second)
	}
	*/

}
