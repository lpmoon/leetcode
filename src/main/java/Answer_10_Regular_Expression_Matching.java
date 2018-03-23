import java.util.HashMap;
import java.util.Map;

public class Answer_10_Regular_Expression_Matching {
    /*
    Implement regular expression matching with support for '.' and '*'.

    '.' Matches any single character.
    '*' Matches zero or more of the preceding element.

    The matching should cover the entire input string (not partial).

    The function prototype should be:
    bool isMatch(const char *s, const char *p)

    Some examples:
    isMatch("aa","a") → false
    isMatch("aa","aa") → true
    isMatch("aaa","aa") → false
    isMatch("aa", "a*") → true
    isMatch("aa", ".*") → true
    isMatch("ab", ".*") → true
    isMatch("aab", "c*a*b") → true

    */

    public static class State {
        private String subReg;

        private Map<String, State> nextStates = new HashMap<String, State>();

        public State getNext(String nextChar) {
            return nextStates.get(nextChar);
        }

        public void addNext(String nextChar, State n) {
            nextStates.put(nextChar, n);
        }
    }

    public State construct(String regex) {
        State nullState = new State();

        int i = regex.length();
        while (i < regex.length()) {
        }

        return null;
    }

    public boolean isMatch(String s, String p) {
        return true;
    }
}
