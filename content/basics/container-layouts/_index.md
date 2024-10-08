+++
title = 'Container and layouts'
date = 2024-10-04T19:45:29+03:00
weight = 1
+++

The library provides a rendering manager `*ebitenui.UI` in 
which you will place your entire UI.

Manager is located in `ebitenui` package.
```go
import "github.com/ebitenui/ebitenui"
```

Manager needs to get constantly updated and drawed.
```go
type Game struct {
	ui *ebitenui.UI
}

func NewGame() *Game {
	return &Game{
		ui: &ebitenui.UI{},
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}
```

Any UI in this library consists of containers that are nested in each other.

Container is located in `widget` package.
```go
import "github.com/ebitenui/ebitenui/widget"
```

The manager is contains a reference to the root container and responsible for delivering events throughout the user interface.
Let's pass our container there so we can interact with it.

There is a single container type which can be created like this.
```go
func NewGame() *Game {
    c := widget.NewContainer()
    return &Game{
        ui: &ebitenui.UI{Container: c},
    }
}
```

The standard library has a package with default colors like `Gainsboro`, `Indianred`, `Steelblue`,  that will be useful to us.

```go
import "golang.org/x/image/colornames"
```

The library draws all interface elements using multiple image tiles also known as [nine-slice](https://en.wikipedia.org/wiki/9-slice_scaling) to prevent image scaling distortion. That package will help us to work with images.

```go
import "github.com/ebitenui/ebitenui/image"
```

The container have several options to setup, like background color.

```go
c := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightblue),
    ),
)
```

Lets run the app. We will see a single container that will take up all the free space.

![image](examples/bas_con_sin.png)

[w](examples/bas_con_sin.txt)
w

{{% expand title="Full example" %}}
s
{{< code source="examples/bas_con_sin.txt" >}}
s
{{% /expand %}}
w

Containers can be composed into each other. The order adding child containers does not matter.

```go
shirt := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightblue),
    ),
)
pants := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightskyblue),
    ),
)
store := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightsteelblue),
    ),
)
)
store.AddChild(shelf)
store.AddChild(pants)
```

To prevent child containers from overlapping each other, we can specify how they are positioned within the parent container using other container properties, such as layout.

```go
store := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
)
```

Positioning within the layout is specified by the `LayoutData` structure inside each container. In addition, at least, each container must have a minimum size.

```go
pants := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionStart,
            StretchVertical:    true,
        }),
        widget.WidgetOpts.MinSize(50, 50),
    ),
)
```

Let's set similar options for other containers.

```go
shirt := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightblue),
    ),
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionEnd,
            StretchVertical:    true,
        }),
        widget.WidgetOpts.MinSize(50, 50),
    ),
)
pants := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightskyblue),
    ),
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionStart,
            StretchVertical:    true,
        }),
        widget.WidgetOpts.MinSize(50, 50),
    ),
)
store := widget.NewContainer(
    widget.ContainerOpts.BackgroundImage(
        image.NewNineSliceColor(colornames.Lightsteelblue),
    ),
    widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
)
store.AddChild(shirt)
store.AddChild(pants)
```

Let's launch the application. We will see two child containers with `different horizontal positions` that will `stretch vertically`.

![image](examples/bas_con_mul.png)


{{% expand title="Full example" %}}
```go
package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

func main() {
	ebiten.SetWindowSize(480, 320)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}

type Game struct {
	ui *ebitenui.UI
}

func NewGame() *Game {
	shirt := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightblue),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				StretchVertical:    true,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)
	pants := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightskyblue),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				StretchVertical:    true,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)
	store := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightsteelblue),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	store.AddChild(shirt)
	store.AddChild(pants)

	return &Game{
		ui: &ebitenui.UI{Container: store},
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}
```
{{% /expand %}}

The library has several different layouts for different situations, you can study each of them in detail on the next pages.
