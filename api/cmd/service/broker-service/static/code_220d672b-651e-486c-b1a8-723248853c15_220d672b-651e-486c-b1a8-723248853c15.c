#include <stdio.h>
#include <ctype.h>

// User function logic goes here
int isPangram(char str[]) {
    strlwr(str);
    int letters[26] = {0};
    for (int i = 0; str[i] != '\0'; i++) {
        if (str[i] >= 'a' && str[i] <= 'z') {
            letters[str[i] - 'a'] = 1;
        }
    }
    for (int i = 0; i < 26; i++) {
        if (letters[i] == 0) {
            return 0;
        }
    }
    return 1;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"The quick brown fox jumps over the lazy dog", "Hello World", "abcdefghijklmnopqrstuvwxyz"};
    int expected[] = {1, 0, 1};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        int result = isPangram(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
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