package main

import "fmt"

type Crate struct {
	Name string
}

type Stack struct {
	Crates []Crate
}

// func main() {
// 	stack1 := Stack{
// 		[]Crate{
// 			{"Z"},
// 			{"N"},
// 		},
// 	}

// 	stack2 := Stack{
// 		[]Crate{
// 			{"M"}, {"C"}, {"D"},
// 		},
// 	}

// 	stack3 := Stack{
// 		[]Crate{
// 			{"P"},
// 		},
// 	}

// 	fmt.Printf("Initial State\n")
// 	printStacks(&stack1, &stack2, &stack3)
// 	fmt.Printf("\n\n")

// 	// Rearrangement procedure
// 	moveStack(1, &stack2, &stack1)
// 	printStacks(&stack1, &stack2, &stack3)
// 	fmt.Printf("\n\n")

// 	moveStack(3, &stack1, &stack3)
// 	printStacks(&stack1, &stack2, &stack3)
// 	fmt.Printf("\n\n")

// 	moveStack(2, &stack2, &stack1)
// 	printStacks(&stack1, &stack2, &stack3)
// 	fmt.Printf("\n\n")

// 	moveStack(1, &stack1, &stack2)

// 	fmt.Printf("After Moves\n")
// 	printStacks(&stack1, &stack2, &stack3)
// 	printMessage(&stack1, &stack2, &stack3)

// }

func moveStack(amount int, fromStack *Stack, toStack *Stack) {
	// Crane version 9000
	// for i := 0; i < amount; i++ {
	// 	toStack.push(fromStack.pop())
	// }

	// Crane version 9001
	tempCrates := []Crate{}
	for i := 0; i < amount; i++ {
		tempCrates = append(tempCrates, *fromStack.pop())
	}

	for i := len(tempCrates) - 1; i >= 0; i-- {
		toStack.push(&tempCrates[i])
	}
}

func (stack *Stack) push(crate *Crate) {
	stack.Crates = append(stack.Crates, *crate)
}

func (stack *Stack) pop() *Crate {
	endCrate := stack.Crates[len(stack.Crates)-1]
	stack.Crates = stack.Crates[:len(stack.Crates)-1]
	return &endCrate
}

func printStacks(stacks ...*Stack) {
	for i, stack := range stacks {
		fmt.Printf("%d", i+1)
		for _, crate := range stack.Crates {
			fmt.Printf("[%v]", crate.Name)
		}
		fmt.Printf("\n")
	}
}

func printMessage(stacks ...*Stack) {
	for _, stack := range stacks {
		topCrate := stack.Crates[len(stack.Crates)-1]
		fmt.Printf("%v", topCrate.Name)
	}
	fmt.Printf("\n")
}

