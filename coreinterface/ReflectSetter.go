package coreinterface

// ReflectSetter
//
// ReflectSetTo
//
//	sets current object to something else by casting,
//	reflection, by unmarshalling or by marshalling
//
// # Set any object from to toPointer object
//
// Valid Inputs or Supported (https://t.ly/SGWUx):
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
//
// Validations:
//   - Check null, if both null no error return quickly.
//   - NotSupported returns as error.
//   - NotSupported: (from, to) - (..., not pointer)
//   - NotSupported: (from, to) - (null, notNull)
//   - NotSupported: (from, to) - (notNull, null)
//   - NotSupported: (from, to) - not same type and not bytes on any
//   - `From` null or nil is not supported and will return error.
//
// Reference:
//   - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
//   - Method document screenshot    : https://prnt.sc/26dmf5g
type ReflectSetter interface {
	// ReflectSetTo
	//
	// ReflectSetter
	//  sets current object to something else by casting,
	//  reflection, by unmarshalling or by marshalling
	//
	// Set any object from to toPointer object
	//
	// Valid Inputs or Supported (https://t.ly/SGWUx):
	//  - From, To: (null, null)                          -- do nothing
	//  - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
	//  - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
	//  - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
	//  - From, To: (otherType, *[]byte)                  -- try marshal, reflect
	//
	// Validations:
	//  - Check null, if both null no error return quickly.
	//  - NotSupported returns as error.
	//      - NotSupported: (from, to) - (..., not pointer)
	//      - NotSupported: (from, to) - (null, notNull)
	//      - NotSupported: (from, to) - (notNull, null)
	//      - NotSupported: (from, to) - not same type and not bytes on any
	//  - `From` null or nil is not supported and will return error.
	//
	// Reference:
	//  - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
	//  - Method document screenshot    : https://prnt.sc/26dmf5g
	ReflectSetTo(toPointer interface{}) error
}
