package main

import (
    ipc "github.com/james-barrow/golang-ipc"
    "log"
    "os"
    "os/exec"
    "time"
)

func startClient(socketName string) {
    c, err := ipc.StartClient(socketName, &ipc.ClientConfig{
        Timeout:    0,
        RetryTimer: 2 * time.Second,
        Encryption: false,
    })
    if err != nil {
        println(err)
        return
    }
    for {
        message, _ := c.Read() // client

        if message.MsgType == -1 {
            if message.Status == "Reconnecting" {
                log.Println("No server connection.. terminating.")
                return
            }
            continue
        }

        switch message.MsgType {
        case 1:
            appendLogMessage(message.Data)
        case 2:
            displayCharacterStats(message.Data)
        case 3:
            displayInventory(message.Data)
        case 4:
            displayEquipment(message.Data)
        }
    }
}

func displayEquipment(data []byte) {
    clearScreen()
    colPrintln(string(data))
}

func appendLogMessage(data []byte) {
    colPrintln(string(data))
}

func displayCharacterStats(data []byte) {
    clearScreen()
    println(string(data))
}
func displayInventory(data []byte) {
    clearScreen()
    colPrintln(string(data))
}

func clearScreen() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}

func main() {
    // first command line arg
    socketName := "rpg.sock"
    if len(os.Args) > 1 {
        socketName = os.Args[1]
    } else {
        println("ERR: No message socket specified")
        return
    }
    clearScreen()

    startClient(socketName)
}
