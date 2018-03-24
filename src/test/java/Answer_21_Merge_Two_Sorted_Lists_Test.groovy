import spock.lang.Specification

class Answer_21_Merge_Two_Sorted_Lists_Test extends Specification {
    def "merge2Lists should work"() {
        def s = new Answer_21_Merge_Two_Sorted_Lists()


        expect:
        s.transferToString(s.mergeTwoLists(l1, l2)) == result

        where:
        l1                          | l2                                  || result
        toListNode([1, 3, 5, 7])    | toListNode([2, 4, 6, 8])             | "1,2,3,4,5,6,7,8"
        toListNode([1, 3, 5, 7, 9]) | toListNode([2, 4, 6, 8])             | "1,2,3,4,5,6,7,8,9"
        toListNode([1, 3, 5, 7])    | toListNode([2, 4, 6, 8, 10, 11, 12]) | "1,2,3,4,5,6,7,8,10,11,12"
        toListNode([])              | toListNode([2, 4, 6, 8])             | "2,4,6,8"
        toListNode([1])             | toListNode([2, 4, 6, 8])             | "1,2,4,6,8"
        toListNode([])              | toListNode([])                       | ""
    }

    def toListNode(List values) {
        def head = new Answer_21_Merge_Two_Sorted_Lists.ListNode(0)
        def tail = head
        values.each { value ->
            def n = new Answer_21_Merge_Two_Sorted_Lists.ListNode(value)
            tail.next = n
            tail = n
        }

        return head.next
    }
}
