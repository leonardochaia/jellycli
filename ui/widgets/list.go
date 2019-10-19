/*
 * Copyright 2019 Tero Vierimaa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package widgets

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"tryffel.net/pkg/jellycli/config"
	"tryffel.net/pkg/jellycli/models"
)

type List struct {
	list    *tview.List
	element models.ListElement
}

func NewList() *List {
	l := &List{
		list: tview.NewList(),
	}
	l.list.SetBorder(true)
	l.list.SetBorderColor(config.ColorBorder)
	l.list.SetTitleColor(config.ColorBorder)
	l.list.SetTitleAlign(tview.AlignLeft)
	l.list.ShowSecondaryText(false)
	l.list.SetShortcutColor(tcell.ColorDefault)
	l.list.SetBackgroundColor(config.ColorBackground)
	l.list.SetSelectedTextColor(config.ColorSecondary)
	l.list.SetMainTextColor(config.ColorPrimary)
	return l
}

func (l *List) Draw(screen tcell.Screen) {
	l.list.Draw(screen)
}

func (l *List) GetRect() (int, int, int, int) {
	return l.list.GetRect()

}

func (l *List) SetRect(x, y, width, height int) {
	l.list.SetRect(x, y, width, height)
}

func (l *List) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return l.list.InputHandler()
}

func (l *List) Focus(delegate func(p tview.Primitive)) {
	l.list.Focus(delegate)
}

func (l *List) Blur() {
	l.list.Blur()
}

func (l *List) GetFocusable() tview.Focusable {
	return l.list.GetFocusable()
}

func (l *List) SetData(element models.ListElement, items []models.Item) {
	l.element = element

	l.list.Clear()
	l.list.SetTitle(string(element))
	for i, v := range items {
		l.list.AddItem(v.GetName(), "", '?', l.namedCb(i, v.GetId()))
	}
}

func (l *List) namedCb(i int, id string) func() {
	return func() {
		l.selectCb(i, id)
	}
}

func (l *List) selectCb(i int, id string) {

}