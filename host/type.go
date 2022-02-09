package host

type Type struct {
	Name string
}

type TypeDependency map[string][]Type

func NewTypeDependency(deps [][]string) TypeDependency {
	res := make(map[string][]Type)
	for _, dep := range deps {
		key := dep[len(dep)-1]
		res[key] = make([]Type, 0, len(dep))
		for _, d := range dep {
			res[key] = append(res[key], Type{d})
		}
	}

	return res
}

func (t TypeDependency) GetDependencyLine(typeName string) []Type {
	return t[typeName]
}
