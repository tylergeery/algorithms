package main

import (
  "bufio"
  "fmt"
  "strconv"
  "os"
  "strings"
)

// basic packet obj
type Packet struct {
    processTime int
    finishTime int
}

// basic buffer obj
type Buffer struct {
    size int
    capacity int
    packets []Packet
    output []Packet
}


type Queue struct {
    currentPackets []Packet
    children []int
}

//
// Queue helper methods
//

// check if queue is empty
func isEmpty(packets []Packet) bool {
    if len(packets) > 0 {
        return false
    }

    return true
}

// Add packet to back of queue
func pushBack(packets []Packet, newPacket Packet) []Packet {
    return append(packets, newPacket)
}

// Remove oldest element from queue
func popFront(packets []Packet) (Packet, []Packet) {
    value, packets := packets[0], packets[1:]

    return value, packets
}

//
// Buffer helper methods
//

// check if buffer is at capacity
func bufferFull(buff *Buffer) bool {
    return buff.size == buff.capacity
}

// Add to the buffer
func bufferAdd(buff *Buffer, p Packet) *Buffer {
    // add to packets queue
    buff.packets = append(buff.packets, p)
    // add to output list
    buff.output = append(buff.output, p)

    return buff
}

// Ignore adding to Buffer, but keep for output
func bufferIgnore(buff *Buffer, p Packet) *Buffer {
    // packet is ignored
    p.processTime = -1
    // add to output list
    buff.output = append(buff.output, p)

    return buff
}

// Process completed packets that can be removed from packets queue
func processCompletedPackets(buff *Buffer, currentTime int) *Buffer {
    for (!isEmpty(buff.packets) && (buff.packets[0].finishTime <= currentTime)) {
        _, buff.packets = popFront(buff.packets)
        buff.size--
    }

    return buff
}

//
// Main script logic
//

func main() {
    // variables needed to parse file
    buff := new(Buffer)

    // open file and defer closure
    file, err := os.Open("/tmp/processpackages.txt")
    if err != nil {
        fmt.Println("Error: %s", err)
    }
    defer file.Close()

    // parse file
    // bufio for reading line by line
    index := 0
    currentTime := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // scanner.Text() has value
        txt := scanner.Text()
        // split on spaces
        props := strings.Fields(txt)

        if index == 0 {
            // first line contains buffer capacity and package count
            buff.size = 0
            buff.capacity, _ = strconv.Atoi(props[0])
        } else {
            // get the new current packet info
            currentTime, _ = strconv.Atoi(props[0])
            processTime, _ := strconv.Atoi(props[1])
            packet := Packet{processTime: currentTime, finishTime: (currentTime + processTime)}

            // process any completed packets
            buff = processCompletedPackets(buff, currentTime)

            // decide if buffer can handle next variable
            if bufferFull(buff) {
                buff = bufferIgnore(buff, packet)
            } else {
                buff = bufferAdd(buff, packet)
                buff.size++
            }
        }

        index++
    }

    // loop through buffer output for results
    for _, output := range buff.output {
        fmt.Println(output.processTime)
    }
}
