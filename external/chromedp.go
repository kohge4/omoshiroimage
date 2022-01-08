package external

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chromedp/chromedp"
)

type ImageGenerator interface {
	GenerateImage()
}

type ChromedpImageGenerator struct {
}

func NewChromedpImageGenerator() ImageGenerator {
	return &ChromedpImageGenerator{}
}

func (g *ChromedpImageGenerator) GenerateImage() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	go screenShot(sigCh)
	<-sigCh
}

func screenShot(sigCh chan os.Signal) {
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

	name := fmt.Sprintf("%s", time.Now())

	text := "おはようございます、よろしくお願いします。"
	url := fmt.Sprintf("http://localhost:8080/fukidashi?message=%s", text)
	if err := chromedp.Run(ctx, elementScreenshot(url, `div.target`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("tmp/%s.png", name), buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
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
