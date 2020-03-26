package room

import (
	"errors"
	"fmt"

	"github.com/roggen-yang/console/peoples"

	"github.com/roggen-yang/console/light"
)

// 房间抽象类
type Room struct {
	move    peoples.IMover
	counter peoples.ICounter
	light   light.ILight
	status  light.IStatus
}

// 构造函数
func NewRoom(move peoples.IMover, counter peoples.ICounter, il light.ILight, is light.IStatus) *Room {
	return &Room{
		move:    move,
		counter: counter,
		light:   il,
		status:  is,
	}
}

// 有人进入房间
func (r *Room) ComeIn(n int) error {
	if n <= 0 {
		return errors.New("come in number should greater than 0")
	}
	if r.counter.GetCounter() == 0 {
		r.light.LightDelay(light.On)
	}
	r.move.ComeIn(uint32(n))
	return nil
}

// 有人离开房间
func (r *Room) GetOut(n int) error {
	if n >= 0 {
		return errors.New("come in number should less than 0")
	}
	if r.counter.GetCounter() < uint32(-n) {
		return errors.New(
			fmt.Sprintf("out number should less than or equal %d",
				r.counter.GetCounter()))
	}
	r.move.GetOut(uint32(-n))
	if r.counter.GetCounter() == 0 {
		r.light.LightDelay(light.Off)
	}
	return nil
}

// 获取灯状态
func (r *Room) GetLightStatus() light.LightStatus {
	return r.status.GetLightStatus()
}

// 获取房间人数
func (r *Room) GetCounter() uint32 {
	return r.counter.GetCounter()
}
