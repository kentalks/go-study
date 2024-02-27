package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

}

func main() {

	/*
		r := utils.SetupRouter()
		// Listen and Server in 0.0.0.0:8080
		r.Run(":8080")
	*/

	/*

		c := make(chan int)
		for i := 0; i < 5; i++ {
			worker := &utils.Worker{Id: i}
			go worker.Process(c)
		}
		for {
			select {
			case c <- rand.Int():
			case <-time.After(time.Millisecond * 100):
				fmt.Println("timed out")
			}
			time.Sleep(time.Millisecond * 50)
		}
	*/

	/*
		var (
			counter = 0
			lock    sync.Mutex
		)
		for i := 0; i < 2; i++ {
			go func() {
				lock.Lock()
				defer lock.Unlock()
				counter++
				fmt.Println(counter)
			}()
			time.Sleep(time.Millisecond * 1)
		}
	*/
	/*
		fmt.Println("start")
		go func() {
			fmt.Println("process")
		}()
		time.Sleep(time.Millisecond * 10)
		fmt.Println("done")
	*/
	/*
		if err := utils.Process(2); err != nil {
			fmt.Println(err)
		}
	*/
	/*
		var input int
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			fmt.Println("no more input")
		}

		n, err := strconv.Atoi("s")
		if err != nil {
			fmt.Println("not a valid No.")
		} else {
			fmt.Println(n)
		}
	*/

	//fmt.Println(utils.PriceCheck(1))
	//utils.HelloMap()
	//utils.HelloSlice()

	//utils.HelloArray()
	/*
		cat := utils.Cat{
			Age: 0,
			Pet: &utils.Pet{Name: "小白"},
		}
		utils.Super(&cat)
		fmt.Println(cat.Age)
		(&cat).Super()
		fmt.Println(cat.Age)
		cat.Super()
		fmt.Println(cat.Age)
	*/
	//cat.Info()

	//utils.PrintHellos([]string{"mei", "ken"})

	// r := utils.RandString(10)
	// fmt.Println(r)
	// fmt.Println(utils.ToUpper(r))

	// s := utils.RandSeq(10)
	// fmt.Println(s)

	// fmt.Println(utils.ProcessStr("asdfasdf", 12))

	/*
		in := [3]string{"a", "b", "c"}
		var out []*string
		for _, v := range in {
			v := v
			out = append(out, &v)
		}
		fmt.Println(*out[0], *out[1], *out[2])
	*/

	/*
		fmt.Println(utils.GCD(-6, 18))
	*/

}
