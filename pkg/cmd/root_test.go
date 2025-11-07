package cmd

import (
	"bytes"
	"strings"
	"testing"
)

// TestRootFlags ensures all expected global flags are registered
func TestRootFlags(t *testing.T) {
	flags := rootCmd.PersistentFlags()

	expectedFlags := []string{"kubeconfig", "remote-context", "all-clusters", "namespace", "all-namespaces"}

	for _, name := range expectedFlags {
		if flags.Lookup(name) == nil {
			t.Errorf("expected persistent flag %q to be registered", name)
		}
	}
}

// TestRootSubcommands ensures all critical subcommands are registered
func TestRootSubcommands(t *testing.T) {
	subcmds := rootCmd.Commands()

	critical := []string{
		"get", "describe", "apply", "delete", "logs", "exec", "install",
	}

	registered := make(map[string]bool)
	for _, cmd := range subcmds {
		registered[cmd.Name()] = true
	}

	for _, name := range critical {
		if !registered[name] {
			t.Errorf("expected critical subcommand %q to be registered", name)
		}
	}
}

// TestRootHelpOutput checks that help text contains expected sections
func TestRootHelpOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing root command: %v", err)
	}

	output := buf.String()

	expectedStrings := []string{
		"kubectl-multi provides multi-cluster operations",
		"Usage:",
		"Examples:",
		"kubectl multi get pods",
	}

	for _, str := range expectedStrings {
		if !strings.Contains(output, str) {
			t.Errorf("help output missing expected string: %q", str)
		}
	}
}

// TestRootExecuteNoPanic ensures Execute() runs without panic for empty args
func TestRootExecuteNoPanic(t *testing.T) {
	rootCmd.SetArgs([]string{})
	err := rootCmd.Execute()
	if err != nil && !strings.Contains(err.Error(), "unknown command") {
		t.Errorf("expected no critical errors, got: %v", err)
	}
}

// TestRootExecuteInvalidSubcommand ensures invalid subcommand returns error
func TestRootExecuteInvalidSubcommand(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"invalidcmd"})
	err := rootCmd.Execute()
	if err == nil {
		t.Errorf("expected error for invalid subcommand, got nil")
	} else if !strings.Contains(err.Error(), "unknown command") {
		t.Errorf("unexpected error message: %v", err)
	}
}
