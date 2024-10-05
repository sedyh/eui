+++
title = "Row layout"
date = 2024-10-04T19:45:29+03:00
weight = 2
+++

Row layout places all child containers in one row or column. It can be useful for creating lists of widgets.

![preview](row-preview.png)

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
	a := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightskyblue),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(50, 80),
		),
	)
	b := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightblue),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(50, 80),
		),
	)
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightcyan),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(50, 80),
		),
	)
	d := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightslategrey),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(50, 80),
		),
	)
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Lightsteelblue),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(
				widget.DirectionVertical,
			),
		)),
	)
	root.AddChild(a)
	root.AddChild(b)
	root.AddChild(c)
	root.AddChild(d)

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

###### Direction

Responsible for whether child containers will follow each other in rows or columns.

{{< tabs title="Direction:" style="transparent" color="#131a22ff" >}}
{{% tab title="Horizontal" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Direction(
            widget.DirectionHorizontal,
        ),
    )),
)
```
![horz](row-horz.png)
{{% /tab %}}
{{% tab title="Vertical" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Direction(
            widget.DirectionVertical,
        ),
    )),
)
```
![vert](row-vert.png)
{{% /tab %}}
{{< /tabs >}}

###### Padding

Layout allows you to specify padding for all child elements but not the itself.

Please note that `Right` and `Bottom` padding works depending on the direction.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}
{{% tab title="Left" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Padding(widget.Insets{
            Left: 50,
        }),
    )),
)
```
![left](row-horz-pad-left.png)
![left](row-vert-pad-left.png)
{{% /tab %}}
{{% tab title="Right" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Padding(widget.Insets{
            Right: 50,
        }),
    )),
)
```
![right](row-horz-pad-right.png)
![right](row-vert-pad-right.png)
{{% /tab %}}
{{% tab title="Top" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Padding(widget.Insets{
            Top: 50,
        }),
    )),
)
```
![top](row-horz-pad-top.png)
![top](row-vert-pad-top.png)
{{% /tab %}}
{{% tab title="Bottom" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Padding(widget.Insets{
            Bottom: 50,
        }),
    )),
)
```
![bottom](row-horz-pad-bottom.png)
![bottom](row-vert-pad-bottom.png)
{{% /tab %}}
{{< /tabs >}}

###### Spacing

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Spacing:" style="transparent" color="#131a22ff" >}}
{{% tab title="75" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Spacing(75),
    )),
)
```
![75](row-spacing-75.png)
{{% /tab %}}
{{% tab title="50" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Spacing(50),
    )),
)
```
![50](row-spacing-50.png)
{{% /tab %}}
{{% tab title="25" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Spacing(25),
    )),
)
```
![25](row-spacing-25.png)
{{% /tab %}}
{{% tab title="0" %}}
```go
root := widget.NewContainer(
    widget.ContainerOpts.Layout(widget.NewRowLayout(
        widget.RowLayoutOpts.Spacing(0),
    )),
)
```
![0](row-spacing-0.png)
{{% /tab %}}
{{< /tabs >}}

### Layout data

###### Stretch

Responsible for stretching the element along the entire length of the opposite axis.

{{< tabs title="Stretch:" style="transparent" color="#131a22ff" >}}
{{% tab title="true" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            Stretch: true,
        }),
    ),
)
```
![true](row-stretch-true.png)
{{% /tab %}}
{{% tab title="false" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            Stretch: false,
        }),
    ),
)
```
![false](row-stretch-false.png)
{{% /tab %}}
{{< /tabs >}}

###### Position

Responsible for aligning the element along the opposite axis if it is not stretched.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}
{{% tab title="Center" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            Position: widget.RowLayoutPositionCenter,
        }),
    ),
)
```
![center](row-position-center.png)
{{% /tab %}}
{{% tab title="Start" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            Position: widget.RowLayoutPositionStart,
        }),
    ),
)
```
![start](row-position-start.png)
{{% /tab %}}
{{% tab title="End" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            Position: widget.RowLayoutPositionEnd,
        }),
    ),
)
```
![end](row-position-end.png)
{{% /tab %}}
{{< /tabs >}}


###### Max size

Responsible for the allowable size of the container.

{{< tabs title="Max:" style="transparent" color="#131a22ff" >}}
{{% tab title="Width" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            MaxWidth: 20,
        }),
    ),
)
```
![width](row-max-width.png)
{{% /tab %}}
{{% tab title="Height" %}}
```go
child := widget.NewContainer(
    widget.ContainerOpts.WidgetOpts(
        widget.WidgetOpts.LayoutData(widget.RowLayoutData{
            MaxHeight: 20,
        }),
    ),
)
```
![height](row-max-height.png)
{{% /tab %}}
{{< /tabs >}}
