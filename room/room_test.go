package room

import (
	"testing"
	"time"

	"github.com/roggen-yang/console/light"
	"github.com/roggen-yang/console/peoples"
	"github.com/smartystreets/goconvey/convey"
)

func TestRoom_ComeIn(t *testing.T) {
	convey.Convey("测试有人进入房间时，灯的状态", t, func() {
		peo := peoples.NewPeoples()
		li := light.NewLight(2*time.Second, 2*time.Second)
		room := NewRoom(peo, peo, li, li)
		err := room.ComeIn(1)
		convey.So(err, convey.ShouldEqual, nil)
		status := room.GetLightStatus()
		convey.So(status, convey.ShouldEqual, light.Off)
		time.Sleep(2 * time.Second)
		status2 := room.GetLightStatus()
		convey.So(status2, convey.ShouldEqual, light.On)
		counter := room.GetCounter()
		convey.So(counter, convey.ShouldEqual, 1)
	})
}

func TestRoom_GetOut(t *testing.T) {
	convey.Convey("测试有人离开房间时，灯的状态", t, func() {
		peo := peoples.NewPeoples()
		li := light.NewLight(2*time.Second, 2*time.Second)
		room := NewRoom(peo, peo, li, li)
		// 有1人进来
		err := room.ComeIn(1)
		convey.So(err, convey.ShouldEqual, nil)
		status := room.GetLightStatus()
		convey.So(status, convey.ShouldEqual, light.Off)
		time.Sleep(2 * time.Second)
		status2 := room.GetLightStatus()
		convey.So(status2, convey.ShouldEqual, light.On)
		counter := room.GetCounter()
		convey.So(counter, convey.ShouldEqual, 1)
		// 离开1人
		err2 := room.GetOut(-1)
		convey.So(err2, convey.ShouldEqual, nil)
		counter2 := room.GetCounter()
		convey.So(counter2, convey.ShouldEqual, 0)
		status3 := room.GetLightStatus()
		convey.So(status3, convey.ShouldEqual, light.On)
		time.Sleep(2 * time.Second)
		status4 := room.GetLightStatus()
		convey.So(status4, convey.ShouldEqual, light.Off)
	})
}

func TestRoom_GetOut2(t *testing.T) {
	convey.Convey("测试有人进入房间，但在延时时间内就离开房间的情况下，灯的状态", t, func() {
		peo := peoples.NewPeoples()
		li := light.NewLight(2*time.Second, 2*time.Second)
		room := NewRoom(peo, peo, li, li)
		// 有1人进来
		err := room.ComeIn(1)
		convey.So(err, convey.ShouldEqual, nil)
		status := room.GetLightStatus()
		convey.So(status, convey.ShouldEqual, light.Off)
		time.Sleep(1 * time.Second)

		// 离开1人
		err2 := room.GetOut(-1)
		convey.So(err2, convey.ShouldEqual, nil)
		counter2 := room.GetCounter()
		convey.So(counter2, convey.ShouldEqual, 0)
		status3 := room.GetLightStatus()
		convey.So(status3, convey.ShouldEqual, light.Off)
	})
}

func TestRoom_ComeIn2(t *testing.T) {
	convey.Convey("测试进入房间后，开灯后立即离开，同时又有人进来的情况下，灯的状态", t, func() {
		peo := peoples.NewPeoples()
		li := light.NewLight(2*time.Second, 2*time.Second)
		room := NewRoom(peo, peo, li, li)
		// 有1人进来
		err := room.ComeIn(1)
		convey.So(err, convey.ShouldEqual, nil)
		status := room.GetLightStatus()
		convey.So(status, convey.ShouldEqual, light.Off)
		time.Sleep(2 * time.Second)
		status2 := room.GetLightStatus()
		convey.So(status2, convey.ShouldEqual, light.On)
		counter := room.GetCounter()
		convey.So(counter, convey.ShouldEqual, 1)

		// 离开1人
		err2 := room.GetOut(-1)
		convey.So(err2, convey.ShouldEqual, nil)
		counter2 := room.GetCounter()
		convey.So(counter2, convey.ShouldEqual, 0)
		status3 := room.GetLightStatus()
		convey.So(status3, convey.ShouldEqual, light.On)
		// 进来2人
		err3 := room.ComeIn(2)
		convey.So(err3, convey.ShouldEqual, nil)
		counter3 := room.GetCounter()
		convey.So(counter3, convey.ShouldEqual, 2)
		status4 := room.GetLightStatus()
		convey.So(status4, convey.ShouldEqual, light.On)
		// 等待2s
		time.Sleep(2 * time.Second)
		status5 := room.GetLightStatus()
		convey.So(status5, convey.ShouldEqual, light.On)
	})
}
