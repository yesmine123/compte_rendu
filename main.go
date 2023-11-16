// main.go

package main

import (
	"ESTIAM/dictionary"
	"bufio"
	"fmt"
	"os"
	"sync"
)

const filename = "dictionary.json"

func main() {
	d := dictionary.New(filename)
	var wg sync.WaitGroup
	userInputDone := make(chan struct{})

	wg.Add(3)

	var reader = bufio.NewReader(os.Stdin)

	go func() {
		defer wg.Done()
		userInput(d, reader, userInputDone)
	}()

	go func() {
		defer wg.Done()
		<-userInputDone // Attend que userInput signale qu'il a terminé
		actionAdd(d, reader)
	}()

	go func() {
		defer wg.Done()
		<-userInputDone // Attend que userInput signale qu'il a terminé
		actionRemove(d, reader)
	}()

	wg.Wait()

	actionList(d)
}

// ...

func userInput(d *dictionary.Dictionary, reader *bufio.Reader, done chan struct{}) {
	for {
		fmt.Print("Choisissez une action (1: Ajouter, 2: Supprimer, 3: Quitter): ")
		var choice int
		_, err := fmt.Fscanf(reader, "%d\n", &choice)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}

		switch choice {
		case 1:
			actionAdd(d, reader)
		case 2:
			actionRemove(d, reader)
		case 3:
			close(done) // Ferme le canal pour signaler que l'entrée utilisateur est terminée
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Mot à ajouter : ")
	word, _ := reader.ReadString('\n')

	fmt.Print("Définition : ")
	definition, _ := reader.ReadString('\n')

	d.Add(word, definition)
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
