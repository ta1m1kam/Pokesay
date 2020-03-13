package pokesay

import "strings"

func Say(options ...Option) (string, error) {
	poke, err := NewPoke(options...)
	if err != nil {
		return "", err
	}

	mon, err := poke.GetPoke(0)
	if err != nil {
		return "", err
	}

	return mon, nil
}

func (poke *Poke) GetPoke(thoughts rune) (string, error) {
	src, err := Asset(poke.typ)
	if err != nil {
		return "", err
	}

	if thoughts == 0 {
		if poke.thinking {
			thoughts = 'o'
		} else {
			thoughts = '\\'
		}
	}

	r := strings.NewReplacer(
		"\\\\", "\\",
		"\\@", "@",
		"\\$", "$",
		"$thoughts", string(thoughts),
		"${thoughts}", string(thoughts),
		)
	newsrc := r.Replace(string(src))
	separate := strings.Split(newsrc, "\n")
	mon := make([]string, 0, len(separate))
	for _, line := range separate {
		if strings.Contains(line, "$the_cow = <<EOC") || strings.HasPrefix(line, "##") {
			continue
		}

		if strings.HasPrefix(line, "EOC") {
			break
		}

		mon = append(mon, line)
	}
	return strings.Join(mon, "\n"), nil
}
