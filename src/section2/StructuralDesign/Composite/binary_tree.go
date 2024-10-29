package composite

/*
A binary tree is a good example to use the Composite pattern
because it naturally fits the part-whole hierarchy that
the Composite pattern is designed to represent.
In a binary tree, each node can be treated as an individual object or
as a composition of other nodes (its children).
This allows for uniform treatment of both individual nodes and subtrees,
making the Composite pattern an ideal choice.
*/

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
