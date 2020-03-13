package pokesay

import "github.com/pkg/errors"

type Poke struct {
	phrase string
	thinking bool
	typ string
}

func NewPoke(options ...Option) (*Poke, error) {
	poke := &Poke{
		phrase: "",
		typ: "pokes/Pikachu.cow",
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
		poke.thinking = true
		return nil
	}
}

func Type(s string) Option {
	if s == "" {
		s = "pokes/Pikachu.cow"
	}

	return func(poke *Poke) error {
		if containPokes(s) {
			poke.typ = s
			return nil
		}
		return errors.Errorf("Could not find %s", s)
	}
}

func containPokes(s string) bool {
	return true
}

//func containPokes(s string) bool {
//	for _, poke := range AssetNames() {
//		if s == poke {
//			return true
//		}
//	}
//	return false
//}
