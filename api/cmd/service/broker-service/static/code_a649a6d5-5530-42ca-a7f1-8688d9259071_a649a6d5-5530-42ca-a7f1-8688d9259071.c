#include <stdio.h>

// User function logic goes here
int countOccurrences(char str[], char ch) {
    int count = 0;
    for (int i = 0; str[i] != '\0'; i++) {
        if (str[i] == ch) {
            count++;
        }
    }
    return count;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"hello world", "goodbye", "abracadabra", "openai"};
    char chars[] = {'o', 'e', 'a', 'b'};
    int expected[] = {2, 2, 5, 2};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int count = countOccurrences(test_cases[i], chars[i]);
        if (count != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: [");
            printf("%s ", test_cases[i]);
            printf("] Expected: %d, Got: %d\n", expected[i], count);
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