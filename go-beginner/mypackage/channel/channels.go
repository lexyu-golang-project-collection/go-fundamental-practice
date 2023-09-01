package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}
func Channels() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	g.PL(<-channel1)
	g.PL(<-channel1)
	g.PL(<-channel1)
	g.PL(<-channel2)
	g.PL(<-channel2)
	g.PL(<-channel2)
}
