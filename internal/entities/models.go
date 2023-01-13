package entities


type DataUser struct{
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Email string    `json:"Email"`
}

type Page struct{
	Title string 
	Body []byte
}

type Notes struct{
	Note string `json:"Note"`
	ID string	`json:"ID"`
}


type Search struct { 
	
}


