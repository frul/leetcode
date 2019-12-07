package com.company;

import java.util.List;

public class Main {

    public static void main(String[] args) {
	    int[] nums = new int[args.length];
		for (int i = 0; i < args.length; i++) {
			nums[i] = Integer.parseInt(args[i]);
		}

		ThreeSum threeSum = new ThreeSum();
		List<List<Integer>> result = threeSum.threeSum(nums);

		System.out.println("Results are:");
		for (List<Integer> part: result)
		{
			for(Integer elem: part)
			{
				System.out.printf("%d, ", elem);
			}
			System.out.println();
		}
    }
}
