package baseactioninf

// ValidationErrorGetter
//
//  Returns validation related error
//  only
type ValidationErrorGetter interface {
	// ValidationError
	//
	//  Returns validation related error
	//  only
	ValidationError() error
}
