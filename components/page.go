package components

import (
	"github.com/KiVirgil/go-echarts/v2/opts"
	"github.com/KiVirgil/go-echarts/v2/render"
)

type Layout string

const (
	PageNoneLayout   Layout = "none"
	PageCenterLayout Layout = "center"
	PageFlexLayout   Layout = "flex"
)

// Charter
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	Validate()
}

// Page represents a page chart.
type Page struct {
	render.Renderer
	opts.Initialization
	opts.Assets

	Charts []interface{}
	Layout Layout
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.Renderer = render.NewPageRender(page, page.Validate)
	page.Layout = PageCenterLayout
	return page
}

func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
	return page
}

// AddCharts adds new charts to the page.
func (page *Page) AddCharts(charts ...Charter) *Page {
	for i := 0; i < len(charts); i++ {
		assets := charts[i].GetAssets()
		for _, v := range assets.JSAssets.Values {
			page.JSAssets.Add(v)
		}

		for _, v := range assets.CSSAssets.Values {
			page.CSSAssets.Add(v)
		}
		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// Validate
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}
