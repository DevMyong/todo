package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	//cfg := config.LoadConfig()
	//log, err := logger.NewZapLogger(cfg.BuildLoggerConfig())
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("Starting the application...")

	//userSvc := user.NewRegisterService(cfg.BuildUserConfig(),
	//	log, nil, nil, nil)

	//for {
	//	userSvc.RegisterLocal("", "", "")
	//	time.Sleep(5 * time.Second)
	//}

	//e := echo.New()
	//
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(200, "To-Do List Backend")
	//})
	//e.Logger.Fatal(e.Start(":8080"))
	//
	//var db database.Database
	//db = database.NewMongoDB()
	//err := db.Connect()
	//if err != nil {
	//	panic(err)
	//}

}

/*
- MSA 환경: 여러 마이크로서비스를 각각 독립된 컨테이너로 실행
- Docker + Kubernetes로 컨테이너 오케스트레이션
- AWS/GCP 등 클라우드 인프라 사용
- logger zap 라이브러리로 로그 생성, 생성된 로그는 k8s에서 파이프 라인으로 수집
*/
