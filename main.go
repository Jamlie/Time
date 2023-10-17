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

    start := time.Now().UnixMilli()
    cmd := exec.Command(args, cmdArgs...)
    out, err := cmd.CombinedOutput()
    end := time.Now().UnixMilli()

    if err != nil {
        if len(out) > 0 {
            fmt.Printf("%s\n", out)
            fmt.Printf("Time: %.3f\n", float64(end - start) / 1000)
            return
        }
        fmt.Printf("Error: %s\n", err)
        return
    }

    fmt.Printf("Time: %.3f\n", float64(end - start) / 1000)
}
