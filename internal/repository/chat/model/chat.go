package model

// Message используется как представление сообщения
type Message struct {
	ChatID    int64  `db:"chat_id"`
	Text      string `db:"text"`
	MemberTag string `db:"user_tag"`
}

// CreateChat используется как представление для создания чата
type CreateChat struct {
	Title string `db:"title"`
}

//// AddUsersToChat используется как представление для добавления пользователя в чат
//type AddUsersToChat struct {
//	ChatID    int64    `db:"chat_id"`
//	UsersTags []string `db:"user_tag"`
//}
//
//// SendMessage используется как представление для отправки сообщения
//type SendMessage struct {
//	UserTag string `db:"user_tag"`
//	ChatID  string `db:"chat_id"`
//	Text    string `db:"text"`
//}
