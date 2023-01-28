package main

import (
    "fmt"
    "os/exec"
)

func main() {
    url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
    cmd := exec.Command("youtube-dl", "-f", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best", "-o", "output.%(ext)s", url)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error : ", err)
    } else {
        fmt.Println("Output : ", string(output))
    }
}
