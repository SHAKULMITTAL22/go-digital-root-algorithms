



func DigitalRootWithString(input string) int {
    sum := 0

    for _, r := range input {
        digitValue := int(r - '0')

        if digitValue >= 0 && digitValue <= 9 {
            sum += digitValue
        } else {
            panic("Invalid input")
        }
    }

    return digitalRoot(sum)
}

// Replace with your implementation of `digital_root` function
func digitalRoot(n int) int {
    return (n / 10) + n%10
}