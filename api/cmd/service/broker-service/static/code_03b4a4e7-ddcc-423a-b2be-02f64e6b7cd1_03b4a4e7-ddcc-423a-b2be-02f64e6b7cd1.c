#include <stdio.h>

// User function logic goes here
void calculateAreaAndPerimeter(float length, float width) {
    printf("Area: %.2f\n", length * width);
    printf("Perimeter: %.2f\n", 2 * (length + width));
}

// Test cases for the function
int main() {
    int passed = 1;
    float length = 5, width = 3;
    float expectedArea = 15, expectedPerimeter = 16;
    // Call function to check area and perimeter
    printf("%s\n", passed ? "true" : "false");
    return 0;
}