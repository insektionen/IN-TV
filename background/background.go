// Package background provides job management for background tasks. It also listens
// for OS Signals and terminates all background tasks when signal is received.
package background

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type jobChan struct {
	name string
	exit chan bool
	done chan bool
}

var (
	wg      = sync.WaitGroup{}
	jobList []*jobChan
)

// Job is a function that will receive data on exit when it should be terminated.
type Job func(exit <-chan bool)

// AddJob add a Job to the background monitor
func AddJob(name string, j Job) {
	wg.Add(1)
	log.Printf("Adding background job %q\n", name)
	job := &jobChan{
		name: name,
		exit: make(chan bool, 1),
		done: make(chan bool, 1),
	}
	jobList = append(jobList, job)

	go func() {
		defer wg.Done()
		j(job.exit)
		log.Printf("Done with background job %q\n", name)
		job.done <- true
	}()
}

// Wait blocks until all background jobs are completed
func Wait() {
	wg.Wait()
}

func signalHandler() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigChannel
	for i := len(jobList) - 1; i >= 0; i-- {
		j := jobList[i]
		select {
		case j.exit <- true:
			log.Printf("Stopping background job %q\n", j.name)
			close(j.exit)
			<-j.done
			close(j.done)
		default:
		}
	}
}

//nolint:gochecknoinits // Always start the background signal handler
func init() {
	go signalHandler()
}

