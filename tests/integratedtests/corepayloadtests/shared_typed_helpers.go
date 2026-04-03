package corepayloadtests

import (
	"fmt"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/errcore"
)

// testUserCov23 is a sample struct used across Coverage23/24/25 test files.
// Moved here from Coverage23_TypedFuncs_test.go so split-recovery subfolders can see it.
type testUserCov23 struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	Age   int    `json:"Age"`
}

func makeTypedWrapperCov23(name, id string, data testUserCov23) *corepayload.TypedPayloadWrapper[testUserCov23] {
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUserCov23](name, id, "testUser", data)
	if err != nil {
		panic(err)
	}
	return tw
}

func makeCollectionCov23() *corepayload.TypedPayloadCollection[testUserCov23] {
	col := corepayload.NewTypedPayloadCollection[testUserCov23](3)
	col.Add(makeTypedWrapperCov23("user", "1", testUserCov23{Name: "Alice", Email: "a@a.com", Age: 30}))
	col.Add(makeTypedWrapperCov23("user", "2", testUserCov23{Name: "Bob", Email: "b@b.com", Age: 25}))
	col.Add(makeTypedWrapperCov23("user", "3", testUserCov23{Name: "Carol", Email: "c@c.com", Age: 35}))
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
