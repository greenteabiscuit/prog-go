package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/subchen/go-xmldom"
)

func main() {
	flag.Parse()
	filename := flag.Args()[0]
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	xml, err := ioutil.ReadAll(fp)
	doc := xmldom.Must(xmldom.ParseXML(string(xml)))
	head := doc.Root

	if head.Name == "num" {
		n1, _ := strconv.ParseFloat(head.Text, 64)
		fmt.Println(n1)
		return
	}

	for {
		// one of mul, sub, add, div
		if head.Name == "num" || len(head.Children) == 1 {
			head = head.Children[0]
		} else {
			break
		}
	}

	p := &Parser{}

	exp := p.parseExpression(head)
	// fmt.Println(exp.String())
	//fmt.Println(exp.Float())
	ans := Eval(exp)
	fmt.Println(ans.Float())
}

type FloatObj struct {
	Value float64
}

func (i *FloatObj) Float() float64 {
	return i.Value
}

func (i *FloatObj) String() string {
	return ""
}

func Eval(exp Expression) Expression {
	switch node := exp.(type) {
	case *InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		// fmt.Println(left.Float())
		// fmt.Println(right.Float())
		leftVal := left.(*FloatObj).Value
		rightVal := right.(*FloatObj).Value
		switch node.Token {
		case "plus":
			return &FloatObj{Value: leftVal + rightVal}
		case "minus":
			return &FloatObj{Value: leftVal - rightVal}
		case "mul":
			return &FloatObj{Value: leftVal * rightVal}
		case "div":
			return &FloatObj{Value: leftVal / rightVal}
		}
	case *FloatLiteral:
		return &FloatObj{Value: node.Value}
	default:
		fmt.Println("failure")
	}
	return nil
}

type Parser struct {
}

func (p *Parser) parseExpression(head *xmldom.Node) Expression {
	if head.Name == "num" {
		n1, _ := strconv.ParseFloat(head.Text, 64)
		return &FloatLiteral{Value: n1}
	}
	sign := head.Name

	expression := &InfixExpression{
		Token: sign,
		Left:  p.parseExpression(head.Children[0]),
		Right: p.parseExpression(head.Children[1]),
	}

	for i := 2; i < len(head.Children); i++ {
		expression = &InfixExpression{
			Token: sign,
			Left:  expression,
			Right: p.parseExpression(head.Children[i]),
		}
	}
	return expression
}

type Node interface {
	String() string
	Float() float64
}

type Expression interface {
	Node
}

type InfixExpression struct {
	Token string // the operator token
	Left  Expression
	Right Expression
}

type FloatLiteral struct {
	Value float64
}

func (il *FloatLiteral) String() string {
	return strconv.FormatFloat(il.Value, 'f', 6, 64)
}
func (il *FloatLiteral) Float() float64 { return il.Value }

func (ie *InfixExpression) Float() float64 { return 0 }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Token + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
