package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	maxSessionPerLine   = 10
	sessionCheckPadding = 4
	sessionCheckHeight  = 600 / (maxSessionPerLine + sessionCheckPadding*2)
	sessionCheckSize    = sessionCheckHeight - (sessionCheckPadding * 2)
)

type mainWindow struct {
	rect         rectLayout
	titleRect    rectLayout
	progressRect rectLayout
	timerRect    rectLayout
	// rect      rectangle
	// timerRect rectangle

	font          *Font
	rectOutline   *ebiten.Image
	outlineConstr constraint

	shouldHighlight bool
}

func (m *mainWindow) init(font *Font, outline *ebiten.Image) {
	const mainWindowPadding = 5
	m.rect = newRectLayout(rectangle{200, 0, 600, windowWidth})
	m.titleRect = m.rect.cut(rectCutUp, 50, mainWindowPadding)

	m.progressRect = m.rect.cut(rectCutUp, 50, mainWindowPadding)
	m.progressRect.cut(rectCutRight, 200, 0)
	m.progressRect.cut(rectCutLeft, 200, 0)

	m.timerRect = m.rect.cut(rectCutUp, 50, mainWindowPadding)
	m.timerRect.cut(rectCutRight, 200, 0)
	m.timerRect.cut(rectCutLeft, 200, 0)

	m.font = font
	m.rectOutline = outline
	m.outlineConstr = constraint{2, 2, 2, 2}
}

func (m *mainWindow) update(mPos point, mLeft bool, selected bool) (startTask bool) {

	m.shouldHighlight = false
	if !isInputHandled(mPos) {
		if selected && m.timerRect.remaining.boundCheck(mPos) {
			m.shouldHighlight = true
			if mLeft {
				startTask = true
			}
		}
	}
	return
}

func (m *mainWindow) draw(dst *ebiten.Image, task *task) {
	ebitenutil.DrawLine(
		dst,
		m.rect.full.x,
		m.rect.full.y,
		m.rect.full.x,
		m.rect.full.y+m.rect.full.height,
		color.White,
	)

	if task != nil {
		if m.shouldHighlight {
			drawRect(dst, m.timerRect.remaining, WhiteA125)
		}
		drawTextCenter(dst, textOptions{
			font: m.font, text: task.name, bounds: m.titleRect.remaining,
			size: largeTextSize, clr: White,
		})
		drawRect(dst, m.progressRect.remaining, White)

		drawImageSlice(dst, m.timerRect.remaining, m.rectOutline, m.outlineConstr, White)
		// checkWidth := float64(task.sessionRequired * sessionCheckHeight)
		// startPos := m.rect.x + (m.rect.width/2 - checkWidth/2)

		// tSize := m.font.MeasureText("- Sessions -", textSize)
		// tPos := point{m.rect.x + (m.rect.width/2 - tSize[0]/2), 70}
		// drawText(dst, textOptions{font: m.font, text: "- Sessions -", pos: tPos, size: textSize, clr: White})
		// for i := 0; i < task.sessionRequired; i += 1 {
		// 	checkBoxRect := rectangle{startPos + float64(i*sessionCheckHeight), 100, sessionCheckSize, sessionCheckSize}
		// 	drawImageSlice(dst, checkBoxRect, m.rectOutline, m.outlineConstr, White)
		// 	if i < task.sessionCompleted {
		// 		checkBoxRect := rectangle{startPos + float64(i*sessionCheckHeight) + 3, 100 + 3, sessionCheckSize - 6, sessionCheckSize - 6}
		// 		drawRect(dst, checkBoxRect, White)
		// 	}
		// }

		// text := string(task.timer.toString())
		// tSize = m.font.MeasureText(text, textSize)
		// tPos = point{m.rect.x + (m.rect.width/2 - tSize[0]/2), 130}
		// drawText(dst, textOptions{font: m.font, text: text, pos: tPos, size: textSize, clr: White})

		// tSize = m.font.MeasureText("Start Timer", textSize)
		// tPos = point{m.timerRect.x + (m.timerRect.width/2 - tSize[0]/2), m.timerRect.y + (m.timerRect.height/2 - m.font.Ascent(textSize)/2 - 2)}
		// drawText(dst, textOptions{font: m.font, text: "Start Timer", pos: tPos, size: textSize, clr: White})
		// drawImageSlice(dst, m.timerRect, m.rectOutline, m.outlineConstr, White)
	}
}
