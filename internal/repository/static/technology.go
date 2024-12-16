package static

import (
	"strings"

	"github.com/mgjules/mgjules-go/internal/entity"
	"github.com/samber/lo"
)

var technologies = []entity.Technology{
	{Name: "Go", Link: "https://go.dev"},
	{Name: "Ruby", Link: "https://www.ruby-lang.org/"},
	{Name: "Bazel", Link: "https://bazel.build/"},
	{Name: "Rest", Link: "https://restfulapi.net/"},
	{Name: "gRPC", Link: "https://grpc.io/"},
	{Name: "Kafka", Link: "https://kafka.apache.org/"},
	{Name: "Kubernetes", Link: "https://kubernetes.io/"},
	{Name: "Docker", Link: "https://www.docker.com/"},
	{Name: "Datadog", Link: "https://www.datadoghq.com/"},
	{Name: "AWS", Link: "https://aws.amazon.com/"},
	{Name: "PostgreSQL", Link: "https://www.postgresql.org/"},
	{Name: "Swagger", Link: "https://swagger.io/"},
	{Name: "InfluxDB", Link: "https://www.influxdata.com/"},
	{Name: "Meilisearch", Link: "https://www.meilisearch.com/"},
	{Name: "Traefik", Link: "https://traefik.io/"},
	{Name: "Docker Swarm", Link: "https://docs.docker.com/engine/swarm/"},
	{Name: "CircleCI", Link: "https://circleci.com/"},
	{Name: "VueJS", Link: "https://vuejs.org/"},
	{Name: "TailwindCSS", Link: "https://tailwindcss.com/"},
	{Name: "DynamoDB", Link: "https://aws.amazon.com/dynamodb/"},
	{Name: "S3", Link: "https://aws.amazon.com/s3/"},
	{Name: "Timestream", Link: "https://aws.amazon.com/timestream/"},
	{Name: "Redis", Link: "https://redis.io/"},
	{Name: "RabbitMQ", Link: "https://www.rabbitmq.com/"},
	{Name: "Blockchain", Link: "https://www.datadoghq.com/"},
	{Name: "Temporal", Link: "https://temporal.io/"},
	{Name: "Pulumi", Link: "https://www.pulumi.com/"},
}

func getTechnologies(tt ...string) []entity.Technology {
	tt = lo.Uniq(tt)
	techs := make([]entity.Technology, 0, len(tt))
	for _, t := range tt {
		tech, found := lo.Find(technologies, func(tech entity.Technology) bool {
			return strings.EqualFold(tech.Name, t)
		})
		if !found {
			continue
		}

		techs = append(techs, tech)
	}
	return lo.Compact(techs)
}
