#include <stdio.h>
#include <string.h>

// User function logic goes here
int areAnagrams(char str1[], char str2[]) {
    if (strlen(str1) != strlen(str2)) return 0;
    int count[256] = {0};
    for (int i = 0; str1[i] != '\0'; i++) {
        count[str1[i]]++;
        count[str2[i]]--;
    }
    for (int i = 0; i < 256; i++) {
        if (count[i] != 0) return 0;
    }
    return 1;
}

// Test cases for the function
int main() {
    char test_cases1[][50] = {"listen", "triangle", "rat", "hello"};
    char test_cases2[][50] = {"silent", "integral", "tar", "bye"};
    int expected[] = {1, 1, 1, 0};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = areAnagrams(test_cases1[i], test_cases2[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Inputs: %s, %s. Expected: %d, Got: %d\n", test_cases1[i], test_cases2[i], expected[i], result);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1; // Return 1 to indicate true
    } else {
        printf("false\n");
        return 0; // Return 0 to indicate false
    }
}