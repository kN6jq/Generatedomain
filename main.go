package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var domainFile string
	var domain string
	flag.StringVar(&domainFile, "f", "subname.txt", "subnameFile")
	flag.StringVar(&domain, "d", "", "dev-FUZZ.website.com")
	flag.Parse()

	if !strings.Contains(domain, "FUZZ") {
		log.Fatal("Domain name requires FUZZ keyword, like dev-FUZZ.website.com")

	}

	f, err := os.Open(domainFile)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	//创建一个新文件
	filePath := domain + ".txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	for scanner.Scan() {
		// do something with a line
		fmt.Println(strings.Replace(domain, "FUZZ", scanner.Text(), -1))
		write.WriteString(strings.Replace(domain, "FUZZ", scanner.Text(), -1) + "\n")
	}
	write.Flush()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
