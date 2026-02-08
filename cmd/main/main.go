package main

func main() {
	application, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	application.Run()
}
