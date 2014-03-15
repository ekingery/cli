package main

import (
	"fmt"
	"github.com/exercism/cli/configuration"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func asserFileDoesNotExist(t *testing.T, filename string) {
	_, err := os.Stat(filename)

	if err == nil {
		t.Errorf("File [%s] already exist.", filename)
	}
}

func TestLogoutDeletesConfigFile(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	assert.NoError(t, err)

	c := configuration.Config{}

	configuration.ToFile(tmpDir, c)

	logout(tmpDir)

	asserFileDoesNotExist(t, configuration.Filename(tmpDir))
}

func TestAskForConfigInfoAllowsSpaces(t *testing.T) {
	oldStdin := os.Stdin
	dirName := "dirname with spaces"
	userName := "TestUsername"
	apiKey := "abc123"

	fakeStdin, err := ioutil.TempFile("", "stdin_mock")
	assert.NoError(t, err)

	fakeStdin.WriteString(fmt.Sprintf("%s\r\n%s\r\n%s\r\n", userName, apiKey, dirName))
	assert.NoError(t, err)

	file, err := os.Open(fakeStdin.Name())
	defer file.Close()

	os.Stdin = file

	c := askForConfigInfo()
	os.Stdin = oldStdin
	absoluteDirName, _ := absolutePath(dirName)
	_, err = os.Stat(absoluteDirName)
	if err != nil {
		t.Errorf("Excercism directory [%s] was not created.", absoluteDirName)
	}
	os.Remove(absoluteDirName)
	os.Remove(fakeStdin.Name())

	assert.Equal(t, c.ExercismDirectory, absoluteDirName)
	assert.Equal(t, c.GithubUsername, userName)
	assert.Equal(t, c.ApiKey, apiKey)
}
