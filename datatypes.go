package main

import "fmt"

func datatypes() {
   /*uint8,uint16,uint64,int8,same*/
   // varible_list can be anyone defined datatype
  // var variable_list optional_data_type;
  // variable_list = m,v int;
  var a, b, c = 3, 4, "abc";


  // var  i, j, k int;
  // var  f, salary float32;
  var d =  42;
  fmt.Println(d);
  fmt.Printf("d is of type %T\n", d);

  fmt.Println(a,b,c)
  fmt.Printf("d is of type %T\t,%T\t,%T\n",a,b,c )
}
