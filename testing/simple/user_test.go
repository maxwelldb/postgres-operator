package smoketest

import (
	//meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os/exec"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	var clientset *kubernetes.Clientset

	// t.Fatal("not implemented")
	t.Run("setup", func(t *testing.T) {
		t.Log("some setup code")
		clientset, _ = SetupKube()
		if clientset == nil {
			t.Error("clientset is nil")
		}
	})

	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"pgo create user", []string{"create", "user", "--managed", "--selector=name=" + TestClusterName, "--username=testuser1"}, ""},
	}

	t.Log("TestCreateUser starts")
	for _, tt := range tests {
		cmd := exec.Command("pgo", tt.args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			//t.Fatal(err)
		}

		actual := string(output)

		t.Logf("actual %s- ", actual)
		found := strings.Contains(actual, "adding new user")
		if !found {
			t.Error("could not find adding new user in output")
		}

	}

	t.Run("teardown", func(t *testing.T) {
		t.Log("some teardown code")
	})
}

func TestUpdateUser(t *testing.T) {
	var clientset *kubernetes.Clientset

	// t.Fatal("not implemented")
	t.Run("setup", func(t *testing.T) {
		t.Log("some setup code")
		clientset, _ = SetupKube()
		if clientset == nil {
			t.Error("clientset is nil")
		}
	})

	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"pgo update user", []string{"update", "user", "--username=testuser1", "--selector=name=" + TestClusterName, "--password=fungo"}, ""},
	}

	t.Log("TestUpdateUser starts")
	for _, tt := range tests {
		cmd := exec.Command("pgo", tt.args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			//t.Fatal(err)
		}

		actual := string(output)

		t.Logf("actual %s- ", actual)
		found := strings.Contains(actual, "changing")
		if !found {
			t.Error("could not find update in output")
		}

	}

	t.Run("teardown", func(t *testing.T) {
		t.Log("some teardown code")
	})
}
func TestShowUser(t *testing.T) {
	var clientset *kubernetes.Clientset

	// t.Fatal("not implemented")
	t.Run("setup", func(t *testing.T) {
		t.Log("some setup code")
		clientset, _ = SetupKube()
		if clientset == nil {
			t.Error("clientset is nil")
		}
	})

	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"pgo show user", []string{"show", "user", "--selector=name=" + TestClusterName}, ""},
	}

	t.Log("TestShowUser starts")
	for _, tt := range tests {
		cmd := exec.Command("pgo", tt.args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			//t.Fatal(err)
		}

		actual := string(output)

		t.Logf("actual %s- ", actual)
		found := strings.Contains(actual, "secret")
		if !found {
			t.Error("could not find secret in output")
		}

	}

	t.Run("teardown", func(t *testing.T) {
		t.Log("some teardown code")
	})
}

func TestDeleteUser(t *testing.T) {
	var clientset *kubernetes.Clientset

	// t.Fatal("not implemented")
	t.Run("setup", func(t *testing.T) {
		t.Log("some setup code")
		clientset, _ = SetupKube()
		if clientset == nil {
			t.Error("clientset is nil")
		}
	})

	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"pgo delete user", []string{"delete", "user", "--username=testuser1", "--selector=name=" + TestClusterName, "--no-prompt"}, ""},
	}

	t.Log("TestCreateUser starts")
	for _, tt := range tests {
		cmd := exec.Command("pgo", tt.args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			//t.Fatal(err)
		}

		actual := string(output)

		t.Logf("actual %s- ", actual)
		found := strings.Contains(actual, "removed")
		if !found {
			t.Error("could not find delete in output")
		}

	}

	t.Run("teardown", func(t *testing.T) {
		t.Log("some teardown code")
	})
}

//TODO: need to add pgo show user
