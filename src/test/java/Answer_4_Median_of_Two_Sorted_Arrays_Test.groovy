import spock.lang.Specification

public class Answer_4_Median_of_Two_Sorted_Arrays_Test extends Specification {
    def "findMedianSortedArrays should work"() {
        def s = new Answer_4_Median_of_Two_Sorted_Arrays()

        expect:
        s.findMedianSortedArrays(nums1, nums2) == result

        where:
        nums1                          | nums2                     || result
        [0] as int[]                   | [2] as int[]              || 1
        [] as int[]                    | [2] as int[]              || 2
        [2] as int[]                   | [] as int[]               || 2
        [0, 1, 2] as int[]             | [3, 4, 5] as int[]        || 2.5
        [1, 4, 6] as int[]             | [2, 3, 5] as int[]        || 3.5
        [1, 4, 6, 10, 11, 23] as int[] | [2, 3, 5, 7, 25] as int[] || 6
    }
}
