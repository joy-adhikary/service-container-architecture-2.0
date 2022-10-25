package Services

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
	Boss        *Boss   `json:"boss"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
