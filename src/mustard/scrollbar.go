package mustard

import (
	"image"
	"image/draw"
	assets "thdwb/assets"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
)

//CreateScrollBarWidget - Creates and returns a new ScrollBar Widget
func CreateScrollBarWidget(orientation ScrollBarOrientation) *ScrollBarWidget {
	var widgets []Widget
	font, _ := truetype.Parse(assets.OpenSans(400))

	return &ScrollBarWidget{
		baseWidget: baseWidget{

			needsRepaint: true,
			widgets:      widgets,

			widgetType: scrollbarWidget,

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",

			font: font,
		},
		orientation: orientation,
	}
}

//SetWidth - Sets the scrollBar width
func (scrollBar *ScrollBarWidget) SetWidth(width int) {
	scrollBar.box.width = width
	scrollBar.fixedWidth = true
	scrollBar.RequestReflow()
}

//SetHeight - Sets the scrollBar height
func (scrollBar *ScrollBarWidget) SetHeight(height int) {
	scrollBar.box.height = height
	scrollBar.fixedHeight = true
	scrollBar.RequestReflow()
}

//SetBackgroundColor - Sets the scrollBar background color
func (scrollBar *ScrollBarWidget) SetTrackColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		scrollBar.backgroundColor = backgroundColor
		scrollBar.needsRepaint = true
	}
}

func (scrollBar *ScrollBarWidget) SetScrollerSize(scrollerSize float64) {
	scrollBar.scrollerSize = scrollerSize
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) SetThumbSize(thumbSize float64) {
	scrollBar.thumbSize = thumbSize
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) SetThumbColor(thumbColor string) {
	scrollBar.thumbColor = thumbColor
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) SetScrollerOffset(scrollerOffset float64) {
	scrollBar.scrollerOffset = scrollerOffset
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) draw() {
	context := scrollBar.window.context

	top := float64(scrollBar.computedBox.top)
	left := float64(scrollBar.computedBox.left)
	width := float64(scrollBar.computedBox.width)
	height := float64(scrollBar.computedBox.height)

	context.SetHexColor(scrollBar.backgroundColor)
	context.DrawRectangle(left, top, width, height)
	context.Fill()

	if scrollBar.scrollerSize > height {
		thumbSize := height * (height / scrollBar.scrollerSize)
		thumbOffset := scrollBar.scrollerOffset

		scrollJump := (scrollBar.scrollerSize - height) / (height - thumbSize)

		context.SetHexColor(scrollBar.thumbColor)
		context.DrawRectangle(left+1, top-(thumbOffset/scrollJump), width-2, thumbSize)
		context.Fill()
	}

	iTop, iLeft, iWidth, iHeight := scrollBar.computedBox.GetCoords()

	if scrollBar.buffer == nil || scrollBar.buffer.Bounds().Max.X != iWidth && scrollBar.buffer.Bounds().Max.Y != iHeight {
		scrollBar.buffer = image.NewRGBA(image.Rectangle{
			image.Point{}, image.Point{iWidth, iHeight},
		})
	}

	draw.Draw(scrollBar.buffer, image.Rectangle{
		image.Point{},
		image.Point{iWidth, iHeight},
	}, context.Image(), image.Point{iLeft, iTop}, draw.Over)
}
