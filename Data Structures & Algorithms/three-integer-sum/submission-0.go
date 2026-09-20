

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
    fmt.Println(nums)
	var triplet [][]int
    
    if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
        triplet = append(triplet, []int{0,0,0})
        return triplet
    }

	for i := 0 ; i < len(nums) - 2 ; i++ {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
		l := i + 1
		r := len(nums) - 1
		for {
			if l >= r {
				break
			}

			res := nums[i] + nums[l] + nums[r]
            
            if res > 0 {
                r--
            }

            if res < 0 {
                l++
            }

			if res == 0 {
				triplet = append(triplet, []int{nums[i], nums[l], nums[r]})
                l++
              for {

                if l >= r || nums[l] != nums[l-1]{
                    break
                }
                l++
              }
			}
            

		}

	}
	return triplet

}