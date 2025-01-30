package main

import (
	"fmt"
	"io"
	"learn-go-with-tests/v2/16-math/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	sh := clockface.SecondHand(t)

	var err error
	_, err = io.WriteString(os.Stdout, svgStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing clock start: %v", err)
	}

	_, err = io.WriteString(os.Stdout, bezel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing clock bezel: %v", err)
	}

	_, err = io.WriteString(os.Stdout, secondHandTag(sh))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing second hand: %v", err)
	}

	_, err = io.WriteString(os.Stdout, svgEnd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing clock end: %v", err)
	}
}

func secondHandTag(p clockface.Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
