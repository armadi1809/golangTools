package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type Watcher struct {
	directoryToWatch     string
	destinationDirectory string
}

func (w *Watcher) Watch() {
	initialStat, err := os.Stat(w.directoryToWatch)
	if err != nil {
		log.Fatal("Could not get the initial stats from directory to watch")
	}

	fmt.Println(initialStat.ModTime())
	file, err := os.Open(w.directoryToWatch)
	if err != nil {
		log.Fatal("Could not open the downloads directory")
	}

	defer file.Close()
	for {
		allFiles, err := file.ReadDir(-1)
		if err != nil {
			log.Fatal("Could not read the directory")
		}

		sort.Slice(allFiles, func(i, j int) bool {
			stat, err := os.Stat(w.directoryToWatch + allFiles[i].Name())
			if err != nil {
				log.Fatal("Can't read stats of file")
			}
			timeFirst := stat.ModTime()

			stat, err = os.Stat(w.directoryToWatch + allFiles[j].Name())
			if err != nil {
				log.Fatal("Can't read stats of file")
			}
			timeSecond := stat.ModTime()

			return timeFirst.After(timeSecond)
		})
		newStat, err := os.Stat(w.directoryToWatch)
		if newStat.ModTime().After(initialStat.ModTime()) {
			fileToMove := allFiles[0]
			os.Rename(w.directoryToWatch+fileToMove.Name(), w.destinationDirectory+fileToMove.Name())
			fmt.Println("Here Now moving file", fileToMove)
			initialStat, err = os.Stat(w.directoryToWatch)
		} else {
			initialStat = newStat
		}

		if _, err := file.Seek(0, 0); err != nil {
			panic(err)
		}
		time.Sleep(10 * time.Second)

	}
}
