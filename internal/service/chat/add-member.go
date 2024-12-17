package chat

import "context"

// AddMembersServ является сервисной прослойкой для добавления пользователя
func (s *serv) AddMembersServ(ctx context.Context, chatID int64, memberTags []string) error {
	err := s.chatServerRepository.AddMembers(ctx, chatID, memberTags)
	if err != nil {
		return err
	}
	return nil
}
