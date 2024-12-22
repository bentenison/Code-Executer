#include <stdio.h>
#include <string.h>

// User function logic goes here
int countDuplicates(char str[]) {
    int count[256] = {0};
    int duplicateCount = 0;
    for (int i = 0; str[i] != '\0'; i++) {
        count[(unsigned char)str[i]]++;
    }
    for (int i = 0; i < 256; i++) {
        if (count[i] > 0
) {
            duplicateCount++;
        }
    }
    return duplicateCount;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"programming", "hello", "abc", "mississippi"};
    int expected[] = {3, 2, 0, 4};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = countDuplicates(test_cases[i]);
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