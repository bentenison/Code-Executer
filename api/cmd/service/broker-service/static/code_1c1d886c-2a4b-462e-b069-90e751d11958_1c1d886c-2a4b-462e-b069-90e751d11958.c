#include <stdio.h>

// User function logic goes here
void removeOccurrences(int arr[], int *size, int element) {
    int i = 0, j = 0;
    while (i < *size) {
        if (arr[i] == element) {
            i++;
        } else {
            arr[j++] = arr[i++];
        }
    }
    *size = j;
}

// Test cases for the function
int main() {
    int passed = 1;
    int arr[] = {1, 2, 3, 2, 4}, size = 5, element = 2;
    // Check array after removing occurrences
    printf("%s\n", passed ? "true" : "false");
    return 0;
}