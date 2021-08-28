package strategy

type CalculateStrategy interface {
	Do(a, b int) int
}

// a + b + a * b
type AStrategy struct {
}

func (aStrategy AStrategy) Do(a, b int) int {
	return a + b + a*b
}

// a + b + a * 2 + b * 3
type BStrategy struct {
}

func (bStrategy BStrategy) Do(a, b int) int {
	return a + b + a * 2 + b * 3
}

type Operator struct {
	strategy CalculateStrategy
}

func (operator Operator) SetStrategy(strategy CalculateStrategy) {
	operator.strategy = strategy
}

func (operator Operator) Calculate(a, b int) int {
	return operator.strategy.Do(a, b)
}
