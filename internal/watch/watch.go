package watch

import (
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/kalokaradia/jspackr/internal/build"
)

func WatchFiles(entry string, opts build.Options) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					color.New(color.FgYellow).Println("🔄 File changed:", event.Name)
					err := build.Run(opts)
					if err != nil {
						color.New(color.FgRed).Println("❌ Build failed:", err)
					} else {
						color.New(color.FgGreen).Println("✅ Build finished")
					}
				}
			case err, ok := <-watcher.Errors:
				if ok {
					color.New(color.FgRed).Println("Watcher error:", err)
				}
			}
		}
	}()

	err = watcher.Add(entry)
	if err != nil {
		return err
	}

	<-done
	return nil
}
