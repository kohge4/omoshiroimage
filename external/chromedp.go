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
	GenerateImage(string) string
}

type ChromedpImageGenerator struct {
}

func NewChromedpImageGenerator() ImageGenerator {
	return &ChromedpImageGenerator{}
}

func (g *ChromedpImageGenerator) GenerateImage(text string) string {

	name := fmt.Sprintf("%d", time.Now().Unix())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	go screenShot(sigCh, name, text)
	<-sigCh

	return name
}

func screenShot(sigCh chan os.Signal, imageName string, text string) {
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

	url := fmt.Sprintf("http://localhost:8080/fukidashi?message=%s", text)
	if err := chromedp.Run(ctx, elementScreenshot(url, `div.target`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(fmt.Sprintf("tmp/%s.png", imageName), buf, 0o644); err != nil {
		log.Fatal(err)
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
