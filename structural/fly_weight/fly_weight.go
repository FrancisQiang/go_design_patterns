package fly_weight

type FlyWeight struct {
	name string
}

type FlyWeightFactory struct {
	pool map[string]FlyWeight
}

func NewFlyWeightFactory() FlyWeightFactory {
	return FlyWeightFactory{
		pool: make(map[string]FlyWeight),
	}
}

func (factory FlyWeightFactory) GetFlyWeight(name string) FlyWeight {
	value, exist := factory.pool[name]
	if !exist {
		value = FlyWeight{
			name: name,
		}
		factory.pool[name] = value
	}
	return value
}
