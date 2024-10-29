package gui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func MainWindowWidget(gtx *layout.Context, list *widget.List, theme *material.Theme, text *string, settings *Settings) {
	layout.Inset{Top: unit.Dp(40), Bottom: unit.Dp(40), Left: unit.Dp(40), Right: unit.Dp(40)}.Layout(
		*gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.End}.Layout(
				gtx,
				layout.Flexed(
					1,
					func(gtx layout.Context) layout.Dimensions {
						return list.Layout(
							gtx,
							1,
							func(gtx layout.Context, index int) layout.Dimensions {
								return material.Label(theme, unit.Sp(settings.Fontsize), *text).Layout(gtx)
							},
						)
					},
				),
			)
		},
	)
}
