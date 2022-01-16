# Quicksort

Here's my implementation of quicksort in Go. I did my best to only read description of the algorithm, rather than code snippets, in order to try to learn what is actually going on. To summarize, quicksort works as such:

1. Pick an element at random to be the pivot. I "randomly" select the final element every time. 
2. Find the final, correctly sorted position for this pivot in the list, and place it there. 
3. Ensure everything to the left of pivot is less than the pivot, and everything to the right is greater. 
4. Recursively apply the proceeding steps to each subset of the list, once to the set to all the items to the left of pivot, and again to right.

The complexity will be O(N * log<sub>2</sub>N). N for each item in the list, and log<sub>2</sub>N to represent the depth of recursion required. They are multiplied because each recursion iterates over it's partitioned list items once. I could be wrong about the specifics of the complexity, but I believe this is correct. 
