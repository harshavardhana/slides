func lockSameResource() {
	// Create distributed mutex to protect resource 'test'
	dm := dsync.NewDRWMutex("test")

	dm.Lock()
	log.Println("first lock granted")

	go func() {
		time.Sleep(5 * time.Second)
		log.Println("first lock unlocked")
		dm.Unlock() // Unblock first lock
	}()

	log.Println("about to lock same resource again...")
	dm.Lock() // Waits here.
	log.Println("second lock granted")

	time.Sleep(2 * time.Second)
	dm.Unlock()
}
