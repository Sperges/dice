package parser

import (
	"fmt"
	"math/rand"
	"roll/src/token"
	"sort"
	"strconv"
)

type Node struct {
	values []int64
	lhs    *Node
	kind   Kind
	rhs    *Node
}

func (n *Node) Kind() Kind {
	return n.kind
}

func ParseStream(tokens *token.Stream, lastKind token.Kind) (*Node, error) {
	var leftNode *Node

	currentToken := tokens.Peek()

	for currentToken.Kind != token.EOF {

		if currentToken.Kind == token.Error {
			return leftNode, fmt.Errorf("error: %q found at %s", currentToken.Text, currentToken.Position)
		} else if currentToken.Kind == token.RParen {
			break
		} else if currentToken.Kind == token.Number {
			tokens.Next()
			value, _ := strconv.ParseInt(currentToken.Text, 0, 64)
			leftNode = &Node{
				values: []int64{value},
				kind:   Number,
			}
		} else if currentToken.Kind == token.LParen {
			tokens.Next()
			if expr, err := ParseStream(tokens, 0); err != nil {
				return expr, err
			} else {
				nextToken := tokens.Next()
				if nextToken.Kind != token.RParen {
					return expr, fmt.Errorf("error: expected right paren at %s", nextToken.Position)
				} else {
					leftNode = expr
				}
			}
		} else if currentToken.Kind <= lastKind {
			break
		} else if currentToken.IsOp() {
			tokens.Next()
			if rightNode, err := ParseStream(tokens, currentToken.Kind); err != nil {
				return leftNode, err
			} else {
				leftNode = &Node{
					lhs:  leftNode,
					kind: FromTokenKind(currentToken.Kind),
					rhs:  rightNode,
				}
			}
		}

		currentToken = tokens.Peek()
	}

	return leftNode, nil
}

func (n *Node) String() string {
	switch n.kind {
	case Number:
		return fmt.Sprint(n.values[0])
	default:
		return fmt.Sprintf("%s {lhs: %s, rhs: %s}", n.kind, n.lhs.String(), n.rhs.String())
	}
}

func (n *Node) EvaluateAndSum() int64 {
	return sum(n.Evaluate())
}

func (n *Node) Evaluate() []int64 {
	switch n.kind {
	case Number:
		return []int64{sum(n.values)}
	case Add:
		return evaluateAdd(n.lhs.Evaluate(), n.rhs.Evaluate())
	case Sub:
		return evaluateSub(n.lhs.Evaluate(), n.rhs.Evaluate())
	case Mul:
		return evaluateMul(n.lhs.Evaluate(), n.rhs.Evaluate())
	case Div:
		return evaluateDiv(n.lhs.Evaluate(), n.rhs.Evaluate())
	case Dice:
		return evaluateDice(sum(n.lhs.Evaluate()), sum(n.rhs.Evaluate()))
	case KeepHigh:
		return evaluateKeepHigh(n.lhs.Evaluate(), sum(n.rhs.Evaluate()))
	case KeepLow:
		return evaluateKeepLow(n.lhs.Evaluate(), sum(n.rhs.Evaluate()))
	case DropHigh:
		return evaluateDropHigh(n.lhs.Evaluate(), sum(n.rhs.Evaluate()))
	case DropLow:
		return evaluateDropLow(n.lhs.Evaluate(), sum(n.rhs.Evaluate()))
	}
	panic("unreachable")
}

func evaluateAdd(lhs, rhs []int64) []int64 {
	return []int64{sum(lhs) + sum(rhs)}
}

func evaluateSub(lhs, rhs []int64) []int64 {
	return []int64{sum(lhs) - sum(rhs)}
}

func evaluateMul(lhs, rhs []int64) []int64 {
	return []int64{sum(lhs) * sum(rhs)}
}

func evaluateDiv(lhs, rhs []int64) []int64 {
	return []int64{sum(lhs) / sum(rhs)}
}

func evaluateDice(dice, faces int64) []int64 {
	results := []int64{}
	var i int64 = 0
	for i < dice {
		result := rand.Int63n(faces) + 1
		fmt.Printf("d%d => %d; ", faces, result)
		results = append(results, result)
		i++
	}
	println()
	return results
}

func evaluateKeepHigh(nums []int64, k int64) []int64 {
	if k > int64(len(nums)) {
		return nums
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[int64(len(nums))-k:]
}

func evaluateKeepLow(nums []int64, k int64) []int64 {
	if k > int64(len(nums)) {
		return nums
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[:k]
}

func evaluateDropHigh(nums []int64, k int64) []int64 {
	if k > int64(len(nums)) {
		return []int64{0}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[:int64(len(nums))-k]
}

func evaluateDropLow(nums []int64, k int64) []int64 {
	if k > int64(len(nums)) {
		return []int64{0}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[k:]
}

func sum(nums []int64) int64 {
	var total int64
	for _, v := range nums {
		total += v
	}
	return total
}
