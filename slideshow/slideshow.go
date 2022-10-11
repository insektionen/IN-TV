package slideshow

import (
	"encoding/json"
	"fmt"
	"github.com/insektionen/IN-TV/mqtt"
	"github.com/insektionen/IN-TV/v1/viewmodels"
	"log"
	"time"
)

var RunningSlideshows = map[string]*Show{}

type Show struct {
	Slideshow *viewmodels.Slideshow
	tick      *time.Ticker
	current   int
	stop      chan bool
}

func NewShow(sl *viewmodels.Slideshow) *Show {
	return &Show{
		Slideshow: sl,
	}
}

func (s *Show) Stop() {
	s.stop <- true
	delete(RunningSlideshows, s.Slideshow.Name)
}

func (s *Show) Run() {
	if len(s.Slideshow.Slides) == 0 {
		return
	}
	if _, ok := RunningSlideshows[s.Slideshow.Name]; ok {
		return
	}
	RunningSlideshows[s.Slideshow.Name] = s
	t := s.Slideshow.Slides[0].Timeout
	s.stop = make(chan bool)
	s.current = 0
	s.tick = time.NewTicker(time.Duration(t) * time.Second)

	go func() {
		for {
			select {
			case <-s.stop:
				return
			case <-s.tick.C:
				next := s.current + 1
				log.Printf("Next screen for \"%s\" is %d\n", s.Slideshow.Name, next)
				if next >= len(s.Slideshow.Slides) {
					next = 0
				}

				message := map[string]int{
					"current": s.current,
					"next":    next,
				}
				data, _ := json.Marshal(message)
				topic := fmt.Sprintf("kistan/in_tv/slideshow/%s/change", s.Slideshow.Name)
				mqtt.Client.Publish(topic, 0, false, data)

				s.current = next
				s.tick.Reset(time.Duration(s.Slideshow.Slides[next].Timeout) * time.Second)
			}
		}
	}()
}
