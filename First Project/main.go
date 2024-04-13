package main

//if the value is not divisible by six evenly please  wthrowing an error
import (
	"errors"
	"fmt"
)

func main(){
	var isIt, err =  isDivisibleBySix(38);
	fmt.Println(isIt)
	if(err!= nil){
		fmt.Printf(err.Error())
	} 

}

func isDivisibleBySix(num int) (bool, error) {
	var err error;
	if (num%6 != 0){
		err = errors.New("No cant divide by 6")
		return false, err;
	}
	err = nil;
	
	return true, nil;
	
}