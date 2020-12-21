package dockerstats

func GetContainer(container_id string) (Status, error) {
	// Get various container parts
	main, _ := getContainerStatus(container_id)
	state, _ := getContainerState(container_id)
	health, _ := getContainerHealth(container_id)

	// Combine the parts into a single final struct
	all, _ := combineStructs(main, state, health)

	return all, nil
}

func GetAllContainers() (map[string]Status, error) {
	// Get container ID's
	ids, _ := getContainerIDs()
	containers := make([string]Status)
	
	for _, id := range ids {
		containers[id], _ = GetContainer(id)
	}

	return containers, nil
}

func GetRunningContainers() (map[string]Status, error) {
	// Get container ID's
	ids, _ := getRunningContainerIDs()
	containers := make(map[string]Status)
	
	for _, id := range ids {
		container, _ := GetContainer(id)
		containers[id] = container
	}

	return containers, nil
}

func GetHealth(container_id string) (Health) {
	out, _ := getContainerHealth(container_id)
	return out
}