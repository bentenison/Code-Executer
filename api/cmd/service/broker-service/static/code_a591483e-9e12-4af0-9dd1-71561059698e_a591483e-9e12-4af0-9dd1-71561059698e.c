#include <stdio.h>
#include <string.h>
#include <stdlib.h>

// User function logic goes here
void sortString(char str[]) {
    qsort(str, strlen(str), sizeof(char), compare);
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"hello", "world", "abcd"};
    char expected[][50] = {"ehllo", "dlorw", "abcd"};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        sortString(test_cases[i]);
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