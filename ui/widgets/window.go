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
	"github.com/sirupsen/logrus"
	"tryffel.net/pkg/jellycli/config"
	"tryffel.net/pkg/jellycli/models"
	"tryffel.net/pkg/jellycli/player"
	"tryffel.net/pkg/jellycli/ui/controller"
)

type Window struct {
	app    *tview.Application
	window *tview.Grid

	// Widgets
	navBar  *NavBar
	status  *Status
	browser *Browser
	search  *Search
	help    *Help

	mediaController controller.MediaController
}

func NewWindow(mc controller.MediaController) Window {
	w := Window{
		app:     tview.NewApplication(),
		status:  newStatus(mc.Control),
		window:  tview.NewGrid(),
		browser: NewBrowser(),
	}

	w.navBar = NewNavBar(w.keyHandlerCb)
	w.mediaController = mc

	w.window.SetTitle(" " + config.AppName + " ")
	w.window.SetTitleColor(config.ColorPrimary)
	w.window.SetBackgroundColor(config.ColorBackground)
	w.window.SetBorderColor(config.ColroMainFrame)
	w.setLayout()
	w.app.SetRoot(w.window, true)

	data := testData()
	w.browser.setData(&data, models.ArtistList)
	w.window.SetInputCapture(w.eventHandler)
	w.search = NewSearch(w.searchCb)
	w.help = NewHelp(w.closeHelp)

	w.status.UpdateState(player.PlayingState{
		State:               player.Play,
		PlayingType:         player.Song,
		Song:                "Song",
		Artist:              "Artist",
		Album:               "Album",
		CurrentSongDuration: 185,
		CurrentSongPast:     92,
		PlaylistDuration:    0,
		PlaylistLeft:        0,
		Volume:              50,
	}, &models.SongInfo{
		Id:       "song1",
		Name:     "Song",
		Length:   185,
		Artist:   "Artist A",
		ArtistId: "a1",
		Album:    "Album B",
		AlbumId:  "ab2",
		Year:     2019,
	})

	return w
}

func (w *Window) Run() error {
	return w.app.Run()
}

func (w *Window) Stop() {
	w.app.Stop()
}

func (w *Window) setLayout() {
	w.window.Clear()
	w.window.SetBorder(true)
	w.window.SetRows(1, -1, 4)
	w.window.SetColumns(-1)

	w.window.AddItem(w.navBar, 0, 0, 1, 1, 1, 30, false)
	w.window.AddItem(w.browser, 1, 0, 1, 1, 15, 10, false)
	w.window.AddItem(w.status, 2, 0, 1, 1, 3, 10, false)
}

func (w *Window) eventHandler(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	out := w.keyHandler(&key)
	if out == nil {
		return nil
	} else {
		return event
	}

}

// Keyhandler that has to react to buttons or drop them completely
func (w *Window) keyHandlerCb(key *tcell.Key) {

}

// Key handler, if match, return nil
func (w *Window) keyHandler(key *tcell.Key) *tcell.Key {
	navbar := config.KeyBinds.NavigationBar
	switch *key {
	case navbar.Quit:
		w.app.Stop()
	case navbar.Search:
		w.browser.AddModal(w.search, 10, 50, true)
		w.app.SetFocus(w.search)
	case tcell.KeyF1:
		w.browser.AddModal(w.help, 25, 50, true)
		w.app.SetFocus(w.help)
	default:
		return key

	}
	return nil
}

func (w *Window) searchCb(query string, doSearch bool) {
	logrus.Debug("In search callback")
	w.browser.RemoveModal(w.search)
	w.app.SetFocus(w.window)

	if doSearch {
		w.mediaController.Search(query)
	}

}

func (w *Window) closeHelp() {
	w.browser.RemoveModal(w.help)
	w.app.SetFocus(w.window)

}