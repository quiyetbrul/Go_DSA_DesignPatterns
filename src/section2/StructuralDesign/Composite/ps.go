package composite

type Parent struct {
	SomeField int
}

type Son struct {
	// Parent inheritance, cannot use son (type Son) as type Parent in in argument to GetParentField
	P Parent
}

func GetParentField(p *Parent) int {
	return p.SomeField
}
