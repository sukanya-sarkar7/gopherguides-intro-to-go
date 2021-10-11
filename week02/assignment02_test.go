//package assignment02_test
package main


import (
	"testing"
	"fmt"
	"github.com/bmizerany/assert"
)	


func TestCollectionArray(t *testing.T) {	

	//assert := assert.New(t)
	//working with array
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	// act as empty variable
	act := [4]string{}
	//act := make([]string, 0, len(exp))
	//iteration in exp
	for i := range exp {
		act[i] = exp[i]
	}
	fmt.Printf("%q\n", act)
	
	for i := range act {
		assert.Equal(t, exp[i], act[i], "The two words should be the same.")

	}

}	

func TestCollectionSlice(t *testing.T) {	

	//assert := assert.New(t)
	//working with Slice
	exp := []string{"John", "Paul", "George", "Ringo"}
	// act as empty variable
	act := make([]string, 0, len(exp))
	//iteration in exp
	for i := range exp {
		act = append(act, exp[i])
	}
	fmt.Printf("%q\n", act)

	for i := range act {
		assert.Equal(t, exp[i], act[i], "content of act and exp are the same.")
		//assert.Equal(t, len(exp), len(act), "The two words should be the same.")

	}
	assert.Equal(t, len(exp), len(act), "length of act and exp are the same.")

}

func TestCollectionMap(t *testing.T) {
	//working with Map
	exp := map[int]string{
		1: "John",
		2: "Paul",
		3: "George",
		4: "Ringo",
	}
	fmt.Println("expected map :", exp)
	act := map[int]string{}
	for i := range exp{
		fmt.Println(i)
		act[i] = exp[i]
	}
	fmt.Println("actual map :", act)

	for v := range act {
		fmt.Println(v)
		fmt.Println(exp[v])
		fmt.Println(act[v])
		assert.Equal(t, exp[v], act[v], "content of act and exp are the same.")
	}	
	//fmt.Println(exp)
	//act := make([]string, 0, len(exp))
	//iteration in exp
	//for i := range exp {
	//	fmt.Println(i)
	//	act = append(act, exp[i])
	//}
	//fmt.Printf("%q\n", act)

	//for v := range act {
	//	 fmt.Println(v)
	//	 fmt.Println(exp[v+1])
	//	 fmt.Println(act[v])
	//	assert.Equal(t, exp[v+1], act[v], "content of act and exp are the same.")
	//}
	assert.Equal(t, len(exp), len(act), "length of act and exp are the same.")

}		
