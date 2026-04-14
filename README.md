# WR

## A Basic weighted selection used for slot machines

### download

```bash
go get github.com/kainkent69/wr
```

### Basic Usage
```go
package main_test

import (
	"github.com/kainkent69/wr/src/wr"
)

func main() {
	list := []*wr.W{
		{
			ID:      1,
			Weights: 5000,
		},
		{
			ID:      2,
			Weights: 5000,
		},

		{
			ID:      3,
			Weights: 5000,
		},

		{
			ID:      4,
			Weights: 10000,
			IsEmpty: true,
		},
	}

	// record the slots
	slot := wr.Slots{
		Lists: list,
		Track: true,
	}

	slot.Init(wr.Default)

	sm := wr.Simulator{
		List:  list,
		Spins: 100000,
	}

	report := sm.Run(wr.Default)
	report.Print()

}

```

### Result 

*Note*: this defer because its random

```bash
{                                                                                                                                   
   "HF": 3,                                                                                                                         
   "SAvg": 3,                                                                                                                       
   "Each": {                                                                                                                        
      "1": {                                                                                                                        
         "HF": 0.13636363636363635,                                                                                                 
         "SAvg": 1,                                                                                                                 
         "Each": null,                                                                                                              
         "Hit": 12,                                                                                                                 
         "Fail": 88,                                                                                                                
         "StreakResult": {                                                                                                          
            "1": 12 
         },               
         "MinStreak": 1,
         "MaxStreak": 1
      },    
      "2": {                     
         "HF": 0.21951219512195125,                               
         "SAvg": 1.2, 
         "Each": null,
         "Hit": 18,              
         "Fail": 82,                                   
         "StreakResult": {
            "1": 13,
            "2": 1, 
            "3": 1        
         },         
         "MinStreak": 1,
         "MaxStreak": 3 
      },               
      "3": {      
         "HF": 0.20481927710843373,
         "SAvg": 1.0625,
         "Each": null, 
         "Hit": 17, 
         "Fail": 83,
         "StreakResult": {
            "1": 15,                     
            "2": 1                       
         },                              
         "MinStreak": 1,                 
         "MaxStreak": 2                  
      },                                               
      "4": {                                           
         "HF": 1.127659574468085,                      
         "SAvg": 1.7586206896551724,                   
         "Each": null,                                            
         "Hit": 53,                                    
         "Fail": 47,                                   
         "StreakResult": {                                                                                    
            "1": 19,                                   
            "2": 5,                                    
            "3": 2,                                               
            "5": 2,                                    
            "6": 1                                                
         },                                                                                                                                                                                                                 
         "MinStreak": 1,                                          
         "MaxStreak": 6                                           
      }                                                           
   },                                                             
   "Hit": 300,                                                    
   "Fail": 100,                                                   
   "StreakResult": {                                              
      "3": 100                                                    
   },                                                             
   "MinStreak": 3,                                                
   "MaxStreak": 3                                                 
}                                
```

