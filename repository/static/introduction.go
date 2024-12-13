package static

import (
	"context"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Static) GetIntroduction(ctx context.Context, id string) (*entity.Introduction, error) {
	return &entity.Introduction{
		ID: id,
		Introduction: `Highly accomplished Senior Software Engineer with over 10 years of experience in 
    backend development and 5+ years specializing in Go. Demonstrated track record of designing and 
    implementing high-performance software solutions, with a focus on optimizing system efficiency 
    and scalability. Proven leadership skills in managing complex projects and collaborating effectively 
    across cross-functional teams. 
    Passionate about continuous learning and staying at the forefront of industry trends and advancements. 
    Last but not least, a definite cat lover.`,
		Avatar: "/img/avatar.webp",
	}, nil
}
