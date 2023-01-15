package entities


type DataUser struct{
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Email string    `json:"Email"`
	ID string 		`json:"ID"`
}

type Notes struct{
	Note string `json:"Note"`
	ID string	`json:"ID"`
}

type Page struct{
	Title string  `json:"Title"`
	Body []byte `json:"Body"`
}

type Search struct { 
	ID string `json:"ID"`
}

type DataBase struct{
	Host string 
	Port string 
	UserName string 
	Password string 
	DBName string 
	SSLMode string 
}


