package static

import (
	"context"
	"time"

	"github.com/mgjules/mgjules-go/internal/entity"
)

func (db *static) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return []entity.Post{
		{
			Title: "Hello World (Lorem Ipsum style)",
			Slug:  "hello-world",
			Summary: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
      Donec id tincidunt tellus. Sed commodo euismod ligula, vel egestas ligula hendrerit vel. 
      Suspendisse sit amet tincidunt est, sed malesuada lacus. Nullam faucibus dui ipsum, ac cursus nulla mattis vel. 
      Vivamus neque leo, tincidunt vel mollis sed, tincidunt et arcu.`,
			CoverImage: "/img/blog/modern-code-screen.webp",
			Content: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec id tincidunt tellus. Sed commodo euismod ligula, vel egestas ligula hendrerit vel. Suspendisse sit amet tincidunt est, sed malesuada lacus. Nullam faucibus dui ipsum, ac cursus nulla mattis vel. Vivamus neque leo, tincidunt vel mollis sed, tincidunt et arcu. Suspendisse mattis mauris eu scelerisque lacinia. Praesent vestibulum massa quis sem varius dapibus. Pellentesque nec auctor sapien, eu laoreet leo. Vivamus tincidunt, est eu sodales imperdiet, purus purus luctus purus, vitae ultrices tortor lectus in elit. Praesent placerat ac nulla vitae tempor.

Fusce id pulvinar est, nec tempor magna. Sed sem ante, finibus at tortor sed, viverra egestas purus. Vestibulum ante orci, scelerisque quis fermentum nec, sagittis vitae neque. Praesent blandit luctus mauris. Nullam dui augue, tristique id viverra eu, sagittis non arcu. Maecenas tortor massa, ultrices quis ornare vel, condimentum et elit. Suspendisse bibendum mollis nulla, a placerat leo ultrices vel. Vivamus rhoncus lacinia viverra. Phasellus ac arcu vitae lorem hendrerit eleifend in vitae erat. Morbi viverra varius sodales. Praesent in libero sapien. Sed mauris eros, consequat nec ullamcorper mattis, pharetra nec lacus. Donec tempor accumsan arcu in tincidunt.

Mauris pellentesque condimentum nisi id feugiat. Mauris orci magna, vulputate in ornare eget, efficitur eleifend sem. Pellentesque et sapien sit amet felis commodo bibendum. In gravida, augue ac malesuada consectetur, ante nunc mattis lectus, eu lobortis dui diam eget nulla. Duis semper tellus id libero iaculis, interdum mollis ante mattis. Nullam at purus magna. Ut ipsum ipsum, hendrerit vitae luctus vel, euismod sit amet eros. Donec sit amet nisi mi. Vivamus condimentum, lectus eu consectetur cursus, justo dolor volutpat libero, ullamcorper scelerisque mauris orci ut massa. Pellentesque vulputate id elit ac sodales.

Curabitur ornare ut magna suscipit varius. Praesent molestie commodo elit, a mattis leo fermentum eu. Donec lacus ligula, vehicula quis mollis non, tristique in metus. Nullam vel commodo felis. Nulla facilisis mauris ut eros feugiat consequat. Cras sit amet fringilla libero, non facilisis dolor. In placerat gravida volutpat. Ut semper, magna condimentum varius dapibus, urna mi viverra justo, tincidunt pellentesque tellus neque a sem.

Nulla sodales viverra massa id iaculis. Integer a rhoncus sem, tincidunt tincidunt sapien. Sed iaculis dapibus felis, at auctor tortor interdum a. Donec consequat viverra nisl id feugiat. Quisque imperdiet, metus sed vulputate pellentesque, mauris purus tempor ligula, quis viverra odio nulla vel sem. Aliquam convallis ante id felis ultrices maximus. Sed eu dolor ac justo efficitur ultrices. Pellentesque ultricies euismod mi sit amet scelerisque. Mauris fringilla ullamcorper purus, id ultrices nibh bibendum nec. Integer ultricies, tellus et rhoncus consectetur, libero tellus aliquet ipsum, quis ultricies leo orci a sapien. Sed vitae tellus gravida nibh finibus interdum. Nullam et tempus dolor, sit amet imperdiet odio. Ut tincidunt, nunc vitae auctor placerat, orci quam malesuada nisi, eu finibus lectus felis ac ex. Mauris tincidunt mi augue, id tempus odio tempus vel. Nulla tortor urna, iaculis eget sapien ut, cursus cursus enim. Ut elementum, quam in viverra dictum, nibh diam rutrum ipsum, sed ullamcorper nisl ligula quis mauris.`,
			Tags: []entity.Tag{
				{
					ID:   "demo-tag",
					Name: "Demo Tag",
					Slug: "demo-tag",
				},
			},
			CreatedAt: time.Now(),
		},
	}, nil
}
