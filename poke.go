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
	if s == "" {
		s = "images/Pikachu.png"
	}

	if !strings.HasSuffix(s, ".png") {
		s += ".png"
	}
	if !strings.HasPrefix(s, "pokes/") {
		s = "images/" + s
	}

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
	imageFiles, err := assetFiles()
	if err != nil {
		return false, err
	}
	for _, image := range imageFiles {
		if strings.Contains(s, image.Name()) {
			return true, nil
		}
	}
	return false, nil
}

func assetFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir("images")
	if err != nil {
		return nil, err
	}

	return files, err
}
