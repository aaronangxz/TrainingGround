package main

type Router struct {
	memoryLimit int
	history     [][3]int
	packets     map[[3]int]bool
	destPackets map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit: memoryLimit,
		history:     [][3]int{},
		packets:     map[[3]int]bool{},
		destPackets: map[int][]int{},
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := [3]int{source, destination, timestamp}
	if _, ok := r.packets[packet]; ok {
		return false
	}

	if len(r.packets) == r.memoryLimit {
		_ = r.ForwardPacket()
	}

	r.packets[packet] = true
	r.history = append(r.history, packet)
	r.destPackets[destination] = append(r.destPackets[destination], timestamp)
	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.history) == 0 {
		return nil
	}

	packet := r.history[0]
	r.history = r.history[1:]
	delete(r.packets, packet)
	r.destPackets[packet[1]] = r.destPackets[packet[1]][1:]
	return []int{packet[0], packet[1], packet[2]}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	timestamps := r.destPackets[destination]
	if len(timestamps) == 0 {
		return 0
	}
	left, right := 0, len(timestamps)-1
	for left < right {
		mid := (left + right) / 2
		if timestamps[mid] < startTime {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if timestamps[left] < startTime {
		return 0
	}

	left2, right2 := left, len(timestamps)-1
	for left2 < right2 {
		mid := (left2 + right2 + 1) / 2
		if timestamps[mid] > endTime {
			right2 = mid - 1
		} else {
			left2 = mid
		}
	}

	if timestamps[left2] > endTime {
		return 0
	}

	return left2 - left + 1
}
