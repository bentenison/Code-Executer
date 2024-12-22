# User function logic goes here
def nth_prime(n):
    sieve = [True] * (n * 10)
    sieve[0] = sieve[1] = False
    for i in range(2, int(len(sieve) ** 0.5) + 1):
        if sieve[i]:
            for j in range(i * i, len(sieve), i):
                sieve[j] = False
    primes = [x for x, prime in enumerate(sieve) if prime]
    return primes[n-1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5),
        (10)
    ]
    expected_outputs = [
        11,
        29
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = nth_prime(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)