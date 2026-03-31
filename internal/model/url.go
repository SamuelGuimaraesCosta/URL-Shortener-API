package model

import "time"

type URL struct {
	Code        string
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	Hits        int
}

func (u *URL) IsExpired() bool {

	if u.ExpiresAt == nil {
		return false
	}

	return time.Now().After(*u.ExpiresAt)
}
