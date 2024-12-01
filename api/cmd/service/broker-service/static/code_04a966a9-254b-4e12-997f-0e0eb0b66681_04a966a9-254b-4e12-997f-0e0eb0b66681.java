public class Main {

    // User function logic goes here
    public static int findLargest(int[] arr) {
    int max = arr[0];
    for (int num : arr) {
        if (num > max) {
            max = num;
        }
    }
    return max;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {3, 5, 1, 9, 2},  // Test case 1: Largest number is 9
        {10, 20, 15, 5}   // Test case 2: Largest number is 20
    };
    for (int[] testCase : testCases) {
        int expectedOutput = Arrays.stream(testCase).max().getAsInt();
        int output = findLargest(testCase);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}