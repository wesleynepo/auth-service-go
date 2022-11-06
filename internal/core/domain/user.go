package domain

type User struct {
    ID       uint `json:"id"`
    Email    string `json:"email"`
    Hash     string `json:"hash"`
    Salt     string `json:"salt"`
}
