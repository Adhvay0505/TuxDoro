package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

func sendNotification(title, message string) {
    exec.Command("notify-send", title, message).Run()
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter work duration in minutes: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    minutes, err := strconv.Atoi(input)
    if err != nil || minutes <= 0 {
        fmt.Println("Invalid input. Using default 25 minutes.")
        minutes = 25
    }

    // Notify start
    sendNotification("Pomodoro Timer", fmt.Sprintf("Starting %d-minute work session. Focus!", minutes))
    fmt.Printf("Pomodoro started for %d minutes. Focus!\n", minutes)

    totalSeconds := minutes * 60
    barLength := 30 // Length of the progress bar in characters

    for elapsed := 0; elapsed <= totalSeconds; elapsed++ {
        percent := float64(elapsed) / float64(totalSeconds)
        filled := int(percent * float64(barLength))
        empty := barLength - filled
        fmt.Printf("\r[%s%s] %3.0f%%", strings.Repeat("=", filled), strings.Repeat(" ", empty), percent*100)
        time.Sleep(1 * time.Second)
    }

    // Notify end
    sendNotification("Pomodoro Timer", "Time's up! Take a break.")
    fmt.Println("\nTime's up! Take a break.")
}

