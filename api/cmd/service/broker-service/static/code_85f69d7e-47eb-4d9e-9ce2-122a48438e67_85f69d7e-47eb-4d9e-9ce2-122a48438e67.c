#include <stdio.h>

// User function logic goes here
char firstNonRepeatingChar(char str[]) {
    int count[256] = {0};
    for (int i = 0; str[i] != '\0'; i++) {
        count[(unsigned char)str[i]]++;
    }
    for (int i = 0; str[i] != '\0'; i++) {
        if (count[(unsigned char)str[i]] == 1) {
            return str[i];
        }
    }
    return '_';
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"swiss", "racecar", "aabbcc", "abacabad"};
    char expected[] = {'w', 'e', '_', 'c'};
    int all_passed = 1;
    for (int i = 0; i < 4; i++) {
        char result = firstNonRepeatingChar(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %c, Got: %c\n", test_cases[i], expected[i], result);
        }
    }
    if (all_passed) {
        printf("true\n");
    } else {
        printf("false\n");
    }
    return 0;
}