
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

// TODO 是否使用多线程？是否使用二叉树来搜索插入节点？
public class Answer_23_Merge_k_Sorted_Lists {

    Comparator c = new Comparator<ListNode>() {
        @Override
        public int compare(ListNode o1, ListNode o2) {
            return Integer.compare(o1.val, o2.val);
        }
    };

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

    /*
        Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
     */
    public ListNode mergeKLists(ListNode[] lists) {
        long start = System.currentTimeMillis();
        if (lists.length == 0) {
            return null;
        }

        ListNode h = lists[0];
        List<ListNode> sortedList = transferToList(h);
        for (int i = 1; i < lists.length; i++) {
            h = merge2Lists(h, lists[i], sortedList);
        }
        long end = System.currentTimeMillis();

        System.out.println(end - start);
        return h;
    }

    public ListNode merge2Lists(ListNode l1, ListNode l2, List<ListNode> l1List) {
        if (l2 == null) {
            return l1;
        }

        if (l1 == null) {
            ListNode t2 = l2;
            while (t2 != null) {
                l1List.add(t2);
                t2 = t2.next;
            }
            return l2;
        }

        ListNode t = l1;
        ListNode t1 = l1;
        ListNode t2 = l2;

//        解法1
//        while (t2 != null) {
//            int index = getInsertAfterNode(l1List, t2);
//            if (index < 0) {
//                ListNode tmp = t2.next;
//                t2.next = t1;
//                t1 = t2;
//                l1List.add(0, t2);
//
//                t2 = tmp;
//            }  else {
//                ListNode tmp = t2.next;
//                ListNode insertNode = l1List.get(index);
//
//                ListNode tmp2 = insertNode.next;
//                insertNode.next = t2;
//                t2.next = tmp2;
//                l1List.add(index + 1, t2);
//
//                t2 = tmp;
//            }
//        }

        // 解法2
        int index = getInsertAfterNode(l1List, t2);
        if (index < 0) {
            ListNode tmp = t2.next;
            t2.next = t1;
            t1 = t2;
            l1List.add(0, t2);

            t2 = tmp;
            index = 0;
            t= t1;
        } else {
            t1 = l1List.get(index);
        }

        ListNode prev = null;

        while (t1 != null && t2 != null) {
            if (t1.val > t2.val) {
                ListNode tmp1 = t2.next;

                prev.next = t2;
                t2.next = t1;
                prev = t2;

                l1List.add(index, t2);
                t2 = tmp1;
                index++;

            } else {
                prev = t1;
                t1 = t1.next;
                index++;
            }
        }

        if (t1 == null && t2 != null) {
            while (t2 != null) {
                prev.next = t2;
                prev = t2;
                l1List.add(index, t2);
                index++;
                t2 = t2.next;
            }
        }

        return t;
    }

    public int getInsertAfterNode(List<ListNode> lists, ListNode n) {
        int index = Collections.binarySearch(lists, n, c);

        if (index < 0) {
            index = -(index + 1) - 1;
        }

        return index;
    }

    public List<ListNode> transferToList(ListNode l) {
        List<ListNode> lists = new ArrayList<>();
        ListNode head = l;
        while (head != null) {
            lists.add(head);
            head = head.next;
        }

        return lists;
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
