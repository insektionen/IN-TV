package viewmodels

type Slide struct {
	Timeout  int         `json:"timeout"`
	Type     string      `json:"type"`
	Position int         `json:"position"`
	Data     interface{} `json:"data"`
}

type Slideshow struct {
	Name   string   `json:"name"`
	Slides []*Slide `json:"slides"`
}

type Status struct {
	RunningSlideshows []string  `json:"running_slideshows"`
	ConnectedScreens  []*Screen `json:"connected_screens"`
}

type Screen struct {
	Name    string `json:"name"`
	LasSeen int64  `json:"las_seen"`
}

type Register struct {
	Name string `json:"name"`
}

type Start struct {
	SlideshowName string   `json:"slideshow_name"`
	ScreenNames   []string `json:"screen_names"`
}

// The struct for the simplified json data from SL API
type SLData struct {
	Name              string `json:"name"`
	Time              string `json:"time"`
	Date              string `json:"date"`
	Direction         string `json:"direction"`
	RTTime            string `json:"rtTime"`
	TimeTillDeparture int    `json:"timeTillDeparture"`
}