func main() {
	stack1 := Stack{
		[]Crate{
			{"W"}, {"M"}, {"L"}, {"F"},
		},
	}

	stack2 := Stack{
		[]Crate{
			{"B"}, {"Z"}, {"V"}, {"M"}, {"F"},
		},
	}

	stack3 := Stack{
		[]Crate{
			{"H"}, {"V"}, {"R"}, {"S"}, {"L"}, {"Q"},
		},
	}

	stack4 := Stack{
		[]Crate{
			{"F"}, {"S"}, {"V"}, {"Q"}, {"P"}, {"M"}, {"T"}, {"J"},
		},
	}

	stack5 := Stack{
		[]Crate{
			{"L"}, {"S"}, {"W"},
		},
	}

	stack6 := Stack{
		[]Crate{
			{"F"}, {"V"}, {"P"}, {"M"}, {"R"}, {"J"}, {"W"},
		},
	}

	stack7 := Stack{
		[]Crate{
			{"J"}, {"Q"}, {"C"}, {"P"}, {"N"}, {"R"}, {"F"},
		},
	}

	stack8 := Stack{
		[]Crate{
			{"V"}, {"H"}, {"P"}, {"S"}, {"Z"}, {"W"}, {"R"}, {"B"},
		},
	}

	stack9 := Stack{
		[]Crate{
			{"B"}, {"M"}, {"J"}, {"C"}, {"G"}, {"H"}, {"Z"}, {"W"},
		},
	}

	fmt.Printf("Initial State\n")
	printStacks(&stack1, &stack2, &stack3, &stack4, &stack5, &stack6, &stack7, &stack8, &stack9)

	// Rearrangement procedure

	moveStack(3, &stack5, &stack7)
	moveStack(2, &stack8, &stack9)
	moveStack(4, &stack3, &stack5)
	moveStack(2, &stack1, &stack7)
	moveStack(1, &stack3, &stack6)
	moveStack(2, &stack1, &stack7)
	moveStack(1, &stack8, &stack7)
	moveStack(4, &stack2, &stack8)
	moveStack(10, &stack9, &stack1)
	moveStack(6, &stack6, &stack2)
	moveStack(1, &stack6, &stack7)
	moveStack(9, &stack8, &stack6)
	moveStack(4, &stack2, &stack4)
	moveStack(2, &stack4, &stack1)
	moveStack(6, &stack1, &stack6)
	moveStack(1, &stack3, &stack2)
	moveStack(2, &stack1, &stack4)
	moveStack(2, &stack4, &stack3)
	moveStack(2, &stack1, &stack3)
	moveStack(4, &stack3, &stack1)
	moveStack(15, &stack7, &stack9)
	moveStack(4, &stack5, &stack9)
	moveStack(13, &stack9, &stack4)
	moveStack(10, &stack4, &stack8)
	moveStack(1, &stack7, &stack4)
	moveStack(6, &stack9, &stack5)
	moveStack(11, &stack6, &stack7)
	moveStack(4, &stack5, &stack7)
	moveStack(3, &stack8, &stack7)
	moveStack(4, &stack2, &stack4)
	moveStack(1, &stack5, &stack1)
	moveStack(5, &stack8, &stack4)
	moveStack(1, &stack5, &stack4)
	moveStack(10, &stack7, &stack1)
	moveStack(8, &stack7, &stack9)
	moveStack(12, &stack1, &stack9)
	moveStack(8, &stack9, &stack1)
	moveStack(2, &stack6, &stack9)
	moveStack(2, &stack8, &stack4)
	moveStack(1, &stack6, &stack9)
	moveStack(13, &stack4, &stack2)
	moveStack(13, &stack4, &stack2)
	moveStack(1, &stack6, &stack1)
	moveStack(1, &stack6, &stack4)
	moveStack(1, &stack4, &stack5)
	moveStack(14, &stack1, &stack8)
	moveStack(1, &stack5, &stack4)
	moveStack(13, &stack9, &stack5)
	moveStack(9, &stack8, &stack2)
	moveStack(8, &stack2, &stack1)
	moveStack(5, &stack8, &stack2)
	moveStack(5, &stack1, &stack6)
	moveStack(3, &stack1, &stack3)
	moveStack(1, &stack4, &stack8)
	moveStack(9, &stack5, &stack9)
	moveStack(18, &stack2, &stack8)
	moveStack(3, &stack3, &stack5)
	moveStack(2, &stack6, &stack4)
	moveStack(14, &stack2, &stack7)
	moveStack(1, &stack4, &stack2)
	moveStack(1, &stack6, &stack9)
	moveStack(1, &stack2, &stack5)
	moveStack(1, &stack6, &stack2)
	moveStack(1, &stack4, &stack6)
	moveStack(6, &stack8, &stack1)
	moveStack(2, &stack6, &stack9)
	moveStack(5, &stack5, &stack3)
	moveStack(1, &stack7, &stack8)
	moveStack(10, &stack9, &stack7)
	moveStack(13, &stack8, &stack5)
	moveStack(5, &stack5, &stack2)
	moveStack(6, &stack5, &stack7)
	moveStack(1, &stack8, &stack5)
	moveStack(5, &stack5, &stack9)
	moveStack(5, &stack9, &stack7)
	moveStack(4, &stack3, &stack8)
	moveStack(6, &stack1, &stack6)
	moveStack(4, &stack2, &stack4)
	moveStack(3, &stack7, &stack5)
	moveStack(2, &stack2, &stack9)
	moveStack(1, &stack3, &stack7)
	moveStack(29, &stack7, &stack9)
	moveStack(4, &stack5, &stack2)
	moveStack(5, &stack6, &stack4)
	moveStack(3, &stack7, &stack9)
	moveStack(3, &stack8, &stack6)
	moveStack(1, &stack2, &stack6)
	moveStack(3, &stack2, &stack5)
	moveStack(1, &stack8, &stack4)
	moveStack(1, &stack5, &stack9)
	moveStack(8, &stack4, &stack9)
	moveStack(15, &stack9, &stack2)
	moveStack(1, &stack5, &stack1)
	moveStack(10, &stack9, &stack4)
	moveStack(5, &stack4, &stack5)
	moveStack(5, &stack5, &stack4)
	moveStack(1, &stack1, &stack9)
	moveStack(1, &stack4, &stack3)
	moveStack(8, &stack2, &stack4)
	moveStack(7, &stack2, &stack7)
	moveStack(1, &stack3, &stack8)
	moveStack(1, &stack5, &stack6)
	moveStack(4, &stack7, &stack3)
	moveStack(1, &stack8, &stack2)
	moveStack(7, &stack4, &stack7)
	moveStack(11, &stack9, &stack7)
	moveStack(5, &stack4, &stack2)
	moveStack(3, &stack9, &stack6)
	moveStack(3, &stack3, &stack8)
	moveStack(4, &stack2, &stack4)
	moveStack(5, &stack9, &stack5)
	moveStack(1, &stack2, &stack1)
	moveStack(3, &stack8, &stack5)
	moveStack(2, &stack9, &stack1)
	moveStack(1, &stack2, &stack5)
	moveStack(2, &stack9, &stack6)
	moveStack(3, &stack7, &stack5)
	moveStack(7, &stack4, &stack1)
	moveStack(4, &stack4, &stack9)
	moveStack(3, &stack7, &stack2)
	moveStack(3, &stack1, &stack9)
	moveStack(1, &stack2, &stack3)
	moveStack(2, &stack7, &stack9)
	moveStack(6, &stack5, &stack4)
	moveStack(6, &stack4, &stack3)
	moveStack(5, &stack5, &stack1)
	moveStack(6, &stack7, &stack8)
	moveStack(1, &stack5, &stack1)
	moveStack(2, &stack9, &stack4)
	moveStack(1, &stack4, &stack3)
	moveStack(10, &stack6, &stack4)
	moveStack(2, &stack2, &stack1)
	moveStack(6, &stack4, &stack1)
	moveStack(5, &stack8, &stack3)
	moveStack(1, &stack8, &stack2)
	moveStack(7, &stack3, &stack9)
	moveStack(1, &stack6, &stack9)
	moveStack(2, &stack7, &stack3)
	moveStack(20, &stack1, &stack6)
	moveStack(7, &stack3, &stack8)
	moveStack(2, &stack9, &stack6)
	moveStack(1, &stack2, &stack3)
	moveStack(2, &stack3, &stack6)
	moveStack(1, &stack1, &stack4)
	moveStack(6, &stack4, &stack7)
	moveStack(5, &stack8, &stack3)
	moveStack(22, &stack6, &stack4)
	moveStack(2, &stack9, &stack7)
	moveStack(3, &stack3, &stack4)
	moveStack(6, &stack4, &stack2)
	moveStack(11, &stack9, &stack3)
	moveStack(9, &stack3, &stack7)
	moveStack(5, &stack4, &stack2)
	moveStack(5, &stack7, &stack2)
	moveStack(5, &stack7, &stack6)
	moveStack(10, &stack2, &stack4)
	moveStack(3, &stack2, &stack1)
	moveStack(1, &stack6, &stack3)
	moveStack(1, &stack1, &stack7)
	moveStack(17, &stack4, &stack1)
	moveStack(1, &stack8, &stack4)
	moveStack(2, &stack7, &stack5)
	moveStack(3, &stack2, &stack5)
	moveStack(3, &stack3, &stack8)
	moveStack(4, &stack5, &stack1)
	moveStack(3, &stack3, &stack7)
	moveStack(1, &stack4, &stack5)
	moveStack(21, &stack1, &stack5)
	moveStack(3, &stack8, &stack3)
	moveStack(4, &stack7, &stack5)
	moveStack(1, &stack1, &stack7)
	moveStack(1, &stack6, &stack3)
	moveStack(4, &stack4, &stack1)
	moveStack(1, &stack8, &stack1)
	moveStack(3, &stack4, &stack9)
	moveStack(5, &stack1, &stack8)
	moveStack(3, &stack9, &stack3)
	moveStack(5, &stack6, &stack1)
	moveStack(5, &stack1, &stack4)
	moveStack(6, &stack3, &stack2)
	moveStack(1, &stack3, &stack2)
	moveStack(3, &stack8, &stack1)
	moveStack(7, &stack2, &stack1)
	moveStack(10, &stack5, &stack2)
	moveStack(12, &stack5, &stack7)
	moveStack(2, &stack8, &stack3)
	moveStack(5, &stack5, &stack8)
	moveStack(8, &stack1, &stack6)
	moveStack(5, &stack4, &stack5)
	moveStack(3, &stack8, &stack6)
	moveStack(1, &stack8, &stack3)
	moveStack(6, &stack6, &stack7)
	moveStack(2, &stack3, &stack8)
	moveStack(3, &stack2, &stack1)
	moveStack(6, &stack2, &stack9)
	moveStack(2, &stack8, &stack4)
	moveStack(1, &stack3, &stack9)
	moveStack(1, &stack8, &stack6)
	moveStack(1, &stack6, &stack9)
	moveStack(7, &stack9, &stack5)
	moveStack(1, &stack9, &stack7)
	moveStack(1, &stack4, &stack6)
	moveStack(2, &stack6, &stack5)
	moveStack(1, &stack4, &stack1)
	moveStack(1, &stack2, &stack7)
	moveStack(5, &stack1, &stack2)
	moveStack(10, &stack7, &stack4)
	moveStack(12, &stack5, &stack7)
	moveStack(6, &stack4, &stack8)
	moveStack(2, &stack5, &stack6)
	moveStack(1, &stack8, &stack9)
	moveStack(1, &stack9, &stack5)
	moveStack(30, &stack7, &stack9)
	moveStack(4, &stack8, &stack4)
	moveStack(1, &stack8, &stack7)
	moveStack(2, &stack1, &stack4)
	moveStack(6, &stack6, &stack3)
	moveStack(1, &stack4, &stack1)
	moveStack(1, &stack1, &stack2)
	moveStack(8, &stack4, &stack8)
	moveStack(1, &stack4, &stack5)
	moveStack(2, &stack5, &stack6)
	moveStack(2, &stack9, &stack8)
	moveStack(3, &stack2, &stack1)
	moveStack(4, &stack3, &stack2)
	moveStack(1, &stack6, &stack4)
	moveStack(1, &stack7, &stack1)
	moveStack(2, &stack8, &stack2)
	moveStack(1, &stack9, &stack2)
	moveStack(2, &stack3, &stack2)
	moveStack(1, &stack4, &stack2)
	moveStack(4, &stack9, &stack6)
	moveStack(3, &stack6, &stack4)
	moveStack(21, &stack9, &stack8)
	moveStack(13, &stack2, &stack7)
	moveStack(9, &stack8, &stack5)
	moveStack(3, &stack1, &stack4)
	moveStack(14, &stack7, &stack2)
	moveStack(5, &stack8, &stack9)
	moveStack(1, &stack1, &stack2)
	moveStack(7, &stack8, &stack6)
	moveStack(2, &stack8, &stack2)
	moveStack(8, &stack6, &stack9)
	moveStack(1, &stack4, &stack5)
	moveStack(5, &stack8, &stack2)
	moveStack(4, &stack5, &stack9)
	moveStack(9, &stack9, &stack6)
	moveStack(2, &stack7, &stack6)
	moveStack(1, &stack8, &stack7)
	moveStack(9, &stack6, &stack4)
	moveStack(1, &stack6, &stack5)
	moveStack(1, &stack7, &stack3)
	moveStack(1, &stack4, &stack7)
	moveStack(1, &stack7, &stack2)
	moveStack(9, &stack2, &stack3)
	moveStack(8, &stack4, &stack1)
	moveStack(8, &stack9, &stack2)
	moveStack(2, &stack6, &stack5)
	moveStack(4, &stack5, &stack2)
	moveStack(2, &stack9, &stack5)
	moveStack(1, &stack4, &stack9)
	moveStack(10, &stack3, &stack7)
	moveStack(1, &stack9, &stack2)
	moveStack(1, &stack5, &stack3)
	moveStack(7, &stack2, &stack8)
	moveStack(7, &stack1, &stack5)
	moveStack(1, &stack1, &stack2)
	moveStack(2, &stack8, &stack2)
	moveStack(1, &stack3, &stack5)
	moveStack(2, &stack8, &stack6)
	moveStack(2, &stack8, &stack9)
	moveStack(2, &stack4, &stack6)
	moveStack(3, &stack2, &stack8)
	moveStack(3, &stack6, &stack7)
	moveStack(7, &stack5, &stack8)
	moveStack(7, &stack2, &stack7)
	moveStack(1, &stack6, &stack8)
	moveStack(5, &stack2, &stack7)
	moveStack(6, &stack8, &stack3)
	moveStack(2, &stack7, &stack1)
	moveStack(7, &stack2, &stack5)
	moveStack(1, &stack3, &stack5)
	moveStack(1, &stack1, &stack5)
	moveStack(2, &stack9, &stack7)
	moveStack(4, &stack3, &stack7)
	moveStack(2, &stack4, &stack6)
	moveStack(1, &stack1, &stack6)
	moveStack(1, &stack2, &stack4)
	moveStack(16, &stack5, &stack6)
	moveStack(1, &stack4, &stack9)
	moveStack(19, &stack6, &stack1)
	moveStack(1, &stack3, &stack5)
	moveStack(1, &stack9, &stack1)
	moveStack(1, &stack8, &stack5)
	moveStack(5, &stack8, &stack3)
	moveStack(5, &stack7, &stack2)
	moveStack(3, &stack2, &stack9)
	moveStack(5, &stack1, &stack7)
	moveStack(2, &stack5, &stack1)
	moveStack(3, &stack9, &stack4)
	moveStack(4, &stack1, &stack9)
	moveStack(2, &stack2, &stack8)
	moveStack(2, &stack8, &stack6)
	moveStack(1, &stack6, &stack9)
	moveStack(4, &stack3, &stack8)
	moveStack(4, &stack8, &stack3)
	moveStack(2, &stack3, &stack8)
	moveStack(1, &stack8, &stack2)
	moveStack(1, &stack9, &stack7)
	moveStack(10, &stack1, &stack7)
	moveStack(26, &stack7, &stack6)
	moveStack(3, &stack9, &stack3)
	moveStack(1, &stack4, &stack6)
	moveStack(2, &stack1, &stack4)
	moveStack(1, &stack1, &stack6)
	moveStack(1, &stack9, &stack3)
	moveStack(1, &stack2, &stack3)
	moveStack(4, &stack4, &stack9)
	moveStack(10, &stack7, &stack8)
	moveStack(3, &stack7, &stack4)
	moveStack(4, &stack9, &stack4)
	moveStack(4, &stack4, &stack7)
	moveStack(4, &stack3, &stack9)
	moveStack(5, &stack7, &stack5)
	moveStack(3, &stack5, &stack1)
	moveStack(3, &stack9, &stack8)
	moveStack(3, &stack1, &stack5)
	moveStack(2, &stack3, &stack5)
	moveStack(7, &stack8, &stack1)
	moveStack(7, &stack8, &stack9)
	moveStack(4, &stack6, &stack3)
	moveStack(3, &stack3, &stack6)
	moveStack(1, &stack3, &stack4)
	moveStack(2, &stack4, &stack1)
	moveStack(1, &stack9, &stack6)
	moveStack(4, &stack1, &stack3)
	moveStack(3, &stack5, &stack1)
	moveStack(1, &stack5, &stack2)
	moveStack(6, &stack1, &stack2)
	moveStack(6, &stack2, &stack7)
	moveStack(2, &stack7, &stack4)
	moveStack(1, &stack2, &stack6)
	moveStack(1, &stack1, &stack4)
	moveStack(3, &stack5, &stack7)
	moveStack(6, &stack7, &stack4)
	moveStack(1, &stack9, &stack3)
	moveStack(1, &stack3, &stack6)
	moveStack(4, &stack4, &stack3)
	moveStack(9, &stack6, &stack1)
	moveStack(10, &stack1, &stack6)
	moveStack(7, &stack4, &stack5)
	moveStack(28, &stack6, &stack4)
	moveStack(3, &stack6, &stack7)
	moveStack(3, &stack3, &stack8)
	moveStack(4, &stack5, &stack7)
	moveStack(1, &stack8, &stack4)
	moveStack(18, &stack4, &stack7)
	moveStack(8, &stack7, &stack6)
	moveStack(6, &stack4, &stack1)
	moveStack(2, &stack5, &stack4)
	moveStack(8, &stack6, &stack1)
	moveStack(2, &stack8, &stack9)
	moveStack(1, &stack5, &stack3)
	moveStack(1, &stack9, &stack1)
	moveStack(5, &stack9, &stack2)
	moveStack(2, &stack9, &stack3)
	moveStack(1, &stack2, &stack5)
	moveStack(2, &stack1, &stack5)
	moveStack(6, &stack7, &stack5)
	moveStack(1, &stack6, &stack4)
	moveStack(6, &stack5, &stack9)
	moveStack(2, &stack4, &stack1)
	moveStack(8, &stack1, &stack8)
	moveStack(4, &stack9, &stack7)
	moveStack(1, &stack5, &stack6)
	moveStack(1, &stack1, &stack6)
	moveStack(2, &stack1, &stack2)
	moveStack(1, &stack9, &stack7)
	moveStack(3, &stack2, &stack4)
	moveStack(2, &stack8, &stack3)
	moveStack(5, &stack8, &stack2)
	moveStack(4, &stack2, &stack5)
	moveStack(1, &stack8, &stack9)
	moveStack(12, &stack3, &stack2)
	moveStack(2, &stack6, &stack2)
	moveStack(12, &stack2, &stack4)
	moveStack(6, &stack2, &stack3)
	moveStack(4, &stack1, &stack9)
	moveStack(8, &stack4, &stack7)
	moveStack(3, &stack3, &stack4)
	moveStack(1, &stack5, &stack4)
	moveStack(5, &stack9, &stack6)
	moveStack(3, &stack5, &stack8)
	moveStack(1, &stack9, &stack1)
	moveStack(2, &stack8, &stack5)
	moveStack(3, &stack5, &stack6)
	moveStack(1, &stack8, &stack4)
	moveStack(4, &stack7, &stack8)
	moveStack(1, &stack1, &stack3)
	moveStack(2, &stack8, &stack3)
	moveStack(7, &stack6, &stack7)
	moveStack(1, &stack3, &stack7)
	moveStack(2, &stack8, &stack6)
	moveStack(22, &stack7, &stack8)
	moveStack(6, &stack4, &stack8)
	moveStack(5, &stack8, &stack6)
	moveStack(5, &stack6, &stack2)
	moveStack(4, &stack2, &stack3)
	moveStack(6, &stack8, &stack5)
	moveStack(4, &stack4, &stack7)
	moveStack(1, &stack3, &stack7)
	moveStack(4, &stack4, &stack5)
	moveStack(1, &stack5, &stack4)
	moveStack(2, &stack6, &stack5)
	moveStack(9, &stack5, &stack6)
	moveStack(10, &stack6, &stack7)
	moveStack(1, &stack2, &stack1)
	moveStack(3, &stack4, &stack8)
	moveStack(16, &stack7, &stack9)
	moveStack(1, &stack7, &stack8)
	moveStack(1, &stack1, &stack8)
	moveStack(1, &stack8, &stack3)
	moveStack(2, &stack7, &stack4)
	moveStack(15, &stack8, &stack1)
	moveStack(1, &stack8, &stack1)
	moveStack(4, &stack8, &stack4)
	moveStack(7, &stack9, &stack7)
	moveStack(3, &stack5, &stack9)
	moveStack(10, &stack9, &stack6)
	moveStack(2, &stack9, &stack2)
	moveStack(7, &stack7, &stack4)
	moveStack(9, &stack3, &stack2)
	moveStack(8, &stack2, &stack7)
	moveStack(1, &stack8, &stack4)
	moveStack(3, &stack2, &stack1)
	moveStack(9, &stack7, &stack1)
	moveStack(9, &stack4, &stack1)
	moveStack(2, &stack7, &stack5)
	moveStack(1, &stack5, &stack4)
	moveStack(1, &stack5, &stack2)
	moveStack(6, &stack1, &stack3)
	moveStack(16, &stack1, &stack2)
	moveStack(9, &stack2, &stack1)
	moveStack(5, &stack6, &stack9)
	moveStack(2, &stack1, &stack9)
	moveStack(1, &stack2, &stack5)
	moveStack(4, &stack4, &stack8)
	moveStack(2, &stack8, &stack2)
	moveStack(2, &stack2, &stack3)
	moveStack(17, &stack1, &stack2)
	moveStack(2, &stack1, &stack9)
	moveStack(13, &stack2, &stack8)
	moveStack(1, &stack2, &stack4)
	moveStack(11, &stack8, &stack3)
	moveStack(3, &stack3, &stack4)
	moveStack(3, &stack9, &stack2)
	moveStack(1, &stack5, &stack2)
	moveStack(1, &stack9, &stack3)
	moveStack(3, &stack4, &stack3)
	moveStack(1, &stack4, &stack9)
	moveStack(3, &stack3, &stack4)
	moveStack(1, &stack8, &stack7)
	moveStack(7, &stack2, &stack9)
	moveStack(3, &stack1, &stack7)
	moveStack(3, &stack2, &stack8)
	moveStack(3, &stack7, &stack9)
	moveStack(10, &stack3, &stack5)
	moveStack(3, &stack6, &stack9)
	moveStack(8, &stack9, &stack4)
	moveStack(1, &stack2, &stack1)
	moveStack(1, &stack7, &stack9)
	moveStack(2, &stack2, &stack3)
	moveStack(4, &stack4, &stack8)
	moveStack(1, &stack6, &stack2)
	moveStack(7, &stack5, &stack3)
	moveStack(1, &stack5, &stack2)
	moveStack(9, &stack8, &stack9)
	moveStack(12, &stack3, &stack8)
	moveStack(1, &stack1, &stack9)
	moveStack(9, &stack8, &stack6)
	moveStack(1, &stack5, &stack7)
	moveStack(1, &stack5, &stack4)
	moveStack(2, &stack2, &stack9)
	moveStack(1, &stack2, &stack6)
	moveStack(2, &stack4, &stack3)
	moveStack(9, &stack4, &stack8)
	moveStack(6, &stack3, &stack6)
	moveStack(12, &stack6, &stack2)
	moveStack(2, &stack6, &stack7)
	moveStack(8, &stack8, &stack3)
	moveStack(5, &stack8, &stack7)
	moveStack(3, &stack6, &stack5)
	moveStack(6, &stack3, &stack7)
	moveStack(6, &stack7, &stack6)
	moveStack(1, &stack4, &stack9)
	moveStack(4, &stack6, &stack5)
	moveStack(20, &stack9, &stack6)
	moveStack(4, &stack9, &stack8)
	moveStack(2, &stack8, &stack7)
	moveStack(4, &stack6, &stack4)
	moveStack(10, &stack6, &stack1)

	fmt.Printf("After Moves\n")
	printStacks(&stack1, &stack2, &stack3, &stack4, &stack5, &stack6, &stack7, &stack8, &stack9)
	printMessage(&stack1, &stack2, &stack3, &stack4, &stack5, &stack6, &stack7, &stack8, &stack9)

}
