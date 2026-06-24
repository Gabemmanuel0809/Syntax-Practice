package main 
import "fmt"

type emails struct {
   username string
}

type wl struct {
   qnum int
   maxq int  
   emails []emails
}

type insertTowl interface {
	insertUser(e emails) string 
}

func (w *wl)insertUser(e emails) string {
   if(w.qnum < w.maxq) {
	   w.qnum += 1 
	   w.emails = append(w.emails, e)
	   return "Email sucesfully added to wait list"
   } else {
	  return "Wait list is full, please try again later"
   }
}

func add(i insertTowl, e emails) {
    fmt.Println(i.insertUser(e))
}

func main() {
	var e1 = emails{username:"johnabc"}
	var e2 = emails{username:"steviewevie"}
	var e3 = emails{username:"villager999hmm"}
	var e4 = emails{username:"kelly999"}
	var e5 = emails{username:"alexxyz"}
	var e6 = emails{username:"siu999"}

    var wl1 = wl{
		qnum:0,
		maxq:5,
		emails: []emails{},
	}

	add(&wl1, e1)
	add(&wl1, e2)
	add(&wl1, e3)
	add(&wl1, e4)
	add(&wl1, e5)
	add(&wl1, e6)

}
