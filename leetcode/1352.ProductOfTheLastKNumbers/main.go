package main

type ProductOfNumbers struct {
	prefixProducts []int // Cumulative product list
}

// Constructor initializes an empty product list
func Constructor() ProductOfNumbers {
	return ProductOfNumbers{prefixProducts: []int{}}
}

// Add inserts a number, updating prefix product logic
func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefixProducts = []int{} // Reset on zero
	} else {
		lastProduct := 1
		if len(this.prefixProducts) > 0 {
			lastProduct = this.prefixProducts[len(this.prefixProducts)-1]
		}
		this.prefixProducts = append(this.prefixProducts, num*lastProduct)
	}
}

// GetProduct returns the product of the last k numbers
func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefixProducts)
	if k > n {
		return 0 // Not enough numbers
	}
	if k == n {
		return this.prefixProducts[n-1] // Full list product
	}
	return this.prefixProducts[n-1] / this.prefixProducts[n-k-1] // Compute via division
}
