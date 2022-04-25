package models

type Report struct {
	Id              int          `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Report          string       `json:"report" gorm:"<-"`
	StudentGradeId  int          `json:"studentGradeId" gorm:"<-"`
	StudentGradeObj StudentGrade `json:"studentGradeObj" gorm:"foreignKey:StudentGradeId"`
}
