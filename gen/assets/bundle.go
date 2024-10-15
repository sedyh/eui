package assets

import (
	"embed"
	"image"
	"io/fs"
	"regexp"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

//go:embed data
var data embed.FS

var (
	Frame    = load(data, "data/frame.png")
	Output   = "assets/examples"
	Parent   = "gen"
	Width    = 480
	Height   = 320
	Pivot    = 28.
	ColorR   = colornames.Indianred
	ColorY   = colornames.Goldenrod
	ColorB   = colornames.Steelblue
	ColorG   = colornames.Mediumseagreen
	ColorD   = colornames.Darkslategray
	ColorL   = colornames.Gainsboro
	ColorMap = map[string]string{
		"assets.ColorR": "colornames.Indianred",
		"assets.ColorY": "colornames.Goldenrod",
		"assets.ColorB": "colornames.Steelblue",
		"assets.ColorG": "colornames.Mediumseagreen",
		"assets.ColorD": "colornames.Darkslategray",
		"assets.ColorL": "colornames.Gainsboro",
	}
	ColorPackage = `"golang.org/x/image/colornames"`
	RegPackage   = regexp.MustCompile(`package .+`)
	RegAssets    = regexp.MustCompile(`(?m)^\s*"\S*assets\S*"\s*\n`)
	RegImports   = regexp.MustCompile(`import \(([\s\S]*?)\n\)`)
	RegNewline   = regexp.MustCompile(`\n+`)
	RegColor     = regexp.MustCompile(`\b(assets.Color[RYBGDL])\b`)
)

func load(fs fs.FS, path string) *ebiten.Image {
	f, err := fs.Open(path)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
