#include <stdio.h>
#include <string.h>

// User function logic goes here
int isPalindrome(char str[]) {
    int len = strlen(str);
    for (int i = 0; i < len / 2; i++) {
        if (str[i] != str[len - i - 1]) {
            return 0;
        }
    }
    return 1;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"madam", "hello", "racecar", "abcd"};
    int expected[] = {1, 0, 1, 0};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = isPalindrome(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
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