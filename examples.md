
	fmt.Println("\nExample 1 - Emit")

	emit := pipe.Of(1, 2, 3, 4, 5)

	for r := range emit {
		fmt.Println(r)
	}

	fmt.Println("\nExample 1 - Emit and sleep")

	emitSleep := pipe.Of(1, 2, 3, 4, 5)

	for r := range pipe.Pipe(getSleepFunc(300), emitSleep) {
		fmt.Println(r)
	}

	fmt.Println("\nExample 2 - Emit array")

	emitArray := pipe.Of([]int{1, 2, 3, 4, 5})

	for r := range emitArray {
		fmt.Println(r)
	}

	fmt.Println("\nExample 3 - Emit unwrapped array")

	emitArrayUnwrapped := pipe.Of([]int{1, 2, 3, 4, 5}...)

	for r := range emitArrayUnwrapped {
		fmt.Println(r)
	}

	fmt.Println("\nExample 4 - Merge")

	merge1 := pipe.Of(1, 2, 3, 4, 5)
	merge2 := pipe.Of(10, 11, 12, 13, 14)

	for r := range pipe.Merge(merge1, merge2) {
		fmt.Println(r)
	}

	fmt.Println("\nExample 5 - Merge with sleep")

	merge1b := pipe.Pipe(getSleepFunc(400), pipe.Of(1, 2, 3, 4, 5))
	merge2b := pipe.Pipe(getSleepFunc(300), pipe.Of(10, 11, 12))

	for r := range pipe.Merge(merge1b, merge2b) {
		fmt.Println(r)
	}