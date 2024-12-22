#include <stdio.h>
#include <string.h>

// User function logic goes here
void removeDuplicates(char str[]) {
    int i, j;
    int len = strlen(str);
    for (i = 0; i < len; i++) {
        for (j = i + 1; j < len; j++) {
            if (str[i] == str[j]) {
                int k;
                for (k = j; k < len; k++) {
                    str[k] = str[k + 1];
                }
                len--;
                j--;
            }
        }
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"programming", "hello", "aabbcc", "abcabc"};
    char expected[][50] = {"progamin", "helo", "abc", "abc"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        removeDuplicates(test_cases[i]);
        if (strcmp(test_cases[i], expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], test_cases[i]);
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