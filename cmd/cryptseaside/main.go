package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rhymof/cryptseaside"
	"github.com/rhymof/cryptseaside/day"
	"github.com/rhymof/cryptseaside/evening"
	"github.com/rhymof/cryptseaside/night"
)

func main() {
	flagVersion := flag.Bool("v", true, "stdout version")
	flagSeed := flag.Bool("s", true, "stdout seed (formatted RFC3339Nano)")
	flagImage := flag.Bool("i", true, "save as png. needs a directory named 'images'")
	flagURI := flag.Bool("u", true, "stdout URI")
	flag.Parse()

	if *flagVersion {
		fmt.Println(cryptseaside.Version)
	}
	t := time.Now()
	title := t.Format(time.RFC3339Nano)
	if *flagSeed {
		fmt.Println(title)
	}
	title = strings.ReplaceAll(title, ":", "_")
	title = strings.ReplaceAll(title, ".", "_")
	title = strings.ReplaceAll(title, "-", "_")
	title = strings.ReplaceAll(title, "+", "_")

	nano := t.UnixNano()

	rgba := image.NewRGBA(image.Rect(0, 0, 256, 256))

	seed := (nano & 255)
	switch {
	case seed < 100:
		rgba = night.Night(rgba, nano)
	case seed < 200:
		rgba = evening.Sunset(rgba, nano)
	default:
		rgba = day.Bluesky(rgba, nano)
	}

	if *flagImage {
		f, err := os.Create("./images/" + title + ".png")
		if err != nil {
			log.Println(err)
			os.Exit(6)
		}
		if err := png.Encode(f, rgba); err != nil {
			if err != nil {
				log.Println(err)
				os.Exit(7)
			}
		}
	}

	var uri string
	if *flagURI {
		uri = generateURI(rgba, title)
	}
	fmt.Println(uri)
}

func generateURI(rgba image.Image, title string) string {
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	if err := png.Encode(buf, rgba); err != nil {
		log.Println(err)
		os.Exit(4)
	}
	head := "data:image/jpeg;base64,"
	b, err := ioutil.ReadAll(buf)
	if err != nil {
		log.Println(err)
		os.Exit(5)
	}
	s := base64.StdEncoding.EncodeToString(b)
	return head + s
}
