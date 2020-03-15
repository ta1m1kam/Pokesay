package pokesay

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

type Poke struct {
	phrase string
	thinking bool
	typ string
}

func NewPoke(options ...Option) (*Poke, error) {
	poke := &Poke{
		phrase: "",
		typ: "images/Pikachu.png",
	}

	for _, o := range options {
		if err := o(poke); err != nil {
			return  nil, err
		}
	}
	return poke, nil
}

type Option func(*Poke) error

func Random() Option {
	return func(poke *Poke) error {
		s := pickPoke()
		s = complementFilePath(s)
		poke.typ = s
		return nil
	}
}

func Phrase(s string) Option {
	return func(poke *Poke) error {
		poke.phrase = s
		return nil
	}
}

func Thinking() Option {
	return func(poke *Poke) error {
		poke.thinking = false
		return nil
	}
}

func Type(s string) Option {
	s = complementFilePath(s)

	return func(poke *Poke) error {
		containPoke, err := containPokes(s)
		if err != nil {
			return err
		}
		if containPoke {
			poke.typ = s
			return nil
		}
		return errors.Errorf("Could not find %s", s)
	}
}

func containPokes(s string) (bool, error) {
	for _, image := range AssetNames() {
		if strings.Contains(s, image) {
			return true, nil
		}
	}
	return false, nil
}

func AssetNames() []string {
	var assetNames []string
	files, err := assetFiles()
	if err != nil {
		return nil
	}

	for _, file := range files {
		assetNames = append(assetNames, file.Name())
	}

	return assetNames
}

func assetFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir("images")
	if err != nil {
		return nil, err
	}

	return files, err
}

func complementFilePath(s string) string {
	if s == "" {
		s = "images/Pikachu.png"
	}

	if !strings.HasSuffix(s, ".png") {
		s += ".png"
	}
	if !strings.HasPrefix(s, "pokes/") {
		s = "images/" + s
	}

	return s
}
