package main

import (
	"ESTIAM/dictionary"
	"bufio"
	"fmt"
	"os"
)

func main() {
	d := dictionary.New()
	reader := bufio.NewReader(os.Stdin)

	actionAdd(d, reader)
	actionDefine(d, reader)
	actionRemove(d, reader)
	actionList(d)
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Mot à ajouter : ")
	word, _ := reader.ReadString('\n')

	fmt.Print("Définition : ")
	definition, _ := reader.ReadString('\n')

	d.Add(word, definition)
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Mot à rechercher : ")
	word, _ := reader.ReadString('\n')

	definition, err := d.Get(word)
	if err == nil {
		fmt.Printf("La définition de %s est: %s\n", word, definition)
	} else {
		fmt.Printf("Mot non trouvé: %s\n", word)
	}
}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Mot à supprimer : ")
	word, _ := reader.ReadString('\n')

	d.Remove(word)
	fmt.Printf("Le mot %s a été supprimé.\n", word)
}

func actionList(d *dictionary.Dictionary) {
	fmt.Println("Liste triée des mots et définitions:")
	d.List()
}
