package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var encoded = base64.StdEncoding.EncodeToString([]byte("Sandy Dwi Handoko Trapsilo"))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	csvString := "sandy,dwi,handoko,trapsilo\n" +
		"budi,pratama,pratama\n" +
		"joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"sandy", "dwi", "handoko"})
	_ = writer.Write([]string{"budi", "pratama", "pratama"})
	_ = writer.Write([]string{"joko", "morro", "diah"})
	writer.Flush()
}
