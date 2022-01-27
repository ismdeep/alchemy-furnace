package config

var Watchers []chan bool

func init() {
	Watchers = make([]chan bool, 0)
}

func GenerateWatcher() chan bool {
	w := make(chan bool)
	Watchers = append(Watchers, w)
	return w
}
