This problem can be solved recursively. 
When solving this problem, the operation of finding the remainder is used to judge whether the units digit meets the condition of good number.And then we use division to figure out whether the tens place, the hundreds place, or the thousands place is a good number.
After n%10, the program determines whether the remainder matches the good number and is valid.
Returns false if n%10 is invalid and not a good number.
If the units digit is valid or good, then the tens digit is judged, then the thousands digit is judged, and so on, up to the highest digit.The Boolean function used for marking is assigned to True if the valid conditions are met from the ones place to the highest place, and one of them is a good number.
