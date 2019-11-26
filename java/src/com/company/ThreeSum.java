package com.company;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class ThreeSum {
    List<List<Integer>> threeSum(int[] nums) {
        ArrayList< List<Integer> > result = new ArrayList<>();
        HashMap<Integer, ArrayList<Integer>> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                Integer value = - (nums[i] + nums[j]);
                if (map.containsKey(value)) {
                    var list = map.get(value);
                    list.add(nums[i]);
                }
                else
                    map.put(value, new ArrayList<>(nums[i]));
            }
        }

        for (int num: nums) {
            if (map.containsKey(num)) {
                var ingredients = map.get(num);
                for (var ingredient: ingredients) {
                    ArrayList<Integer> resultPart = new ArrayList<>();
                    resultPart.add(num);
                    resultPart.add(ingredient);
                    resultPart.add(-(ingredient + num));
                    result.add(resultPart);
                }
            }
        }

        return result;
    }
}
