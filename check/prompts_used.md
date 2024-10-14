{
  "title": "Palindrome Checker",
  "description": "Write a function that checks whether a given input is a palindrome.",
  "template_code": "def main(input_str):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    all_passed = True\n    test_cases = [\n        (\"A man a plan a canal Panama\",),\n        (\"racecar\",),\n        (\"hello\",)\n    ]\n    expected_outputs = [\n        True,\n        True,\n        False\n    ]\n    for test_input, expected in zip(test_cases, expected_outputs):\n        result = main(test_input[0])\n        if result != expected:\n            all_passed = False\n            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')\n    print(all_passed)",
  "language": "python",
  "language_code": "py",
  "difficulty": "easy",
  "tags": [
    "string",
    "palindrome",
    "beginner"
  ]
}. Create 10 python questions which follows given structure to insert in mongodb collection.change the template code accordingly.there should be no syntax error in template when executed with user logic.


ideas
make the executor code generic per language 
add the code stats from container
add the stats and metrics in db
implement get containers
check health of the containers and return stats
add the service discovery to identify the health of service
add metrics exporters
