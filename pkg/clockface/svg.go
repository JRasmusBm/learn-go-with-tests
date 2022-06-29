package clockface

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

func tee[T any](val T) T {
	fmt.Printf("%v", val)

	return val
}

func (c *Clockface) WriteSVG(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, tee(c.bezel("#000")))
	io.WriteString(w, c.handTag(c.SecondHand(t), "#f00"))
	io.WriteString(w, c.handTag(c.MinuteHand(t), "#000"))
	io.WriteString(w, c.handTag(c.HourHand(t), "#000"))
	io.WriteString(w, svgEnd)
}

func (c *Clockface) handTag(p Point, color string) string {
	return fmt.Sprintf(`<line x1="%.00f" y1="%.00f" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`, c.origin, c.origin, p.X, p.Y, color)
}

func (c *Clockface) bezel(color string) string {
	return fmt.Sprintf(`<circle cx="%.00f" cy="%.00f" r="%.00f" style="fill:#fff;stroke:%s;stroke-width:5px;"/>`, c.origin, c.origin, c.scale, color)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const svgEnd = `</svg>`
