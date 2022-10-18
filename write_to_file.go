func WriteFile(output string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(Strip(output))
		fmt.Println("Done")
	}
	file.Close()
}