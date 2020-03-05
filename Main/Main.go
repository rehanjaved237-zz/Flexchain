package main

import (
  "fmt"
  "flag"
  "log"
  "net"
  "net/http"
  "html/template"
  a1 "../Account"
  cour1 "../Course"
  n1 "../Network"
  p1 "../PersInfo"
  b1 "../Block"
  c1 "../Blockchain"
  w1 "../Web"
  buff1 "../BlockBuffer"
)

// IMPORTANT PROBLEMS TO SOLVE

// Correct GenerateBlockHash function in ../Block
// Remove blk.PrevHash = ""
// Blockchain printing issue


var ownAddr = flag.String("addr", ":10000", "Own Address")
var defaultPeer = flag.String("dpeer", ":10000", "Default Peer")
var frontEnd = flag.String("fend", ":12000", "Front End")
var defaultDatabase = flag.String("db", "localhost:8000", "Database is running on this IP")

var (
	templates = template.Must(
    template.ParseFiles(
      "index.html",
      "myHtml.html",
      "Home.html",
      "Login.html",
     "Teacher.html",
     "Student.html",
     "HOD.html",
     "HODAddStudent.html",
     "HODAddInstructor.html",
     "HODOfferCourse.html",
     ))
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "Login.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    OwnAddr string
    LoginAddr string
  }{
    "http://localhost" + w1.FrontEnd + "/Home/",
    "http://localhost" + w1.FrontEnd + "/Login/",
  }

  err := templates.ExecuteTemplate(w, "Home.html", &data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func myHtmlHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "myHtml.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HODHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    Port string
    OwnAddr string
    LoginAddr string
  }{
    w1.FrontEnd,
    "http://localhost" + w1.FrontEnd + "/Home/",
    "http://localhost" + w1.FrontEnd + "/Login/",
  }
	err := templates.ExecuteTemplate(w, "HOD.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "Teacher.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "Student.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HodAddStudentHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    Port string
  }{
    w1.FrontEnd,
  }
  err := templates.ExecuteTemplate(w, "HODAddStudent.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HodAddInstructorHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    Port string
  }{
    w1.FrontEnd,
  }
  err := templates.ExecuteTemplate(w, "HODAddInstructor.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HodOfferCourseHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    Port string
  }{
    w1.FrontEnd,
  }
  err := templates.ExecuteTemplate(w, "HODOfferCourse.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HodOfferCourse1Handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
    r.ParseForm()

		courseName := r.FormValue("CourseName")
    courseDescription := r.FormValue("CourseDescription")
		creditHrs := r.FormValue("CreditHrs")
		semester := r.FormValue("Semester")
    assignedTeacher := r.FormValue("AssignedTeacher")

    fmt.Println("Course Name:", courseName)
    fmt.Println("Course Description:", courseDescription)
    fmt.Println("Credit Hours:", creditHrs)
    fmt.Println("Semester:", semester)
    fmt.Println("Assigned Teacher:", assignedTeacher)

    newCourse := cour1.Course {
      CourseName: courseName,
      CourseDescription: courseDescription,
      CreditHours: creditHrs,
      Semester: semester,
      AssignedTeacher: assignedTeacher,
    }

    blk := b1.GenerateBlock("Course", newCourse)

    n1.BroadCastBlock(blk)
	}
}

func loginRequestHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Yahan tak chala hai")
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
		r.ParseForm()
		role := r.FormValue("role")
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		fmt.Println("role:", role)

		conn, err := net.Dial("tcp", n1.DefaultDatabase)
		if err != nil {
			log.Println("Error while connecting to the Default Database")
			log.Println("Error:", err)
		}

		login := a1.Account{Role: role, Username: username, Password: password}

		byteLogin := n1.GobEncode(login)
		data := append(n1.CmdToBytes("login"), byteLogin...)

		n, err := conn.Write(data)
		if err != nil {
			log.Println("Error while writing on network in login reqest handler")
		}
		fmt.Println("No of lines written on default database:", n)

		input := make([]byte, 8192)
		n, err = conn.Read(input)
		if err != nil {
			log.Println("Error while reading from network in login request handler")
		}
		fmt.Println("No of bytes read from default database:", n)
		role = ""
		role = n1.BytesToCmd(input)
		fmt.Println("Role:", role)
		if role == "student" {
			fmt.Println("student1")
			http.Redirect(w, r, "/student", http.StatusSeeOther)
		} else if role == "teacher" {
			fmt.Println("teacher1")
			http.Redirect(w, r, "/teacher", http.StatusSeeOther)
		} else if role == "hod" {
			fmt.Println("hod1")
			http.Redirect(w, r, "/hod", http.StatusSeeOther)
		} else {
			fmt.Println("Record Not Found")
			http.Redirect(w, r, "/Login", http.StatusSeeOther)
		}
	}
}

func runHandlers()  {

  http.HandleFunc("/", DefaultHandler)
  http.HandleFunc("/Home/", HomeHandler)
  http.HandleFunc("/Login/", LoginHandler)
  http.HandleFunc("/loginRequest/", loginRequestHandler)
  http.HandleFunc("/myHtml/", myHtmlHandler)
  http.HandleFunc("/teacher/", TeacherHandler)
  http.HandleFunc("/student/", StudentHandler)
  http.HandleFunc("/hod/", HODHandler)
  http.HandleFunc("/hodaddstudent/", HodAddStudentHandler)
  http.HandleFunc("/hodaddinstructor/", HodAddInstructorHandler)
  http.HandleFunc("/hodoffercourse/", HodOfferCourseHandler)
  http.HandleFunc("/hodoffercourse1/", HodOfferCourse1Handler)

  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  log.Fatal(http.ListenAndServe(w1.FrontEnd , nil))
}

func main() {
  flag.Parse()

  n1.OwnAddress = *ownAddr
  n1.DefaultPeer = *defaultPeer
  w1.FrontEnd = *frontEnd
  n1.DefaultDatabase = *defaultDatabase

  fmt.Println("Own Address:", n1.OwnAddress)
  fmt.Println("Default Peer:", n1.DefaultPeer)
  fmt.Println("Front End:", w1.FrontEnd)
  fmt.Println("Default Database:", n1.DefaultDatabase)

  go runHandlers()
  n1.StartServer()

  input := 0
  for {
    fmt.Println("Choose one of the following scenario:")
    fmt.Println("Enter 1 to view add a new node")
    fmt.Println("Enter 2 to view all known nodes")
    fmt.Println("Enter 3 to add a new PersInfo Block")
    fmt.Println("Enter 4 to view PersChain")
    fmt.Println("Enter 5 to view the BlockBuffer")
    fmt.Scanln(&input)

    switch input {
    case 1:
      var ip string
      fmt.Println("Enter the new ip")
      fmt.Scanln(&ip)
      n1.AddToKnownNode(ip)
    case 2:
      n1.PrintKnownNodes()
    case 3:
      var input p1.PersInfo
      input.PersInfoInput()
      blk := b1.GenerateBlock("PersInfo", input)
//      blk.PrintBlock()
      n1.BroadCastBlock(blk)
    case 4:
      c1.PrintBlockchain(c1.PersInfoChain)
    case 5:
      buff1.BlkBuffer.PrintBlockBuffer()
    default:
      fmt.Println("Invalid Command Entered")
    }
  }

}

/*type Fun1 struct {
  A int
}

type Fun2 struct {
  A string
  B float64
}

func main() {
  blk1 := b1.GenerateBlock("Rehan", Fun1{A: 5})
  blk1.PrintBlock()
  blk2 := b1.GenerateBlock("Rehan", Fun1{A: 6})
  blk2.PrintBlock()
  blk3 := b1.GenerateBlock("Rehan", Fun2{A: "Alhamdulillah", B: 4.7})
  blk3.PrintBlock()
}*/
