package model

// CreateChat используется как представление для создания чата в сервисе
type CreateChat struct {
	Title string
}

// Message используется как представление сообщения
type Message struct {
	ChatID    int64  `db:"chat_id"`
	Text      string `db:"text"`
	MemberTag string `db:"user_tag"`
}
