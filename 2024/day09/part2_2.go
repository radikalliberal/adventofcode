// part2.go
package day09

import (
    "container/list"
    "fmt"
)

type file struct {
	id    int
	size  int
	free  bool
	tried bool
}

func ParseInputPart2(input string) *list.List {
    files := list.New()
    for i, c := range input {
        if i % 2 == 1 {
            files.PushBack(file{id: -1, size: int(c - '0'), free: true, tried: false})
        } else {
            files.PushBack(file{id: i / 2, size: int(c - '0'), free: false, tried: false})
        }
    }
    return files
}

func PrintFiles(files *list.List) {
    printString := ""
    for e := files.Front(); e != nil; e = e.Next() {
        f := e.Value.(file)
        for i := 0; i < f.size; i++ {
            if f.free {
                printString += "."
            } else {
                printString += fmt.Sprintf("%d", f.id)
            }
        }
    }
    fmt.Println(printString)
}


func ComputeChecksumPart2(files *list.List) int {
    sum := 0
    multiplier := 0
    for e := files.Front(); e != nil; e = e.Next() {
        f := e.Value.(file)
        if f.free {
            multiplier += f.size
        } else {
            for i := 0; i < f.size; i++ {
                sum += multiplier * f.id
                multiplier++
            }
        }
    }
    return sum
}

func SizeOfFiles(files *list.List) int {
    sum := 0
    for e := files.Front(); e != nil; e = e.Next() {
        f := e.Value.(file)
        sum += f.size
    }
    return sum
}

func CompactFiles_(files *list.List) *list.List {
    initialSize := SizeOfFiles(files)
    for {
        for e := files.Back(); e != nil; e = e.Prev() {
            f := e.Value.(file)
            if !f.free && !f.tried {
                f.tried = true
                for e2 := files.Front(); e2 != e && e2 != nil; e2 = e2.Next() {
                    f2 := e2.Value.(file)
                    if f2.free && f2.size >= f.size {
                        // PrintFiles(files)
                        files.InsertBefore(f, e2)
                        files.InsertBefore(file{id: -1, size: f.size, free: true, tried: false}, e)
                        files.Remove(e)
                        if f2.size > f.size {
                            e2.Value = file{id: -1, size: f2.size - f.size, free: true, tried: false}
                        } else {
                            files.Remove(e2)
                        }
                        if SizeOfFiles(files) != initialSize {
                            PrintFiles(files)
                            panic("Size of files changed")
                        }
                        break
                    }
                }
            }
            if e == files.Front() {
                return files
            }
        }
    }
}

func Part2_2(input []string) int {
    filesystem := ParseInputPart2(input[0])
    files := CompactFiles_(filesystem)
    // PrintFiles(files)
    checksum := ComputeChecksumPart2(files)
    return checksum
}
