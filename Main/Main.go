package main

import (
  "fmt"
  "flag"
  "log"
//  "net"
  "net/http"
  "html/template"
//  a1 "../Account"
  std1 "../Student"
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
var fileName = flag.String("filename", "blockchain.json", "File where blockchain data is saved")
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
     "HODMineHistory.html",
     "BlockList.html",
     "TeacherCourses.html",
     "StudentEnrolledCourses.html",
     "StudentEnrollCourse.html",
     "TeacherGradeStudents.html",
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

func HodAddStudent1Handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
    r.ParseForm()

    photo := r.FormValue("Photo")
    rollNo := r.FormValue("RollNo")
    name := r.FormValue("Name")
		fatherName := r.FormValue("FatherName")
    cnic := r.FormValue("CNIC")
		phone := r.FormValue("Phone")
		department := r.FormValue("Department")
    email := r.FormValue("Email")

    fmt.Println("Photo:", photo)
    fmt.Println("Rollno:", rollNo)
    fmt.Println("Name:", name)
    fmt.Println("Father Name:", fatherName)
    fmt.Println("CNIC:", cnic)
    fmt.Println("Phone:", phone)
    fmt.Println("Department:", department)
    fmt.Println("Email:", email)

    newStudent := std1.Student {
      Photo: photo,
      RollNo: rollNo,
      Name: name,
      FatherName: fatherName,
      CNIC: cnic,
      Phone: phone,
      Department: department,
      Email: email,
    }

    blk := b1.GenerateBlock("Student", newStudent)

    n1.BroadCastBlock(blk)
	}
  http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func HodOfferCourse1Handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
		fmt.Fprintf(w, "Error aa gaya bhai.. yaar ye get function call kar raha hai.. html mein yakeenan koi masla hai")
		fmt.Fprintf(w, "Get function called")
	} else {
    r.ParseForm()

    courseCode := r.FormValue("CourseCode")
		courseName := r.FormValue("CourseName")
    courseDescription := r.FormValue("CourseDescription")
		creditHrs := r.FormValue("CreditHrs")
		semester := r.FormValue("Semester")
    assignedTeacher := r.FormValue("AssignedTeacher")

    fmt.Println("Course Code:", courseCode)
    fmt.Println("Course Name:", courseName)
    fmt.Println("Course Description:", courseDescription)
    fmt.Println("Credit Hours:", creditHrs)
    fmt.Println("Semester:", semester)
    fmt.Println("Assigned Teacher:", assignedTeacher)

    newCourse := cour1.Course {
      CourseCode: courseCode,
      CourseName: courseName,
      CourseDescription: courseDescription,
      CreditHours: creditHrs,
      Semester: semester,
      AssignedTeacher: assignedTeacher,
    }

    blk := b1.GenerateBlock("Course", newCourse)

    n1.BroadCastBlock(blk)
	}
  http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func HodMineHistoryHandler(w http.ResponseWriter, r *http.Request) {
  data := struct {
    Port string
    BBuffer buff1.BlockBuffer
  }{
    w1.FrontEnd,
    buff1.BlkBuffer,
  }
  err := templates.ExecuteTemplate(w, "HODMineHistory.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TeacherCoursesHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("BlockListHandler Run Successfully")
  data := struct {
    Port string
    Chain []b1.Block
  }{
    w1.FrontEnd,
    c1.Chain1.FilterBlockchain("Course"),
  }
  err := templates.ExecuteTemplate(w, "TeacherCourses.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func AddToBlockchainHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("AddToBlockchainHandler Run")
  hash := r.URL.Path[len("/addtoblockchain/"):]

  status, index := buff1.BlkBuffer.FindBlock(hash)
  if status == true {

    _, blk := buff1.BlkBuffer.RemoveBlock(index)
    n1.BroadCastRemoveBlockBuffer(hash)
    blk.Status = true
    n1.BroadCastBlock(blk)

  } else {
    fmt.Println("Block Not Found")
  }
  http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func RemoveFromBlockBufferHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("RemoveFromBlockBuffer Run")
  hash := r.URL.Path[len("/removefromblockbuffer/"):]

  status, index := buff1.BlkBuffer.FindBlock(hash)
  if status == true {
    fmt.Println("Block Found")
    _, _ = buff1.BlkBuffer.RemoveBlock(index)
    n1.BroadCastRemoveBlockBuffer(hash)
  } else {
    fmt.Println("Block Not Found")
  }
  http.Redirect(w, r, "/hod/", http.StatusSeeOther)
}

func BlocklistHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("BlockListHandler Run Successfully")
  data := struct {
    Port string
    Chain []b1.Block
  }{
    w1.FrontEnd,
    c1.Chain1.SliceBlockchain(),
  }
  err := templates.ExecuteTemplate(w, "BlockList.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func StudentEnrollCourseHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("StudentEnrollCourseHandler Successfully")
  data := struct {
    Port string
    Chain []b1.Block
  }{
    w1.FrontEnd,
    c1.Chain1.FilterBlockchain("Course"),
  }
  err := templates.ExecuteTemplate(w, "StudentEnrollCourse.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TeacherGradeStudentsHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("TeacherGradeStudentsHandler Executed")
  data := struct {
    Port string
    Chain []b1.Block
  }{
    w1.FrontEnd,
    c1.Chain1.FilterBlockchain("Student"),
  }
  err := templates.ExecuteTemplate(w, "TeacherGradeStudents.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func StudentEnrolledCoursesHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("StudentEnrolledCoursesHandler Successfully")
  data := struct {
    Port string
    Chain []b1.Block
  }{
    w1.FrontEnd,
    c1.Chain1.FilterBlockchain("Course"),
  }
  err := templates.ExecuteTemplate(w, "StudentEnrolledCourses.html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
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
  http.HandleFunc("/hodaddstudent1/", HodAddStudent1Handler)
  http.HandleFunc("/hodaddinstructor/", HodAddInstructorHandler)
  http.HandleFunc("/hodoffercourse/", HodOfferCourseHandler)
  http.HandleFunc("/hodoffercourse1/", HodOfferCourse1Handler)
  http.HandleFunc("/hodminehistory/", HodMineHistoryHandler)
  http.HandleFunc("/addtoblockchain/", AddToBlockchainHandler)
  http.HandleFunc("/removefromblockbuffer/", RemoveFromBlockBufferHandler)
  http.HandleFunc("/blocklist/", BlocklistHandler)
  http.HandleFunc("/teachercourses/", TeacherCoursesHandler)
  http.HandleFunc("/studentenrollcourse/", StudentEnrollCourseHandler)
  http.HandleFunc("/studentenrolledcourses/", StudentEnrolledCoursesHandler)
  http.HandleFunc("/teachergradestudents/", TeacherGradeStudentsHandler)

  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  log.Fatal(http.ListenAndServe(w1.FrontEnd , nil))
}

func main() {
  flag.Parse()

  n1.OwnAddress = *ownAddr
  n1.DefaultPeer = *defaultPeer
  w1.FrontEnd = *frontEnd
  n1.DefaultDatabase = *defaultDatabase
  c1.FileName = *fileName

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
      c1.PrintBlockchain(c1.Chain1)
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
