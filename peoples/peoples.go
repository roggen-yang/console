package peoples

// 人行为接口
type IMover interface {
	// 进入房间，n为进入的人数
	ComeIn(n uint32)
	// 离开房间，n 为离开的人数
	GetOut(n uint32)
}

// 房间里面人数查看接口
type ICounter interface {
	GetCounter() uint32
}

// 计数器
type Peoples struct {
	// counter：表示房间里面人数，>0 表示有人， = 0 表示没人
	counter uint32
}

// Counter构造函数
func NewPeoples() *Peoples {
	return &Peoples{}
}

// 实现有人进入房间的逻辑
func (c *Peoples) ComeIn(n uint32) {
	c.counter += n
}

// 实现有人离开房间逻辑
func (c *Peoples) GetOut(n uint32) {
	c.counter -= n
}

// 获取当前房间人数
func (c *Peoples) GetCounter() uint32 {
	return c.counter
}
