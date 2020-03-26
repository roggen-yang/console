package light

import (
	"time"
)

// 定义等的状态
type LightStatus int

const (
	// 表示关灯状态
	Off LightStatus = iota
	// 表示开灯状态
	On
)

// 灯操作接口
type ILight interface {
	// 开灯操作
	LightOn()
	// 关灯操作
	LightOff()
	// 延时操作
	LightDelay(opt LightStatus)
}

// 状态查看接口
type IStatus interface {
	// 获取当前灯的状态
	GetLightStatus() LightStatus
}

// 灯结构
type Light struct {
	// 灯当前状态
	status LightStatus
	// 进入房间后的延迟时间
	comeDelay time.Duration
	// 离开房间后的延迟时间
	outDelay time.Duration
	// 下一次开灯时间
	nextOn *time.Time
	// 下一次关灯时间
	nextOff *time.Time
	// 是否处于开灯延时中
	onDelaying bool
	// 是否处于关灯延时中
	offDelaying bool
}

// 构造函数
func NewLight(cd, od time.Duration) *Light {
	l := &Light{
		status:      Off,
		comeDelay:   cd,
		outDelay:    od,
		onDelaying:  false,
		offDelaying: false,
	}
	go l.lightOnSync()
	go l.lightOffSync()
	return l
}

// 实现开灯逻辑
func (l *Light) LightOn() {
	l.status = On
	l.nextOn = nil
	l.onDelaying = false
}

// 实现关灯逻辑
func (l *Light) LightOff() {
	l.status = Off
	l.nextOff = nil
	l.offDelaying = false
}

// 延时操作
func (l *Light) LightDelay(opt LightStatus) {
	switch opt {
	case On:
		l.setNextLightOnTime(&l.comeDelay)
		l.onDelaying = true
		l.offDelaying = false
	case Off:
		l.setNextLightOffTime(&l.outDelay)
		l.offDelaying = true
		l.onDelaying = false
	}
}

// 获取当前灯的状态
func (l *Light) GetLightStatus() LightStatus {
	return l.status
}

// 设置下次开灯时间
func (l *Light) setNextLightOnTime(ti *time.Duration) {
	if ti == nil {
		l.nextOn = nil
	} else {
		next := time.Now().Add(*ti)
		l.nextOn = &next
	}
}

// 设置下次关灯时间
func (l *Light) setNextLightOffTime(ti *time.Duration) {
	if ti == nil {
		l.nextOff = nil
	} else {
		next := time.Now().Add(*ti)
		l.nextOff = &next
	}
}

// 同步开灯操作
func (l *Light) lightOnSync() {
	for {
		if l.onDelaying {
			if time.Now().After(*l.nextOn) {
				l.LightOn()
			}
		}
	}
}

// 同步关灯动作
func (l *Light) lightOffSync() {
	for {
		if l.offDelaying {
			if time.Now().After(*l.nextOff) {
				l.LightOff()
			}
		}
	}
}
