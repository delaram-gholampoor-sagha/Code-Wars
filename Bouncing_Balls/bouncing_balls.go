package kata

func BouncingBall(h, bounce, window float64) int {
	// we have to return -1 if non of these conditions are true
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	// the height of the bounce of a ball
	bounceHight := h * bounce
	// number of times that the mother can see the ball
	count := 1

	// if bounceheight was greater than the window parameter then the mother can see the ball
	for i := 0; bounceHight > window; i++ {
		count += 2
		bounceHight = bounceHight * bounce
	}
	return count
}
