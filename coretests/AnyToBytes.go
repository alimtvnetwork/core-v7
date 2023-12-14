package coretests

import "encoding/json"

// AnyToBytes
//
// ## Steps:
//   - If already in  []byte then return as is.
//   - If already in *[]byte then return as []byte without pointer by checking if not null.
//   - If already in  string then return as []byte(string).
//   - For rest of the cases, convert to json using Marshal and then returns the bytes
//
// Panic if json marshal has error.
func AnyToBytes(anyItem interface{}) []byte {
	switch expectedAs := anyItem.(type) {
	case []byte:
		if expectedAs == nil {
			return nil
		}

		return expectedAs
	case *[]byte:
		if expectedAs == nil || *expectedAs == nil {
			return nil
		}

		return *expectedAs
	case string:
		return []byte(expectedAs)
	default:
		toBytes, err := json.Marshal(expectedAs)

		if err != nil {
			panic(err)
		}

		return toBytes
	}
}
