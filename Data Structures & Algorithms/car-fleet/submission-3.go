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

	  var maxTime float64
	  var fleets int
	  for _, car := range cars {
          tta := float64(target-car.Position) / float64(car.Speed)
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
