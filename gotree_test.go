package gotree

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleTree() {
	artist := newTestTree("Pantera")
	album := artist.Add(Str("Far Beyond Driven\nsee https://en.wikipedia.org/wiki/Pantera\n(1994)"))
	five := album.Add(Str("5 minutes Alone"))
	five.Add(Str("song by American\ngroove metal"))
	album.Add(Str("I’m Broken"))
	album.Add(Str("Good Friends and a Bottle of Pills"))

	artist.Add(Str("Power Metal\n(1988)"))
	artist.Add(Str("Cowboys from Hell\n(1990)"))
	fmt.Println(artist.Print())

	// Output:
	// Pantera
	// ├── Far Beyond Driven
	// │   see https://en.wikipedia.org/wiki/Pantera
	// │   (1994)
	// │   ├── 5 minutes Alone
	// │   │   └── song by American
	// │   │       groove metal
	// │   ├── I’m Broken
	// │   └── Good Friends and a Bottle of Pills
	// ├── Power Metal
	// │   (1988)
	// └── Cowboys from Hell
	//     (1990)
}

func newTestTree(s string) Tree {
	return New(Str(s))
}

func TestnewTestTree(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			name: "Create new Tree",
			args: args{
				text: "new tree",
			},
			want: &tree{
				item:     Str("new tree"),
				children: []Tree{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTestTree(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTestTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Add(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	type args struct {
		text string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        Tree
		parentCount int
	}{
		{
			name: "Adding a new item into an empty tree",
			args: args{
				text: "child item",
			},
			fields: fields{
				items: []Tree{},
			},
			want: &tree{
				item:     Str("child item"),
				children: []Tree{},
			},
			parentCount: 1,
		},
		{
			name: "Adding a new item into a full tree",
			args: args{
				text: "fourth item",
			},
			fields: fields{
				items: []Tree{
					newTestTree("test"),
					newTestTree("test2"),
					newTestTree("test3"),
				},
			},
			want: &tree{
				item:     Str("fourth item"),
				children: []Tree{},
			},
			parentCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				item:     Str(tt.fields.text),
				children: tt.fields.items,
			}
			got := tree.Add(Str(tt.args.text))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tree.Add() = %v, want %v", got, tt.want)
			}
			if tt.parentCount != len(tree.Items()) {
				t.Errorf("tree total children = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_AddTree(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	type args struct {
		tree Tree
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		itemCount int
	}{
		{
			name: "Adding a new item into an empty tree",
			args: args{
				tree: newTestTree("child item"),
			},
			fields: fields{
				items: []Tree{},
			},
			itemCount: 1,
		},
		{
			name: "Adding a new item into a full tree",
			args: args{
				tree: newTestTree("fourth item"),
			},
			fields: fields{
				items: []Tree{
					newTestTree("test"),
					newTestTree("test2"),
					newTestTree("test3"),
				},
			},
			itemCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				item:     Str(tt.fields.text),
				children: tt.fields.items,
			}
			tree.AddTree(tt.args.tree)
		})
	}
}

func Test_tree_Text(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return the correct value",
			fields: fields{
				text: "item",
			},
			want: "item",
		},
		{
			name: "Return the correct value while empty",
			fields: fields{
				text: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				item:     Str(tt.fields.text),
				children: tt.fields.items,
			}
			if got := tree.Item().TreeItemText(); got != tt.want {
				t.Errorf("tree.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Items(t *testing.T) {
	type fields struct {
		text  string
		items []Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   []Tree
	}{
		{
			name: "Return empty if there is no children under the tree",
			fields: fields{
				text:  "top level item",
				items: []Tree{},
			},
			want: []Tree{},
		},
		{
			name: "Return all children under the tree",
			fields: fields{
				text: "top level item",
				items: []Tree{
					newTestTree("first child"),
					newTestTree("second child"),
				},
			},
			want: []Tree{
				newTestTree("first child"),
				newTestTree("second child"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &tree{
				item:     Str(tt.fields.text),
				children: tt.fields.items,
			}
			if got := tree.Items(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tree.Items() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Print(t *testing.T) {
	threeLevelTree := newTestTree("First Level")
	threeLevelTree.Add(Str("Second level")).Add(Str("Third Level"))

	complexTree := newTestTree("Daft Punk")
	ram := complexTree.Add(Str("Random Access Memories"))
	complexTree.Add(Str("Humam After All"))
	alive := complexTree.Add(Str("Alive 2007"))

	ram.Add(Str("Give Life Back to Music"))
	ram.Add(Str("Giorgio by Moroder"))
	ram.Add(Str("Within"))

	alive.Add(Str("Touch It/Technologic"))
	alive.Add(Str("Face to Face/Too Long"))

	type fields struct {
		tree Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Print a single item tree",
			fields: fields{
				tree: newTestTree("single item"),
			},
			want: `single item
`,
		},
		{
			name: "Print a three level tree",
			fields: fields{
				tree: threeLevelTree,
			},
			want: `First Level
└── Second level
    └── Third Level
`,
		},
		{
			name: "Print a three level tree",
			fields: fields{
				tree: complexTree,
			},
			want: `Daft Punk
├── Random Access Memories
│   ├── Give Life Back to Music
│   ├── Giorgio by Moroder
│   └── Within
├── Humam After All
└── Alive 2007
    ├── Touch It/Technologic
    └── Face to Face/Too Long
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.tree.Print(); got != tt.want {
				t.Errorf("tree.Print() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
