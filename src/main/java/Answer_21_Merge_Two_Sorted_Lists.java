public class Answer_21_Merge_Two_Sorted_Lists {

    public static class ListNode {
        int val;
        ListNode next;

        public ListNode(int x) {
            val = x;
        }

        public String toString() {
            return String.valueOf(val);
        }
    }

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode head = new ListNode(0);
        ListNode tail = head;

        ListNode t1 = l1;
        ListNode t2 = l2;

        while (t1 != null && t2 != null) {
            if (t1.val < t2.val) {
                tail.next = t1;
                tail = t1;
                t1 = t1.next;
            } else {
                tail.next = t2;
                tail = t2;
                t2 = t2.next;
            }
        }

        if (t1 != null) {
            while (t1 != null) {
                tail.next = t1;
                tail = t1;
                t1 = t1.next;
            }
        }

        if (t2 != null) {
            while (t2 != null) {
                tail.next = t2;
                tail = t2;
                t2 = t2.next;
            }
        }

        return head.next;
    }

    public String transferToString(ListNode l) {
        if (l == null) {
            return "";
        }

        StringBuilder sb = new StringBuilder();

        ListNode h = l;
        while (h != null) {
            if (h.next == null) {
                sb.append(h.val);
            } else {
                sb.append(h.val + ",");
            }
            h = h.next;
        }

        return sb.toString();
    }
}
