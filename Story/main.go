package main

import(
	"fmt"
	"bufio"
	"os"
	//"log"
	"time"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input Resto Name: ")
	name, _ := reader.ReadString('\n')

	ti := time.Now()
	tim := ti.Format("15:04:05")

    fmt.Println("Input date of print: ")
	//date, _ := reader.ReadString('\n')
	var date string
	fmt.Scanln(&date)

	fmt.Println("Input cashier name: ")
    cashier, _ := reader.ReadString('\n')


	fmt.Println("Input n item: ")
    var n int
    fmt.Scanln(&n)
	
	var item map[string]int
	item = make(map[string]int)
	var total int = 0
    for i:=0; i<n; i++{
		
		fmt.Println("input item name: ")
		//s, _ := reader.ReadString('\n')
		var s string
		fmt.Scanln(&s)
		fmt.Println("input item price: ")
		var p int
		fmt.Scanln(&p)
	
		item[s] = p
		total += p
	}
	s := "\t" + name + "\t"
    s1 := "Nama Kasir :\t" + cashier

	fmt.Printf("\n\n")
	
	if (len(s) <= 10){
		fmt.Printf(s)
	}
	if ((len(s) <= 20) && (len(s) > 10)){
		s = strings.TrimLeft(s, "\t")
		fmt.Printf("     %s\n",s)
	}
	if ((len(s) <= 30) && (len(s) > 20)){
		s = strings.TrimLeft(s, "\t")
		fmt.Printf("   %s\n",s)
	}
	if (len(s) > 30){
		s := strings.SplitN(name, " ", 2)
		var v []string
		for i := range s{
			v = append(v,s[i])
		}
			fmt.Printf("\t %s\n", v[0])
			fmt.Printf("   %s \n", v[1])		
	}


	st := "Tanggal : " + date + " " + tim
	fmt.Println(st)
	fmt.Println()

	if (len(s1) <= 10){
		fmt.Printf(s1)
	}
	if ((len(s1) <= 20) && (len(s1) > 10)){
		
		fmt.Println(s1)
	}
	if ((len(s1) <= 30) && (len(s1) > 20)){
	
		fmt.Println(s1)
	}
	if (len(s1) > 30){
		s1 := strings.SplitN(cashier, " ", 2)
		var v []string
		for i := range s1{
			v = append(v,s1[i])
		}
			fmt.Printf("Nama Kasir :\t    %s\n", v[0])
			fmt.Printf("   %s ", v[1])		
	}


	
	fmt.Printf("==============================\n\n")
    

	for k,v := range item{
		
		s2 := k +  "........................Rp "
		if (len(s2) <= 10){
			fmt.Printf("%s%d\n",s2,v)
		} 
		if ((len(s2) <= 20) && (len(s2) > 10)){
			s2 = k +  "........................Rp "
			fmt.Printf("%s%d\n",s2,v)
		}
		if ((len(s2) <= 30) && (len(s2) > 20)){
			s2 = k +  ".....................Rp "
			fmt.Printf("%s%d\n",s2,v)
		}
		if (len(s2) > 30){
		
				fmt.Printf("%s..................Rp %d \n", k, v)	
		}
	}

	fmt.Printf("\n\nTotal.................Rp %d ", total)
	time.Sleep(60 * time.Second)
}