Implement int sqrt(int x) function.
Calculate and return the square root of x, where x is a non-negative integer.
Since the return type is an integer, the result retains only the integer part, and the decimal part is dropped.

Example 1:
Input: 4

Output: 2

Example 2:
Input: 8

Output: 2
Instructions:The square root of 8 is 2.82842...Since the return type is an integer, the decimal part will be omitted.

Solution to this problem:
1. Determine the data type. Since the problem requires that the return type be an integer, the decimal part will be omitted, so I assign all values computed in the function to the integer type.

2. Determine the range of values to be calculated. It's not hard to find that:
	(x/2 +1)^2¡Ýx
   So, the maximum value of the root should not exceed 1/2.
3. Consider a critical case: when the square of the value of the root is just above or equal to the value of x, the value of that root is returned.