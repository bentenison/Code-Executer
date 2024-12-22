#include <stdio.h>
#include <string.h>

// User function logic goes here
char firstNonRepeating(char str[]) {
    int freq[256] = {0};
    for (int i = 0; str[i] != '\0'; i++) {
        freq[str[i]]++;
    }
    for (int i = 0; str[i] != '\0'; i++) {
        if (freq[str[i]] == 1) {
            return str[i];
        }
    }
    return -1;
}

// Test cases for the function
int main() {
    int passed = 1;
    char test_cases[3][100] = {"swiss", "hello", "aabbcc"};
    char expected[] = {'w', 'h', -1};
    for (int i = 0; i < 3; i++) {
        char result = firstNonRepeating(test_cases[i]);
        if (result != expected[i]) {
            passed = 0;
            printf("Failed for Input: %s. Expected: %c, Got: %c\n", test_cases[i], expected[i], result);
        }
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}