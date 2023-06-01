# 

## Time complexity :
O(N) - worst case 
This happens when all the elements are the same and we search for some different element. At each step, we will only be able to reduce the search space by 1 since arr[mid] equals arr[start] and it's not possible to decide the relative position of target from arr[mid]. Example: [1, 1, 1, 1, 1, 1, 1], target = 2.

O(log n)
Best case: This happens when all the elements are distinct. At each step, we will be able to divide our search space into half just like a normal binary search.
## Space complexity :

O(1)



https://leetcode.com/problems/search-in-rotated-sorted-array-ii/