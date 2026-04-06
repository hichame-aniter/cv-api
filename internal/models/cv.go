package models

type CV struct {
	Basics       Basics        `json:"basics"`
	Work         []Experience  `json:"work"`
	Volunteer    []Volunteer   `json:"volunteer"`
	Education    []Education   `json:"education"`
	Awards       []Award       `json:"awards"`
	Certificates []Certificate `json:"certificates"`
	Publications []Publication `json:"publications"`
	Skills       []Skill       `json:"skills"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	References   []Reference   `json:"references"`
	Projects     []Project     `json:"projects"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Image    *string   `json:"image"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Url      string    `json:"url"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}
type Location struct {
	Address     string `json:"address"`
	Postalcode  string `json:"postalCode"`
	City        string `json:"city"`
	Countrycode string `json:"countryCode"`
	Region      string `json:"region"`
}
type Profile struct {
	ID       string `json:"id"`
	Network  string `json:"network"`
	Username string `json:"username"`
	Url      string `json:"url"`
}
type Experience struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	Url        string   `json:"url"`
	Startdate  string   `json:"startDate"`
	Enddate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}
type Volunteer struct {
	ID           string   `json:"id"`
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Url          string   `json:"url"`
	Startdate    string   `json:"startDate"`
	Enddate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlight    []string `json:"highlight"`
}
type Education struct {
	ID          string   `json:"id"`
	Institution string   `json:"institution"`
	Url         string   `json:"url"`
	Area        string   `json:"area"`
	Studytype   string   `json:"studyType"`
	Startdate   string   `json:"startDate"`
	Enddate     string   `json:"endDate"`
	Score       *string  `json:"score"`
	Courses     []string `json:"courses"`
}
type Award struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}
type Certificate struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Issuer string `json:"issuer"`
	Url    string `json:"url"`
}
type Publication struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	Releasedate string `json:"releaseDate"`
	Url         string `json:"url"`
	Summary     string `json:"summary"`
}
type Skill struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Level    *string  `json:"level"`
	Keywords []string `json:"keywords"`
}
type Language struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}
type Interest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Reference struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Reference string `json:"reference"`
}
type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Startdate   string   `json:"startDate"`
	Enddate     string   `json:"endDate"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	Url         string   `json:"url"`
}
