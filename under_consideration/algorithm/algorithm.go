package main

import "fmt"

type Node struct {
	Name     string
	Children []*Node
}

var shogun = []string{"家康", "秀忠", "家光", "家綱", "綱吉", "家宣", "家継", "吉宗", "家重", "家治", "家斉", "家慶", "家定", "家茂", "慶喜"}

func findShogunNumber(name string) string {
	for index, n := range shogun {
		if n == name {
			return fmt.Sprintf(" %d", index+1)
		}
	}
	return ""
}

func main() {
	parentChilds := [][2]string{
		{"家康", "信康"}, {"家康", "秀康"}, {"家康", "秀忠"}, {"家康", "忠輝"}, {"家康", "義直"},
		{"秀忠", "家光"}, {"秀忠", "忠長"}, {"秀忠", "和子"}, {"秀忠", "正之"}, {"家光", "家綱"},
		{"家光", "綱重"}, {"家光", "綱吉"}, {"綱重", "家宣"}, {"家宣", "家継"}, {"家康", "頼信"},
		{"頼信", "光貞"}, {"光貞", "吉宗"}, {"吉宗", "家重"}, {"家重", "家治"}, {"家重", "重好"},
		{"吉宗", "宗武"}, {"宗武", "定信"}, {"吉宗", "宗尹"}, {"宗尹", "治済"}, {"治済", "家斉"},
		{"治済", "斉敦"}, {"治済", "斉匡"}, {"家斉", "家慶"}, {"家慶", "家定"}, {"家斉", "斉順"},
		{"斉順", "家茂"}, {"斉匡", "慶頼"}, {"慶頼", "家達"}, {"家康", "頼房"}, {"頼房", "頼重"},
		{"頼重", "頼侯"}, {"頼侯", "頼豊"}, {"頼豊", "宗堯"}, {"宗堯", "宗翰"}, {"宗翰", "治保"},
		{"治保", "治紀"}, {"治紀", "斉昭"}, {"斉昭", "慶喜"},
	}

	root := &Node{Name: "家康"}
	buildTree(root, parentChilds)
	printTree(root, "")
}

func buildTree(parent *Node, relations [][2]string) {
	for _, relation := range relations {
		if relation[0] == parent.Name {
			child := &Node{Name: relation[1]}
			parent.Children = append(parent.Children, child)
			buildTree(child, relations)
		}
	}
}

func printTree(node *Node, prefix string) {
	fmt.Printf("%s%s%s\n", prefix, node.Name, findShogunNumber(node.Name))
	for _, child := range node.Children {
		printTree(child, prefix+"+-")
	}
}
