package fetcher

import (
	"context"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/mgjules/mgjules-go/entity"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/repository"
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
	"go.uber.org/multierr"
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
	for _, s := range f.subscribers {
		if s == nil {
			continue
		}

		f.pool.Submit(s)
	}
}

func (f *Fetcher) Fetch() {
	operation := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var (
			wg   sync.WaitGroup
			errs error
		)
		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching meta data...")
			meta, err := f.repo.GetMeta(ctx, "bd99e066-440b-11ed-924c-9fd15527df84")
			if err != nil {
				logger.L.Errorf("failed to get meta: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.meta = meta
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching links data...")
			links, err := f.repo.GetLinks(ctx)
			if err != nil {
				logger.L.Errorf("failed to get links: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.links = links
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching introduction data...")
			intro, err := f.repo.GetIntroduction(ctx, "a4296eac-441b-11ed-924c-830c8fd1144c")
			if err != nil {
				logger.L.Errorf("failed to get introduction: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.intro = intro
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching section data...")
			sections, err := f.repo.GetSections(ctx)
			if err != nil {
				logger.L.Errorf("failed to get sections: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.sections = sections
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching experience data...")
			experiences, err := f.repo.GetExperiences(ctx)
			if err != nil {
				logger.L.Errorf("failed to get experiences: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.experiences = experiences
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching project data...")
			projects, err := f.repo.GetProjects(ctx)
			if err != nil {
				logger.L.Errorf("failed to get projects: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.projects = projects
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching contribution data...")
			contributions, err := f.repo.GetContributions(ctx)
			if err != nil {
				logger.L.Errorf("failed to get contributions: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.contributions = contributions
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching award data...")
			awards, err := f.repo.GetAwards(ctx)
			if err != nil {
				logger.L.Errorf("failed to get awards: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.awards = awards
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching interest data...")
			interests, err := f.repo.GetInterests(ctx)
			if err != nil {
				logger.L.Errorf("failed to get interests: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.interests = interests
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching language data...")
			languages, err := f.repo.GetLanguages(ctx)
			if err != nil {
				logger.L.Errorf("failed to get languages: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.languages = languages
				f.dataMu.Unlock()
			}
		})

		wg.Add(1)
		f.pool.Submit(func() {
			defer wg.Done()

			logger.L.Debug("Fetching post data...")
			posts, err := f.repo.GetPosts(ctx)
			if err != nil {
				logger.L.Errorf("failed to get posts: %v", err)
				errs = multierr.Append(errs, err)
			} else {
				f.dataMu.Lock()
				f.posts = posts
				f.dataMu.Unlock()
			}
		})

		wg.Wait()

		return errs
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
		logger.L.Errorf("failed to fetched data: %v", err)
		return
	}

	f.dataMu.Lock()
	f.fetchedAt = time.Now()
	f.dataMu.Unlock()

	f.notify()
}

func (f *Fetcher) Meta() entity.Meta {
	f.dataMu.RLock()
	meta := DeepCopy(f.meta)
	f.dataMu.RUnlock()

	return meta
}

func (f *Fetcher) Links() []entity.Link {
	f.dataMu.RLock()
	links := DeepCopy(&f.links)
	f.dataMu.RUnlock()

	return links
}

func (f *Fetcher) Intro() entity.Introduction {
	f.dataMu.RLock()
	intro := DeepCopy(f.intro)
	f.dataMu.RUnlock()

	return intro
}

func (f *Fetcher) Sections() []entity.Section {
	f.dataMu.RLock()
	sections := DeepCopy(&f.sections)
	f.dataMu.RUnlock()

	return sections
}

func (f *Fetcher) Experiences() []entity.Experience {
	f.dataMu.RLock()
	experiences := DeepCopy(&f.experiences)
	f.dataMu.RUnlock()

	return experiences
}

func (f *Fetcher) Projects() []entity.Project {
	f.dataMu.RLock()
	projects := DeepCopy(&f.projects)
	f.dataMu.RUnlock()

	return projects
}

func (f *Fetcher) Contributions() []entity.Contribution {
	f.dataMu.RLock()
	contributions := DeepCopy(&f.contributions)
	f.dataMu.RUnlock()

	return contributions
}

func (f *Fetcher) Awards() []entity.Award {
	f.dataMu.RLock()
	awards := DeepCopy(&f.awards)
	f.dataMu.RUnlock()

	return awards
}

func (f *Fetcher) Interests() []entity.Interest {
	f.dataMu.RLock()
	interests := DeepCopy(&f.interests)
	f.dataMu.RUnlock()

	return interests
}

func (f *Fetcher) Languages() []entity.Language {
	f.dataMu.RLock()
	languages := DeepCopy(&f.languages)
	f.dataMu.RUnlock()

	return languages
}

func (f *Fetcher) Posts() []entity.Post {
	f.dataMu.RLock()
	posts := DeepCopy(&f.posts)
	f.dataMu.RUnlock()

	return posts
}

func (f *Fetcher) FetchedAt() time.Time {
	f.dataMu.RLock()
	fetchedAt := f.fetchedAt
	f.dataMu.RUnlock()

	return fetchedAt
}
