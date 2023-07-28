package main

import "flag"

func main() {

	directoryToWatch := flag.String("watch", "", "The Directory To monitor changes on")
	destinationDirectory := flag.String("destination", "", "The Directory To move files to")

	flag.Parse()
	watcher := &Watcher{directoryToWatch: *directoryToWatch, destinationDirectory: *destinationDirectory}
	watcher.Watch()
}
