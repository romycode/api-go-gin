package bootstrap

import (
	"context"
	"time"

	_ "github.com/lib/pq"

	"github.com/romycode/go-api-template/internal/platform/server"
)

func Run() error {
	//db, err := sql.Open(
	//	"postgres",
	//	fmt.Sprintf(
	//		"postgresql://%s:%s@%s:%s/%s",
	//		os.Getenv("DB_USER"),
	//		os.Getenv("DB_PASS"),
	//		os.Getenv("DB_HOST"),
	//		os.Getenv("DB_PORT"),
	//		os.Getenv("DB_NAME")),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}

	ctx, srv := server.New(context.Background(), "0.0.0.0", 8080, time.Minute)
	return srv.Run(ctx)
}
