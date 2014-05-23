package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type Result struct{ time float64 }

func upload(url, file string) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	fw, err := w.CreateFormFile("image", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	// Add the other fields
	if fw, err = w.CreateFormFile("key", file); err != nil {
		return
	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		return
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("PUT", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func lauchOneRocket(i int, file string) {
	t0 := time.Now()
	err := upload("http://127.0.0.1:9090", file)
	t1 := time.Now()
	if err == nil {
		fmt.Printf("%d rocket launched from %v\n", i, file)
		fmt.Printf("%v time\n", t1.Sub(t0))
	} else {
		fmt.Printf("%d rocket failed from %v with err = %v\n", i, file, err)
	}
}

func main() {
	file, err := os.Open("rockets.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i = i + 1
		if i < 1000 {
			go lauchOneRocket(i, scanner.Text())
		}
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
