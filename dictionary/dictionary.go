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
}

func New(filename string) *Dictionary {
	return &Dictionary{
		filename: filename,
		entries:  make(map[string]Entry),
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
	d.load() // Charger les données depuis le fichier

	entry := Entry{
		Definition: definition,
	}
	d.entries[word] = entry

	err := d.save() // Sauvegarder les données dans le fichier
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde dans le fichier:", err)
	}

	fmt.Println("Mot ajouté avec succès.")
}

func (d *Dictionary) Get(word string) (Entry, error) {
	d.load() // Charger les données depuis le fichier

	entry, found := d.entries[word]
	if found {
		return entry, nil
	}
	return Entry{}, fmt.Errorf("Mot non trouvé: %s", word)
}

func (d *Dictionary) Remove(word string) {
	d.load() // Charger les données depuis le fichier

	delete(d.entries, word)

	err := d.save() // Sauvegarder les données dans le fichier
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde dans le fichier:", err)
	}

	fmt.Printf("Le mot %s a été supprimé.\n", word)
}

func (d *Dictionary) List() {
	d.load() // Charger les données depuis le fichier

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
