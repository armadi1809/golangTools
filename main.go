package main

func main() {

	watcher := &Watcher{directoryToWatch: "/Users/azizrmadi/Downloads/", destinationDirectory: "/Users/azizrmadi/golangTools/"}
	watcher.Watch()
}
