package main

import (
	"errors"
	"fmt"
	"image/png"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"uir/assets"
	"uir/examples/bas_con_mul"
	"uir/examples/bas_con_sin"
	"uir/examples/lay_anc_dir_hor"
	"uir/examples/lay_row_pre"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type State int

const (
	Switched State = iota
	Updated
)

type Game struct {
	mu        sync.Mutex
	scenes    []Scene
	state     State
	current   int
	offscreen *ebiten.Image
}

func NewGame() *Game {
	root := assets.Output
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil
		}
		name := d.Name()
		if len(name) > 0 && name[0] == '.' {
			return nil
		}
		return os.Remove(path)
	})

	return &Game{
		scenes: []Scene{
			bas_con_sin.NewGame(),
			bas_con_mul.NewGame(),
			lay_anc_dir_hor.NewGame(),
			lay_row_pre.NewGame(),
		},
		state:   Switched,
		current: 0,
		offscreen: ebiten.NewImage(
			assets.Frame.Bounds().Dx(),
			assets.Frame.Bounds().Dy(),
		),
	}
}

func (g *Game) Update() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.current == len(g.scenes) {
		return ebiten.Termination
	}

	if g.state == Switched {
		err := g.scenes[g.current].Update()
		if err != nil {
			return err
		}
		g.state = Updated
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.state == Updated {
		g.scenes[g.current].Draw(screen)

		g.SaveImage(screen)
		g.SaveCode()

		g.current++
		g.state = Switched
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func (g *Game) SaveImage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, assets.Pivot)
	g.offscreen.DrawImage(assets.Frame, nil)
	g.offscreen.DrawImage(screen, op)

	name := FindPackage(g.scenes[g.current])
	base := filepath.Base(name)
	path := filepath.Join(assets.Output, base+".png")
	if err := WriteImage(g.offscreen, path); err != nil {
		log.Fatal(err)
	}
	log.Println("saved", path)

	g.offscreen.Clear()
}

func (g *Game) SaveCode() {
	name := FindPackage(g.scenes[g.current])

	in := RemoveRoot(name)
	in = filepath.Join(in, "main.go")
	source, err := ReadSource(in)
	if err != nil {
		log.Fatal(err)
	}

	source = assets.RegPackage.ReplaceAllString(source, "package main")
	source = assets.RegAssets.ReplaceAllString(source, "")
	source = assets.RegImports.ReplaceAllStringFunc(source, func(s string) string {
		return assets.RegNewline.ReplaceAllString(s, "\n")
	})
	source = assets.RegImports.ReplaceAllStringFunc(source, func(match string) string {
		return match[:len(match)-1] + "\t" + assets.ColorPackage + "\n)"
	})
	source = assets.RegColor.ReplaceAllStringFunc(source, func(match string) string {
		return assets.ColorMap[match]
	})

	out := filepath.Base(name)
	out = filepath.Join(assets.Output, out+".txt")
	err = WriteSource(source, out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("saved", name)
}

func FindPackage(i any) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath()
}

func WriteImage(img *ebiten.Image, path string) (e error) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer Close(&e, f)
	if err = png.Encode(f, img); err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	return nil
}

func ReadSource(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}
	return string(data), err
}

func WriteSource(source, path string) error {
	err := os.WriteFile(path, []byte(source), 0666)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func RemoveRoot(path string) string {
	sep := string(filepath.Separator)
	slash := filepath.ToSlash(path)
	parts := strings.Split(slash, sep)
	if len(parts) > 1 && parts[0] == "" {
		parts = parts[1:]
	}
	return strings.Join(parts[1:], sep)
}

func Close(dest *error, c io.Closer) {
	*dest = errors.Join(*dest, c.Close())
}

func main() {
	log.SetFlags(0)
	ebiten.SetWindowSize(assets.Width, assets.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
