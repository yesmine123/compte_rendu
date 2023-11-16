// dictionary.go

package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	filename string
	entries  map[string]Entry
	addCh    chan entryOperation
	removeCh chan string
}

type entryOperation struct {
	word       string
	definition string
}

func New(filename string) *Dictionary {
	d := &Dictionary{
		filename: filename,
		entries:  make(map[string]Entry),
		addCh:    make(chan entryOperation),
		removeCh: make(chan string),
	}

	go d.startWorker()
	return d
}

func (d *Dictionary) startWorker() {
	for {
		select {
		case op := <-d.addCh:
			d.load()
			entry := Entry{Definition: op.definition}
			d.entries[op.word] = entry
			d.save()

		case word := <-d.removeCh:
			d.load()
			delete(d.entries, word)
			d.save()
		}
	}
}

func (d *Dictionary) load() error {
	file, err := ioutil.ReadFile(d.filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &d.entries)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) save() error {
	file, err := json.MarshalIndent(d.entries, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(d.filename, file, 0644)
}

func (d *Dictionary) Add(word, definition string) {
	d.addCh <- entryOperation{word, definition}
}

func (d *Dictionary) Remove(word string) {
	d.removeCh <- word
}

func (d *Dictionary) Get(word string) (Entry, error) {
	d.load()

	entry, found := d.entries[word]
	if found {
		return entry, nil
	}
	return Entry{}, fmt.Errorf("Mot non trouvé: %s", word)
}

func (d *Dictionary) List() {
	d.load()

	var words []string
	for word := range d.entries {
		words = append(words, word)
	}

	sort.Strings(words)

	fmt.Println("Liste triée des mots et définitions:")
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, d.entries[word])
	}
}
