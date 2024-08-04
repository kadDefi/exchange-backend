package service

import (
	"context"
	"encoding/json"
	"exchange-backend/internal/domain"
	"exchange-backend/internal/pkg/log"
	"fmt"
	"os"
	"syscall"
)

func (s *Service) ProcessTaskRefreshTokenURL(ctx context.Context, arg string) error {
	var t domain.TaskRefreshTokenURL
	if err := json.Unmarshal([]byte(arg), &t); err != nil {
		return err
	}

	log.FromContext(ctx).Sugar().Infof("Processing task refresh token,CollectionAddress: %s", t.CollectionAddress)
	log.FromContext(ctx).Sugar().Infof("Processing task refresh token,CollectionAddress: %s", t.TokenID)
	log.FromContext(ctx).Sugar().Info("==========`Processing task refresh=========")
	panic(fmt.Sprintf("Task refresh"))

	pid := os.Getpid()
	if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
		fmt.Println("Error sending signal:", err)
		return err
	}

	log.FromContext(ctx).Sugar().Info("==========`Processing task refresh end=========")

	return nil
}
