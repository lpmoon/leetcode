import spock.lang.Specification

class Answer_23_Merge_k_Sorted_Lists_Test extends Specification {
    def "mergeKLists should work"() {
        def s = new Answer_23_Merge_k_Sorted_Lists()


        expect:
        s.transferToString(s.mergeKLists(l)) == result

        where:
        l || result
//        params() | ""
        (Answer_23_Merge_k_Sorted_Lists.ListNode[]) [toListNode([1, 3, 5, 7]), toListNode([2, 4, 6, 8]), toListNode([99])].toArray() | "1,2,3,4,5,6,7,8,99"
        (Answer_23_Merge_k_Sorted_Lists.ListNode[]) [toListNode([1, 3, 5, 7, 9]), toListNode([2, 4, 6, 8]), toListNode([3, 3, 10])] | "1,2,3,3,3,4,5,6,7,8,9,10"
        (Answer_23_Merge_k_Sorted_Lists.ListNode[]) [toListNode([1, 3, 5, 7]), toListNode([2, 4, 6, 8, 10, 11, 12, 19, 20]), toListNode([13]), toListNode([14]), toListNode([15])] | "1,2,3,4,5,6,7,8,10,11,12,13,14,15,19,20"
        (Answer_23_Merge_k_Sorted_Lists.ListNode[]) [toListNode([]), toListNode([2, 4, 6, 8])] | "2,4,6,8"
        (Answer_23_Merge_k_Sorted_Lists.ListNode[]) [toListNode([1]), toListNode([2, 4, 6, 8])] | "1,2,4,6,8"
        (Answer_23_Merge_k_Sorted_Lists.ListNode[]) [toListNode([]), toListNode([])] | ""
    }

    def toListNode(List values) {
        def head = new Answer_23_Merge_k_Sorted_Lists.ListNode(0)
        def tail = head
        values.each { value ->
            def n = new Answer_23_Merge_k_Sorted_Lists.ListNode(value)
            tail.next = n
            tail = n
        }

        return head.next
    }

    def params() {
        def lists = []
        for (int i = 0; i < 10000; i++) {
            lists.add(new Answer_23_Merge_k_Sorted_Lists.ListNode((int)(Math.random() * 100)))
        }

        return (Answer_23_Merge_k_Sorted_Lists.ListNode[]) lists.toArray()
    }
}
