func carFleet(target int, position []int, speed []int) int {
      var cars []Pair
	  for idx := range position {
		  pair := Pair{
			Position : position[idx],
			Speed: speed[idx],
		  }
		  cars = append(cars, pair)
	  }
      
	  sort.Slice(cars, func (i,j int) bool {
		 return cars[i].Position > cars[j].Position
	  })
      
	  var maxTime int
	  var fleets int
	  for _, car := range cars {
          tta := (target - car.Position) / car.Speed
          if tta > maxTime {
			maxTime = tta
			fleets++
		  }
	  }
	    
	  return fleets
}

type Pair struct{
	Position int
	Speed int
}
