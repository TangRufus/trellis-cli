package cmd

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/roots/trellis-cli/trellis"
)

func TestDeployRunValidations(t *testing.T) {
	ui := cli.NewMockUi()

	cases := []struct {
		name            string
		projectDetected bool
		args            []string
		out             string
		code            int
	}{
		{
			"no_project",
			false,
			nil,
			"No Trellis project detected",
			1,
		},
		{
			"no_args",
			true,
			nil,
			"Usage: trellis",
			1,
		},
		{
			"missing_site_arg",
			true,
			[]string{"development"},
			"Error: missing SITE argument",
			1,
		},
		{
			"too_many_args",
			true,
			[]string{"development", "site", "foo"},
			"Error: too many arguments",
			1,
		},
	}

	for _, tc := range cases {
		mockProject := &MockProject{tc.projectDetected}
		trellis := trellis.NewTrellis(mockProject)
		deployCommand := &DeployCommand{ui, trellis}

		code := deployCommand.Run(tc.args)

		if code != tc.code {
			t.Errorf("expected code %d to be %d", code, tc.code)
		}

		combined := ui.OutputWriter.String() + ui.ErrorWriter.String()

		if !strings.Contains(combined, tc.out) {
			t.Errorf("expected output %q to contain %q", combined, tc.out)
		}
	}
}

func TestDeployRun(t *testing.T) {
	ui := cli.NewMockUi()
	mockProject := &MockProject{true}
	trellis := trellis.NewTrellis(mockProject)
	deployCommand := &DeployCommand{ui, trellis}

	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	cases := []struct {
		name string
		args []string
		out  string
		code int
	}{
		{
			"default",
			[]string{"development", "example.com"},
			"./bin/deploy.sh development example.com",
			0,
		},
	}

	for _, tc := range cases {
		code := deployCommand.Run(tc.args)

		if code != tc.code {
			t.Errorf("expected code %d to be %d", code, tc.code)
		}

		combined := ui.OutputWriter.String() + ui.ErrorWriter.String()

		if !strings.Contains(combined, tc.out) {
			t.Errorf("expected output %q to contain %q", combined, tc.out)
		}
	}
}
