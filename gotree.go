// Package gotree create and print tree.
package gotree

import (
	"strings"
)

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

type (
	tree struct {
		item     TreeItemText
		children []Tree
	}

	TreeItemText interface {
		TreeItemText() string
	}

	// Tree is tree interface
	Tree interface {
		Add(s TreeItemText) Tree
		AddTree(tree Tree)
		Items() []Tree
		Item() TreeItemText
		Print() string
	}

	printer struct {
	}

	// Printer is printer interface
	Printer interface {
		Print(Tree) string
	}
)

// New returns a new GoTree.Tree
func New(s TreeItemText) Tree {
	return &tree{
		item:     s,
		children: []Tree{},
	}
}

// Add adds a node to the tree
func (t *tree) Add(s TreeItemText) Tree {
	n := New(s)
	t.children = append(t.children, n)
	return n
}

// AddTree adds a tree as an item
func (t *tree) AddTree(tree Tree) {
	t.children = append(t.children, tree)
}

// Item returns the node's value
func (t *tree) Item() TreeItemText {
	return t.item
}

// Items returns all children in the tree
func (t *tree) Items() []Tree {
	return t.children
}

// Print returns an visual representation of the tree
func (t *tree) Print() string {
	return newPrinter().Print(t)
}

func newPrinter() Printer {
	return &printer{}
}

// Print prints a tree to a string
func (p *printer) Print(t Tree) string {
	return t.Item().TreeItemText() + newLine + p.printItems(t.Items(), []bool{})
}

func (p *printer) printText(text string, spaces []bool, last bool) string {
	var result string
	for _, space := range spaces {
		if space {
			result += emptySpace
		} else {
			result += continueItem
		}
	}

	indicator := middleItem
	if last {
		indicator = lastItem
	}

	var out string
	lines := strings.Split(text, "\n")
	for i := range lines {
		text := lines[i]
		if i == 0 {
			out += result + indicator + text + newLine
			continue
		}
		if last {
			indicator = emptySpace
		} else {
			indicator = continueItem
		}
		out += result + indicator + text + newLine
	}

	return out
}

func (p *printer) printItems(t []Tree, spaces []bool) string {
	var result string
	for i, f := range t {
		last := i == len(t)-1
		result += p.printText(f.Item().TreeItemText(), spaces, last)
		if len(f.Items()) > 0 {
			spacesChild := append(spaces, last)
			result += p.printItems(f.Items(), spacesChild)
		}
	}
	return result
}
