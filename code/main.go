package main

import (
	"errors"
	"fmt"

	// "math/rand"
	"sync"
	"time"
)

type person struct{
	name string
	age uint8
	address string
	job string
}

type doctor struct {
	person
	hospital string
	phone string
}


func (p person) sayHi(){
	fmt.Printf("I am %v\n", p.name)
	fmt.Printf("I work as a %v\n", p.job)
	fmt.Printf("I am %v\n", p.age)
	fmt.Printf("I live at %v\n", p.address)
}

func main() {
	
/* //******** Functions and Errors ********** 
	
	//* print hello world
	name := "Mostafa"
	sayHello(name) // Hello Mostafa!

	//* Errors 
	res, remin, err  := getIntDivision(20,3)
	
	if err == nil {
		println(res,remin) // 6 2
	}

	res, remin, err = getIntDivision(10,0)
	if err != nil{
		println(err.Error()) // You Can't divide by 0
	}

	println(getIntDivision(10,3)) // 3 1 (0x0,0x0)

	println(getIntDivision(-10,3)) // -3 -1 (0x0,0x0)
	println(getIntDivision(3,2))  // 1 1 (0x0,0x0)

	println(getIntDivision(-3,2)) // -1 -1 (0x0,0x0)
	println(getIntDivision(5,2)) // 2 1 (0x0,0x0)

	println(getIntDivision(-5,2)) // -2 -1 (0x0,0x0)
	*/

/*	//******** Arrays, Slices, Maps **********

		//* Array declarations
		intArr := [...]int32{9,8,7,6,5,4}
		var intArr = [...]int32 {4,1,2,3}
		var intArr [3]int32 = [3]int32 {4,1,2}



		//* for loops
		for i := range intArr{
			fmt.Println(intArr[i])
		}

		//* Slices
		intSlice := []int{}
		// intSlice := make([]int,5)
		for i := 0 ; i < 10; i++{
			intSlice = append(intSlice,i)
		}

		for i := range intSlice{
			fmt.Println(intSlice[i])
		}


		//* Maps

		map1 := map[int]string {1:"a",2:"b",3:"c",4:"d"}

		fmt.Println(map1[1])
		fmt.Println(map1[2])
		res, ok := map1[3]
		print(res,ok)

		delete(map1,1)
		delete(map1,3)

		for i,v := range map1{
			fmt.Printf("Key: %v => value: %v\n",i,v)
		}
*/

/* //******** Structs, Interfaces **********

		mostafa := person {"Mostafa Hesham", 22, "Cairo" , "SDE"}
		Eslam := doctor { person{"Eslam Hesham", 18, "Cairo", "Doctor"} , "Al-Ahram", "123456789"}

		mostafa.sayHi()
		Eslam.sayHi()

*/


/* //******** Pointer And References *********
	
	//? Note:
	//! You Don't need to use pointer with slices
	//* it's passed by reference auto.
	
	arr := [5]int{1,2,3,4,5}
	
	var res [5]int
	
	// modify array without pointers
	res = doubleArray(arr)
	fmt.Println(arr)
	fmt.Println(res)
	
	// modify array using pointers
	res = doubleArrayInPlace(&arr)	
	fmt.Println(arr)
	fmt.Println(res)
	*/

	//******** Go Routines, Concurrency, Locks, and Channels In Go *********

	
		//* Serial/Sequential Code 
		t0 := time.Now()

		// for i := range dbData{
		// 	callDB(i)

		// }

		fmt.Printf("Total Time To finish db calls: %v \n", time.Since(t0) ) // 7-8 seconds 


		//* Concurrent/Parallel Code 

		
		t1 := time.Now()
		
		// for i := range dbData{
		// 	wg.Add(1)
		// 	go callDB(i)
		// }

		i := 0
		for i < 2_000{
			wg.Add(1)
			go callDB(i)
			i++
		}
		
		wg.Wait()
		fmt.Printf("Total Time To finish db calls: %v \n", time.Since(t1) ) // 1.4-1.9 seconds 
		fmt.Println(results2)
	}
	
	var wg = sync.WaitGroup{}
	var mx = sync.RWMutex{}
	var dbData = [...]string {"q","w","e","r","t","y"}
	var results []string
	var results2 []int
	// var results = make([]string,6)
	
	func callDB(i int) {
		// delay := rand.Float32() * 2000 // delay up to: 2_000 milliSeconds
		// duration := time.Duration(delay) * time.Millisecond // convert to time duration type
		duration := time.Duration(2000) * time.Millisecond // convert to time duration type

		//* slow DB call here
		time.Sleep(duration)
		// res := dbData[i]

		//? writing concurrently without Locks/Mutexes ?
		//! => unexpected Results ! :( 
		
			// saveToDB(res)
			// logDB()

		// fmt.Printf("the DB call Result is: %v \n", res)

			saveToResults(i)
		// saveToResults2(i)
		// logResults2()

		wg.Done()
	}
	

	func saveToResults(i interface{}){
		switch v := i.(type) {
		case int:
			results2 = append(results2, v)
		case string:
			results = append(results, v)
		}
	}

	func saveToResults2(s int){
		mx.Lock()
		results2 = append(results2, s)
		mx.Unlock()
	}

	func saveToResults1(s string){
		mx.Lock()
		results = append(results, s)
		mx.Unlock()
	}

	func logResults(){
		mx.RLock()
		fmt.Println(results)
		mx.RUnlock()
	}

	func logResults2(){
		mx.RLock()
		fmt.Println(results2)
		mx.RUnlock()
	}


func doubleArray(arr [5]int) [5]int {
	for i := range arr {
		arr[i] *= 2
	}
	return arr
}
	
func doubleArrayInPlace(arr *[5]int) [5]int {
	for i := range *arr {
		(*arr)[i] *= 2
	}
	return *arr
}

func sayHello(name string) {
	fmt.Printf("Hello %v!\n", name)
}


func getIntDivision(enum int ,denum int ) (int,int,error) {
	var err error  
	var result, reminder int

	if denum == 0{
		err=errors.New("You Can't divide by 0")
		return result , reminder, err 
	}

	result = enum / denum
	reminder = enum % denum
	return result , reminder, err 
	// return 10
}
