package set

// Set 集合 主要用于去重
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet 创建集合
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// Add 添加元素
func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

// ToSlice 转为切片
func (s *Set[T]) ToSlide() []T {
	var slide []T
	for k := range s.m {
		slide = append(slide, k)
	}
	return slide
}
