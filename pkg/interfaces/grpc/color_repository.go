package grpc

import (
	"context"
	service "github.com/ademsa/go-grpc-color-service/pkg/service"
	"go-issue-tracker/pkg/domain"
	"google.golang.org/grpc"
	"time"
)

// ColorRepository is a repository
type ColorRepository struct {
	Endpoint string
}

// NewColorRepository to create ColorRepository
func NewColorRepository(endpoint string) *ColorRepository {
	return &ColorRepository{
		Endpoint: endpoint,
	}
}

// GetColor to get color
func (r *ColorRepository) GetColor() (domain.Color, error) {
	var color domain.Color

	c, err := grpc.Dial(r.Endpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return color, err
	}
	defer c.Close()

	client := service.NewColorServiceClient(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetColor(ctx, &service.GetColorRequest{})
	if err != nil {
		return color, err
	}

	color.HexCode = response.Color

	return color, nil
}
