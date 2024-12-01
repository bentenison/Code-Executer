public class Main {

    // User function logic goes here
    public static boolean isPrime(int n) {
    if (n <= 1) return false;
    for (int i = 2; i <= Math.sqrt(n); i++) {
        if (n % i == 0) return false;
    }
    return true;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[] testCases = {
        7,  // Test case 1: Prime number
        10   // Test case 2: Not a prime number
    };
    for (int testCase : testCases) {
        boolean expectedOutput = (testCase == 7);
        boolean output = isPrime(testCase);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}