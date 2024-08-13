package main

var someVar string

func Init() error {
	someVar = "Initialized"
	return nil
}

func GetVar() string {
	return someVar
}
