package dockerstats

import (
	"os/exec"
)

// ----- Container IDs -----
// This function returns the ID's of all containers that currently exist on the host
// TODO: add better error handling
func getAllContainerIDs() ([]string, error) {
	raw, err := exec.Command("docker", "ps", "-aq").Output()
	arr, _ := stringToList(string(raw))
	return arr, err
}

// This function returns the ID's of all containers that are currently running on the host
// TODO: add better error handling
func getRunningContainerIDs() ([]string, error) {
	raw, err := exec.Command("docker", "ps", "-q").Output()
	arr, _ := stringToList(string(raw))
	return arr, err
}

// ----- Container Information -----

// Get the main status for the container
func getContainerStatus(container_id string) (Status, error) {
	// JSON ready format
	format := "{ \"Name\":\"{{.Name}}\", \"ID\":\"{{.ID}}\", \"CPUPerc\":\"{{.CPUPerc}}\", \"MemUsage\":\"{{.MemUsage}}\", \"NetIO\":\"{{.NetIO}}\", \"BlockIO\":\"{{.BlockIO}}\", \"MemPerc\":\"{{.MemPerc}}\", \"PIDs\":{{.PIDs}} }"
	dockerRes, err := exec.Command("docker", "stats", "--format", format, "--no-stream", container_id).Output()

	out, _ := buildStatusStruct(dockerRes)

	return out, err
}

// Get the state for the container from docker inspection
func getContainerState(container_id string) (Status, error) {
	format := "{{json .State}}"
	dockerRes, err := exec.Command("docker", "inspect", "--format", format, container_id).Output()

	out, _ := buildStateStruct(dockerRes)

	return out, err
}

// Get container healthcheck results if it is enabled for the container
func getContainerHealth(container_id string) (Health, error) {
	format := "{{json .State.Health}}"
	dockerRes, err := exec.Command("docker", "inspect", "--format", format, container_id).Output()

	out, _ := buildHealthStruct(dockerRes)

	return out, err
}