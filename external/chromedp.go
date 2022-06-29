package external

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"omoshiroimg/external/gcp"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chromedp/chromedp"
)

type ImageGenerator interface {
	GenerateImage(string,string,string) string
}

type ChromedpImageGenerator struct {
}

func NewChromedpImageGenerator() ImageGenerator {
	return &ChromedpImageGenerator{}
}

func (g *ChromedpImageGenerator) GenerateImage(text string, event string, room string ) string {

	name := fmt.Sprintf("%d", time.Now().Unix())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	go screenShot(sigCh, name, text, event, room)
	<-sigCh

	return name
}

func screenShot(sigCh chan os.Signal, imageName string, text string, event string, room string ) {
	// create context
	defer func() {
		sigCh <- syscall.SIGTERM
	}()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.WindowSize(1500, 800),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	//if err := chromedp.Run(ctx, elementScreenshot(`https://pkg.go.dev/`, `img.Homepage-logo`, &buf)); err != nil {
	url := fmt.Sprintf("http://localhost:8080/fukidashi?message=%s&event=%s&room=%s", text, event, room)
	if err := chromedp.Run(ctx, elementScreenshot(url, `div.target`, &buf)); err != nil {
		log.Println(err)
		return
	}
	if err := ioutil.WriteFile(fmt.Sprintf("tmp/%s.png", imageName), buf, 0o644); err != nil {
		log.Println(err)
		return
	}

	if err := gcp.StreamFileUpload(buf, fmt.Sprintf("tmp/%s.png", imageName)); err != nil {
		log.Println(err)
		return
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")

	return
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
