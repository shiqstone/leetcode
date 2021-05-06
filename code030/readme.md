30. Substring with Concatenation of All Words
---
You are given a string **s**, and a list of **words**, that are all of the same length. Find all starting indices of substring(s) in **s** that is concatenation of each word in **words** exactly once and without any intervening characters.

## Example 1
```
Input:
  s = "barfoothefoobarman"
  words = ["foo", "bar"]
Output: [0, 9]
Explaination: Substrings starting at index 0 and index 9 are "barfoo" and "foobar" respectively.
The output order does not matter, return [9,0] if fine too
```
## Example 2
```
Input:
  s = "wordgoodgoodgoodbestword"
  words = ["word", "good", "best", "word"]
Output: []
