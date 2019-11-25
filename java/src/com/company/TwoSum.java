package com.company;

import java.util.HashMap;

class TwoSum {
    int[] twoSum(int[] nums, int target) {
        int[] result = new int[2];
        HashMap<Integer, Integer> map = new HashMap<>();
        int i = 0;
        for (int num: nums) {
            if (map.containsKey(num)){
                result[0] = i;
                result[1] = map.get(num);
                break;
            }
            else{
                map.put(target - num, i);
            }
            ++i;
        }
        return result;
    }
}
