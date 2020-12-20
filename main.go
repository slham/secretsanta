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

type node struct {
	name string
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) prettyPrint() []byte {
	out := make([]byte, 0)

	toPrint := l.head
	for l.length != 0 {
		log.Printf("name: %s", toPrint.name)
		out = append(out, fmt.Sprintf("name: %s", toPrint.name)...)
		out = append(out, "\n"...)

		toPrint = toPrint.next
		l.length--
	}

	return out
}

func (l linkedList) extractValues(delimiter string) []byte {
	out, delimByte := make([]byte, 0), []byte(delimiter)
	toExtract := l.head

	for l.length != 0 {
		out = append(out, toExtract.name...)
		if l.length > 1 {
			out = append(out, delimByte...)
		}
		toExtract = toExtract.next
		l.length--
	}

	return out
}

func (l *linkedList) popWithValue(value string) *node {
	if l.length == 0 {
		return nil
	}

	if l.head.name == value {
		out := l.head
		l.head = l.head.next
		l.length--
		return out
	}

	prev := l.head
	for prev.next != nil {
		if prev.next.name == value {
			out := prev.next
			prev.next = prev.next.next
			l.length--
			return out
		}
		prev = prev.next
	}

	return nil
}

func shuffleSantas(l *linkedList, r *rand.Rand, names []string) {
	total := len(names)
	for i := 0; i < total; i++ {
		spot := r.Intn(total)
		name := names[spot]
		n := l.popWithValue(name)
		l.prepend(n)
	}
}

func main() {
	fileName := flag.String("names", "names.txt", "location of file containing comma delimited names for secret santa. default location is 'names.txt'.")
	debug := flag.Bool("debug", false, "indicates running binary in debug mode")
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

	santas := &linkedList{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	for _, name := range names {
		santas.prepend(&node{name: name})
	}

	if *debug {
		log.Println(string(santas.prettyPrint()))
	}

	shuffleSantas(santas, r, names)

	if *debug {
		log.Println(string(santas.prettyPrint()))
	}

	extractedytes := santas.extractValues(",")

	if *debug {
		log.Println(string(extractedytes))
	}

	err = ioutil.WriteFile("out.txt", extractedytes, 0644)
	if err != nil {
		log.Fatal("could not output results", err.Error())
		return
	}

	log.Println("done!")
}
