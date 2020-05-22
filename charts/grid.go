package charts

// GridOpts is the option set for a tooltip component.
type GridOpts struct {
	// 是否显示直角坐标系网格
	Show bool `json:"show,omitempty"`
	// grid组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// grid组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// grid组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// grid组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// grid 区域是否包含坐标轴的刻度标签。
	// containLabel 为 false 的时候：
	// grid.left grid.right grid.top grid.bottom grid.width grid.height 决定的是由坐标轴形成的矩形的尺寸和位置。
	// 这比较适用于多个 grid 进行对齐的场景，因为往往多个 grid 对齐的时候，是依据坐标轴来对齐的。
	// containLabel 为 true 的时候：
	// grid.left grid.right grid.top grid.bottom grid.width grid.height 决定的是包括了坐标轴标签在内的所有内容所形成的矩形的位置。
	// 这常用于『防止标签溢出』的场景，标签溢出指的是，标签长度动态变化时，可能会溢出容器或者覆盖其他组件。
	ContainLabel bool `json:"containLabel,omitempty"`
}

func (GridOpts) markGlobal() {}
