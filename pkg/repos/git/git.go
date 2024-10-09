package git

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sanjay920/gptscript/pkg/hash"
)

func exists(dir string) (bool, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func Checkout(ctx context.Context, base, repo, commit, toDir string) error {
	if found, err := exists(toDir); err != nil {
		return err
	} else if found {
		return fmt.Errorf("%s already exists, can not create repo", toDir)
	}

	if err := os.MkdirAll(filepath.Dir(toDir), 0755); err != nil {
		return err
	}

	if usePureGo() {
		return checkoutPureGo(ctx, base, repo, commit, toDir)
	}

	if err := fetch(ctx, base, repo, commit); err != nil {
		return err
	}

	log.InfofCtx(ctx, "Checking out %s to %s", commit, toDir)
	return gitWorktreeAdd(ctx, gitDir(base, repo), toDir, commit)
}

func gitDir(base, repo string) string {
	return filepath.Join(base, "repos", hash.Digest(repo))
}

func fetch(ctx context.Context, base, repo, commit string) error {
	gitDir := gitDir(base, repo)
	if found, err := exists(gitDir); err != nil {
		return err
	} else if !found {
		log.InfofCtx(ctx, "Cloning %s", repo)
		if err := cloneBare(ctx, repo, gitDir); err != nil {
			return err
		}
	}
	log.InfofCtx(ctx, "Fetching %s at %s", commit, repo)
	return fetchCommit(ctx, gitDir, commit)
}
