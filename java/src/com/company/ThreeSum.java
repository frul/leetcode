package com.company;

import java.util.*;

class ThreeSum {
    List< List<Integer> > twoSum(int[] nums, int i) {
        int target = -nums[i];
        HashSet<Integer> set = new HashSet<>();
        List<List<Integer>> result = new ArrayList<>();
        for (int pos = 0; pos < nums.length; ++pos) {
            if (pos == i) continue;
            int x = nums[pos];
            int y = target - x;
            if (set.contains(x)) {
                List<Integer> list = new ArrayList<>(Arrays.asList(
                        x, y, -target));
                Collections.sort(list);
                result.add(list);
            }
            else {
                set.add(y);
            }
        }

        return result;
    }

    List<List<Integer>> threeSum(int[] nums) {
        HashSet< List<Integer> > result = new HashSet<>();
        for (int i = 0; i < nums.length; ++i) {
            result.addAll(twoSum(nums, i));
        }

        return new ArrayList<>(result);
    }
}
