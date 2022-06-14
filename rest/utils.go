package main

import (
	"time"
	"encoding/json"
)

func FilterPeople(arr []Person, cond func(Person) bool) []Person {
	result := []Person{}
	for i := range arr {
	  if cond(arr[i]) {
		result = append(result, arr[i])
	  }
	}
	return result
 }

 func FilterPlanets(arr []Planet, cond func(Planet) bool) []Planet {
	 result := []Planet{}
	 for i := range arr {
	   if cond(arr[i]) {
		 result = append(result, arr[i])
	   }
	 }
	 return result
  }
  
   func SetTimeout(fn func(), millis int) {
	  time.AfterFunc(time.Duration(millis) * time.Millisecond, func () {
		  fn()
	  })
   }
  

func Stringify(obj interface{}) string {
	jsonValue, _ := json.Marshal(obj)
	return string(jsonValue)
}
