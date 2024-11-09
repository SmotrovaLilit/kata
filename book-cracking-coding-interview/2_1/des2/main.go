package des2

func RemoveDuplicates(list LinkedList) {
	if list.IsEmpty() {
		return
	}
	// Start with the first node
	currentNode := list.head

	// Iterate through each node in the list
	for currentNode != nil {
		// Start comparing with the next nodes
		compareNode := currentNode.Next
		previousNode := currentNode

		// Loop through the rest of the list to find duplicates
		for compareNode != nil {
			if currentNode.Value == compareNode.Value {
				// Remove duplicate by skipping the current compareNode
				previousNode.Next = compareNode.Next
			} else {
				// Move previousNode and compareNode forward
				previousNode = compareNode
			}
			// Move to the next node to compare
			compareNode = compareNode.Next
		}
		// Move currentNode forward to the next node
		currentNode = currentNode.Next
	}
}
