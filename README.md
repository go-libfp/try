# try 


## An idiomatic error monad for go, use this if you hate if err != nil 

## Usage 


``` 
	import (
		"github.com/go-libfp/try"
		"errors" 
		"fmt" 
	) 








	// constructors 


	f_off  := errors.New("fuck off m8")

	// This Function is hella useful in practice since, everything returns t, error in go
	okWrapped := try.Wrap("hello", nil) 
	errWrapped := try.Wrap("", f_off) 

	ok1 := try.Ok("hello") 
	err1 := try.Err[string](f_off) 




	// closures
	
	errFun := func(s String) (string, error) { 
		e := errors.New("fuck off m8") 
		return s, e 
	}


	bindErr := func( s string) Try[string] {
		return try.Wrap( errFun(s) )
	} 

	
	
	


	errPrint := func(e error) {
		fmt.Printf("Error: %s\n", e.Error()) 
	} 


	printer := func(s string) {
		fmt.Printf("Success: %s\n", s) 
	}

	world := func(s string) string {return fmt.Sprintf("%s world!", s)}






	// combinators 

	ok2 := try.Map(ok, world)



	ok1.OnSuccess(printer)
	err1.OnErr(errPrint) 
	ok2.OnSuccess(printer) 



	// Getters 
	x, e := ok1.Get() 
	e = ok.GetErr() 
	x = ok.GetVal() 

```