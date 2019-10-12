package gtk

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"log"

	"github.com/disintegration/imaging"

	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/chrischdi/gbooth/pkg/proto"
	"github.com/chrischdi/gbooth/pkg/ui/bindata"
)

const (
	Title = "Photobox"
)

func NewGTK() (*GTK, error) {
	ui := &GTK{}
	return ui, nil
}

type GTK struct {
	window         *gtk.Window
	overlay        *gtk.Overlay
	image          *gtk.Image
	imageArrows    *gtk.Image
	countdownLabel *gtk.Label
}

func (ui *GTK) Init() error {
	var err error
	// initialize gtk
	gtk.Init(nil)

	// create window
	ui.window, err = createWindow(Title, 1280, 800)
	if err != nil {
		return fmt.Errorf("error creating gtk.Window: %v", err)
	}

	// load overlay image
	b, err := bindata.Asset("arrows.png")
	if err != nil {
		return err
	}

	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		return err
	}

	_, err = loader.Write(b)
	if err != nil {
		return err
	}

	pb, err := loader.GetPixbuf()
	if err != nil {
		return err
	}

	// create window content
	ui.image, ui.imageArrows, ui.overlay, ui.countdownLabel, err = createContent(pb)
	if err != nil {
		return fmt.Errorf("error creating content: %v", err)
	}
	// ui.image.SetSizeRequest(ui.window.GetAllocatedWidth(), ui.window.GetAllocatedHeight())

	ui.window.Add(ui.overlay)

	ui.window.ShowAll()

	go ui.background()

	b, err = bindata.Asset("background.jpg")
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	return ui.Publish(img)

	return nil
}

func (ui *GTK) Free() error {
	return nil
}

func (ui *GTK) Background(ctx context.Context, in *proto.Image) (*proto.UIResponse, error) {
	img, _, err := image.Decode(bytes.NewReader(in.GetImage()))
	if err != nil {
		return nil, err
	}

	if err := ui.Publish(img); err != nil {
		return nil, err
	}
	return &proto.UIResponse{}, nil
}

func (ui *GTK) Timer(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error) {
	var err error
	if in.GetText() == "" {
		_, err = glib.IdleAdd(func() {
			ui.imageArrows.SetVisible(false)
			gtkSetCountdownLabel(ui.countdownLabel, in.GetText())
		})
		if err != nil {
			log.Printf("error on idleAdd for countdownLabel: %v", err)
		}
	} else {
		_, err = glib.IdleAdd(func() {
			gtkEnableArrows(ui.imageArrows)
			gtkSetCountdownLabel(ui.countdownLabel, in.GetText())
		})
	}
	if err != nil {
		return nil, fmt.Errorf("error on idleAdd for gtkSetContdownLabel: %v", err)
	}
	return &proto.UIResponse{}, nil

	// return &proto.UIResponse{}, nil
}

func (ui *GTK) Error(ctx context.Context, in *proto.TextRequest) (*proto.UIResponse, error) {
	var err error
	if in.GetText() == "" {
		_, err = glib.IdleAdd(func() {
			ui.imageArrows.SetVisible(false)
			gtkSetCountdownLabel(ui.countdownLabel, fmt.Errorf(in.GetText()))
		})
		if err != nil {
			log.Printf("error on idleAdd for countdownLabel: %v", err)
		}
	} else {
		_, err = glib.IdleAdd(func() {
			ui.imageArrows.SetVisible(false)
			gtkSetCountdownLabel(ui.countdownLabel, fmt.Errorf(in.GetText()))
		})
	}
	if err != nil {
		return nil, fmt.Errorf("error on idleAdd for gtkSetContdownLabel: %v", err)
	}
	return &proto.UIResponse{}, nil
}

