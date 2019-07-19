package websocket

// Worker will process any job in parallel
func Worker(id int, jobs <-chan func()) {
	for j := range jobs {
		j()
	}
}
