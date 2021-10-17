package main

import (
	"fmt"
	//"errors"
)

type Movie struct {
	Length int
	Name   string
	rating float32
	play 	int
	viewers int
}
type Theatre struct {
	//viewers int
	movies 	[]Movie
	//play	int
}
//pointer receiver
func (x *Movie) Rate(p float32) error {
	x.rating = p
	if x.play == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}
	return fmt.Errorf("error from Rate")
}
func (y *Movie) Play(n int) {
	y.play = y.play + 1
	y.viewers = y.viewers + n
}
// value receiver
func (z Movie) Viewers() int{
	//z.viewers = p
	//fmt.Println(z.viewers)
	return z.viewers

}
func (l Movie) Plays() int{
	return l.play
}
func (a Movie) Rating() float64 {
	var c float32 = a.rating
	//var b float64 = float64(c)
	fmt.Printf("Underlying Type of b: %T\n", c)
	fmt.Println(c)
	b2 := float64(c)
	fmt.Printf("Underlying Type of b2: %T\n", b2)
	fmt.Println(b2)
	var x int = a.play
	//var y float64 = float64(x)
	fmt.Printf("Underlying Type of x: %T\n", x)
	fmt.Println(x)
	y2 := float64(x)
	fmt.Printf("Underlying Type of y: %T\n", y2)
	fmt.Println(y2)

	//fmt.Println(p)
	return b2/y2
}
func (k Movie) String(n string, l int, r float32) string { 
	k.Name = n
	k.Length =l
	k.rating = r
	return fmt.Sprintf("%s (%dm) %.1f%%", n, l, r)
	//return k.Name, k.Length
}
//CritiqueFn
func CritiqueFn(x *Movie) (float32,error){
	fmt.Println("========CritiqueFn========")
	fmt.Println(x.Name)
	if x.Name == "" {						//if there is no movie exists it'll return 0 and error message
		return 0.0,fmt.Errorf("no movies to play")  //res =0 and error="no movies to play"
	}
	return 6.89,nil					//res=6.89 and error=nil
}


// 3rd part
func (b *Theatre) Play(n int, m ...Movie) error { 
	//b.viewers = n
	//m.Play(n)
	b.movies = m
	if len(m) >= 0 {
		for _, v := range m {
			v.Play(n)
		}
	}
	fmt.Println(len(m))
	if len(m) == 0 {
		return fmt.Errorf("no movies to play from theatre play")
	}
	//var x string = m.Name
	//fmt.Println(m)
	//fmt.Println(n) 
	return fmt.Errorf("error from theatre Play")
}

//4th part
func (c Theatre) Critique(x []Movie) error {
	c.movies = x 
	for _, v := range x {
		fmt.Println("========Critique========")
		fmt.Println(v)
		v.Play(1)
		res,err:=CritiqueFn(&v)
		fmt.Println(res)
		fmt.Println(err)
		if res !=0 {
			fmt.Println(v.Rate(res))
		}
		
	}
	return fmt.Errorf("error from Critique")
}


func main() {
		m := &Movie{
			Length: 3,
			Name: "Titanic",
			rating: 1.2,
			play: 0,
			viewers: 0,
		}
		t := &Theatre{
			//viewers :0,
			movies : []Movie{},
		}
		// call rating
		fmt.Println(m.Rate(1.2))
		//fmt.Println(m.rating)

		//call viewers and play
		m.Play(50)
		//fmt.Println(m.Play(50))
		fmt.Println("number of plays:", m.play)
		fmt.Println("number of viewers:", m.viewers)

		//call viewers
		s1 := m.Viewers()
		fmt.Println("people who have viewed:", s1)
		//call play
		s2 := m.Plays()
		fmt.Println("times the movie has been played:", s2)
		//call rating
		s3 := m.Rating()
		fmt.Println("rating of the movie in float64:", s3)

		// call name, length and rating
		fmt.Println(m.String("Wizard of Oz", 107, 99.56))
		//s3 := m.String("Wizard of Oz (102m) 99.0%")
		//fmt.Println(s3)

		//3rd part
		movies := []Movie{}
		//t.Play(100, movies...)
		fmt.Println(t.Play(100, movies...))

		//4th part
		movies1 := []Movie{{1,"abc",1,1,1},{2,"xyz",2,2,2}, {3, "mnp",3,3,3}, {4, "rst",4,4,4}}
		fmt.Println(t.Critique(movies1))
}


