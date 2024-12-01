public class Main {

    // User function logic goes here
    public static String checkEvenOrOdd(int num) {
    if (num % 2 == 0) {
        return "Even";
    } else {
        return "Odd";
    }
}

    // Debugging execution of user function
    public static void main(String[] args) {
    // Call the function with a test case
    System.out.println(checkEvenOrOdd(3));  // Test case 1
}
}