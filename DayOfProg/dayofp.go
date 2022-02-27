package main

import(

	"fmt"
)


func main(){

	var year []int32 {1700,1704,1722,1800,1812,1830,1900,1917,1918,1919,2000,2100,2200,2300,2400,2500,2600,2700}


	if year >= 1700 && year <=2700{
        if year % 4 == 0{
            return fmt.Sprintf("12.09%d",+ year)
        
        }else {
        
            return fmt.Sprintf("13.09%d",+year)
        }
    } else if year >=1919 && year<=2700{
        if year %4==0 && year % 100 !=0 || year%400==0{
            return fmt.Sprintf("12.09%d",+year)
        }
    } 
    return fmt.Sprintf("13.09%d",+year)



}