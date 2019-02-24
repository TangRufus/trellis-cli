package cmd

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/roots/trellis-cli/trellis"
)

func TestVaultDecryptRunValidations(t *testing.T) {
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
			"too_many_args",
			true,
			[]string{"production", "foo"},
			"Error: too many arguments",
			1,
		},
	}

	for _, tc := range cases {
		mockProject := &MockProject{tc.projectDetected}
		trellis := trellis.NewTrellis(mockProject)
		vaultDecryptCommand := NewVaultDecryptCommand(ui, trellis)

		code := vaultDecryptCommand.Run(tc.args)

		if code != tc.code {
			t.Errorf("expected code %d to be %d", code, tc.code)
		}

		combined := ui.OutputWriter.String() + ui.ErrorWriter.String()

		if !strings.Contains(combined, tc.out) {
			t.Errorf("expected output %q to contain %q", combined, tc.out)
		}
	}
}

func TestVaultDecryptRun(t *testing.T) {
	ui := cli.NewMockUi()
	mockProject := &MockProject{true}
	trellis := trellis.NewTrellis(mockProject)
	vaultDecryptCommand := NewVaultDecryptCommand(ui, trellis)

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
			[]string{"production"},
			"ansible-vault decrypt group_vars/all/vault.yml group_vars/production/vault.yml",
			0,
		},
		{
			"files_flag_single_file",
			[]string{"--files=foo", "production"},
			"ansible-vault decrypt foo",
			0,
		},
		{
			"files_flag_multiple_file",
			[]string{"--files=foo,bar", "production"},
			"ansible-vault decrypt foo bar",
			0,
		},
	}

	for _, tc := range cases {
		code := vaultDecryptCommand.Run(tc.args)

		if code != tc.code {
			t.Errorf("expected code %d to be %d", code, tc.code)
		}

		combined := ui.OutputWriter.String() + ui.ErrorWriter.String()

		if !strings.Contains(combined, tc.out) {
			t.Errorf("expected output %q to contain %q", combined, tc.out)
		}
	}
}
