package main

import (
	"flag"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

var (
	pageSize      = 100    // Number of keys to display per page
	currentPage   = 0      // Current page index
	filteredKeys  [][]byte // Filtered keys based on search or prefix
	currentPrefix string   // Current prefix filter
	showHelp      = false  // Flag to show/hide help window
)

func main() {
	// Command-line flags
	dbPath := flag.String("db", "", "Path to the LevelDB database")
	flag.Parse()

	// Open the LevelDB database
	db, err := leveldb.OpenFile(*dbPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check if dbPath is not empty
	if *dbPath == "" {
		log.Fatal("dbPath is empty, add -db flag to specify the path to the LevelDB database")
	}

	defer db.Close()

	// Create a new tview application
	app := tview.NewApplication()

	// Create a list to display keys
	keyList := tview.NewList().
		SetWrapAround(false)
	keyList.SetBorder(true).
		SetTitle("Keys")

	// Create a text view to display the value of the selected key
	valueList := tview.NewTextView()
	valueList.SetDynamicColors(true).
		SetBorder(true).
		SetTitle("Value")

	// Create an input field for searching and prefix filtering
	searchBox := tview.NewInputField()
	searchBox.
		SetLabel("Search/Prefix: ").
		SetFieldWidth(20).
		SetChangedFunc(func(text string) {
			currentPrefix = text
			filterKeys(db, keyList, valueList)
		}).
		SetDoneFunc(func(key tcell.Key) {
			app.SetFocus(keyList)
		})
	keyList.SetDoneFunc(func() {
		app.SetFocus(searchBox)
	})

	// Create a help window
	helpText := `Use Arrow keys to navigate
'n' for next page
'p' for previous page
'esc' to change focus
'enter' when on search go to keys
'enter' when on keys shows value
'q' to quit
'h' to toggle this help window`
	helpWindow := tview.NewTextView().
		SetDynamicColors(true).
		SetText(helpText)
	helpWindow.
		SetBorder(true).
		SetTitle("Help")

	// Create a flexbox layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(searchBox, 1, 1, false).
		AddItem(tview.NewFlex().
			AddItem(keyList, 0, 1, true).
			AddItem(valueList, 0, 2, false), 0, 1, true).
		AddItem(tview.NewTextView().SetText("Use Arrow keys to navigate, 'n' for next page, 'p' for previous page, 'h' for help"), 1, 1, false)

	// Populate the initial key list
	filterKeys(db, keyList, valueList)

	// Set input capture for pagination and help window
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'n':
			nextPage(db, keyList, valueList)
		case 'p':
			prevPage(db, keyList, valueList)
		case 'q', 'Q':
			app.Stop()
		case 'h', 'H':
			showHelp = !showHelp
			if showHelp {
				flex.AddItem(helpWindow, 0, 1, false)
			} else {
				flex.RemoveItem(helpWindow)
			}
		}
		return event
	})

	// Set up and run the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		log.Fatal(err)
	}
}

// filterKeys filters keys based on the current prefix and repopulates the key list
func filterKeys(db *leveldb.DB, keyList *tview.List, valueList *tview.TextView) {
	keyList.Clear()
	currentPage = 0

	iter := db.NewIterator(util.BytesPrefix([]byte(currentPrefix)), nil)
	defer iter.Release()

	filteredKeys = [][]byte{}
	for iter.Next() {
		key := iter.Key()
		filteredKeys = append(filteredKeys, append([]byte{}, key...)) // Copy key to avoid reference issues
	}

	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}

	displayPage(db, keyList, valueList)
}

// displayPage displays the current page of keys in the key list
func displayPage(db *leveldb.DB, keyList *tview.List, valueList *tview.TextView) {
	keyList.Clear()

	start := currentPage * pageSize
	end := start + pageSize
	if end > len(filteredKeys) {
		end = len(filteredKeys)
	}

	for i := start; i < end; i++ {
		key := filteredKeys[i]
		keyList.AddItem(string(key), "", 0, func() {
			value, err := db.Get(key, nil)
			if err != nil {
				valueList.SetText(fmt.Sprintf("[red]Error: %v", err))
			} else {
				valueList.SetText(fmt.Sprintf("Key: %s\n\nValue: %s", key, value))
			}
		})
	}
}

// nextPage moves to the next page of keys
func nextPage(db *leveldb.DB, keyList *tview.List, valueList *tview.TextView) {
	if (currentPage+1)*pageSize < len(filteredKeys) {
		currentPage++
		displayPage(db, keyList, valueList)
	}
}

// prevPage moves to the previous page of keys
func prevPage(db *leveldb.DB, keyList *tview.List, valueList *tview.TextView) {
	if currentPage > 0 {
		currentPage--
		displayPage(db, keyList, valueList)
	}
}
