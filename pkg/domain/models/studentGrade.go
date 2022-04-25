package models

type StudentGrade struct {
	Id             int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId         string `json:"userId" gorm:"index:idx_submission,unique"`
	ExamId         string `json:"examId" gorm:"index:idx_submission,unique"`
	CourseId       string `json:"courseId" gorm:"<-"`
	Grade          int    `json:"grade" gorm:"<-"`
	CheatingStatus string `json:"cheatingStatus" gorm:"<-"`
}

type StudentGradeRepository interface {
	Add()
}
