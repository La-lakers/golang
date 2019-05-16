package main


import "fmt"


func main(){
   var x uint8 = 1<<2 | 1 << 4  // 100 OR 10000 =  10100
   var y uint8 = 1<<3 | 1<<5 //  1000 OR 100000 = 101000
   fmt.Printf("%08b\n", x)  // 
   fmt.Printf("%08b\n", y)
   fmt.Printf("%08b\n", x | y) //OR 00111100
   fmt.Printf("%08b\n", x^y)  //异或 00111100
   fmt.Printf("%08b\n", x&y)  //与 00000000
   fmt.Printf("%08b\n", x&^y) // 按位清零  00101000
   // 找到位是1的位置
   for i := uint8(0); i < 8; i++{
       if x&(1<<i) != 0 {
           fmt.Println(i)
       }
   }

}
