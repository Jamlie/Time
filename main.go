package main

import (
    "fmt"
    "time"
    "flag"
    "os/exec"
)

func main() {
    flag.Parse()
    argsArray := flag.Args()
    args := argsArray[0]
    cmdArgs := argsArray[1:]

    if len(cmdArgs) == 0 {
        cmdArgs = []string{""}
    }

    start := time.Now().UnixMilli()
    cmd := exec.Command(args, cmdArgs...)
    end := time.Now().UnixMilli()

    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }

    fmt.Printf("%s\n", out)
    fmt.Printf("Time: %.3f\n", float64(end - start) / 1000)
}