func convertImageToPixbufAtSize(img image.Image, width, height int) (*gdk.Pixbuf, error) {
	resized := imaging.Fill(img, width, height, imaging.Center, imaging.Box)

	// write image to buffer
	var buf bytes.Buffer
	err := imaging.Encode(&buf, resized, imaging.JPEG, imaging.JPEGQuality(95))
	if err != nil {
		return nil, err
	}

	// load buffer to pixbuf
	loader, err := gdk.PixbufLoaderNewWithType("jpeg")
	if err != nil {
		return nil, err
	}
	_, err = loader.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	pb, err := loader.GetPixbuf()
	if err != nil {
		return nil, err
	}

	return pb, nil
}

func (ui *GTK) Publish(img image.Image) error {
	width, height := ui.window.GetSize()
	pb, err := convertImageToPixbufAtSize(img, width, height)
	if err != nil {
		return err
	}

	// Publish the pixbuf
	_, err = glib.IdleAdd(func() { gtkPublish(ui, pb) })
	if err != nil {
		log.Printf("error on idleAdd for image: %v", err)
	}
	log.Println("publish image done")
	return nil
}

func (ui *GTK) ShowError(err error, duration time.Duration) {
	log.Printf("trying to show error: %v", err)
	_, idleaddErr := glib.IdleAdd(func() {
		// disable overlay image
		ui.imageArrows.SetVisible(false)
		gtkSetCountdownLabel(ui.countdownLabel, err)
	})
	if idleaddErr != nil {
		log.Printf("error on idleAdd for countdownLabel: %v", err)
	}
	time.Sleep(duration)
}

func (ui *GTK) background() {
	gtk.Main()
	panic("gtk.Main did return")
}

func createWindow(title string, width, height int) (*gtk.Window, error) {
	w, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, fmt.Errorf("error creating gtk.Window: %v", err)
	}
	w.SetTitle(title)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.SetDefaultSize(width, height)

	// set to full screen
	w.Fullscreen()
	return w, nil
}

func createContent(arrows *gdk.Pixbuf) (*gtk.Image, *gtk.Image, *gtk.Overlay, *gtk.Label, error) {
	o, err := gtk.OverlayNew()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error creating gtk.Overlay: %v", err)
	}
	o.SetHExpand(true)
	o.SetVExpand(true)

	i, err := gtk.ImageNew()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error creating gtk.Image: %v", err)
	}
	i.SetHExpand(false)
	i.SetVExpand(false)
	o.Add(i)

	iArrows, err := gtk.ImageNewFromPixbuf(arrows)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error creating gtk.Image: %v", err)
	}
	iArrows.SetHExpand(false)
	iArrows.SetVExpand(false)
	iArrows.SetVisible(false)
	o.AddOverlay(iArrows)

	l, err := gtk.LabelNew("")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error creating gtk.Label: %v", err)
	}
	// set position
	l.SetHAlign(gtk.ALIGN_CENTER)
	l.SetVAlign(gtk.ALIGN_CENTER)
	l.SetLineWrap(true)

	o.AddOverlay(l)

	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error creating draw handler: %v", err)
	}

	return i, iArrows, o, l, nil
}

func gtkPublish(ui *GTK, pixbuf *gdk.Pixbuf) {
	// set background image
	ui.image.SetFromPixbuf(pixbuf)
	// clear countdown label
	gtkSetCountdownLabel(ui.countdownLabel, "")
	// disable overlay image
	ui.imageArrows.SetVisible(false)
	// draw
	ui.overlay.QueueDraw()
}

func gtkEnableArrows(image *gtk.Image) {
	image.Show()
	image.QueueDraw()
}

func gtkSetCountdownLabel(label *gtk.Label, i interface{}) {
	var tpl string
	switch i.(type) {
	case error:
		tpl = "<span font_desc='Tahoma 30' color='#f44248'>%v</span>"
	case int:
		tpl = "<span font_desc='Tahoma 120' color='#f44248'>%d</span>"
	case string:
		tpl = "<span font_desc='Tahoma 120' color='#f44248'>%s</span>"
	default:
		tpl = "<span font_desc='Tahoma 120' color='#f44248'>%v</span>"
	}
	label.SetMarkup(fmt.Sprintf(tpl, i))

	label.QueueDraw()
}
