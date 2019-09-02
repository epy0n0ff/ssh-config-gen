package util

import (
	"fmt"

	"github.com/spf13/pflag"
)

func ValidateGenCommandArguments(flag *pflag.FlagSet) error {
	if !flag.Changed("toml") {
		return fmt.Errorf("toml argument not set")
	}

	if !flag.Changed("tpl") {
		return fmt.Errorf("tpl argument not set")
	}

	return nil
}
