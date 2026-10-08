import "slices"

type TimeMap struct {
     entries  map[string][]ValueTuple
}

type ValueTuple struct { 
	Value string
	Timestamp int
}

func Constructor() TimeMap {
     return TimeMap{
		entries: map[string][]ValueTuple{},
	 }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	
	v := ValueTuple{
		Value: value,
		Timestamp: timestamp,
	}
    
	_, ok := this.entries[key]

	if !ok{
		this.entries[key] = append(this.entries[key], v) 
		return
	}

	i := sort.Search(len(this.entries[key]), func(i int) bool {
		return this.entries[key][i].Timestamp >= timestamp
	})

	this.entries[key] = slices.Insert(this.entries[key], i, v)

}

func (this *TimeMap) Get(key string, timestamp int) string {

	low, high := 0, len(this.entries[key])
    // [1,2,4,5,6]
	for low < high {
		mid := low + (high - low) / 2
        
		if this.entries[key][mid].Timestamp <= timestamp {
           low = mid + 1
		} else  {
			high = mid
		}
	}

	if low == 0 {
		return ""
	} 

	return this.entries[key][low-1].Value
}

