package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Result struct{ time float64 }

func newfileUploadRequest(uri string, file string) (*http.Request, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(file))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, f)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func lauchOneRocket(i int, file string) {
	t0 := time.Now()

	request, err := newfileUploadRequest("http://127.0.0.1:9090", file)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		// fmt.Println(resp.StatusCode)
		// fmt.Println(resp.Header)
		// fmt.Println(body)
	}

	t1 := time.Now()
	if err == nil {
		//fmt.Printf("%d, rocket %v, time=%v\n", i, file, t1.Sub(t0))
		fmt.Printf("%d,%.6f\n", i, float64(t1.Sub(t0))/float64(time.Millisecond))
	} else {
		fmt.Printf("%d rocket failed from %v with err = %v\n", i, file, err)
	}
}

func lauchKatyusha(rockets int) {
	file, err := os.Open("rockets.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i = i + 1
		if i <= rockets {
			go lauchOneRocket(i, scanner.Text())
		}
	}
}

func main() {
	flag.Parse()
	rockets, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	lauchKatyusha(rockets)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
