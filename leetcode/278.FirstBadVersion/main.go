package main

import "fmt"

type badVersion struct {
	bad []int
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
func (b *badVersion) isBadVersion(version int) bool {
	for _, bad := range b.bad {
		if version == bad {
			return true
		}
	}
	return false
}

func setBadVersions(start, end int) *badVersion {
	b := &badVersion{}
	for i := start; i <= end; i++ {
		b.bad = append(b.bad, i)
	}
	return b
}

func (b *badVersion) binarySearch(l, r int) int {
	if l <= r {
		mid := l + (r-l)/2
		fmt.Println(mid)

		//first bad version is mid
		if !b.isBadVersion(mid) && b.isBadVersion(mid+1) {
			return mid + 1
		}
		//first bad version is before mid
		if b.isBadVersion(mid) {
			return b.binarySearch(l, mid-1)
		}

		//first bad version is after mid
		return b.binarySearch(mid+1, r)
	}
	return l
}

func (b *badVersion) firstBadVersion(n int) int {
	return b.binarySearch(0, n)
}

func main() {
	b := setBadVersions(33, 34)
	fmt.Println(b.bad)
	fmt.Println(b.firstBadVersion(100))
}
