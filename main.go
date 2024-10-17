package main

import (
	"github.com/go-co-op/gocron/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

const (
	url       = "amqp://guest:guest@rabbitmq:5672"
	xc        = "exchange"
	q         = "exclusive-queue"
	startup   = 10 * time.Second
	interval  = 5 * time.Second
	max       = 60 * time.Second
	lockSleep = 1 * time.Second
)

func main() {
	time.Sleep(startup)
	log.Print("starting app")

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("failed to connect: %s", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("terminating amqp connection")
	}()

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	// add a job to the scheduler
	_, err = s.NewJob(
		gocron.CronJob("0/5 * * * * *", true),
		gocron.NewTask(doWork, conn),
	)
	if err != nil {
		log.Fatal(err)
	}
	// each job has a unique id
	//log.Print(j.ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(max):
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
		log.Fatal(err)
	}
}

func doWork(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("failed to open a channel: %s", err)
		return
	}
	defer func() {
		err := ch.Close()
		if err != nil {
			log.Print(err)
			return
		}
	}()

	queue, err := ch.QueueDeclare(q, false, true, true, false, nil)
	if err != nil {
		log.Printf("i don't have the lock :(")
	} else {
		log.Printf("i have the lock :)")
		time.Sleep(lockSleep)
		_, err := ch.QueueDelete(queue.Name, false, false, false)
		if err != nil {
			log.Printf("trouble while deleting queue: %s", err)
		}
	}
}
