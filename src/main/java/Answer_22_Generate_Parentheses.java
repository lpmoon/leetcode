import java.util.ArrayList;
import java.util.List;

public class Answer_22_Generate_Parentheses {
    /*

    Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
    For example, given n = 3, a solution set is:
    [
        "((()))",
        "(()())",
        "(())()",
        "()(())",
        "()()()"
    ]
    */

    public List<String> generateParenthesis(int n) {
        List<String> all = new ArrayList<String>();
        generate("", n, n, all);
        System.out.println(all);
        return all;
    }

    private int generate(String current, int leaveLeft, int leaveRight, List<String> all) {
        if (leaveLeft  == 0) {
            current += repeat(")", leaveRight);
            all.add(current);
            return 1;
        }

        if (leaveRight  == 0) {
            current += repeat("(", leaveLeft);
            all.add(current);
            return 1;
        }

        if (leaveLeft >= leaveRight) {
            return generate(current + "(", leaveLeft - 1, leaveRight, all);
        } else {
            return generate(current + "(", leaveLeft - 1, leaveRight, all) + generate(current + ")", leaveLeft, leaveRight - 1, all);
        }
    }

    private String repeat(String s, int time) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < time; i++) {
            sb.append(s);
        }

        return sb.toString();
    }
}
