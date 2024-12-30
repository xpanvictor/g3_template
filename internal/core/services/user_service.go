package services

import (
	"fmt"
	"go_graphql_gin/internal/core/domain"
	"go_graphql_gin/internal/core/ports"
	"log"
	"time"
)

type userService struct {
	userRepo        ports.UserRepository
	profileRepo     ports.ProfileRepository
	notificationSvc ports.NotificationService
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	// Start transaction
	err := s.userRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Create user
		if err := s.userRepo.Create(txCtx, user); err != nil {
			return err
		}

		// Create associated profile
		profile := &domain.Profile{
			UserID:      user.ID,
			DisplayName: user.Username,
			CreatedAt:   time.Now(),
		}
		if err := s.profileRepo.Create(txCtx, profile); err != nil {
			return err
		}

		// Update user with profile ID
		user.ProfileID = profile.ID
		if err := s.userRepo.Update(txCtx, user); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Notify after successful transaction
	go func() {
		if err := s.notificationSvc.NotifyUserCreation(context.Background(), user.ID); err != nil {
			log.Printf("failed to send notification: %v", err)
		}
	}()

	return nil
}
