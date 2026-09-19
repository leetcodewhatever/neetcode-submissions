impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    
    let n = nums.len() - 1;
    let mut suffix = vec![1;nums.len()+1];
    let mut prefix = vec![1;nums.len()+1];
    let mut result:Vec<i32> = Vec::new();

    for (l,r) in nums.iter().enumerate().zip(nums.iter().rev().enumerate()){
        let rr = n - r.0;
        
        if l.0 == 0 || rr  == nums.len() - 1{
            continue
        }

        prefix[l.0 as usize] = nums[l.0 - 1 as usize] * prefix[l.0 - 1 as usize];
        suffix[rr] = nums[rr + 1] * suffix[rr + 1];
        
    }


    for i in 0..nums.len(){
        result.push(prefix[i as usize] * suffix[i as usize]);
    }

    return result
    }
}
