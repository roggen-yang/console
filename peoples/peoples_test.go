package peoples

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestPeoples_ComeIn(t *testing.T) {
	convey.Convey("测试有人进入房间", t, func() {
		peo := NewPeoples()
		peo.ComeIn(1)
		c := peo.GetCounter()
		convey.So(c, convey.ShouldEqual, 1)
	})
}

func TestPeoples_GetOut(t *testing.T) {
	convey.Convey("测试有人离开房间", t, func() {
		peo := NewPeoples()
		peo.ComeIn(1)
		c1 := peo.GetCounter()
		convey.So(c1, convey.ShouldEqual, 1)
		peo.GetOut(1)
		c2 := peo.GetCounter()
		convey.So(c2, convey.ShouldEqual, 0)
	})
}
