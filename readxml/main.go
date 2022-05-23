package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/subchen/go-xmldom"
)

func main() {
	fp, err := os.Open("./a.xml")
	if err != nil {
		panic(err)
	}
	xml, err := ioutil.ReadAll(fp)
	doc := xmldom.Must(xmldom.ParseXML(string(xml)))
	// fmt.Println(doc.Root)
	// 再帰か、、
	// interpreterのあれでコピる
	/*
		for _, child := range doc.Root.Children {
			fmt.Println("===")
			fmt.Println(child.Name)
			fmt.Println(child.Text)
			fmt.Println("===")
		}
	*/
	head := doc.Root
	stack := Stack{
		arr: []string{head.Name},
	}
	// helper(head)
	stack.helper(head)
	for _, item := range stack.arr {
		fmt.Println(item)
	}

	pointer := len(stack.arr) - 1
	for pointer > 0 {
		n1, _ := strconv.ParseFloat(stack.arr[pointer], 64)
		pointer--
		n2, _ := strconv.ParseFloat(stack.arr[pointer], 64)
		pointer--
		sign := stack.arr[pointer]

		switch sign {
		case "plus":
			stack.arr[pointer] = strconv.FormatFloat(n2+n1, 'f', 2, 64)
		case "minus":
			stack.arr[pointer] = strconv.FormatFloat(n2-n1, 'f', 2, 64)
		case "mul":
			stack.arr[pointer] = strconv.FormatFloat(n2*n1, 'f', 2, 64)
		case "div":
			stack.arr[pointer] = strconv.FormatFloat(n2/n1, 'f', 2, 64)
		}
	}
	fmt.Println(stack.arr)
}

// こっからはあのコンパイラと同じ、はず

type Stack struct {
	arr []string
}

func (s *Stack) helper(head *xmldom.Node) {
	if head == nil {
		return
	}
	for _, child := range head.Children {
		if child.Name == "num" {
			s.arr = append(s.arr, child.Text)
		} else {
			if child.Text == "" {
				s.arr = append(s.arr, child.Name)
			} else {
				s.arr = append(s.arr, child.Name, child.Text)
			}
		}
		s.helper(child)
	}
}

// define a stack and process it?

/*
func helper(head *xmldom.Node) {
	if head == nil {
		return
	}
	for _, child := range head.Children {
		fmt.Println("===")
		fmt.Println(child.Name)
		fmt.Println(child.Text)
		helper(child)
	}
}
*/
