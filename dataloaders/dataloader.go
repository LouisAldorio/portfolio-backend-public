package dataloaders

import (
	"context"
	"myapp/service"
	"net/http"
	"time"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type loaders struct {
	Education     *EducationLoader
	Photo         *PhotoLoader
	ImageCategory *ImagesCategoryLoader
	Subscriber    *SubsciberLoader
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		wait := 5 * time.Millisecond

		ldrs.Education = &EducationLoader{
			wait:     wait,
			maxBatch: 100,
			fetch:    service.GetEducationByWhereInCategoryID,
		}

		ldrs.Photo = &PhotoLoader{
			wait:     wait,
			maxBatch: 100,
			fetch:    service.GetPhotosByWhereInCategoryID,
		}

		ldrs.ImageCategory = &ImagesCategoryLoader{
			wait:     wait,
			maxBatch: 100,
			fetch:    service.GetImagesCategoryByWhereIn,
		}

		ldrs.Subscriber = &SubsciberLoader{
			wait: wait,
			maxBatch: 100,
			fetch: service.GetSubscriberByWhereIn,
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
