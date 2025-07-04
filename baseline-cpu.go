package main
import (
	"fmt"
	"math/rand"
	"time"
)
type Node struct {
	Name     string
	CPUUsage float64
}
func generateNodes(n int) []Node {
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			Name:     fmt.Sprintf("node-%d", i+1),
			CPUUsage: 70 + rand.Float64()*20,
		}
	}
	return nodes
}
func printHeader() {
	fmt.Printf("| %-10s | %-22s |\n", "Node", "Baseline CPU Usage (%)")
	printLine()
}
func printNodes(nodes []Node) {
	for _, node := range nodes {
		fmt.Printf("| %-10s | %-22.2f |\n", node.Name, node.CPUUsage)
	}
}
func baseline3Nodes() {
	nodes := generateNodes(3)
	printHeader()
	printNodes(nodes)
	printLine()
}
func baseline5Nodes() {
	nodes := generateNodes(5)
	printHeader()
	printNodes(nodes)
	printLine()
}
func baseline7Nodes() {
	nodes := generateNodes(7)
	printHeader()
	printNodes(nodes)
	printLine()
}
func baseline9Nodes() {
	nodes := generateNodes(9)
	printHeader()
	printNodes(nodes)
	printLine()
}
func baseline11Nodes() {
	nodes := generateNodes(11)
	printHeader()
	printNodes(nodes)
	printLine()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	baseline3Nodes()
	baseline5Nodes()
	baseline7Nodes()
	baseline9Nodes()
	baseline11Nodes()
}
