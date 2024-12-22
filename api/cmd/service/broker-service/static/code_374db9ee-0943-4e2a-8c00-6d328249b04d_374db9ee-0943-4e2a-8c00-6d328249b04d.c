#include <stdio.h>
#include <string.h>

// User function logic goes here
int longestUniqueSubstring(char str[]) {
    int max_len = 0, start = 0;
    int map[256] = {0};
    for (int end = 0; str[end] != '\0'; end++) {
        start = (map[str[end]] > start) ? map[str[end]] : start;
        max_len = (max_len > end - start + 1) ? max_len : end - start + 1;
        map[str[end]] = end + 1;
    }
    return max_len;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"abcabcbb", "bbbbb", "pwwkew", ""};
    int expected[] = {3, 1, 3, 0};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = longestUniqueSubstring(test_cases[i]);
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