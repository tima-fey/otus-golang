package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func successfulTask() error {
	return nil
}

func failedTask() error {
	return errors.New("error")
}
func successfulTaskSlow() error {
	time.Sleep(2 * time.Millisecond)
	return nil
}
func failedTaskSlow() error {
	time.Sleep(3 * time.Millisecond)
	return errors.New("error")
}
func TestTaskExecutorSuccess(t *testing.T) {
	taskChannel := make(chan func() error, 2)
	errorsChannel := make(chan func() error, 2)
	excexutorChannel := make(chan bool, 2)
	successfulChannel := make(chan func() error, 2)
	taskChannel <- successfulTask
	close(taskChannel)
	taskExecutor(taskChannel, errorsChannel, successfulChannel, excexutorChannel, 2)
	require.Equal(t, 1, len(successfulChannel))
	require.Equal(t, 0, len(errorsChannel))
}

func TestTaskExecutorFailed(t *testing.T) {
	taskChannel := make(chan func() error, 2)
	errorsChannel := make(chan func() error, 2)
	excexutorChannel := make(chan bool, 2)
	successfulChannel := make(chan func() error, 2)
	taskChannel <- failedTask
	close(taskChannel)
	taskExecutor(taskChannel, errorsChannel, successfulChannel, excexutorChannel, 2)
	require.Equal(t, 0, len(successfulChannel))
	require.Equal(t, 1, len(errorsChannel))
}

func TestTaskExecutorMixed(t *testing.T) {
	taskChannel := make(chan func() error, 3)
	errorsChannel := make(chan func() error, 3)
	excexutorChannel := make(chan bool, 3)
	successfulChannel := make(chan func() error, 3)
	taskChannel <- successfulTask
	taskChannel <- failedTask
	taskChannel <- successfulTask
	close(taskChannel)
	taskExecutor(taskChannel, errorsChannel, successfulChannel, excexutorChannel, 2)
	require.Equal(t, 2, len(successfulChannel))
	require.Equal(t, 1, len(errorsChannel))
}

func TestRun(t *testing.T) {
	test := []func() error{successfulTask, failedTask, successfulTask}
	suc, fail, err := run(test, 2, 2)
	require.Equal(t, 2, suc)
	require.Equal(t, 1, fail)
	require.Equal(t, nil, err)
}
func TestRunSlow(t *testing.T) {
	test := []func() error{successfulTaskSlow, failedTaskSlow, successfulTaskSlow}
	suc, fail, err := run(test, 2, 2)
	require.Equal(t, 2, suc)
	require.Equal(t, 1, fail)
	require.Equal(t, nil, err)
}
func TestRunFailed(t *testing.T) {
	test := []func() error{successfulTaskSlow, failedTaskSlow, successfulTaskSlow, successfulTaskSlow, failedTask, successfulTaskSlow}
	suc, fail, err := run(test, 2, 2)
	require.Equal(t, 3, suc)
	require.Equal(t, 2, fail)
	require.Error(t, err)
}

func TestRunBigN(t *testing.T) {
	test := []func() error{successfulTaskSlow, failedTaskSlow, successfulTaskSlow, successfulTaskSlow, failedTask, successfulTaskSlow}
	suc, fail, err := run(test, 8, 2)
	require.Equal(t, 4, suc)
	require.Equal(t, 2, fail)
	require.Error(t, err)
}

func TestRunBigN2(t *testing.T) {
	test := []func() error{successfulTaskSlow}
	suc, fail, err := run(test, 8, 2)
	require.Equal(t, 1, suc)
	require.Equal(t, 0, fail)
	require.Equal(t, nil, err)
}
