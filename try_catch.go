package goutils

/*
	https://dzone.com/articles/try-and-catch-in-golang
*/

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

/*
	Throw
*/
func Throw(up Exception) {
	panic(up)
}

/*
	Do
*/
func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
