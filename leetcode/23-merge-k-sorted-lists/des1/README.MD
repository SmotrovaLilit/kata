# Compare one by one 

## Time complexity : 
O(kN)
where k\text{k}k is the number of linked lists.

Almost every selection of node in final linked costs O(k)O(k)O(k) (k-1\text{k-1}k-1 times comparison).
There are NNN nodes in the final linked list.
## Space complexity :

- O(n) Creating a new linked list costs O(n) space.
- O(1) It's not hard to apply in-place method - connect selected nodes instead of creating new nodes to fill the new linked list.

https://leetcode.com/problems/merge-k-sorted-lists/solutions/127466/merge-k-sorted-list/