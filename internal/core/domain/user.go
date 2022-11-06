package domain

type User struct {
    Id       float64 `json:"id"`
    Email    string `json:"email"`
    Hash     string `json:"hash"`
    Salt     string `json:"salt"`
}
