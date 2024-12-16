package static

import (
	"context"
	"time"

	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/samber/lo"
)

func (db *static) GetExperiences(ctx context.Context) ([]entity.Experience, error) {
	return []entity.Experience{
		{
			ID:           "livestorm",
			Company:      "Livestorm",
			Position:     "Senior Software Engineer",
			From:         time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
			Link:         "https://livestorm.co",
			Technologies: getTechnologies("Go", "Ruby", "Bazel", "Rest", "gRPC", "Kafka", "Kubernetes", "Docker", "Datadog", "AWS", "PostgreSQL"),
			Tasks: []string{
				`Migrating Ruby on Rails services to Go so as to improve the performance, reliability, 
        development velocity and maintainability of the platform.`,
				`Migrating internal REST endpoints to gRPC so as to improve the performance and 
        reliability of intraservice communication.`,
			},
		},
		{
			ID:       "atellio",
			Company:  "Atellio",
			Position: "Senior Software Engineer",
			From:     time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			To:       lo.ToPtr(time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)),
			Link:     "https://livestorm.co",
			Technologies: getTechnologies("Go", "Rest", "gRPC", "Swagger", "PostgreSQL", "InfluxDB", "Meilisearch", "Traefik", "Docker Swarm", "Datadog",
				"CircleCI", "S3"),
			Tasks: []string{
				`Worked on and optimised the *activityFeed*, *notifications*, *templates* and 
        *searchIndexer* engines to perform more reliably under heavy load.`,
				`Developed the *metrics* engine using [Mixpanel](https://mixpanel.com/) for product analytics.`,
				`Devised and implemented a testing strategy for the various engines (i.e *activityFeed*, *notifications*
        , *templates*, *searchIndexer* and *metrics*) to ensure their resiliency.`,
				`Worked on and optimised the internal event system to scale from less than a hundred to 
        thousands of clients on a single instance of the platform.`,
				`Devised and implemented a service-wide type-safe validation framework to ensure data validity and consistency.`,
			},
		},
		{
			ID:       "ringier-sa",
			Company:  "Ringier SA",
			Position: "Senior Full-Stack Developer",
			From:     time.Date(2021, 02, 01, 0, 0, 0, 0, time.UTC),
			To:       lo.ToPtr(time.Date(2022, 06, 01, 0, 0, 0, 0, time.UTC)),
			Link:     "https://ringier.com/about-us/south-africa",
			Technologies: getTechnologies("Go", "Rest", "Swagger", "VueJS", "TailwindCSS", "DynamoDB", "Timestream", "Docker", "Redis", "RabbitMQ", "Datadog",
				"Blockchain", "Temporal", "S3", "AWS", "Pulumi"),
			Tasks: []string{
				`Worked on the core Go microservices powering the [Ringier Event Bus](https://docs.bus.ritdu.net/) serving millions of events monthly 
        across more than a dozen of companies and teams.`,
				`Developed an Infrastructure as Code Go microservice codenamed *Atlas* to fully automate provisioning of the Ringier Event
        Bus on AWS with automated integration testing.</br>Infrastructure cost reduced by 35% and new environment provisioning time reduced by 30%.`,
				`Developed a Go CLI application codenamed *Timestream Travel* as an AWS Timestream backup solution for [webvitalize.io](https://webvitalize.io/).`,
				`Developed and showcased a payment Go microservice using Web3 technologies during a play-week.`,
				`Developed and showcased a MVP Event Bus using Temporal at its core during a play-week.`,
				`Organised company-wide Go workshops to onboard and train new Gophers.`,
			},
		},
	}, nil
}
