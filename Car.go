package main

// Car :
// Creates a car instance.
type Car struct {
	// id               int
	originPos        Point // Origin point
	currentPos       Point // Current point
	destinationPos   Point // Destination point
	velocity         float32
	cMap             *City
	currentSemaphore *Semaphore
}

func (c *Car) move() {

	// is my velocity enough to move?
	// is next position a semaphor?
	// is it green or red? -> green: go || red: STOP ->v:0
	// is someone in front of me?
	// is their velocity slower? yes -> v:0

	// move
}