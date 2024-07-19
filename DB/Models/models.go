package Models

type User struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Deleted        bool   `json:"deleted"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	PhoneNumber    string `json:"phoneNumber"`
	ProfilePicture string `json:"profilePicture"`
}

type UsersFile []User
