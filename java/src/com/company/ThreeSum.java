package com.company;

import javafx.util.Pair;

import java.util.*;

class ThreeSum {
    boolean isSolutionValid(Pair<Integer, Integer>[] pairs) {
        HashMap<Integer, Integer> actual = new HashMap<>();
        for (Pair<Integer, Integer> pair: pairs) {
            if (actual.containsKey(pair.getKey())) actual.put(pair.getKey(), actual.get(pair.getKey()) + 1);
            else actual.put(pair.getKey(), 1);
        }

        HashMap<Integer, Integer> real = new HashMap<>();
        for (Pair<Integer, Integer> pair: pairs) {
            if (!real.containsKey(pair.getKey()))
                real.put(pair.getKey(), pair.getValue());
        }

        boolean valid = true;
        for (Map.Entry<Integer, Integer> entry : actual.entrySet()) {
            Integer key = entry.getKey();
            Integer value = real.get(key);
            if (entry.getValue() > value) {
                valid = false;
                break;
            }
        }

        return valid;
    }

    List<List<Integer>> threeSum(int[] nums) {
        HashSet< List<Integer> > result = new HashSet<>();
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int num: nums) {
            if (map.containsKey(num)) map.put(num, map.get(num) + 1);
            else map.put(num, 1);
        }

        for(Map.Entry<Integer, Integer> pair : map.entrySet()) {
            for(Map.Entry<Integer, Integer> pair2 : map.entrySet()) {
                int y = -pair.getKey() - pair2.getKey();
                if (map.containsKey(y))
                {
                    Pair<Integer, Integer>[] solution = new Pair[3];
                    solution[0] = new Pair<> ( pair.getKey(), pair.getValue() );
                    solution[1] = new Pair<> ( pair2.getKey(), pair.getValue() );
                    solution[2] = new Pair<> ( y, map.get(y) );
                    if (isSolutionValid(solution)) {
                        List<Integer> solutionList = Arrays.asList(
                                pair.getKey(), pair2.getKey(), y);
                        Collections.sort(solutionList);
                        result.add(solutionList);
                    }
                }
            }
        }

        return new ArrayList<>(result);
    }
}
