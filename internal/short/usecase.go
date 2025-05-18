package short

import (
	"context"
	"net/url"

	"github.com/dalmow/sdalm/internal/config"
	"github.com/dalmow/sdalm/pkg/utils"
)

type ShortenUseCase interface {
	ShortenURL(ctx context.Context, originalUrl string) (string, error)
	Resolve(ctx context.Context, identifier string) (string, error)
}

type shortenUseCase struct {
	repo ShortsRepository
	conf *config.Config
}

func NewShortenUseCase(repo ShortsRepository, c *config.Config) ShortenUseCase {
	return &shortenUseCase{
		repo: repo,
		conf: c,
	}
}

func (u *shortenUseCase) ShortenURL(ctx context.Context, originalUrl string) (string, error) {
	generated, err := utils.RandomAliasGen(ctx, u.repo)
	if err != nil {
		return "", err
	}

	u.repo.Create(ctx, &Short{
		ID:          generated,
		OriginalURL: originalUrl,
	})

	finalUrl, err := buildFinalUrl(u.conf.BasePath, generated)
	return finalUrl, err
}

func (u *shortenUseCase) Resolve(ctx context.Context, identifier string) (string, error) {
	short, err := u.repo.Find(ctx, identifier)
	if err != nil {
		return "", err
	}

	return short.OriginalURL, nil
}

func buildFinalUrl(baseUrl, alias string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	u.Path = alias
	return u.String(), nil
}
