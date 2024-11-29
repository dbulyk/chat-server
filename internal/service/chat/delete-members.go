package chat

import "context"

// RemoveMembersServ является сервисной прослойкой для удаления пользователей из чата
func (s *serv) RemoveMembersServ(ctx context.Context, chatID int64, memberTags []string) error {
	err := s.chatServerRepository.RemoveMembers(ctx, chatID, memberTags)
	if err != nil {
		return err
	}
	return nil
}
