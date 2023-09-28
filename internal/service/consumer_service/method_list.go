package consumer_service

import (
	"context"
)

func (s *ConsumerService) MethodList(ctx context.Context) []string {
	return Methods
}
