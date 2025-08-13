// 一致性哈希算法的简单实现
package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// Node 表示一个物理节点
type Node struct {
	Name string // 节点名称
}

type ConsistentHash struct {
	// 哈希环
	hashRing map[uint32]*Node
	// 虚拟节点的数量
	numVirtualNodes int
	// 排序后的哈希值
	sortedHashes []uint32
}

func NewConsistentHash(numVirtualNodes int) *ConsistentHash {
	return &ConsistentHash{
		hashRing:        make(map[uint32]*Node),
		numVirtualNodes: numVirtualNodes,
	}
}

func hash(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

// AddNode 向一致性哈希环中添加一个节点
func (c *ConsistentHash) AddNode(node *Node) {
	// 向哈希环中添加虚拟节点
	for i := 0; i < c.numVirtualNodes; i++ {
		// 生成虚拟节点的名称
		virtualNodeName := fmt.Sprintf("%s-%d", node.Name, i)
		// 计算虚拟节点的哈希值
		virtualNodeHash := hash(virtualNodeName)
		// 将虚拟节点添加到哈希环中
		c.hashRing[virtualNodeHash] = node
		// 将虚拟节点的哈希值添加到排序数组中
		c.sortedHashes = append(c.sortedHashes, virtualNodeHash)
	}
	// 对哈希值进行排序
	sort.Slice(c.sortedHashes, func(i, j int) bool {
		return c.sortedHashes[i] < c.sortedHashes[j]
	})
}

// RemoveNode 从一致性哈希环中删除一个节点
func (c *ConsistentHash) RemoveNode(node *Node) {
	for i := 0; i < c.numVirtualNodes; i++ {
		// 生成虚拟节点的名称
		virtualNodeName := fmt.Sprintf("%s-%d", node.Name, i)
		// 计算虚拟节点的哈希值
		virtualNodeHash := hash(virtualNodeName)
		// 从哈希环中删除虚拟节点
		delete(c.hashRing, virtualNodeHash)
		// 从排序数组中删除虚拟节点的哈希值
		for j, hash := range c.sortedHashes {
			if hash == virtualNodeHash {
				c.sortedHashes = append(c.sortedHashes[:j], c.sortedHashes[j+1:]...)
				break
			}
		}
	}
}

// GetNode 获取一个数组的哈希值对应的节点
func (c *ConsistentHash) GetNode(key string) *Node {
	// 计算哈希值
	hashValue := hash(key)
	// 在哈希环中查找哈希值对应的节点
	for _, hash := range c.sortedHashes {
		if hashValue <= hash {
			return c.hashRing[hash]
		}
	}
	// 如果哈希值大于哈希环中的所有哈希值，则返回哈希环中的第一个节点
	return c.hashRing[c.sortedHashes[0]]
}

func main() {

	//创建一个一致性哈希对象，假设每个节点有3个虚拟节点
	ch := NewConsistentHash(3)
	// 添加节点到哈希环
	node1 := &Node{Name: "Node1"}
	node2 := &Node{Name: "Node2"}
	node3 := &Node{Name: "Node3"}

	ch.AddNode(node1)
	ch.AddNode(node2)
	ch.AddNode(node3)

	// 获取一些数据的哈希值对应的节点
	data1 := "data13234"
	data2 := "data2sdfsdfsfds"
	data3 := "data32323fffdsfd"

	fmt.Printf("Data %s is stored in %s\n", data1, ch.GetNode(data1).Name)
	fmt.Printf("Data %s is stored in %s\n", data2, ch.GetNode(data2).Name)
	fmt.Printf("Data %s is stored in %s\n", data3, ch.GetNode(data3).Name)

	// 移除一个节点
	ch.RemoveNode(node3)

	// 再次获取数据的节点
	fmt.Printf("Data %s is now stored in %s\n", data1, ch.GetNode(data1).Name)
	fmt.Printf("Data %s is now stored in %s\n", data2, ch.GetNode(data2).Name)
	fmt.Printf("Data %s is now stored in %s\n", data3, ch.GetNode(data3).Name)

}
