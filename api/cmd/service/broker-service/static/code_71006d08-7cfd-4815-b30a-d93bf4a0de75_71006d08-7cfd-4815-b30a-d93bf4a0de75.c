#include <stdio.h>

// User function logic goes here
int countDistinctSubstrings(char* str) {
    int len = strlen(str);
    int count = 0;
    for (int i = 0; i < len; i++) {
        for (int j = i + 1; j <= len; j++) {
            char* substr = (char*)malloc(j - i + 1);
            strncpy(substr, str + i, j - i);
            substr[j - i] = '\0';
            // Store the substring in a set to count distinct substrings
            if (substr) {
                count++; // Placeholder for distinct counting
            }
            free(substr);
        }
    }
    return count;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"abab", "abc", "aaa"};
    int expected[] = {6, 6, 3};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        int result = countDistinctSubstrings(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s\nExpected: %d, Got: %d\n", test_cases[i], expected[i], result);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1;
    } else {
        printf("false\n");
        return 0;
    }
}