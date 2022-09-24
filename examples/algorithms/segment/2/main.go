package main

// взято отсюда
// https://rtoch.com/posts/golang-segment-tree/

// Поиск подотрезка с максимальной суммой
// https://e-maxx.ru/upload/e-maxx_algo.pdf

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SegmentTreeData struct {
	prefix int
	suffix int
	best   int
	total  int
}

/*
Иными словами, для каждого отрезка дерева отрезков ответ на нём уже предпосчитан, а также дополнительно
ответ посчитан среди всех отрезков, упирающихся в левую границу отрезка, а также среди всех отрезков, упирающихся
в правую границу.
*/
func createData(value int) SegmentTreeData {
	return SegmentTreeData{
		prefix: value, // максимальная сумму среди всех	префиксов
		suffix: value, // максимальная сумму среди всех	суффиксов
		best:   value, // максимальная сумму подотрезка на отрезке
		total:  value, // сумма на отрезке
	}
}

func merge(left, right SegmentTreeData) SegmentTreeData {
	total := left.total + right.total
	prefix := max(left.prefix, left.total+right.prefix)
	suffix := max(right.suffix, right.total+left.suffix)
	best := max(
		left.best,
		right.best,
		prefix,
		suffix,
		left.suffix+right.prefix,
	)

	return SegmentTreeData{
		prefix,
		suffix,
		best,
		total,
	}
}

type SegmentTree struct {
	n    int
	data []SegmentTreeData
}

func Build(arr []int) *SegmentTree {
	n := len(arr)
	length := n * 4
	data := make([]SegmentTreeData, length)
	tree := &SegmentTree{
		n,
		data,
	}
	tree.build(arr, 1, 1, n)

	return tree
}

func (tree *SegmentTree) build(arr []int, index, left, right int) {

	if left > right {
		return
	} else if left == right {
		tree.data[index] = createData(arr[left-1])
	} else {
		middle := (left + right) / 2

		tree.build(arr, index*2, left, middle)
		tree.build(arr, index*2+1, middle+1, right)
		tree.data[index] = merge(tree.data[index*2], tree.data[index*2+1])
	}
}

func (tree *SegmentTree) Update(x, y int) {
	tree.update(1, 1, tree.n, x, y)
}

func (tree *SegmentTree) update(index int, left int, right int, updateIndex int, updateValue int) {
	if left > right || left > updateIndex || right < updateIndex {
		return
	} else if left == right {
		tree.data[index] = createData(updateValue)
	} else {
		middle := (left + right) / 2

		tree.update(index*2, left, middle, updateIndex, updateValue)
		tree.update(index*2+1, middle+1, right, updateIndex, updateValue)
		tree.data[index] = merge(tree.data[index*2], tree.data[index*2+1])
	}
}

func (tree *SegmentTree) Find(x, y int) int {
	return tree.find(1, 1, tree.n, x, y).best
}

func (tree *SegmentTree) find(index int, left int, right int, findLeft int, findRight int) SegmentTreeData {
	if left == findLeft && right == findRight {
		return tree.data[index]
	} else {
		middle := (left + right) / 2

		if findRight <= middle {
			return tree.find(index*2, left, middle, findLeft, findRight)
		} else if findLeft > middle {
			return tree.find(index*2+1, middle+1, right, findLeft, findRight)
		} else {
			leftResult := tree.find(index*2, left, middle, findLeft, min(middle, findRight))
			rightResult := tree.find(index*2+1, middle+1, right, max(findLeft, middle+1), findRight)
			mergedResult := merge(leftResult, rightResult)
			return mergedResult
		}
	}
}

func max(x int, rest ...int) int {
	mx := x

	for _, value := range rest {
		if mx < value {
			mx = value
		}
	}

	return mx
}

func min(x int, rest ...int) int {
	mn := x

	for _, value := range rest {
		if mn > value {
			mn = value
		}
	}

	return mn
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func readInt() int {
	var value int

	fmt.Fscanf(reader, "%d\n", &value)

	return value
}

func writeInt(value int) {
	fmt.Fprintln(writer, value)
}

func readArray(n int) []int {
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	stringArray := strings.Split(strings.TrimSpace(line), " ")
	if len(stringArray) != n {
		panic(fmt.Errorf("Expected input array to be of size %d, but was %d", n, len(stringArray)))
	}

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		value, err := strconv.Atoi(stringArray[i])
		if err != nil {
			panic(err)
		}

		arr[i] = value
	}

	return arr
}

func main() {
	// defer writer.Flush()

	// n := readInt()
	// arr := readArray(n)
	arr := []int{18, 21, -3, 7, 14, -5, 2}
	tree := Build(arr)

	//tree.Update(1, 3)
	value := tree.Find(1, 4)
	fmt.Printf("%+v\n", value)

	value = tree.Find(3, 4)
	fmt.Printf("%+v\n", value)

	value = tree.Find(3, 5)
	fmt.Printf("%+v\n", value)

	value = tree.Find(5, 7)
	fmt.Printf("%+v\n", value)

	fmt.Println()
	// m := readInt()
	// for i := 0; i < m; i++ {
	// 	query := readArray(3)
	// 	t := query[0]
	// 	x := query[1]
	// 	y := query[2]

	// 	if t == 0 {
	// 		tree.Update(x, y)
	// 	} else {
	// 		value := tree.Find(x, y)
	// 		writeInt(value)
	// 	}
	// }
}
