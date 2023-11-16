package dictionary

import (
	"fmt"
	"sort"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {

	return e.Definition
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{
		Definition: definition,
	}
	d.entries[word] = entry
	fmt.Println("Mot ajouté avec succès.")
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entry, found := d.entries[word]
	if found {
		return entry, nil
	}
	return Entry{}, fmt.Errorf("Mot non trouvé: %s", word)
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
	fmt.Printf("Le mot %s a été supprimé.\n", word)
}

func (d *Dictionary) List() {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}

	// Trier les mots.
	sort.Strings(words)

	// Afficher les mots et leurs définitions.
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, d.entries[word])
	}
}
