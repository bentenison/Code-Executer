#include <stdio.h>

// User function logic goes here
void reverseString(char str[]) {
    int len = 0;
    while (str[len] != '\0') {
        len++;
    }
    for (int i = 0; i < len / 2; i++) {
        char temp = str[i];
        str[i] = str[len - i - 1];
        str[len - i - 1] = temp;
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Programming", "Hello", "World", "C"};
    char expected[][50] = {"gnimmargorP", "olleH", "dlroW", "C"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
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