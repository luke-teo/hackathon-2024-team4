package console

import (
	"first_move/internal/app/task"
)

func (c *Console) ParseTextChatFromCsv(path string) error {
	return task.DispatchParseTextChatTask(c.app, path)
}
