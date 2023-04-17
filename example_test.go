package try_test 

import (
	"github.com/go-libfp/try"
	"encoding/json"
	"fmt"

	"testing"
	"log"
	"errors"
)


type Foo struct { Bar string}


func jsonDecode[T any](data []byte) (T, error)  {
	var out T 
	err := json.Unmarshal(data, &out)
	return out, err 
} 


func ExampleUsage() {
	t := try.WrapErr( json.Marshal(Foo{"hello"}) ) 
	

	// bind ET will run a function that returns err tuple 
	t1 := try.BindET(t, jsonDecode[Foo])

	t2 := try.Map(t1, func(x Foo) string {
		return x.Bar
	}  ) 


	t2.OnSuccess(func(x string) {
		fmt.Println(x) 
	} ) 


	t2.OnErr(func(e error) {
		log.Println(e)
	})

}




func ExampleConstructorsGetters() {
	t := try.Ok("hello buddeh")
	x, e := t.Get() 

	fmt.Println(x)
	

	//maybe rename try.Err to Fail in the future
	dam := errors.New("damn")
	fail_t := try.Err[string](dam)

	t.IsOk() 
	t.IsErr() 

	e = t.GetErr() 	
	x = t.GetValue() 

	if e != nil {
		log.Println(e) 
	}


	fail_t.OnErr(
		func(e error) { log.Println(e) }, 
	) 


} 





func TestExamples(t *testing.T) {
	ExampleUsage()
	ExampleConstructorsGetters() 
}

