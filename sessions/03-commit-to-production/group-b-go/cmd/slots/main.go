// clinic-slots: lists bookable appointment slots.
//
//	go run ./cmd/slots -date 2026-06-01 -clinic Leeds
//	go run ./cmd/slots -health
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"brightcare/clinic-slots/internal/slots"
)

func main() {
	date := flag.String("date", time.Now().In(slots.London).Format("2006-01-02"), "day to list, YYYY-MM-DD")
	clinic := flag.String("clinic", "Leeds", "clinic name")
	health := flag.Bool("health", false, "print ok and exit (used by the smoke test)")
	flag.Parse()

	if *health {
		fmt.Println("ok")
		return
	}
	day, err := time.ParseInLocation("2006-01-02", *date, slots.London)
	if err != nil {
		fmt.Fprintln(os.Stderr, "date must look like 2026-06-01")
		os.Exit(2)
	}
	for _, s := range slots.ForDay(*clinic, day) {
		fmt.Println(s.Label())
	}
}
