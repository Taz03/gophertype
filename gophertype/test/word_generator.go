package test

import (
	"embed"
	"encoding/json"
	"math/rand"
)

type Generator struct {
	Language string   `json:"name"`
	Words    []string `json:"words"`
}

func NewGenerator(resourcesFs embed.FS, language string) (g *Generator) {
	data, _ := resourcesFs.ReadFile("languages/" + language + ".json")
	json.Unmarshal(data, &g)

	return
}

func (g *Generator) GenerateWord() string {
	return g.Words[rand.Intn(len(g.Words))]
}

func (g *Generator) GenerateWords(n int) []string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = g.GenerateWord()
	}

	return words
}
