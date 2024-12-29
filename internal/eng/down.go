package eng

func Down() error {
	if err := DeletePidFile(); err != nil {
		return err
	}
	if err := DeleteSockFile(); err != nil {
		return err
	}
	return nil
}
