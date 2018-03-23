import java.util.Arrays;

public class Answer_43_Multiply_Strings {
    /*
    Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.
    Note:

    The length of both num1 and num2 is < 110.
    Both num1 and num2 contains only digits 0-9.
    Both num1 and num2 does not contain any leading zero.
    You must not use any built-in BigInteger library or convert the inputs to integer directly
     */

    public String multiply(String num1, String num2) {
        int length1 = num1.length();
        int length2 = num2.length();

        int row = Math.min(length1, length2);
        int column = length1 + length2;

        int record[][] = new int[row][column];

        String multi1;
        String multi2;

        if (length1 > length2) {
            multi1 = num1;
            multi2 = num2;
        } else {
            multi1 = num2;
            multi2 = num1;
        }

        int p = column - 1;

        for (int i = multi2.length() - 1; i >=0 ; i--) {
            // 计算num2的第i位乘以num1的结果
            int j = multi1.length() - 1;
            int p1 = p;

            int tmp = 0;
            while (j >= 0) {
                int m = Integer.parseInt(multi2.substring(i, i + 1)) * Integer.parseInt(multi1.substring(j, j + 1));
                m += tmp;
                tmp = m / 10;
                record[multi2.length() - 1 - i][p1] = m % 10;
                j--;
                p1--;
            }

            p--;

            if (tmp != 0 ) {
                record[multi2.length() - 1 - i][p1] = tmp;
            }
        }

        int result[] = new int[column];

        int tmp = 0;

        for (int i = 0; i < row; i++) {
            tmp = 0;
            for (int j = column - 1; j >= 0; j--) {
                int s = record[i][j] + result[j] + tmp;
                result[j] = s % 10;
                tmp = s / 10;
            }
        }

        String resultS = tmp > 0 ? String.valueOf(tmp) : "";

        System.out.println(Arrays.toString(result));

        int j = result.length;

        for (int i = 0; i < result.length; i++) {
            if (result[i] != 0) {
                j = i;
                break;
            }
        }

        for (int i = j; i < result.length; i++) {
            resultS += String.valueOf(result[i]);
        }

        return resultS.equals("") ? "0" : resultS;
    }

    public static void main(String[] args) {
        Answer_43_Multiply_Strings s = new Answer_43_Multiply_Strings();
        System.out.println(s.multiply("999", "999"));
    }

}
