68. Text Justification
---
Given an array of words and a width *maxWidth*, format the text such that each line has exactly *maxWidth* characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces `' '` when nessesary so that each line has exactly *maxWidth* characters.

Extra spaces between words should be destribute as evenly as possible. If the number of spaces on a line do not devide evenly between words, the empty slot on the left will be assigned more than the slot the right.

For the last line of text, it should be left justified and no extra space insert between words.

Note:
- A word is define as character sequence consisting of non-space character only
- Each word's length is guaranteed greater than 0 and not exceed maxWidth
- The input words array contain at least one word.

Example 1:
```
Input: 
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
Output:
[
    "This    is    an",
    "example  of text",
    "justification.  "
]
```

Example 2:
```
Input:
words = ["What", "must", "be", "acknowledgment", "shall", "be"]
maxWidth = 16
Output:
[
    "What   must  be",
    "acknowledgment ",
    "shall be       ",
]
```

Example 3:
```
Input:
words = ["Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "compute.", "Art", "is", "everything", "we", "do"]
maxWidth = 20
Output:
[
    "Science  is  what we",
    "understand      well",
    "enough to explain to",
    "a  compute.  Art  is",
    "everything  else  we",
    "do                  ", 
]
```
