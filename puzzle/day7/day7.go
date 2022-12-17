package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"regexp"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

type Node struct {
	Name     string
	Size     int
	Dir      bool
	Children map[string]*Node
}

func (n Node) String() string {
	return n.Name
}

func (d Day) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	tree := &Node{
		Name: "/", Dir: true, Children: make(map[string]*Node),
	}

	var path []string

	for buf.Scan() {
		l := buf.Text()
		// log.Printf("line %s\n", l)

		switch {
		case l == "$ cd /":
			path = []string{}
		case l == "$ cd ..":
			path = path[:len(path)-1]
		case cdToChild(l):
			var x string
			fmt.Sscanf(l, "$ cd %s", &x)
			path = append(path, x)
		case l == "$ ls":
			break
		case l[:3] == "dir":
			var d Node
			fmt.Sscanf(l, "dir %s", &d.Name)
			d.Dir = true
			d.Children = make(map[string]*Node)
			addNode(tree, path, &d)
		default:
			var f Node
			fmt.Sscanf(l, "%d %s", &f.Size, &f.Name)
			addNode(tree, path, &f)
		}
	}

	var sum int
	walkTree(tree, &sum, func(n *Node, i *int) {
		if n.Dir && n.Size <= 100000 {
			*i += n.Size
		}
	})

	del := math.MaxInt
	walkTree(tree, &del, func(n *Node, i *int) {
		spaceReq := 30000000 - (70000000 - tree.Size)
		if n.Dir && n.Size >= spaceReq && n.Size < del {
			*i = n.Size
		}
	})

	// spaceReq := 30000000 - (70000000 - tree.Size)
	// log.Printf("Used space %d, free-up %d\n", tree.Size, spaceReq)

	log.Printf("Answer part I: %d\n", sum)
	log.Printf("Answer part II: %d\n", del)
}

func cdToChild(l string) bool {
	match, _ := regexp.MatchString("^\\$ cd [a-z\\.]+$", l)
	return match
}

func addNode(tree *Node, path []string, node *Node) {
	if !node.Dir {
		tree.Size += node.Size
	}

	if len(path) == 0 {
		tree.Children[node.Name] = node
		return
	}

	child, ok := tree.Children[path[0]]
	if ok {
		addNode(child, path[1:], node)
	}
}

// /a/b/c
func walkTree(tree *Node, agg *int, f func(*Node, *int)) {
	// fmt.Printf("%s: %d (%v)\n", tree.Name, tree.Size, tree.Children)
	f(tree, agg)

	for _, node := range tree.Children {
		walkTree(node, agg, f)
	}
}
