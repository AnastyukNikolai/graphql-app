package repositories

import (
	"context"
	"graphql-app/ent"
	"graphql-app/ent/reminder"
	"graphql-app/graph/model"
)

type ReminderRepositoryEnt struct {
	DBClient *ent.Client
}

func NewReminderRepositoryEnt(DBClient *ent.Client) *ReminderRepositoryEnt {
	return &ReminderRepositoryEnt{DBClient: DBClient}
}

func (r *ReminderRepositoryEnt) CreateReminder(reminder *model.NewReminder) (*model.Reminder, error) {
	reminderDB, err := r.DBClient.Reminder.
		Create().
		SetText(reminder.Text).
		SetTodoID(reminder.TodoID).
		Save(context.Background())
	if err != nil {
		return &model.Reminder{}, err
	}
	return &model.Reminder{
		ID:   reminderDB.ID,
		Text: reminder.Text,
	}, nil
}

func (r *ReminderRepositoryEnt) GetAllReminders() ([]*model.Reminder, error) {
	var reminders []*model.Reminder
	remindersDB, err := r.DBClient.Reminder.Query().All(context.Background())
	if err != nil {
		return reminders, err
	}
	for _, reminderDB := range remindersDB {
		var reminderModel *model.Reminder
		reminderModel = &model.Reminder{
			ID:   reminderDB.ID,
			Text: reminderDB.Text,
		}
		reminders = append(reminders, reminderModel)
	}
	return reminders, nil
}

func (r *ReminderRepositoryEnt) GetReminderById(reminderId int) (*model.Reminder, error) {
	var reminderModel *model.Reminder
	reminderDB, err := r.DBClient.Reminder.Query().Where(reminder.ID(reminderId)).Only(context.Background())
	if err != nil {
		return reminderModel, err
	}
	return &model.Reminder{
		ID:   reminderDB.ID,
		Text: reminderDB.Text,
	}, nil
}
