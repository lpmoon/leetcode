import java.util.Arrays;

public class Answer_4_Median_of_Two_Sorted_Arrays {
    /*
     *  There are two sorted arrays nums1 and nums2 of size m and n respectively.
     *   Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
     */
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int len = nums1.length + nums2.length;
        int[] all = new int[len];

        int i = 0;
        int j = 0;
        int p = 0;

        if (nums1.length == 0) {
            all = nums2;
        } else if (nums2.length == 0){
            all = nums1;
        } else {
            while (i <= nums1.length - 1 && j <= nums2.length - 1) {
                if (nums1[i] <= nums2[j]) {
                    all[p++] = nums1[i++];
                } else {
                    all[p++] = nums2[j++];
                }

                if (i == nums1.length) {
                    for (int t = j; t < nums2.length; t++) {
                        all[p++] = nums2[t];
                    }

                    break;
                }

                if (j == nums2.length) {
                    for (int t = i; t < nums1.length; t++) {
                        all[p++] = nums1[t];
                    }

                    break;
                }
            }
        }

        System.out.println(Arrays.toString(all));

        if (len % 2 == 0) {
            return ((double) (all[len / 2] + all[len / 2 - 1])) / 2;
        } else {
            return (double) (all[len / 2]);
        }

    }
}
