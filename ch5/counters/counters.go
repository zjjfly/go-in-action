package counters

type alertCounter int

//把工厂函数命名为New是一个好的实践
func New(value int) alertCounter {
	return alertCounter(value)
}
