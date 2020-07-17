package main

type Complex struct {
	Name string
}

func fetch() ([]Complex, error) {
	return []Complex{
		{Name: "Dorian"},
		{Name: "Gustavo"},
	}, nil
}
