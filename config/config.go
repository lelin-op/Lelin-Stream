// Updated loadFromEnvFile function to make fsb.env optional.

func loadFromEnvFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("INFO: Env file does not exist, falling back to environment variables.")
		return
	}
	// Existing loading logic...
}
