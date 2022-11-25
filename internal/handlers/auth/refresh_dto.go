package auth

type BodyRefresh struct {
    Refresh string `json:"refresh" binding:"required"`
}
