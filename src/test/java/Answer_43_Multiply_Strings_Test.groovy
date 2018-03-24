import spock.lang.Specification

class Answer_43_Multiply_Strings_Test extends Specification {
    def "multiply should work"() {
        def s = new Answer_43_Multiply_Strings()

        expect:
        s.multiply(num1, num2) == result

        where:
        num1 | num2 || result
        "2"  | "2"  || "4"
        "9"  | "9"  || "81"
        "44" | "99" || "4356"
        "99" | "99" || "9801"
    }
}
