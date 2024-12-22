#include <stdio.h>
#include <string.h>

// User function logic goes here
void reverseWords(char str[]) {
    int start = 0, end = 0, len = strlen(str);
    while (end <= len) {
        if (str[end] == ' ' || str[end] == '\0') {
            int word_end = end - 1;
            while (start < word_end) {
                char temp = str[start];
                str[start] = str[word_end];
                str[word_end] = temp;
                start++;
                word_end--;
            }
            start = end + 1;
        }
        end++;
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"hello world", "the quick brown fox", "a b c"};
    char expected[][50] = {"olleh dlrow", "eht kciuq nworb xof", "a b c"};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        reverseWords(test_cases[i]);
        if (strcmp(test_cases[i], expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], test_cases[i]);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1; // Return 1 to indicate true
    } else {
        printf("false\n");
        return 0; // Return 0 to indicate false
    }
}