package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	titlePtr := flag.String("title", "Notification", "The title to display in the notification")
	bodyPtr := flag.String("body", "Consider yourself notified", "The body of the notification")
	flag.Parse()

	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, *bodyPtr, *titlePtr))

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
