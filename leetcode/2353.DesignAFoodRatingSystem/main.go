package main

import "container/heap"

type FoodEntry struct {
	food   string
	rating int
}

type MaxHeap []FoodEntry

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].rating != h[j].rating {
		return h[i].rating > h[j].rating
	}
	return h[i].food < h[j].food
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(FoodEntry))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineHeaps  map[string]*MaxHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineHeaps:  make(map[string]*MaxHeap),
	}

	for i := range foods {
		fr.foodToCuisine[foods[i]] = cuisines[i]
		fr.foodToRating[foods[i]] = ratings[i]

		if _, exists := fr.cuisineHeaps[cuisines[i]]; !exists {
			fr.cuisineHeaps[cuisines[i]] = &MaxHeap{}
			heap.Init(fr.cuisineHeaps[cuisines[i]])
		}

		heap.Push(fr.cuisineHeaps[cuisines[i]], FoodEntry{foods[i], ratings[i]})
	}

	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := this.foodToCuisine[food]
	this.foodToRating[food] = newRating
	heap.Push(this.cuisineHeaps[cuisine], FoodEntry{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.cuisineHeaps[cuisine]

	for h.Len() > 0 {
		entry := (*h)[0]
		if this.foodToRating[entry.food] == entry.rating {
			return entry.food
		}
		heap.Pop(h) // Remove outdated entry
	}
	return ""
}
