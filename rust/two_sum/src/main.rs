struct Solution;

impl Solution {

    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut left: usize = 0;
        let mut right: usize = numbers.len() - 1;
        let mut result: Vec<i32> = vec![];
        while left < right {
            let sum: i32 = numbers[left] + numbers[right];
            match sum.cmp(&target) {
                std::cmp::Ordering::Equal => {
                    result.push(left as i32 + 1);
                    result.push(right as i32 + 1);
                    break;
                }
                std::cmp::Ordering::Less => {
                    left += 1;
                }
                std::cmp::Ordering::Greater => {
                    right -= 1;
                }
            }
        }
        result
    }
}

fn main() {
    let numbers = vec![2, 7, 11, 15];
    let target = 9;
    println!("{:?}", Solution::two_sum(numbers, target));
}
