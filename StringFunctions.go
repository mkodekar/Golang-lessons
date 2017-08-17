package main

import ("fmt"
       "strings"
        "sort"
        "os"
        "log"
        "io/ioutil"
/*       "math"
       "sort"
       "os"
       "io/ioutil"
       "strconv"*/)

func main() {
    
    sampleString := "Hello world"
    
    fmt.Println(strings.Contains(sampleString, "lo"))
    fmt.Println(strings.Index(sampleString, "l"))
    fmt.Println(strings.Count(sampleString, "l"))
    fmt.Println(strings.Replace(sampleString, "l", "x", 3))
    
    csvString := "1,2,3,4,5,6"
    fmt.Println(strings.Split(csvString, ","))
    
    listOfLetters := []string{"c", "a", "b"}
    sort.Strings(listOfLetters)
    
    fmt.Println("Letters: ", listOfLetters)
    
    listOfNums := strings.Join([]string{"3", "2", "1"}, ",")
    
    fmt.Println(listOfNums)
    
    file, err := os.Create("sample.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    
    
    file.WriteString(" This is some random text")
    file.Close()
    
    stream, err := ioutil.ReadFile("sample.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    
    readString := string(stream)
    fmt.Println(readString)
}