#include <stdio.h>

// User function logic goes here
char* longestPalindrome(char* str) {
    int len = strlen(str);
    if (len == 0) return str;
    int start = 0, max_len = 1;
    for (int i = 0; i < len; i++) {
        for (int j = i; j < len; j++) {
            int is_palindrome = 1;
            for (int k = i, l = j; k < l; k++, l--) {
                if (str[k] != str[l]) {
                    is_palindrome = 0;
                    break;
                }
            }
            if (is_palindrome && (j - i + 1) > max_len) {
                max_len = j - i + 1;
                start = i;
            }
        }
    }
    char* result = (char*)malloc(max_len + 1);
    strncpy(result, str + start, max_len);
    result[max_len] = '\0';
    return result;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"babad", "cbbd", "a", "racecar"};
    char* expected[] = {"bab", "bb", "a", "racecar"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        char* result = longestPalindrome(test_cases[i]);
        if (strcmp(result, expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s\nExpected: %s, Got: %s\n", test_cases[i], expected[i], result);
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