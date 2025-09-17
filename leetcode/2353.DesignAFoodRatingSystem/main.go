package main

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineFoods  map[string][]string
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineFoods:  make(map[string][]string),
	}

	for i := range foods {
		fr.foodToCuisine[foods[i]] = cuisines[i]
		fr.foodToRating[foods[i]] = ratings[i]
		fr.cuisineFoods[cuisines[i]] = append(fr.cuisineFoods[cuisines[i]], foods[i])
	}

	return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	fr.foodToRating[food] = newRating
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	bestFood := ""
	bestRating := -1

	for _, food := range fr.cuisineFoods[cuisine] {
		rating := fr.foodToRating[food]
		if rating > bestRating || (rating == bestRating && food < bestFood) {
			bestRating = rating
			bestFood = food
		}
	}
	return bestFood
}
