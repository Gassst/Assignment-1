package model

type Person struct {
	ID                string `json:"id"`
	StudentCode       string `json:"student_code"`
	StudentName       string `json:"student_name"`
	StudentAddress    string `json:"student_address"`
	StudentOccupation string `json:"student_occupation"`
	JoiningReason     string `json:"joining_reason"`
}

type Participants struct {
	Participants []Person `json:"participants"`
}
