package main 
import(
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"math/rand"
)

type Employee struct{
	 //here json is used in order for encoding so that it appears as specified in json output
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	College *College `json:"college"`
}
type College struct{
	Collegename string `json:"collegename"`
	Address string `json:"address"`
}
var employees[] Employee


func getEmployees(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(employees)

}
func getEmployee(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	params:=mux.Vars(r)
	for _,item:=range employees{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func deleteEmployee(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","application/json")
		params:=mux.Vars(r)
		for index,item:=range employees{
			if item.ID==params["id"]{
				employees=append(employees[:index],employees[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(employees)
}

func createEmployee(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var employee Employee
	_=json.NewDecoder(r.Body).Decode(&employee)
	employee.ID=strconv.Itoa(rand.Intn(1000000))
	employees=append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

func main(){
	r:=mux.NewRouter();
	employees=append(employees,Employee{ID:"1",Name: "vishal",Email: "kamathvishal26@gmail.com",College: &College{Collegename: "joseph",Address: "vamanjoor"}})
	employees=append(employees,Employee{ID:"2",Name: "rahul",Email: "rahulgandhi25@gmail.com",College: &College{Collegename: "alosius",Address: "moodubidre"}})
	r.HandleFunc("/employees",getEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}",getEmployee).Methods("GET")
	r.HandleFunc("/employees",createEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}",deleteEmployee).Methods("DELETE")
	fmt.Printf("server running successfully on port 8002\n")
	log.Fatal(http.ListenAndServe(":8002",r))	
}



