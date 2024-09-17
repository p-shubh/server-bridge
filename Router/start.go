package router

func Start() {

	r := SetupRouter()
	r.Run(":8080")
}
