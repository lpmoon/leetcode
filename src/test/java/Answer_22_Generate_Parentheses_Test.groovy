import spock.lang.Specification

class Answer_22_Generate_Parentheses_Test extends Specification {
    def "generateParenthesis should work"() {
        def s = new Answer_22_Generate_Parentheses()

        expect:
        def result = s.generateParenthesis(num)
        Collections.sort(result)
        Collections.sort(expect)
        result == expect


        where:
        num | expect
        1   | ["()"]
        2   | ["(())", "()()"]
        3   | ["((()))", "(()())", "(())()", "()()()", "()(())"]
        4   | ["(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"]
    }
}
