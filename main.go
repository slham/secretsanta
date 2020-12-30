package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func swapElements(arr []string, a, b, l int) {
	if a < 0 || a >= l || b < 0 || b >= l {
		return
	}

	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func shuffle(arr []string, r *rand.Rand) {
	total := len(arr)
	for i := 0; i < total; i++ {
		a, b := r.Intn(total), r.Intn(total)
		swapElements(arr, a, b, total)
	}
}

func main() {
	fileName := flag.String("names", "names.txt", "location of file containing comma delimited names for secret santa. default location is 'names.txt'.")
	flag.Parse()

	nameBytes, err := ioutil.ReadFile(*fileName)
	if err != nil {
		missing := fmt.Sprintf("open %s: no such file or directory", *fileName)
		if err.Error() == missing {
			log.Println("please create a file of comma delimited names. please name the file 'names.txt'")
		}
		log.Fatal(err.Error())
		return
	}

	nameString := string(nameBytes)
	names := strings.Split(strings.Trim(nameString, "\n\r"), ",")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	shuffle(names, r)

	out, delimBytes := make([]byte, 0), []byte(",")
	length := len(names)

	for i, name := range names {
		out = append(out, name...)
		if i < length-1 {
			out = append(out, delimBytes...)
		}
	}

	err = ioutil.WriteFile("out.txt", out, 0644)
	if err != nil {
		log.Fatal("could not output results", err.Error())
		return
	}

	log.Println("done!")
}
