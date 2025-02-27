package main

type TimeMap struct {
	KeyMap map[string][]Pair
}

type Pair struct {
	val       string
	timeStamp int
}

func Constructor() TimeMap {
	return TimeMap{
		KeyMap: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	pairObj := Pair{
		val:       value,
		timeStamp: timestamp,
	}

	// If key is empty, init new Pair slice
	if _, ok := this.KeyMap[key]; !ok {
		this.KeyMap[key] = make([]Pair, 0)
	}
	this.KeyMap[key] = append(this.KeyMap[key], pairObj)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	// No key found, return empty
	if _, ok := this.KeyMap[key]; !ok {
		return ""
	}
	arr := this.KeyMap[key]
	// prev latest ts cannot be greater than the requested ts
	if arr[0].timeStamp > timestamp {
		return ""
	}

	// Binary Search
	left, right := 0, len(arr)-1
	res := ""
	for left <= right {
		mid := (left + right) / 2
		if arr[mid].timeStamp == timestamp {
			return arr[mid].val
		} else if arr[mid].timeStamp < timestamp {
			res = arr[mid].val
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
