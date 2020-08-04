package main

import (
	"sort"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func task3(startx int, increment int, values []int) (t map[int]int) {

	cv, a, b, c, d := getcommonValue(values) //get common value 240
	//creating table
	table2 := make(map[int]int)
	for i, value := range values {
		table2[startx+i*increment] = value
	}

	addBelow := func() {
		d = append(d, cv+d[len(d)-1]) //new 4th diff is commonvalue + last element of 4th diff
		c = append(c, d[len(d)-1]+c[len(c)-1])
		b = append(b, c[len(c)-1]+b[len(b)-1]) //second diff
		a = append(a, b[len(b)-1]+a[len(a)-1]) //first diff
		table2[startx+len(table2)*increment] = table2[startx+len(table2)-1] + a[len(a)-1]
	}

	//for substracting we need to track saperate counter
	decrementTracker := 0
	addAvobe := func() {
		d = append([]int{d[0] - cv}, d...)
		c = append([]int{c[0] - d[0]}, d...)
		b = append([]int{b[0] - c[0]}, d...)
		a = append([]int{a[0] - b[0]}, d...)
		table2[startx-decrementTracker*increment] = table2[startx] - a[0]
		decrementTracker++ //update counter
	}

	//go below
	for i := 0; i < 7; i++ {
		addBelow()
	}
	//add above
	for i := 0; i < 10; i++ {
		addAvobe()
	}
	//because table2 map is not sorted we need to use other slice to sort our key
	sortx := make([]int, 0, len(table2))
	//migrading keys from table2
	for k := range table2 {
		sortx = append(sortx, k)
	}
	//sorting
	sort.Ints(sortx)
	//table creator
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("x", "f(x)")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	//using our sorted key value inserting output in each row
	for _, k := range sortx {
		tbl.AddRow(k, table2[k])
	}
	//pringout table
	tbl.Print()
	return table2

}

//returns common number and the rest of the differences
func getcommonValue(values []int) (cv int, a []int, b []int, c []int, d []int) {
	firstdiv := []int{
		values[1] - values[0],
		values[2] - values[1],
		values[3] - values[2],
		values[4] - values[3],
		values[5] - values[4],
	}
	seconddiv := []int{
		firstdiv[1] - firstdiv[0],
		firstdiv[2] - firstdiv[1],
		firstdiv[3] - firstdiv[2],
		firstdiv[4] - firstdiv[3],
	}
	thirddiv := []int{
		seconddiv[1] - seconddiv[0],
		seconddiv[2] - seconddiv[1],
		seconddiv[3] - seconddiv[2],
	}
	forthdiv := []int{
		thirddiv[1] - thirddiv[0],
		thirddiv[2] - thirddiv[1],
	}
	fifhdiv := []int{
		forthdiv[1] - forthdiv[0],
	}
	return fifhdiv[0], firstdiv, seconddiv, thirddiv, forthdiv
}

func main() {
	input := []int{-15, -5, -5, -3, 109, 775}

	task3(-2, 1, input)
}
