package comments

import (
	"github.com/Omelman/task-management/internal/models"
	"github.com/Omelman/task-management/internal/repo"
)

func UpdateComment(newComment models.CommentRequest) (models.Comment, error) {
	comment, err := repo.Get().Comments().Create(models.Comment{
		ID:          newComment.Data.ID,
		CommentText: newComment.Data.CommentText,
		TaskID:      newComment.Data.TaskID,
	})
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
