package memory

type Memory struct {
	F     Function
	Cache map[int]FunctionResult
}

type Function func(cache *Memory, key int) (interface{}, error)

type FunctionResult struct {
	Value interface{}
	Err   error
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.Cache[key]
	if !exists {
		result.Value, result.Err = m.F(m, key)
		m.Cache[key] = result
	}
	return result.Value, result.Err
}
