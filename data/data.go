package data

// TODO: Skill struct sec
type Skill struct {
	Name string
}

// TODO: exp struct sec
type Experience struct {
	Title       string
	Company     string
	Description string
}

// TODO: educations struct sec
type Education struct {
	School, Degree, Major string
	Year                  int
}

// TODO: Skill sec
var Skills [100]Skill
var SkillCount int = 0

// TODO: Exp sec
var Experiences [100]Experience
var ExperienceCount int = 0

// TODO: Educations sec
var Educations [100]Education
var EducationCount int = 0
