# Recursion 

## Time complexity : 

O(n+m)

Because each recursive call increments the pointer to l1 or l2 by one (approaching the dangling null at the end of each list), there will be exactly one call to mergeTwoLists per element in each list. Therefore, the time complexity is linear in the combined size of the lists.

## Space complexity : 

O(n+m)

The first call to mergeTwoLists does not return until the ends of both l1 and l2 have been reached, so n+mn + mn+m stack frames consume O(n+m)O(n + m)O(n+m) space.


https://leetcode.com/problems/merge-two-sorted-lists/description/