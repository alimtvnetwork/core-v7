package corepayloadtests

import (
	"fmt"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/errcore"
)

// testUserTyped is a sample struct used across TypedPayloadWrapper test files.
// Moved here from TypedPayloadWrapper_test.go so split-recovery subfolders can see it.
type testUserTyped struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Age   int    `json:"Age"`
}

func makeTypedWrapper(name, id string, data testUserTyped) *corepayload.TypedPayloadWrapper[testUserTyped] {
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUserTyped](name, id, "testUser", data)
	if err != nil {
		panic(err)
	}
	return tw
}

func makeTypedCollection() *corepayload.TypedPayloadCollection[testUserTyped] {
	col := corepayload.NewTypedPayloadCollection[testUserTyped](3)
	col.Add(makeTypedWrapper("user", "1", testUserTyped{Name: "Alice", Email: "a@a.com", Age: 30}))
	col.Add(makeTypedWrapper("user", "2", testUserTyped{Name: "Bob", Email: "b@b.com", Age: 25}))
	col.Add(makeTypedWrapper("user", "3", testUserTyped{Name: "Carol", Email: "c@c.com", Age: 35}))
	return col
}

// createNumberedUsers creates a typed collection with N numbered users.
// Moved here from TypedCollectionPaging_test.go so split-recovery subfolders can see it.
func createNumberedUsers(count int) *corepayload.TypedPayloadCollection[testUser] {
	wrappers := make([]*corepayload.TypedPayloadWrapper[testUser], 0, count)

	for i := 0; i < count; i++ {
		user := testUser{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@test.com", i),
			Age:   20 + i,
		}

		typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
			user.Name,
			fmt.Sprintf("user-%d", i),
			user,
		)
		errcore.HandleErr(err)

		wrappers = append(wrappers, typed)
	}

	return corepayload.TypedPayloadCollectionFrom[testUser](wrappers)
}

// testUserCov23 is a sample struct for TypedPayloadWrapper_SetName tests.
type testUserCov23 struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func makeTypedWrapperCov23(name, id string, data testUserCov23) *corepayload.TypedPayloadWrapper[testUserCov23] {
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUserCov23](name, id, "testUserCov23", data)
	if err != nil {
		panic(err)
	}
	return tw
}

func makeCollectionCov23() *corepayload.TypedPayloadCollection[testUserCov23] {
	col := corepayload.NewTypedPayloadCollection[testUserCov23](3)
	col.Add(makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice"}))
	col.Add(makeTypedWrapperCov23("user", "2", testUserCov23{Name: "Bob"}))
	col.Add(makeTypedWrapperCov23("user", "3", testUserCov23{Name: "Carol"}))
	return col
}
