public class Main {

    // User function logic goes here
    public static String checkAllEven(int[] nums) {
    for (int num : nums) {
        if (num % 2 != 0) {
            return "False";
        }
    }
    return "True";
}

    // Debugging execution of user function
    public static void main(String[] args) {
    // Test case 1
    int[] testCase1 = {2, 4, 6};
    System.out.println(checkAllEven(testCase1));  // Expected output: True
    // Test case 2
    int[] testCase2 = {1, 2, 3};
    System.out.println(checkAllEven(testCase2));  // Expected output: False
    // Test case 3
    int[] testCase3 = {6, 8, 10};
    System.out.println(checkAllEven(testCase3));  // Expected output: True
    // Test case 4
    int[] testCase4 = {2, 4, 5};
    System.out.println(checkAllEven(testCase4));  // Expected output: False
}
}