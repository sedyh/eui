+++
title = "Anchor layout"
date = 2024-10-04T19:45:29+03:00
weight = 1
+++

Anchor layout allows the child containers to be anchored to a specific constraint.

![preview](anchor-preview.png)

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
	left := widget.NewContainer(
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
	right := widget.NewContainer(
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
	up := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightcyan),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:  widget.AnchorLayoutPositionStart,
				StretchHorizontal: true,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)
	down := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightslategrey),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:  widget.AnchorLayoutPositionEnd,
				StretchHorizontal: true,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightsteelblue),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	root.AddChild(left)
	root.AddChild(right)
	root.AddChild(up)
	root.AddChild(down)

	return &Game{
		ui: &ebitenui.UI{Container: root},
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

### Layout options

###### Padding

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}
{{% tab title="Left" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewAnchorLayout(
        widget.AnchorLayoutOpts.Padding(widget.Insets{
            Left: 50,
        }),
    )),
)
```
![left](anchor-pad-left.png)
{{% /tab %}}
{{% tab title="Right" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewAnchorLayout(
        widget.AnchorLayoutOpts.Padding(widget.Insets{
            Right: 50,
        }),
    )),
)
```
![right](anchor-pad-right.png)
{{% /tab %}}
{{% tab title="Top" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewAnchorLayout(
        widget.AnchorLayoutOpts.Padding(widget.Insets{
            Top: 50,
        }),
    )),
)
```
![top](anchor-pad-top.png)
{{% /tab %}}
{{% tab title="Bottom" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewAnchorLayout(
        widget.AnchorLayoutOpts.Padding(widget.Insets{
            Bottom: 50,
        }),
    )),
)
```
![bottom](anchor-pad-bottom.png)
{{% /tab %}}
{{< /tabs >}}

### Layout data

###### Horizontal position

Responsible for aligning the element along the horizontal axis.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}
{{% tab title="Center" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionCenter,
        }),
    ),
)
```
![center](anchor-horz-center.png)
{{% /tab %}}
{{% tab title="Start" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionStart,
        }),
    ),
)
```
![start](anchor-horz-start.png)
{{% /tab %}}
{{% tab title="End" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionEnd,
        }),
    ),
)
```
![end](anchor-horz-end.png)
{{% /tab %}}
{{< /tabs >}}

###### Vertical position

Responsible for aligning the element along the vertical axis.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}
{{% tab title="Center" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionCenter,
        }),
    ),
)
```
![center](anchor-vert-center.png)
{{% /tab %}}
{{% tab title="Start" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionStart,
        }),
    ),
)
```
![start](anchor-vert-start.png)
{{% /tab %}}
{{% tab title="End" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
            HorizontalPosition: widget.AnchorLayoutPositionEnd,
        }),
    ),
)
```
![end](anchor-vert-end.png)
{{% /tab %}}
{{< /tabs >}}
