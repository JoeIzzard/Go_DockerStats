package dockerstats

import (
	"strings"
	"encoding/json"
	"strconv"
	"fmt"
)

// This function builds the exported Struct from the raw byte array that was captured
func buildStatusStruct(rawIn []byte) (Status, error) {
	// Unmarshal the JSON into a rawStruct for use in the next step
	var rawJson rawStatus
	json.Unmarshal(rawIn, &rawJson)

	// Construct the new struct
	var out Status

	// Basic Info
	out.ID = rawJson.ID
	out.Name = rawJson.Name

	// CPU
	out.CPU, _ = strconv.ParseFloat(strings.ReplaceAll(rawJson.CPUPerc, "%", ""), 8)

	// Memory
	var parts []string
	parts = strings.Split(string(rawJson.MemUsage), "/")
	out.Memory.Using.Base, out.Memory.Using.Value, _ = parseData(parts[0])
	out.Memory.Limit.Base, out.Memory.Limit.Value, _ = parseData(parts[1])
	out.Memory.Percent, _ = strconv.ParseFloat(strings.ReplaceAll(rawJson.MemPerc, "%", ""), 8)

	// Net IO
	parts = strings.Split(string(rawJson.NetIO), "/")
	out.NetIO.In.Base, out.NetIO.In.Value, _ = parseData(parts[0])
	out.NetIO.Out.Base, out.NetIO.Out.Value, _ = parseData(parts[1])

	// Block IO
	parts = strings.Split(string(rawJson.BlockIO), "/")
	out.BlockIO.In.Base, out.BlockIO.In.Value, _ = parseData(parts[0])
	out.BlockIO.Out.Base, out.BlockIO.Out.Value, _ = parseData(parts[1])

	// PIDs
	out.PIDs = rawJson.PIDs

	// TODO: Handle Errors better than this!	
	return out, nil
}

// This function builds the exported Struct from the raw byte array that was captured
func buildStateStruct(rawIn []byte) (Status, error) {
	// Unmarshal the JSON into a rawStruct for use in the next step
	var rawJson rawState
	json.Unmarshal(rawIn, &rawJson)


	// Construct the new struct
	var out Status

	// Status
	out.State.Message = rawJson.Status

	// Running
	out.State.Running = rawJson.Running

	// Paused
	out.State.Paused = rawJson.Paused

	// Restarting
	out.State.Restarting = rawJson.Restarting

	// OOMKilled
	out.State.OOMKilled = rawJson.OOMKilled

	// Dead
	out.State.Dead = rawJson.Dead

	// PID
	out.PID = rawJson.Pid

	// ExitCode
	out.ExitCode = rawJson.ExitCode

	// StartedAt
	out.StartedAt = rawJson.StartedAt

	// FinishedAt
	out.FinishedAt = rawJson.FinishedAt

	// TODO: Better error handling
	return out, nil
}

// This function builds the exported Struct from the raw byte array that was captured
func buildHealthStruct(rawIn []byte) (Health, error) {
	// Convert it to a string
	asString := string(rawIn)
	test := asString == "null"
	fmt.Print(test)

	// Test if HealthCheck is enabled
	if (!test) {
		var disabled Health
		disabled.Enabled = false
		disabled.Status = "Disabled"

		return disabled, nil
	}

	// Unmarshal the JSON into a rawStruct for use in the next step
	var rawJson rawHealth
	json.Unmarshal(rawIn, &rawJson)

	var out Health

	// Convert!
	out.Enabled = true
	out.FailStreak = rawJson.FailingStreak
	out.Status = rawJson.Status
	out.Log = rawJson.Log

	// TODO: Better error handling
	return out, nil
}

// This function combines the tree different Structs into a single main struct
func combineStructs(primary Status, state Status, health Health) (Status, error) {
	new := primary

	new.State = state.State
	new.PID = state.PID
	new.ExitCode = state.ExitCode
	new.StartedAt = state.StartedAt
	new.FinishedAt = state.FinishedAt
	new.Health = health
	
	// TODO: Better error handling
	return new, nil
}