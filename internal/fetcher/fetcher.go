package fetcher

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/mgjules/mgjules-go/internal/repository"
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
)

const maxFreshFetchAttempts = 2

type Fetcher struct {
	repo repository.Repository
	pool *ants.Pool
	cron *cron.Cron

	dataMu        sync.RWMutex // guards the data
	meta          *entity.Meta
	links         []entity.Link
	intro         *entity.Introduction
	sections      []entity.Section
	experiences   []entity.Experience
	projects      []entity.Project
	contributions []entity.Contribution
	awards        []entity.Award
	interests     []entity.Interest
	languages     []entity.Language
	posts         []entity.Post
	fetchedAt     time.Time

	subscribersMu sync.RWMutex // guards the subscribers
	subscribers   []func()
}

func New(repo repository.Repository, pool *ants.Pool) *Fetcher {
	if repo == nil {
		panic("repo cannot be nil")
	}
	if pool == nil {
		panic("pool cannot be nil")
	}

	f := &Fetcher{
		repo: repo,
		pool: pool,
		cron: cron.New(),
	}

	f.cron.AddFunc("@hourly", func() {
		f.Fetch()
	})

	return f
}

func (f *Fetcher) Start() {
	f.cron.Start()
}

func (f *Fetcher) Stop() {
	f.cron.Stop()
}

func (f *Fetcher) AddSubscriber(s func()) {
	f.subscribersMu.Lock()
	f.subscribers = append(f.subscribers, s)
	f.subscribersMu.Unlock()
}

func (f *Fetcher) notify() {
	f.subscribersMu.RLock()
	defer f.subscribersMu.RUnlock()
	for _, s := range f.subscribers {
		if s == nil {
			continue
		}
		f.pool.Submit(s)
	}
}

func fetchEntity[T any](
	f *Fetcher,
	ctx context.Context,
	wg *sync.WaitGroup,
	name string,
	fetch func(context.Context) (T, error),
	store func(T),
) {
	wg.Add(1)
	f.pool.Submit(func() {
		defer wg.Done()
		slog.Debug("Fetching data...", "entity", name)
		result, err := fetch(ctx)
		if err != nil {
			slog.Error("failed to fetch", "entity", name, "error", err)
			return
		}
		f.dataMu.Lock()
		store(result)
		f.dataMu.Unlock()
	})
}

func (f *Fetcher) Fetch() {
	operation := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		var wg sync.WaitGroup

		fetchEntity(f, ctx, &wg, "meta",
			func(ctx context.Context) (*entity.Meta, error) {
				return f.repo.GetMeta(ctx, "bd99e066-440b-11ed-924c-9fd15527df84")
			},
			func(v *entity.Meta) { f.meta = v })
		fetchEntity(f, ctx, &wg, "links",
			func(ctx context.Context) ([]entity.Link, error) { return f.repo.GetLinks(ctx) },
			func(v []entity.Link) { f.links = v })
		fetchEntity(f, ctx, &wg, "intro",
			func(ctx context.Context) (*entity.Introduction, error) {
				return f.repo.GetIntroduction(ctx, "a4296eac-441b-11ed-924c-830c8fd1144c")
			},
			func(v *entity.Introduction) { f.intro = v })
		fetchEntity(f, ctx, &wg, "sections",
			func(ctx context.Context) ([]entity.Section, error) { return f.repo.GetSections(ctx) },
			func(v []entity.Section) { f.sections = v })
		fetchEntity(f, ctx, &wg, "experiences",
			func(ctx context.Context) ([]entity.Experience, error) { return f.repo.GetExperiences(ctx) },
			func(v []entity.Experience) { f.experiences = v })
		fetchEntity(f, ctx, &wg, "projects",
			func(ctx context.Context) ([]entity.Project, error) { return f.repo.GetProjects(ctx) },
			func(v []entity.Project) { f.projects = v })
		fetchEntity(f, ctx, &wg, "contributions",
			func(ctx context.Context) ([]entity.Contribution, error) { return f.repo.GetContributions(ctx) },
			func(v []entity.Contribution) { f.contributions = v })
		fetchEntity(f, ctx, &wg, "awards",
			func(ctx context.Context) ([]entity.Award, error) { return f.repo.GetAwards(ctx) },
			func(v []entity.Award) { f.awards = v })
		fetchEntity(f, ctx, &wg, "interests",
			func(ctx context.Context) ([]entity.Interest, error) { return f.repo.GetInterests(ctx) },
			func(v []entity.Interest) { f.interests = v })
		fetchEntity(f, ctx, &wg, "languages",
			func(ctx context.Context) ([]entity.Language, error) { return f.repo.GetLanguages(ctx) },
			func(v []entity.Language) { f.languages = v })
		fetchEntity(f, ctx, &wg, "posts",
			func(ctx context.Context) ([]entity.Post, error) { return f.repo.GetPosts(ctx) },
			func(v []entity.Post) { f.posts = v })

		wg.Wait()
		return nil
	}

	exp := backoff.NewExponentialBackOff()
	exp.InitialInterval = 5 * time.Second
	exp.Multiplier = 2.0
	exp.MaxInterval = 5 * time.Minute
	exp.MaxElapsedTime = 15 * time.Minute

	var bkf backoff.BackOff = exp
	if !f.fetchedAt.IsZero() {
		bkf = backoff.WithMaxRetries(bkf, maxFreshFetchAttempts)
	}

	err := backoff.Retry(operation, bkf)
	if err != nil {
		slog.Error("failed to fetched data", "error", err)
		return
	}

	f.dataMu.Lock()
	f.fetchedAt = time.Now()
	f.dataMu.Unlock()

	f.notify()
}

func (f *Fetcher) Meta() entity.Meta {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	if f.meta == nil {
		return entity.Meta{}
	}
	return *f.meta
}

func (f *Fetcher) Links() []entity.Link {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.links
}

func (f *Fetcher) Introduction() entity.Introduction {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	if f.intro == nil {
		return entity.Introduction{}
	}
	return *f.intro
}

func (f *Fetcher) Sections() []entity.Section {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.sections
}

func (f *Fetcher) Experiences() []entity.Experience {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.experiences
}

func (f *Fetcher) Projects() []entity.Project {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.projects
}

func (f *Fetcher) Contributions() []entity.Contribution {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.contributions
}

func (f *Fetcher) Awards() []entity.Award {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.awards
}

func (f *Fetcher) Interests() []entity.Interest {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.interests
}

func (f *Fetcher) Languages() []entity.Language {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.languages
}

func (f *Fetcher) Posts() []entity.Post {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.posts
}

func (f *Fetcher) FetchedAt() time.Time {
	f.dataMu.RLock()
	defer f.dataMu.RUnlock()
	return f.fetchedAt
}
