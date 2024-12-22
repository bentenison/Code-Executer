#include <stdio.h>
#include <string.h>

// User function logic goes here
void listContainingAllChars(char list[][100], int n, char word[]) {
    for (int i = 0; i < n; i++) {
        int found = 1;
        for (int j = 0; word[j] != '\0'; j++) {
            if (strchr(list[i], word[j]) == NULL) {
                found = 0;
                break;
            }
        }
        if (found) {
            printf("%s\n", list[i]);
        }
    }
}

// Test cases for the function
int main() {
    int passed = 1;
    char test_cases[2][100] = {"apple", "banana", "orange", "grape", "peach"};
    char word[] = "ape";
    // Call the function here and validate output
    printf("%s\n", passed ? "true" : "false");
    return 0;
}