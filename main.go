package main

import (
	"fmt"
	"time"

	proofofbids "github.com/shajalahamedcse/newBid/proof-of-bids"
)

const biderationtimeInSeconds = 10 // this is the bidration time
const numberOfBiderator = 5        // number of user will perticipate in bideration
const softwarePackPower = 100      // it will deduct 20 bids/sec. we can make it dymamic later
const bids = 1000                  // number of bids biderator is using in bideration. will be dynamic
const bidsNeeded = 1300

func main() {

	fmt.Println("Lets build the linked list first")

	bl := proofofbids.BiderList{
		Head:           nil,
		Tail:           nil,
		Length:         0,
		PowerSummation: 0,
	}

	for i := 0; i < numberOfBiderator; i++ {

		bl.Append(bids, softwarePackPower, bidsNeeded)
	}

	bb := biderationtimeInSeconds
	fmt.Printf("Lets start the bideration process.%d\n", bl.Length)

	q := proofofbids.New(1)

	for i := 0; i < biderationtimeInSeconds; i++ {

		bb--
		<-time.After(1 * time.Second) // Will do this in every second
		q.Add()
		go func() {
			defer q.Done()
			bl.UpdatePerSecond(bb)
		}()
		q.Wait()
	}

}
