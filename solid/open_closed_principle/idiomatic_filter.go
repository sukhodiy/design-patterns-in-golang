// Idiomatic filter
package main

// Idiomatic filter
type IdiomaticFilter struct{}

// Specification function
type SpecFunc func(p *Product) bool

// Filter
func (f *IdiomaticFilter) Filter(
	products []Product, spec SpecFunc,
) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

// IdiomaticFuncFilter
func IdiomaticFuncFilter[T any](ss []T, test func(*T) bool) (ret []T) {
	for _, s := range ss {
		if test(&s) {
			ret = append(ret, s)
		}
	}
	return
}
