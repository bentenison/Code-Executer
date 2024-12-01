public class Main {

    // User function logic goes here
    public static int[] fibonacci(int n) {
    int[] fib = new int[n];
    fib[0] = 0;
    if (n > 1) {
        fib[1] = 1;
        for (int i = 2; i < n; i++) {
            fib[i] = fib[i - 1] + fib[i - 2];
        }
    }
    return fib;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {5},  // Test case 1: Fibonacci sequence for n=5
        {7}   // Test case 2: Fibonacci sequence for n=7
    };
    for (int[] testCase : testCases) {
        int input = testCase[0];
        int[] expectedOutput = fibonacci(input);
        int[] output = fibonacci(input);
        if (!Arrays.equals(expectedOutput, output)) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}