#include <stdio.h>

// User function logic goes here
void reverseString(char str[]) {
    int start = 0, end = strlen(str) - 1;
    while (start < end) {
        char temp = str[start];
        str[start] = str[end];
        str[end] = temp;
        start++;
        end--;
    }
}

// Test cases for the function
int main() {
    char test_cases[][10] = {"hello", "world", "abcd"};
    char expected[][10] = {"olleh", "dlrow", "dcba"};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        reverseString(test_cases[i]);
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