impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {

     let n =  nums.len();
     let mut result:Vec<i32> = Vec::new();
     for i in 0..n{
        for j in i+1..n{
            if nums[i as usize] + nums[j as usize] == target {
                result.push(i as i32);
                result.push(j as i32);
                return result
            }
     }
     }
    return vec![];
    }
}
