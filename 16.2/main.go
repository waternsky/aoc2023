package main

import (
	"fmt"
	"os"
	"strings"
)

func energisedTiles(inp string, start node) []node {
	grid := strings.Split(inp, "\n")
	stack := stack[node]{}
	stack.Push(start)
	var visited []node
	for true {
		if len(stack) == 0 {
			break
		}
		n := stack.Pop()
		visited = append(visited, n)
		nxt := nextNode(grid, n)
		//fmt.Println("Next: ", nxt)
		for _, v := range nxt {
			if !isIn[node](v, visited) {
				stack.Push(v)
			}
		}
	}
	return visited
}

func nextNode(grid []string, curr node) []node {
	s := string(grid[curr.position[0]][curr.position[1]])
	//fmt.Println(s)
	switch s {
	case ".":
		n := node{position: curr.position.Add(curr.dir), dir: curr.dir}
		if validNode(n, grid) {
			return []node{n}
		}
	case "|":
		if curr.dir.X == 0 {
			n := node{position: curr.position.Add(curr.dir), dir: curr.dir}
			if validNode(n, grid) {
				return []node{n}
			}
		} else {
			up := node{position: curr.position.Add(Vector{0, 1}), dir: Vector{0, 1}}
			down := node{position: curr.position.Add(Vector{0, -1}), dir: Vector{0, -1}}
			var total []node
			if validNode(up, grid) {
				total = append(total, up)
			}
			if validNode(down, grid) {
				total = append(total, down)
			}
			return total
		}
	case "-":
		if curr.dir.Y == 0 {
			n := node{position: curr.position.Add(curr.dir), dir: curr.dir}
			if validNode(n, grid) {
				return []node{n}
			}
		} else {
			up := node{position: curr.position.Add(Vector{1, 0}), dir: Vector{1, 0}}
			down := node{position: curr.position.Add(Vector{-1, 0}), dir: Vector{-1, 0}}
			var total []node
			if validNode(up, grid) {
				total = append(total, up)
			}
			if validNode(down, grid) {
				total = append(total, down)
			}
			return total
		}
	case "\\":
		if curr.dir.X == -1 {
			n := node{position: curr.position.Add(Vector{0, 1}), dir: Vector{0, 1}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
		if curr.dir.X == 1 {
			n := node{position: curr.position.Add(Vector{0, -1}), dir: Vector{0, -1}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
		if curr.dir.Y == -1 {
			n := node{position: curr.position.Add(Vector{1, 0}), dir: Vector{1, 0}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
		if curr.dir.Y == 1 {
			n := node{position: curr.position.Add(Vector{-1, 0}), dir: Vector{-1, 0}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
	case "/":
		if curr.dir.X == -1 {
			n := node{position: curr.position.Add(Vector{0, -1}), dir: Vector{0, -1}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
		if curr.dir.X == 1 {
			n := node{position: curr.position.Add(Vector{0, 1}), dir: Vector{0, 1}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
		if curr.dir.Y == -1 {
			n := node{position: curr.position.Add(Vector{-1, 0}), dir: Vector{-1, 0}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
		if curr.dir.Y == 1 {
			n := node{position: curr.position.Add(Vector{1, 0}), dir: Vector{1, 0}}
			if validNode(n, grid) {
				return []node{n}
			}
		}
	}
	return nil
}

func isIn[T comparable](elm T, arr []T) bool {
	for _, v := range arr {
		if v == elm {
			return true
		}
	}
	return false
}

func validNode(n node, grid []string) bool {
	if n.position[0] >= 0 && n.position[0] < len(grid) && n.position[1] >= 0 && n.position[1] < len(grid[0]) {
		return true
	}
	return false
}

type stack[T any] []T

func (s *stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *stack[T]) Pop() T {
	ans := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ans
}

type Vector struct {
	X int
	Y int
}

type cell [2]int

func (c cell) Add(v Vector) cell {
	var ans cell
	ans[1] = c[1] + v.X
	ans[0] = c[0] - v.Y // down needs to add to lineIdx
	return ans
}

type node struct {
	position cell
	dir      Vector
}

func maxEnergizer(inp string) int {
    ans := 0
    grid := strings.Split(inp, "\n")
    var startPos []node
    for i := 0; i < len(grid[0]); i++ {
        startPos = append(startPos, node{position: cell{0, i}, dir: Vector{0, -1}})
        startPos = append(startPos, node{position: cell{len(grid)-1, i}, dir: Vector{0, 1}})
    }
    for i := 0; i < len(grid); i++ {
        startPos = append(startPos, node{position: cell{i, 0}, dir: Vector{1, 0}})
        startPos = append(startPos, node{position: cell{i, len(grid[0])-1}, dir: Vector{-1, 0}})
    }
    for _, v := range startPos {
        chk := totalEnergisedTiles(inp, v)
        if ans < chk {
            ans = chk
        }
    }
    return ans
}

func totalEnergisedTiles(inp string, start node) int {
    nodes := energisedTiles(inp, start)
	var ans []cell
	for _, n := range nodes {
		if !isIn[cell](n.position, ans) {
			ans = append(ans, n.position)
		}
	}
    return len(ans)
}

func main() {

    parseData, _ := os.ReadFile("../16.1/input.16.1.txt")
    parseAsStr := string(parseData)
    parseAsStr = strings.TrimSpace(parseAsStr)

    fmt.Println(maxEnergizer(parseAsStr))
}
