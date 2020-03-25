package room

import (
	"errors"
	"fmt"

	"github.com/roggen-yang/console/peoples"

	"github.com/roggen-yang/console/light"
)

// 上下文封装类
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
func (c *Room) ComeIn(n int) error {
	if n <= 0 {
		return errors.New("come in number should greater than 0")
	}
	c.move.ComeIn(uint32(n))
	c.light.LightDelay(light.On)
	return nil
}

// 有人离开房间
func (c *Room) GetOut(n int) error {
	if n >= 0 {
		return errors.New("come in number should less than 0")
	}
	if c.counter.GetCounter() < uint32(-n) {
		return errors.New(
			fmt.Sprintf("out number should less than or equal %d",
				c.counter.GetCounter()))
	}
	c.move.GetOut(uint32(-n))
	c.light.LightDelay(light.Off)
	return nil
}

// 获取灯状态
func (c *Room) GetLightStatus() light.LightStatus {
	return c.status.GetLightStatus()
}

// 获取房间人数
func (c *Room) GetCounter() uint32 {
	return c.counter.GetCounter()
}
