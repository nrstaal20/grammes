package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimes(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Times' is called with an int", func() {
			result := g.Times(123)
			Convey("Then result should equal 'g.times(123)'", func() {
				So(result.String(), ShouldEqual, "g.times(123)")
			})
		})

	})

}
