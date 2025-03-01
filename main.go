package main

import (
	"fmt"
	"github.com/tfajarama/go-logic-exercise/logic1"
	"github.com/tfajarama/go-logic-exercise/logic2"
	"github.com/tfajarama/go-logic-exercise/logic3"
	"github.com/tfajarama/go-print-slice"
)

func main() {
	fmt.Println("\n----------------------------------------[LOGIC-01]----------------------------------------")

	fmt.Println("\nLogic 1 Number 1:")
	go_print_slice.PrintSlice1D(logic1.Logic1AscStep(10, 1, 2))

	fmt.Println("\nLogic 1 Number 2:")
	go_print_slice.PrintSlice1D(logic1.Logic1AscStep(10, 2, 2))

	fmt.Println("\nLogic 1 Number 3:")
	go_print_slice.PrintSlice1D(logic1.Logic1AscStep(10, 3, 3))

	fmt.Println("\nLogic 1 Number 4:")
	go_print_slice.PrintSlice1D(logic1.Logic1DescStep(10, 1, 2))

	fmt.Println("\nLogic 1 Number 5:")
	go_print_slice.PrintSlice1D(logic1.Logic1DescStep(10, 2, 2))

	fmt.Println("\nLogic 1 Number 6:")
	go_print_slice.PrintSlice1D(logic1.Logic1DescStep(10, 3, 3))

	fmt.Println("\nLogic 1 Number 7:")
	go_print_slice.PrintSlice1D(logic1.Logic1AscDescStep(10, 1, 2))
	go_print_slice.PrintSlice1D(logic1.Logic1AscDescStep(11, 1, 2))

	fmt.Println("\nLogic 1 Number 8:")
	go_print_slice.PrintSlice1D(logic1.Logic1AscDescStep(10, 2, 2))
	go_print_slice.PrintSlice1D(logic1.Logic1AscDescStep(11, 2, 2))

	fmt.Println("\nLogic 1 Number 9:")
	go_print_slice.PrintSlice1D(logic1.Logic1AscDescStep(10, 3, 3))
	go_print_slice.PrintSlice1D(logic1.Logic1AscDescStep(11, 3, 3))

	fmt.Println("\nLogic 1 Number 10:")
	go_print_slice.PrintSlice1DStr(logic1.Logic1OddNumEvenFizz(10, 2, 2))

	fmt.Println("\nLogic 1 Number 11:")
	go_print_slice.PrintSlice1DStr(logic1.Logic1OddBuzzEvenNum(10, 3, 2))

	fmt.Println("\nLogic 1 Number 12:")
	go_print_slice.PrintSlice1D(logic1.Logic1Repeat4Elements(12, 1, 2))

	fmt.Println("\n----------------------------------------[LOGIC-02]----------------------------------------")

	fmt.Println("\nLogic 2 Number 1:")
	go_print_slice.PrintSlice2D(logic2.Logic2AscRows(9, 1, 2))

	fmt.Println("\nLogic 2 Number 2:")
	go_print_slice.PrintSlice2D(logic2.Logic2AscRows(9, 2, 2))

	fmt.Println("\nLogic 2 Number 3:")
	go_print_slice.PrintSlice2D(logic2.Logic2AscAll(9, 1, 2))

	fmt.Println("\nLogic 2 Number 4:")
	go_print_slice.PrintSlice2D(logic2.Number4(9, 1, 3, 3, 2))

	fmt.Println("\nLogic 2 Number 5:")
	go_print_slice.PrintSlice2D(logic2.Logic2AscAllAlternating(9, 1, 2))

	fmt.Println("\nLogic 2 Number 6:")
	go_print_slice.PrintSlice2D(logic2.Logic2AscAllAlternatingCustomOddEvenRow(9, 1, 3, 2, 3))

	fmt.Println("\nLogic 2 Number 7:")
	go_print_slice.PrintSlice2D(logic2.Logic2StairsDown(9, 1, 2))

	fmt.Println("\nLogic 2 Number 8:")
	go_print_slice.PrintSlice2D(logic2.Logic2StairsUp(9, 1, 2))

	fmt.Println("\nLogic 2 Number 9:")
	go_print_slice.PrintSlice2D(logic2.Logic2XSign(9, 1, 2))

	fmt.Println("\nLogic 2 Number 10:")
	go_print_slice.PrintSlice2D(logic2.Logic2TriangleDownLeft(9, 1, 2))

	fmt.Println("\nLogic 2 Number 11:")
	go_print_slice.PrintSlice2D(logic2.Logic2TriangleUpRight(9, 1, 2))

	fmt.Println("\nLogic 2 Number 12:")
	go_print_slice.PrintSlice2D(logic2.Logic2TriangleLeftRight(9, 1, 2))

	fmt.Println("\nLogic 2 Number 13:")
	go_print_slice.PrintSlice2D(logic2.Logic2TriangleUpDown(9, 1, 2))

	fmt.Println("\n----------------------------------------[LOGIC-03]----------------------------------------")

	fmt.Println("\nLogic 3 Number 1:")
	go_print_slice.PrintSlice2D(logic3.Number1(9))

	fmt.Println("\nLogic 3 Number 2:")
	go_print_slice.PrintSlice2D(logic3.Number2(9))

	fmt.Println("\nLogic 3 Number 3:")
	go_print_slice.PrintSlice2D(logic3.Number3(9))

	fmt.Println("\nLogic 3 Number 4:")
	go_print_slice.PrintSlice2D(logic3.Number4(9))

	fmt.Println("\nLogic 3 Number 5:")
	go_print_slice.PrintSlice2D(logic3.Number5(9))

	fmt.Println("\nLogic 3 Number 6:")
	go_print_slice.PrintSlice2D(logic3.Number6(9))

	fmt.Println("\nLogic 3 Number 7:")
	go_print_slice.PrintSlice2D(logic3.Number7(9))

	fmt.Println("\nLogic 3 Number 8:")
	go_print_slice.PrintSlice2D(logic3.Number8(9))

	fmt.Println("\nLogic 3 Number 9:")
	go_print_slice.PrintSlice2D(logic3.Number9(9))

	fmt.Println("\nLogic 3 Number 10:")
	go_print_slice.PrintSlice2D(logic3.Number10(9))

	fmt.Println("\nLogic 3 Number 11:")
	go_print_slice.PrintSlice2D(logic3.Number11(9))

	fmt.Println("\nLogic 3 Number 11b:")
	go_print_slice.PrintSlice2D(logic3.Number11b(9))

	fmt.Println("\nLogic 3 Number 12:")
	go_print_slice.PrintSlice2D(logic3.Number12SWStyle(9))

	fmt.Println("\nLogic 3 Number 12b:")
	go_print_slice.PrintSlice2D(logic3.Number12bSWStyle(9))

	//fmt.Println("\nLogic 3 Number 12:")
	//go_print_slice.PrintSlice2D(logic3.Number12HWStyle(9))

	//fmt.Println("\nLogic 3 Number 12b:")
	//go_print_slice.PrintSlice2D(logic3.Number12bHWStyle(9))

	fmt.Println("\nLogic 3 Number 13:")
	go_print_slice.PrintSlice2D(logic3.Number13(9))

	fmt.Println("\nLogic 3 Number 14:")
	go_print_slice.PrintSlice2D(logic3.Number14(5))
}
