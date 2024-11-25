package test

import (
	"embed"
	"strconv"
	"strings"

	"github.com/taz03/gophertype/config"
)

type TestWords struct {
    generator *Generator

    Words []string
}

func NewTestWords(cfg config.Config, resourcesFs embed.FS) *TestWords {
    generator := NewGenerator(resourcesFs, cfg.Language)

    if strings.HasPrefix(cfg.Mode, "time") {
        return &TestWords{
            generator: generator,
            Words: generator.GenerateWords(1000),
        }
    } else if strings.HasPrefix(cfg.Mode, "words") {
        words, _ := strconv.ParseInt(strings.Split(cfg.Mode, " ")[1], 0, 64)
        return &TestWords{
            generator: generator,
            Words: generator.GenerateWords(int(words)),
        }
    }

    return &TestWords{
        Words: []string{"hello", "world"},
    }
}

func (t *TestWords) AddWords(n int) {
    if n == 0 {
        n = 1000
    }

    words := t.generator.GenerateWords(n)
    t.Words = append(t.Words, words...)
}
