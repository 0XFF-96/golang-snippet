package channel

// close stream
// 如何避免被重复 close ?
func (s *sendLockServerStream) Close() {
	if _, isClose := <-s.DoneCh; isClose {
		return
	}
	close(s.DoneCh)
}
