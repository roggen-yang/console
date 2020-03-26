package light

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestLight_LightOn(t *testing.T) {
	convey.Convey("测试开灯动作", t, func() {
		li := NewLight(2*time.Second, 2*time.Second)
		li.LightOn()
		status := li.GetLightStatus()
		convey.So(status, convey.ShouldEqual, On)
	})
}

func TestLight_LightOff(t *testing.T) {
	convey.Convey("测试关灯动作", t, func() {
		li := NewLight(2*time.Second, 2*time.Second)
		li.LightOff()
		status := li.GetLightStatus()
		convey.So(status, convey.ShouldEqual, Off)
	})
}

func TestLight_LightDelay(t *testing.T) {
	convey.Convey("测试延迟动作", t, func() {
		li := NewLight(2*time.Second, 2*time.Second)
		li.LightDelay(On)
		time.Sleep(2 * time.Second)
		status1 := li.GetLightStatus()
		convey.So(status1, convey.ShouldEqual, On)
		li.LightDelay(Off)
		time.Sleep(2 * time.Second)
		status2 := li.GetLightStatus()
		convey.So(status2, convey.ShouldEqual, Off)
	})
}
