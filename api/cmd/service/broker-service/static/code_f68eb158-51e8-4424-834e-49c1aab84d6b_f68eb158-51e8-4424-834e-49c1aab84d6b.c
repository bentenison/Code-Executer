#include <stdio.h>

// User function logic goes here
char* removeChar(char* str, char ch) {
    int i = 0, j = 0;
    while (str[i]) {
        if (str[i] != ch) {
            str[j++] = str[i];
        }
        i++;
    }
    str[j] = '\0';
    return str;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello World", "abracadabra", "apple", "banana"};
    char chars[] = {'o', 'a', 'p', 'n'};
    char* expected[] = {"Hell Wrld", "brcdbr", "le", "baa"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        char* result = removeChar(test_cases[i], chars[i]);
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