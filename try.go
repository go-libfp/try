
package try

type Try[T any] struct {
	val T 
	err error 
}


func Ok[T any](t T) Try[T] {
	return Try[T]{t, nil}
}


func Err[T any](e error) Try[T] {
	var v T
	return Try[T]{v, e}
}



func Map[T any, U any](t Try[T], f func(x T) U) Try[U] {
	if t.err == nil {
	
		out := f(t.val)
		return Ok(out) 
	}
	
	return Err[U](t.err)
}

func Bind[T, U any](t Try[T], f func(x T) Try[U]) Try[U] {
	if t.err == nil {
		return f(t.val)
	}

	return Err[U](t.err) 

}

func WrapErr[T any](t T, err error) Try[T] {
	return Try[T]{t, err}
}


func (t *Try[T]) Get() (T, error) {
	return t.val, t.err 
}


func (t *Try[T]) OnErr(f func(e error)) {
	if t.err != nil {
		f(t.err)
	}
}




func (t *Try[T]) OnSuccess(f func(x T) ) {
	if err == nil {
		f(t.val)
	}
}


func (t *Try[T]) IsErr() bool {
	return err != nil
} 

func (T *Try[T]) IsOk() bool {
	return e == nil 
}
