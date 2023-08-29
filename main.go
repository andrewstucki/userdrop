package main

/*
#include <stdlib.h>
#include <stdio.h>

void hello() {
  printf("%s\n", "Called into C function!");
}
*/
import "C"

import (
	"fmt"
	"log"
	"net/http"
	"os/user"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Google returned a '%d' status code.\n", resp.StatusCode)

	C.hello()

	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Running as: %s\n", current.Username)
}
