// Bad Filter
package main

// Bad Filter
type BadFilter struct {
	//
}

// FilterByColor
func (f *BadFilter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// FilterBySize
func (f *BadFilter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

// FilterBySizeAndColor
func (f *BadFilter) FilterBySizeAndColor(
	products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}
