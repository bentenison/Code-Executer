To include test inputs and expected outputs directly within the `questions` collection, we can modify the document structure to store this information as part of each question. This allows you to define the test cases along with the question itself, making it easier to manage and execute tests.

### Updated Questions Collection Structure

#### Example Document Structure for `questions`

```json
{
  "_id": ObjectId("..."),
  "title": "Example Question Title",
  "description": "Describe the problem the user needs to solve.",
  "creator_id": ObjectId("..."), // Reference to the creator (user) who added this question
  "language": "Python", // Programming language for the question
  "created_at": ISODate("2023-10-05T12:00:00Z"),
  "updated_at": ISODate("2023-10-05T12:00:00Z"),
  "difficulty": "Medium", // Difficulty level (e.g., Easy, Medium, Hard)
  "tags": ["array", "string"], // Tags for categorization
  "test_template_id": ObjectId("..."), // Reference to the test template used for this question
  "test_cases": [ // Array of test cases
    {
      "input": "test input 1", // Example input for the test case
      "expected_output": "expected output 1" // Expected output for the given input
    },
    {
      "input": "test input 2",
      "expected_output": "expected output 2"
    }
  ]
}
```

### Explanation of the Fields

- **test_cases:** This field is an array that stores multiple test case objects. Each test case consists of:
  - **input:** The specific input value that the user’s solution will process.
  - **expected_output:** The output that should be produced for the given input.

### Example of a Complete Question Document

```json
{
  "_id": ObjectId("..."),
  "title": "Sum of Two Numbers",
  "description": "Write a function that returns the sum of two numbers.",
  "creator_id": ObjectId("..."),
  "language": "Python",
  "created_at": ISODate("2023-10-05T12:00:00Z"),
  "updated_at": ISODate("2023-10-05T12:00:00Z"),
  "difficulty": "Easy",
  "tags": ["math", "addition"],
  "test_template_id": ObjectId("..."),
  "test_cases": [
    {
      "input": "1, 2",
      "expected_output": "3"
    },
    {
      "input": "-1, 1",
      "expected_output": "0"
    },
    {
      "input": "100, 200",
      "expected_output": "300"
    }
  ]
}
```

### Benefits of This Structure

1. **Self-Contained Questions:** Each question contains all the necessary information, including test cases, making it easier to manage and retrieve data.

2. **Ease of Execution:** When executing tests, the relevant inputs and expected outputs can be accessed directly from the question document, streamlining the testing process.

3. **Flexibility:** You can easily add or modify test cases without altering the overall structure of the question.

### Summary

This updated structure for the `questions` collection allows for a comprehensive representation of programming questions, including defined test inputs and expected outputs. This setup facilitates easier management and execution of code tests, enhancing the overall functionality of your code execution platform. Adjust the structure further based on your specific needs!

```{
  "_id": ObjectId("..."),
  "template": `
def main(input):
    # User's main logic starts here
    %s
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "%s" # User input for the test case
    expected_output = "%s" # Expected output
    
    result = main(test_input)
    print("Pass" if result == expected_output else "Fail")
`,
  "language": "Python", // Programming language the template is designed for
  "created_at": ISODate("2023-10-05T12:00:00Z"),
  "updated_at": ISODate("2023-10-05T12:00:00Z"),
  "description": "Template for testing programming questions in Python."
}```

