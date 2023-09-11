package corejsontests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/errcore"
)

func Test_FromTo(t *testing.T) {
	type Example struct {
		A       string
		B       int
		SomeMap map[string]string
	}

	exampleFrom := &Example{
		A:       "Something",
		B:       1,
		SomeMap: map[string]string{},
	}

	exampleTo := &Example{}

	err := corejson.Deserialize.FromTo(
		exampleFrom,
		exampleTo)

	errcore.HandleErr(err)

	to := stringutil.AnyToStringNameField(exampleTo)
	from := stringutil.AnyToStringNameField(exampleFrom)

	convey.Convey("corejson.Deserializer.FromTo - should matches from to casting", t, func() {
		convey.So(to, convey.ShouldEqual, from)
	})
}