[{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1c8"
  },
  "title": "Prime Number Checker",
  "description": "Write a function that checks whether a given number is prime.",
  "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (2,),\n        (11,),\n        (15,),\n        (1,)\n    ]\n    expected_outputs = [\n        True,\n        True,\n        False,\n        False\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "math",
    "prime",
    "beginner"
  ],
  "testcases": [
    {
      "input": 2,
      "expectedOutput": true
    },
    {
      "input": 11,
      "expectedOutput": true
    },
    {
      "input": 15,
      "expectedOutput": false
    },
    {
      "input": 1,
      "expectedOutput": false
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1c9"
  },
  "title": "Factorial Calculator",
  "description": "Write a function to compute the factorial of a given number.",
  "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (5,),\n        (0,),\n        (1,),\n        (3,)\n    ]\n    expected_outputs = [\n        120,\n        1,\n        1,\n        6\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "math",
    "factorial",
    "beginner"
  ],
  "testcases": [
    {
      "input": 5,
      "expectedOutput": 120
    },
    {
      "input": 0,
      "expectedOutput": 1
    },
    {
      "input": 1,
      "expectedOutput": 1
    },
    {
      "input": 3,
      "expectedOutput": 6
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1ca"
  },
  "title": "Fibonacci Sequence Generator",
  "description": "Write a function that returns the n-th Fibonacci number.",
  "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (0,),\n        (1,),\n        (5,),\n        (10,)\n    ]\n    expected_outputs = [\n        0,\n        1,\n        5,\n        55\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "math",
    "fibonacci",
    "beginner"
  ],
  "testcases": [
    {
      "input": 0,
      "expectedOutput": 0
    },
    {
      "input": 1,
      "expectedOutput": 1
    },
    {
      "input": 5,
      "expectedOutput": 5
    },
    {
      "input": 10,
      "expectedOutput": 55
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1cb"
  },
  "title": "Palindrome Checker",
  "description": "Write a function to check if a given string is a palindrome.",
  "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('racecar',),\n        ('hello',),\n        ('madam',),\n        ('world',)\n    ]\n    expected_outputs = [\n        True,\n        False,\n        True,\n        False\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "palindrome",
    "beginner"
  ],
  "testcases": [
    {
      "input": "racecar",
      "expectedOutput": true
    },
    {
      "input": "hello",
      "expectedOutput": false
    },
    {
      "input": "madam",
      "expectedOutput": true
    },
    {
      "input": "world",
      "expectedOutput": false
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1cc"
  },
  "title": "Sum of Digits",
  "description": "Write a function that calculates the sum of the digits of a given number.",
  "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (123,),\n        (4567,),\n        (89,),\n        (0,)\n    ]\n    expected_outputs = [\n        6,\n        22,\n        17,\n        0\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "math",
    "digits",
    "beginner"
  ],
  "testcases": [
    {
      "input": 123,
      "expectedOutput": 6
    },
    {
      "input": 4567,
      "expectedOutput": 22
    },
    {
      "input": 89,
      "expectedOutput": 17
    },
    {
      "input": 0,
      "expectedOutput": 0
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1cd"
  },
  "title": "Reverse a String",
  "description": "Write a function that reverses a given string.",
  "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('hello',),\n        ('world',),\n        ('python',),\n        ('!dlrow',)\n    ]\n    expected_outputs = [\n        'olleh',\n        'dlrow',\n        'nohtyp',\n        'world!'\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "reverse",
    "beginner"
  ],
  "testcases": [
    {
      "input": "hello",
      "expectedOutput": "olleh"
    },
    {
      "input": "world",
      "expectedOutput": "dlrow"
    },
    {
      "input": "python",
      "expectedOutput": "nohtyp"
    },
    {
      "input": "!dlrow",
      "expectedOutput": "world!"
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1ce"
  },
  "title": "Find the Largest Element",
  "description": "Write a function that finds the largest element in a list.",
  "template_code": "def main(lst):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ([1, 2, 3, 4, 5],),\n        ([10, 5, 20, 15],),\n        ([7, 8, 9, 1],),\n        ([3],)\n    ]\n    expected_outputs = [\n        5,\n        20,\n        9,\n        3\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "lists",
    "largest",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        1,
        2,
        3,
        4,
        5
      ],
      "expectedOutput": 5
    },
    {
      "input": [
        10,
        5,
        20,
        15
      ],
      "expectedOutput": 20
    },
    {
      "input": [
        7,
        8,
        9,
        1
      ],
      "expectedOutput": 9
    },
    {
      "input": [
        3
      ],
      "expectedOutput": 3
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1cf"
  },
  "title": "Count Vowels",
  "description": "Write a function to count the number of vowels in a given string.",
  "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('hello',),\n        ('world',),\n        ('aeiou',),\n        ('xyz',)\n    ]\n    expected_outputs = [\n        2,\n        1,\n        5,\n        0\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "vowels",
    "beginner"
  ],
  "testcases": [
    {
      "input": "hello",
      "expectedOutput": 2
    },
    {
      "input": "world",
      "expectedOutput": 1
    },
    {
      "input": "aeiou",
      "expectedOutput": 5
    },
    {
      "input": "xyz",
      "expectedOutput": 0
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d0"
  },
  "title": "Anagram Checker",
  "description": "Write a function to check if two strings are anagrams of each other.",
  "template_code": "def main(s1, s2):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('listen', 'silent'),\n        ('hello', 'world'),\n        ('evil', 'vile'),\n        ('python', 'java')\n    ]\n    expected_outputs = [\n        True,\n        False,\n        True,\n        False\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0], test_input[1])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "anagrams",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        "listen",
        "silent"
      ],
      "expectedOutput": true
    },
    {
      "input": [
        "hello",
        "world"
      ],
      "expectedOutput": false
    },
    {
      "input": [
        "evil",
        "vile"
      ],
      "expectedOutput": true
    },
    {
      "input": [
        "python",
        "java"
      ],
      "expectedOutput": false
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d1"
  },
  "title": "Remove Duplicates",
  "description": "Write a function that removes duplicates from a list.",
  "template_code": "def main(lst):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ([1, 2, 2, 3, 4],),\n        ([5, 5, 5, 5],),\n        ([1, 2, 3, 4],),\n        ([],)\n    ]\n    expected_outputs = [\n        [1, 2, 3, 4],\n        [5],\n        [1, 2, 3, 4],\n        []\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "lists",
    "duplicates",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        1,
        2,
        2,
        3,
        4
      ],
      "expectedOutput": [
        1,
        2,
        3,
        4
      ]
    },
    {
      "input": [
        5,
        5,
        5,
        5
      ],
      "expectedOutput": [
        5
      ]
    },
    {
      "input": [
        1,
        2,
        3,
        4
      ],
      "expectedOutput": [
        1,
        2,
        3,
        4
      ]
    },
    {
      "input": [],
      "expectedOutput": []
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d2"
  },
  "title": "Binary Search",
  "description": "Write a function that performs binary search on a sorted list.",
  "template_code": "def main(lst, target):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ([1, 2, 3, 4, 5], 3),\n        ([10, 20, 30, 40], 20),\n        ([5, 6, 7, 8], 1),\n        ([1, 3, 5, 7, 9], 7)\n    ]\n    expected_outputs = [\n        2,\n        1,\n        -1,\n        3\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0], test_input[1])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "searching",
    "binary",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        [
          1,
          2,
          3,
          4,
          5
        ],
        3
      ],
      "expectedOutput": 2
    },
    {
      "input": [
        [
          10,
          20,
          30,
          40
        ],
        20
      ],
      "expectedOutput": 1
    },
    {
      "input": [
        [
          5,
          6,
          7,
          8
        ],
        1
      ],
      "expectedOutput": -1
    },
    {
      "input": [
        [
          1,
          3,
          5,
          7,
          9
        ],
        7
      ],
      "expectedOutput": 3
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d3"
  },
  "title": "Merge Two Sorted Lists",
  "description": "Write a function to merge two sorted lists into a single sorted list.",
  "template_code": "def main(lst1, lst2):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ([1, 3, 5], [2, 4, 6]),\n        ([10, 20], [5, 15]),\n        ([1], []),\n        ([], [7, 8])\n    ]\n    expected_outputs = [\n        [1, 2, 3, 4, 5, 6],\n        [5, 10, 15, 20],\n        [1],\n        [7, 8]\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0], test_input[1])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "lists",
    "merging",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        [
          1,
          3,
          5
        ],
        [
          2,
          4,
          6
        ]
      ],
      "expectedOutput": [
        1,
        2,
        3,
        4,
        5,
        6
      ]
    },
    {
      "input": [
        [
          10,
          20
        ],
        [
          5,
          15
        ]
      ],
      "expectedOutput": [
        5,
        10,
        15,
        20
      ]
    },
    {
      "input": [
        [
          1
        ],
        []
      ],
      "expectedOutput": [
        1
      ]
    },
    {
      "input": [
        [],
        [
          7,
          8
        ]
      ],
      "expectedOutput": [
        7,
        8
      ]
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d4"
  },
  "title": "Fibonacci Sequence",
  "description": "Write a function to return the nth Fibonacci number.",
  "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (1,),\n        (2,),\n        (5,),\n        (10,)\n    ]\n    expected_outputs = [\n        1,\n        1,\n        5,\n        55\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "numbers",
    "fibonacci",
    "beginner"
  ],
  "testcases": [
    {
      "input": 1,
      "expectedOutput": 1
    },
    {
      "input": 2,
      "expectedOutput": 1
    },
    {
      "input": 5,
      "expectedOutput": 5
    },
    {
      "input": 10,
      "expectedOutput": 55
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d5"
  },
  "title": "Palindrome Checker",
  "description": "Write a function to check if a string is a palindrome.",
  "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('racecar',),\n        ('hello',),\n        ('madam',),\n        ('world',)\n    ]\n    expected_outputs = [\n        True,\n        False,\n        True,\n        False\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "palindrome",
    "beginner"
  ],
  "testcases": [
    {
      "input": "racecar",
      "expectedOutput": true
    },
    {
      "input": "hello",
      "expectedOutput": false
    },
    {
      "input": "madam",
      "expectedOutput": true
    },
    {
      "input": "world",
      "expectedOutput": false
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d6"
  },
  "title": "Sum of Even Numbers",
  "description": "Write a function to return the sum of all even numbers in a list.",
  "template_code": "def main(lst):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ([1, 2, 3, 4, 5],),\n        ([10, 11, 12],),\n        ([2, 4, 6],),\n        ([1, 3, 5],)\n    ]\n    expected_outputs = [\n        6,\n        22,\n        12,\n        0\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "numbers",
    "even",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        1,
        2,
        3,
        4,
        5
      ],
      "expectedOutput": 6
    },
    {
      "input": [
        10,
        11,
        12
      ],
      "expectedOutput": 22
    },
    {
      "input": [
        2,
        4,
        6
      ],
      "expectedOutput": 12
    },
    {
      "input": [
        1,
        3,
        5
      ],
      "expectedOutput": 0
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d7"
  },
  "title": "Factorial Calculator",
  "description": "Write a function to calculate the factorial of a number.",
  "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (0,),\n        (1,),\n        (5,),\n        (3,)\n    ]\n    expected_outputs = [\n        1,\n        1,\n        120,\n        6\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "numbers",
    "factorial",
    "beginner"
  ],
  "testcases": [
    {
      "input": 0,
      "expectedOutput": 1
    },
    {
      "input": 1,
      "expectedOutput": 1
    },
    {
      "input": 5,
      "expectedOutput": 120
    },
    {
      "input": 3,
      "expectedOutput": 6
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d8"
  },
  "title": "Reverse a String",
  "description": "Write a function to reverse a string.",
  "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('hello',),\n        ('world',),\n        ('python',),\n        ('racecar',)\n    ]\n    expected_outputs = [\n        'olleh',\n        'dlrow',\n        'nohtyp',\n        'racecar'\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "reverse",
    "beginner"
  ],
  "testcases": [
    {
      "input": "hello",
      "expectedOutput": "olleh"
    },
    {
      "input": "world",
      "expectedOutput": "dlrow"
    },
    {
      "input": "python",
      "expectedOutput": "nohtyp"
    },
    {
      "input": "racecar",
      "expectedOutput": "racecar"
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1d9"
  },
  "title": "Find the Maximum Number",
  "description": "Write a function to find the maximum number in a list.",
  "template_code": "def main(lst):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ([1, 2, 3],),\n        ([5, 10, 15],),\n        ([2, 4, 6, 8],),\n        ([7, 3, 9],)\n    ]\n    expected_outputs = [\n        3,\n        15,\n        8,\n        9\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "numbers",
    "maximum",
    "beginner"
  ],
  "testcases": [
    {
      "input": [
        1,
        2,
        3
      ],
      "expectedOutput": 3
    },
    {
      "input": [
        5,
        10,
        15
      ],
      "expectedOutput": 15
    },
    {
      "input": [
        2,
        4,
        6,
        8
      ],
      "expectedOutput": 8
    },
    {
      "input": [
        7,
        3,
        9
      ],
      "expectedOutput": 9
    }
  ]
},
{
  "_id": {
    "$oid": "670a7f4f862371b638a4f1da"
  },
  "title": "Count Vowels in a String",
  "description": "Write a function to count the number of vowels in a string.",
  "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        ('hello',),\n        ('world',),\n        ('python',),\n        ('racecar',)\n    ]\n    expected_outputs = [\n        2,\n        1,\n        1,\n        3\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "strings",
    "vowels",
    "beginner"
  ],
  "testcases": [
    {
      "input": "hello",
      "expectedOutput": 2
    },
    {
      "input": "world",
      "expectedOutput": 1
    },
    {
      "input": "python",
      "expectedOutput": 1
    },
    {
      "input": "racecar",
      "expectedOutput": 3
    }
  ]
}]